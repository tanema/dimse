package conn

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"time"

	"github.com/suyashkumar/dicom"
	"github.com/tanema/dimse/src/chunkreader"
	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/pdu"
	"github.com/tanema/dimse/src/serviceobjectpair"
	"github.com/tanema/dimse/src/transfersyntax"
)

type (
	Conn struct {
		ctx    context.Context
		conn   net.Conn
		events chan readResult
	}
	readResult struct {
		evt pdu.PDU
		err error
	}
)

func Check(addr string) error {
	testconn, err := net.Dial("tcp", addr)
	defer testconn.Close()
	return err
}

func Connect(ctx context.Context, addr string, timeout time.Duration) (*Conn, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}
	c := &Conn{
		ctx:    ctx,
		conn:   conn,
		events: make(chan readResult),
	}
	go c.listen()
	return c, nil
}

func (c *Conn) listen() {
	for {
		data := make([]byte, pdu.DefaultMaxPDUSize)
		n, err := c.conn.Read(data)
		if err != nil {
			return
		}
		if n > 0 {
			pdu, err := pdu.ReadPDU(bytes.NewBuffer(data[:n]))
			// fmt.Printf("recv: %vb %s\n", n, pdu)
			c.events <- readResult{evt: pdu, err: err}
		}
	}
}

func (c *Conn) Read() (pdu.PDU, error) {
	select {
	case evt := <-c.events:
		return evt.evt, evt.err
	case <-c.ctx.Done():
		return nil, c.ctx.Err()
	}
}

func (c *Conn) Send(msg pdu.PDU) error {
	data, err := pdu.EncodePDU(msg)
	if err != nil {
		return err
	}
	// fmt.Printf("send: %vb %s\n", len(data), msg)
	_, err = c.conn.Write(data)
	return err
}

func (c *Conn) Close() error {
	close(c.events)
	return c.conn.Close()
}

func (c *Conn) Associate(sopsClasses []serviceobjectpair.UID, transfersyntaxes []transfersyntax.UID) (*pdu.ContextManager, error) {
	if len(transfersyntaxes) == 0 {
		transfersyntaxes = transfersyntax.StandardSyntaxes
	}

	assocPDI, ctxManager := pdu.CreateAssoc(sopsClasses, transfersyntaxes)
	if err := c.Send(assocPDI); err != nil {
		return nil, err
	}
	evt, err := c.Read()
	if err != nil {
		return nil, err
	}
	switch pt := evt.(type) {
	case *pdu.AAssociate:
		for _, item := range pt.Items {
			if pcu, ok := item.(*pdu.PresentationContextItem); ok && pcu.Result == pdu.PresentationContextAccepted {
				ts := pcu.Items[0].(*pdu.TransferSyntaxSubItem)
				ctxManager.Accept(pcu.ContextID, transfersyntax.UID(ts.Name))
			}
		}
		return ctxManager, nil
	case *pdu.AAssociateRj:
		return nil, fmt.Errorf("association rejected %v", pt.Reason)
	default:
		return nil, fmt.Errorf("unexpected message %T after sending associate", evt)
	}
}

func (c *Conn) Realease() error {
	if err := c.Send(pdu.CreateRelease()); err != nil {
		return err
	}
	evt, err := c.Read()
	if err != nil {
		return err
	}
	switch evt.(type) {
	case *pdu.AReleaseRp:
		return nil
	default:
		return fmt.Errorf("unexpected message %T after sending release", evt)
	}
}

func (c *Conn) Abort() { c.Send(pdu.CreateAbort()) }

func (c *Conn) Pdata(ctxMan *pdu.ContextManager, cmd *commands.Command, payload []byte) (*commands.Command, []dicom.Dataset, error) {
	var ctxID uint8
	//var ts transfersyntax.UID
	for _, classUID := range cmd.AffectedSOPClassUID {
		if pctx, err := ctxMan.GetWithSOP(classUID); err == nil {
			ctxID = pctx.ContextID
			//ts = pctx.AcceptedTransferSyntax
			break
		}
	}

	// This should not really ever happen because we collect sops from the command
	// and feed it into the associate call. The only time this could happen is if
	// the server rejected a specific presentation context.
	if ctxID == 0 {
		return nil, nil, fmt.Errorf("Could not find an associated presentation context item for command which means the server rejected the AffectedSOPClassUID you requested.")
	}

	value, err := commands.Encode(cmd, transfersyntax.ImplicitVRLittleEndian)
	if err != nil {
		return nil, nil, err
	}
	// encode the command first and then send data along
	pdatas := pdu.CreatePdata(ctxID, true, value)
	if cmd.CommandDataSetType != commands.Null {
		pdatas = append(pdatas, pdu.CreatePdata(ctxID, false, payload)...)
	}
	for _, pd := range pdatas {
		if err := c.Send(pd); err != nil {
			return nil, nil, err
		}
	}

	return c.readPData()
}

func (c *Conn) readPData() (*commands.Command, []dicom.Dataset, error) {
	var cmd *commands.Command
	sets := []dicom.Dataset{}
	for {
		evt, err := c.Read()
		if err != nil {
			return nil, nil, err
		}
		switch tevt := evt.(type) {
		case *pdu.PDataTf:
			for _, item := range tevt.Items {
				if item.Command {
					cmd, err = commands.Decode(item.Value, transfersyntax.ImplicitVRLittleEndian)
					if err != nil {
						return nil, nil, err
					} else if cmd.CommandDataSetType == commands.Null {
						return cmd, nil, nil
					}
				} else {
					payload, err := decode(item.Value, false)
					if err != nil {
						return nil, nil, err
					}
					sets = append(sets, payload)
					if item.Last {
						return cmd, sets, nil
					}
				}
			}
		case *pdu.AAbort:
			return nil, nil, fmt.Errorf("aborted pdata. Reason: %s Source: %s", tevt.Reason, tevt.Source)
		default:
			return nil, nil, fmt.Errorf("unexpected message %T after sending release", evt)
		}
	}
}

func decode(data []byte, implicit bool) (dicom.Dataset, error) {
	r := chunkreader.New()
	if err := r.Decode(data, implicit); err != nil {
		return dicom.Dataset{}, err
	}
	return r.Dataset(), nil
}
