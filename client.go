package dimse

import (
	"bytes"
	"fmt"
	"net"
	"slices"
	"sync/atomic"

	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/pdu"
	"github.com/tanema/dimse/src/serviceobjectpair"
	"github.com/tanema/dimse/src/transfersyntax"
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
			c.events <- readResult{evt: pdu, err: err}
		}
	}
}

func (c *Client) sendPDU(msg pdu.PDU) error {
	data, err := pdu.EncodePDU(msg)
	if err != nil {
		return err
	}
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

func (c *Client) Dispatch(cmd *Command) (*Command, error) {
	sops := collectSOPs(cmd)
	ctxMan, err := c.associate(sops, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.pdata(ctxMan, cmd)
	if err != nil {
		return nil, err
	}
	return resp, c.realease()
}

func (c *Client) associate(sopsClasses []string, transfersyntaxes []string) (*pdu.ContextManager, error) {
	if len(transfersyntaxes) == 0 {
		transfersyntaxes = transfersyntax.StandardSyntaxes
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

func (c *Client) pdata(ctxMan *pdu.ContextManager, cmd *Command) (*Command, error) {
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
		return nil, fmt.Errorf("Could not find an associated presentation context item for command which means the server rejected the AffectedSOPClassUID you requested.")
	}

	value, err := EncodeCmd(cmd)
	if err != nil {
		return nil, err
	}
	// TODO need ContextID, and to split up into multiple
	pdatas := pdu.CreatePdata(ctxID, value)
	for _, pd := range pdatas {
		if err := c.sendPDU(pd); err != nil {
			return nil, err
		}
	}
	for {
		evt := <-c.events
		if evt.err != nil {
			return nil, evt.err
		}
		switch tevt := evt.evt.(type) {
		case *pdu.PDataTf:
			for _, item := range tevt.Items {
				cmd, err := DecodeCmd(item.Value)
				return cmd, err
			}
			return nil, nil
		case *pdu.AAbort:
			return nil, fmt.Errorf("aborted pdata. Reason: %s Source: %s", tevt.Reason, tevt.Source)
		default:
			return nil, fmt.Errorf("unexpected message %T after sending release", evt.evt)
		}
	}
}

func (c *Client) Echo() error {
	msgID := int(c.nextMsgID())
	resp, err := c.Dispatch(&Command{
		CommandField:        commands.CECHORQ,
		MessageID:           msgID,
		AffectedSOPClassUID: serviceobjectpair.VerificationClasses,
		CommandDataSetType:  commands.Null,
	})
	if err != nil {
		return err
	}
	if resp.CommandField != commands.CECHORSP {
		return fmt.Errorf("received %s in response to echo", resp.CommandField)
	} else if resp.MessageID != msgID {
		return fmt.Errorf("received %v message id but sent %v", resp.MessageID, msgID)
	}
	return nil
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
