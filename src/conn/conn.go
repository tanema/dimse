package conn

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/suyashkumar/dicom"

	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/encoding"
	"github.com/tanema/dimse/src/pdu"
	"github.com/tanema/dimse/src/serviceobjectpair"
	"github.com/tanema/dimse/src/transfersyntax"
)

type (
	Conn struct {
		ctx    context.Context
		conn   net.Conn
		events chan readResult
		cfg    Config
	}
	Config struct {
		MaxConnections    int
		ConnectionTimeout time.Duration
		ChunkSize         uint32
		AETitle           string
	}
	readResult struct {
		evt any
		err error
	}
)

var DefaultConfig = &Config{
	MaxConnections:    10,
	ConnectionTimeout: time.Second,
	ChunkSize:         pdu.DefaultMaxPDUSize,
	AETitle:           "anon-ae",
}

func Connect(ctx context.Context, addr string, cfg Config) (*Conn, error) {
	conn, err := net.DialTimeout("tcp", addr, cfg.ConnectionTimeout)
	if err != nil {
		return nil, err
	}
	c := &Conn{
		ctx:    ctx,
		conn:   conn,
		events: make(chan readResult),
		cfg:    cfg,
	}
	return c, nil
}

func (c *Conn) Read() (any, error) {
	return pdu.NewReader(c.conn).Next()
}

func (c *Conn) Send(msg any) error {
	data, err := pdu.EncodePDU(msg)
	if err != nil {
		return err
	}
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

	assocPDI, ctxManager := pdu.CreateAssoc(c.cfg.AETitle, c.cfg.ChunkSize, sopsClasses, transfersyntaxes)
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

	allDataSets := []dicom.Dataset{}
	for {
		cmd, ds, err := c.readPData()
		if err != nil {
			return nil, nil, err
		}
		allDataSets = append(allDataSets, ds...)
		if cmd.Status != commands.Pending {
			return cmd, allDataSets, nil
		}
	}
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
	return encoding.NewReader(bytes.NewBuffer(data), binary.LittleEndian).Decode(implicit)
}
