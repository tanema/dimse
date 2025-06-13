package dimse

import (
	"bytes"
	"fmt"
	"net"
	"slices"
	"sync/atomic"

	"github.com/tanema/dimse/obj"
	"github.com/tanema/dimse/obj/commands"
	"github.com/tanema/dimse/pdu"
)

type (
	Client struct {
		addr   string
		events chan readResult
		conn   net.Conn
		msgID  int32
	}
	readResult struct {
		evt pdu.PDU
		err error
	}
)

func Connect(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	client := &Client{
		conn:   conn,
		events: make(chan readResult),
	}
	go client.listen()
	return client, nil
}

func (c *Client) nextMsgID() int32 {
	return atomic.AddInt32(&c.msgID, 1)
}

func (c *Client) listen() {
	for {
		data := make([]byte, 4096)
		n, err := c.conn.Read(data)
		if err != nil {
			return
		}
		if n > 0 {
			pdu, err := pdu.ReadPDU(bytes.NewBuffer(data[:n]))
			fmt.Printf("recv: %T %v\n", pdu, err)
			c.events <- readResult{evt: pdu, err: err}
		}
	}
}

func (c *Client) sendPDU(msg pdu.PDU) error {
	data, err := pdu.EncodePDU(msg)
	if err != nil {
		return err
	}
	fmt.Printf("send: %T\n", msg)
	return c.send(data)
}

func (c *Client) send(body []byte) error {
	_, err := c.conn.Write(body)
	return err
}

func (c *Client) Close() error {
	close(c.events)
	return c.conn.Close()
}

func (c *Client) Dispatch(cmd *Command) error {
	sops := collectSOPs(cmd)
	if ctxMan, err := c.associate(sops, nil); err != nil {
		return err
	} else if err := c.pdata(ctxMan, cmd); err != nil {
		return err
	}
	return c.realease()
}

func (c *Client) associate(sopsClasses []string, transfersyntaxes []string) (*pdu.ContextManager, error) {
	if len(transfersyntaxes) == 0 {
		transfersyntaxes = obj.StandardTransferSyntaxes
	}

	assocPDI, ctxManager := pdu.CreateAssoc(sopsClasses, transfersyntaxes)
	if err := c.sendPDU(assocPDI); err != nil {
		return nil, err
	}
	evt := <-c.events
	if evt.err != nil {
		return nil, evt.err
	}
	switch pt := evt.evt.(type) {
	case *pdu.AAssociate:
		for _, item := range pt.Items {
			if pcu, ok := item.(*pdu.PresentationContextItem); ok && pcu.Result == pdu.PresentationContextAccepted {
				ctxManager.Accept(pcu.ContextID)
			}
		}
		return ctxManager, nil
	case *pdu.AAssociateRj:
		return nil, fmt.Errorf("association rejected %v", pt.Reason)
	default:
		return nil, fmt.Errorf("unexpected message %T after sending associate", evt.evt)
	}
}

func (c *Client) realease() error {
	if err := c.sendPDU(pdu.CreateRelease()); err != nil {
		return err
	}
	evt := <-c.events
	if evt.err != nil {
		return evt.err
	}
	switch evt.evt.(type) {
	case *pdu.AReleaseRp:
		return nil
	default:
		return fmt.Errorf("unexpected message %T after sending release", evt.evt)
	}
}

func (c *Client) pdata(ctxMan *pdu.ContextManager, cmd *Command) error {
	var ctxID uint8
	for _, classUID := range cmd.AffectedSOPClassUID {
		if pctx, err := ctxMan.GetWithSOP(classUID); err == nil {
			ctxID = pctx.ContextID
			break
		}
	}

	// This should not really ever happen because we collect sops from the command
	// and feed it into the associate call. The only time this could happen is if
	// the server rejected a specific presentation context.
	if ctxID == 0 {
		return fmt.Errorf("Could not find an associated presentation context item for command which means the server rejected the AffectedSOPClassUID you requested.")
	}

	value, err := EncodeCmd(cmd)
	if err != nil {
		return err
	}
	// TODO need ContextID, and to split up into multiple
	pdatas := pdu.CreatePdata(ctxID, value)
	for _, pd := range pdatas {
		if err := c.sendPDU(pd); err != nil {
			return err
		}
	}
	for {
		evt := <-c.events
		if evt.err != nil {
			return evt.err
		}
		switch tevt := evt.evt.(type) {
		case *pdu.PDataTf:
			for _, item := range tevt.Items {
				fmt.Printf("ContextID: %v Command: %v Last: %v Value: %v\n", item.ContextID, item.Command, item.Last, len(item.Value))
				cmd, err := DecodeCmd(item.Value)
				fmt.Println(cmd, err)
			}
			return nil
		case *pdu.AAbort:
			return fmt.Errorf("aborted pdata. Reason: %s Source: %s", tevt.Reason, tevt.Source)
		default:
			return fmt.Errorf("unexpected message %T after sending release", evt.evt)
		}
	}
}

func (c *Client) Echo() error {
	return c.Dispatch(&Command{
		CommandField:        commands.CECHORQ,
		MessageID:           int(c.nextMsgID()),
		AffectedSOPClassUID: obj.VerificationClasses,
		CommandDataSetType:  commands.Null,
	})
}

// collectSOPs will collect SOPs from all commands to be put into the association
// request, and ensure that they are unique.
func collectSOPs(cmds ...*Command) []string {
	sops := []string{}
	for _, cmd := range cmds {
		sops = append(sops, cmd.AffectedSOPClassUID...)
	}
	slices.Sort(sops)
	return slices.Compact(sops)
}
