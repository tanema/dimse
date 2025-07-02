package dimse

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"sync/atomic"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"

	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/defn/abort"
	"github.com/tanema/dimse/src/defn/presentationctx"
	"github.com/tanema/dimse/src/defn/serviceobjectpair"
	"github.com/tanema/dimse/src/defn/source"
	"github.com/tanema/dimse/src/defn/status"
	"github.com/tanema/dimse/src/defn/transfersyntax"
	"github.com/tanema/dimse/src/encoding"
	"github.com/tanema/dimse/src/pdu"
)

type Conn struct {
	conn       net.Conn
	msgID      int32
	aeTitle    string
	entity     *Entity
	cfg        *ConnectionConfig
	ctxManager *pdu.ContextManager
}

func Connect(ctx context.Context, aetitle string, entity Entity, cfg *ConnectionConfig) (*Conn, error) {
	addr := fmt.Sprintf("%v:%v", entity.Host, entity.Port)
	dialer := net.Dialer{Timeout: cfg.Timeout}
	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Conn{
		aeTitle: aetitle,
		entity:  &entity,
		conn:    conn,
		cfg:     cfg,
	}, nil
}

func (c *Conn) Read() (any, error) {
	return pdu.NewReader(c.conn).Next()
}

func (c *Conn) Send(msg any, bo binary.ByteOrder) error {
	data, err := pdu.EncodePDU(msg, bo)
	if err != nil {
		return err
	}
	_, err = c.conn.Write(data)
	return err
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

func (c *Conn) Associate(sopsClasses []serviceobjectpair.UID, ts []transfersyntax.UID) error {
	assocPDI, ctxManager := pdu.CreateAssoc(c.aeTitle, c.entity.Title, c.cfg.ChunkSize, sopsClasses, ts)
	if err := c.Send(assocPDI, binary.BigEndian); err != nil {
		return err
	}
	evt, err := c.Read()
	if err != nil {
		return err
	}
	switch pt := evt.(type) {
	case *pdu.AAssociate:
		for _, pcu := range pt.PresentationItems {
			if pcu.Result == presentationctx.Accepted {
				ctxManager.Accept(pcu.ContextID, transfersyntax.UID(pcu.TransferSyntaxes[0]))
			}
		}
		c.cfg.ChunkSize = pt.MaximumLengthReceived
		c.ctxManager = ctxManager
		return nil
	case *pdu.AAssociateRj:
		return fmt.Errorf("association rejected source: %s reason: %s", pt.Source, pt.Reason)
	case *pdu.AAbort:
		return fmt.Errorf("association aborted source: %s reason: %s", pt.Source, pt.Reason)
	default:
		return fmt.Errorf("unexpected message %T after sending associate", evt)
	}
}

func (c *Conn) Realease() error {
	if err := c.Send(&pdu.AReleaseRq{}, binary.BigEndian); err != nil {
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

func (c *Conn) Abort() {
	c.Send(&pdu.AAbort{
		Source: source.ServiceUser,
		Reason: abort.NotSpecified,
	}, binary.BigEndian)
}

func (c *Conn) Pdata(cmd *commands.Command, ds *dicom.Dataset) (*commands.Command, []dicom.Dataset, error) {
	ctxID, ts, err := c.ctxManager.GetAccepted(cmd.AffectedSOPClassUID...)
	if err != nil {
		return nil, nil, err
	} else if err := c.sendCmd(ctxID, cmd, ts, ds); err != nil {
		return nil, nil, err
	}

	allDataSets := []dicom.Dataset{}
	bo, implicit := transfersyntax.Info(ts)
	for {
		cmd, ds, err := c.readPData(ts, ctxID, bo, implicit)
		if err != nil {
			return nil, nil, err
		}
		allDataSets = append(allDataSets, ds...)
		switch cmd.Status {
		case status.Pending:
			continue
		case status.Successful:
			return cmd, allDataSets, nil
		default:
			if status.StatusLevel(cmd.Status) == status.Failure {
				return cmd, allDataSets, fmt.Errorf("received %s status %s from server: %s", cmd.CommandField, cmd.Status, cmd.ErrorComment)
			}
			return cmd, allDataSets, nil
		}
	}
}

func (c *Conn) readPData(ts transfersyntax.UID, ctxID uint8, bo binary.ByteOrder, implicit bool) (*commands.Command, []dicom.Dataset, error) {
	var cmd *commands.Command
	sets := []dicom.Dataset{}
	buffer := []byte{}
	for {
		evt, err := c.Read()
		if err != nil {
			return nil, nil, err
		}
		switch pd := evt.(type) {
		case *pdu.PDataTf:
			if pd.Command {
				cmd, err = commands.Decode(pd.Value, ts)
				if err != nil {
					return nil, nil, err
				} else if !cmd.HasData {
					return cmd, sets, nil
				}
				continue
			}

			switch cmd.CommandField {
			case commands.CFINDRSP, commands.CGETRSP, commands.CMOVERSP:
				payload, err := encoding.NewReader(bytes.NewBuffer(pd.Value), bo, implicit).Decode()
				if err != nil {
					return nil, nil, err
				}
				sets = append(sets, payload)
				if pd.Last {
					return cmd, sets, nil
				}
			case commands.CSTORERQ:
				buffer = append(buffer, pd.Value...)
				if pd.Last {
					d, err := dicom.Parse(bytes.NewBuffer(buffer), int64(len(buffer)), nil)
					if err != nil {
						return nil, nil, err
					} else if err := c.cstoreRsp(ctxID, cmd); err != nil {
						return nil, nil, err
					}
					sets = append(sets, d)
				}
			default:
				return nil, nil, fmt.Errorf("unhandled message type %s", cmd.CommandField)
			}
		case *pdu.AAbort:
			return nil, nil, fmt.Errorf("aborted pdata. Reason: %s Source: %s", pd.Reason, pd.Source)
		default:
			return nil, nil, fmt.Errorf("unexpected message %T after sending release", evt)
		}
	}
}

func (c *Conn) cstoreRsp(ctxID uint8, cmd *commands.Command) error {
	return c.sendCmd(ctxID, &commands.Command{
		CommandField:              commands.CSTORERSP,
		MessageIDBeingRespondedTo: cmd.MessageID,
		Status:                    status.Successful,
		AffectedSOPClassUID:       cmd.AffectedSOPClassUID,
		AffectedSOPInstanceUID:    cmd.AffectedSOPInstanceUID,
	}, transfersyntax.ImplicitVRLittleEndian, nil)
}

func (c *Conn) sendCmd(ctxID uint8, cmd *commands.Command, ts transfersyntax.UID, ds *dicom.Dataset) error {
	cmd.MessageID = int(atomic.AddInt32(&c.msgID, 1))
	cmd.HasData = ds != nil

	value, err := commands.Encode(cmd, transfersyntax.ImplicitVRLittleEndian)
	if err != nil {
		return err
	}

	// encode the command first and then send data along
	pdatas := pdu.CreatePdata(ctxID, true, int(c.cfg.ChunkSize), value)
	if ds != nil {
		buf := bytes.NewBuffer([]byte{})
		writer, err := dicom.NewWriter(buf)
		if err != nil {
			return err
		}
		writer.SetTransferSyntax(transfersyntax.Info(ts))

		for _, elem := range ds.Elements {
			if elem.Tag.Group == tag.MetadataGroup {
				continue
			}

			if err := writer.WriteElement(elem); err != nil {
				return err
			}
		}
		pdatas = append(pdatas, pdu.CreatePdata(ctxID, false, int(c.cfg.ChunkSize), buf.Bytes())...)
	}

	for _, pd := range pdatas {
		if err := c.Send(pd, binary.BigEndian); err != nil {
			return err
		}
	}
	return nil
}
