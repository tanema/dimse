package dimse

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"slices"
	"sync/atomic"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/pdu"
	"github.com/tanema/dimse/src/query"
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

func NewClient(addr string) (*Client, error) {
	testconn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	testconn.Close()

	return &Client{
		addr:   addr,
		events: make(chan readResult),
	}, nil
}

func (c *Client) connect() error {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return err
	}
	c.conn = conn
	go c.listen(conn)
	return nil
}

func (c *Client) disconnect() error {
	err := c.conn.Close()
	c.conn = nil
	return err
}

func (c *Client) nextMsgID() int32 {
	return atomic.AddInt32(&c.msgID, 1)
}

func (c *Client) listen(conn net.Conn) {
	for {
		data := make([]byte, pdu.DefaultMaxPDUSize)
		n, err := conn.Read(data)
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

func (c *Client) sendPDU(msg pdu.PDU) error {
	data, err := pdu.EncodePDU(msg)
	if err != nil {
		return err
	}
	// fmt.Printf("send: %vb %s\n", len(data), msg)
	return c.send(data)
}

func (c *Client) send(body []byte) error {
	if c.conn == nil {
		return fmt.Errorf("client is not connected to server")
	}
	_, err := c.conn.Write(body)
	return err
}

func (c *Client) Close() error {
	close(c.events)
	return nil
}

func (c *Client) dispatch(cmd *Command, payload []byte) (*Command, []dicom.Dataset, error) {
	if cmd.CommandDataSetType != commands.Null && len(payload) == 0 {
		return nil, nil, fmt.Errorf("empty payload provided to a command that requires a payload")
	} else if err := c.connect(); err != nil {
		return nil, nil, err
	}
	defer c.disconnect()
	ctxMan, err := c.associate(collectSOPs(cmd), nil)
	if err != nil {
		return nil, nil, err
	}
	respCmd, data, err := c.pdata(ctxMan, cmd, payload)
	if err != nil {
		c.abort()
		return nil, nil, err
	} else if err := c.realease(); err != nil {
		return nil, nil, err
	} else if cmd.MessageID != respCmd.MessageID {
		return nil, nil, fmt.Errorf("received %v message id but sent %v", cmd.MessageID, respCmd.MessageID)
	}
	return respCmd, data, nil
}

func (c *Client) associate(sopsClasses []serviceobjectpair.UID, transfersyntaxes []transfersyntax.UID) (*pdu.ContextManager, error) {
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
				ts := pcu.Items[0].(*pdu.TransferSyntaxSubItem)
				ctxManager.Accept(pcu.ContextID, transfersyntax.UID(ts.Name))
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

func (c *Client) abort() {
	c.sendPDU(pdu.CreateAbort())
}

func (c *Client) pdata(ctxMan *pdu.ContextManager, cmd *Command, payload []byte) (*Command, []dicom.Dataset, error) {
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

	value, err := EncodeCmd(cmd, transfersyntax.ImplicitVRLittleEndian)
	if err != nil {
		return nil, nil, err
	}
	// encode the command first and then send data along
	pdatas := pdu.CreatePdata(ctxID, true, value)
	if cmd.CommandDataSetType != commands.Null {
		pdatas = append(pdatas, pdu.CreatePdata(ctxID, false, payload)...)
	}
	for _, pd := range pdatas {
		if err := c.sendPDU(pd); err != nil {
			return nil, nil, err
		}
	}

	return c.readPData()
}

func (c *Client) readPData() (*Command, []dicom.Dataset, error) {
	var cmd *Command
	var err error
	sets := []dicom.Dataset{}
	for {
		evt := <-c.events
		if evt.err != nil {
			return nil, nil, evt.err
		}
		switch tevt := evt.evt.(type) {
		case *pdu.PDataTf:
			for _, item := range tevt.Items {
				if item.Command {
					cmd, err = DecodeCmd(item.Value, transfersyntax.ImplicitVRLittleEndian)
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
			return nil, nil, fmt.Errorf("unexpected message %T after sending release", evt.evt)
		}
	}
}

// Echo will issue an echo command and will return an error if something went wrong.
// No error will be returned if the error command returned successfully.
func (c *Client) Echo() error {
	msgID := int(c.nextMsgID())
	resp, _, err := c.dispatch(&Command{
		CommandField:        commands.CECHORQ,
		MessageID:           msgID,
		AffectedSOPClassUID: serviceobjectpair.VerificationClasses,
		CommandDataSetType:  commands.Null,
	}, nil)
	if err != nil {
		return err
	} else if resp.CommandField != commands.CECHORSP {
		return fmt.Errorf("received %s in response to echo", resp.CommandField)
	}
	return nil
}

func (c *Client) Query(level query.Level, q []*dicom.Element) (*Query, error) {
	if len(q) == 0 {
		return nil, fmt.Errorf("Query: empty query")
	}
	query := &Query{
		client: c,
		Level:  level,
		Filter: q,
	}
	return query, query.encodePayload()
}

func (q *Query) SetPriority(p int) *Query {
	q.Priority = p
	return q
}

func (q *Query) Find() ([]dicom.Dataset, error) {
	msgID := int(q.client.nextMsgID())
	resp, data, err := q.client.dispatch(&Command{
		CommandField:        commands.CFINDRQ,
		MessageID:           msgID,
		AffectedSOPClassUID: []serviceobjectpair.UID{q.sopForCmd(commands.CFINDRQ)},
		CommandDataSetType:  commands.NonNull,
		Priority:            q.Priority,
	}, q.payload)
	if err != nil {
		return nil, err
	} else if resp.CommandField != commands.CFINDRSP {
		return nil, fmt.Errorf("received %s in response to find", resp.CommandField)
	}
	return data, nil
}

func (q *Query) Get() ([]dicom.Dataset, error) {
	resp, data, err := q.client.dispatch(&Command{
		CommandField:        commands.CGETRQ,
		MessageID:           int(q.client.nextMsgID()),
		AffectedSOPClassUID: []serviceobjectpair.UID{q.sopForCmd(commands.CGETRQ)},
		CommandDataSetType:  commands.NonNull,
		Priority:            q.Priority,
	}, q.payload)
	if err != nil {
		return nil, err
	} else if resp.CommandField != commands.CGETRSP {
		return nil, fmt.Errorf("received %s in response to find", resp.CommandField)
	}
	return data, nil
}

func (q *Query) Move(dst string) ([]dicom.Dataset, error) {
	resp, data, err := q.client.dispatch(&Command{
		CommandField:        commands.CMOVERQ,
		MessageID:           int(q.client.nextMsgID()),
		AffectedSOPClassUID: []serviceobjectpair.UID{q.sopForCmd(commands.CMOVERQ)},
		Priority:            q.Priority,
		MoveDestination:     dst,
		CommandDataSetType:  commands.NonNull,
	}, q.payload)
	if err != nil {
		return nil, err
	} else if resp.CommandField != commands.CMOVERSP {
		return nil, fmt.Errorf("received %s in response to find", resp.CommandField)
	}
	return data, nil
}

func (q *Query) Store(inst []serviceobjectpair.UID, id int, dst, title string) ([]dicom.Dataset, error) {
	resp, data, err := q.client.dispatch(&Command{
		CommandField:                         commands.CSTORERQ,
		MessageID:                            int(q.client.nextMsgID()),
		AffectedSOPClassUID:                  []serviceobjectpair.UID{q.sopForCmd(commands.CMOVERQ)},
		CommandDataSetType:                   commands.NonNull,
		Priority:                             q.Priority,
		MoveDestination:                      dst,
		AffectedSOPInstanceUID:               inst,
		MoveOriginatorApplicationEntityTitle: title,
		MoveOriginatorMessageID:              id,
	}, q.payload)
	if err != nil {
		return nil, err
	} else if resp.CommandField != commands.CMOVERSP {
		return nil, fmt.Errorf("received %s in response to find", resp.CommandField)
	}
	return data, nil
}

// collectSOPs will collect SOPs from all commands to be put into the association
// request, and ensure that they are unique.
func collectSOPs(cmds ...*Command) []serviceobjectpair.UID {
	sops := []serviceobjectpair.UID{}
	for _, cmd := range cmds {
		sops = append(sops, cmd.AffectedSOPClassUID...)
	}
	slices.Sort(sops)
	return slices.Compact(sops)
}

func (q *Query) sopForCmd(kind commands.Kind) serviceobjectpair.UID {
	switch q.Level {
	case query.Patient:
		switch kind {
		case commands.CFINDRQ:
			return serviceobjectpair.PatientRootQueryRetrieveInformationModelFind
		case commands.CGETRQ:
			return serviceobjectpair.PatientRootQueryRetrieveInformationModelGet
		case commands.CMOVERQ:
			return serviceobjectpair.PatientRootQueryRetrieveInformationModelMove
		}
	case query.Study, query.Series:
		switch kind {
		case commands.CFINDRQ:
			return serviceobjectpair.StudyRootQueryRetrieveInformationModelFind
		case commands.CGETRQ:
			return serviceobjectpair.StudyRootQueryRetrieveInformationModelGet
		case commands.CMOVERQ:
			return serviceobjectpair.StudyRootQueryRetrieveInformationModelMove
		}
	}
	return ""
}

func (q *Query) encodePayload() error {
	foundQRLevel := false
	buf := bytes.NewBuffer([]byte{})
	w, err := dicom.NewWriter(buf)
	if err != nil {
		return err
	}
	w.SetTransferSyntax(binary.LittleEndian, true)
	for _, elem := range q.Filter {
		if elem.Tag == tag.QueryRetrieveLevel {
			foundQRLevel = true
		}
		if err := w.WriteElement(elem); err != nil {
			return err
		}
	}
	if !foundQRLevel {
		var qrLevelString string
		switch q.Level {
		case query.Patient:
			qrLevelString = "PATIENT"
		case query.Study:
			qrLevelString = "STUDY"
		case query.Series:
			qrLevelString = "SERIES"
		}
		elem, err := dicom.NewElement(tag.QueryRetrieveLevel, []string{qrLevelString})
		if err != nil {
			return err
		}
		if err := w.WriteElement(elem); err != nil {
			return err
		}
	}
	q.payload = buf.Bytes()
	return nil
}
