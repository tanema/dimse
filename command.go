package dimse

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/tanema/dimse/obj/tags"
)

type (
	Command struct {
		Dataset             dicom.Dataset
		CommandField        CommandKind
		AffectedSOPClassUID []string
		MessageID           int
		CommandDataSetType  int
		Status              int
		ErrorComment        string
		Priority            int    // CStore CMove CGet CFind
		MoveDestination     string // CMove
		// CMove CGet
		NumberOfRemainingSuboperations int
		NumberOfCompletedSuboperations int
		NumberOfFailedSuboperations    int
		NumberOfWarningSuboperations   int
		// CStore
		AffectedSOPInstanceUID               []string
		MoveOriginatorApplicationEntityTitle string
		MoveOriginatorMessageID              int
	}
	CommandKind int
)

const (
	CSTORERQ        CommandKind = 0x0001
	CSTORERSP       CommandKind = 0x8001
	CGETRQ          CommandKind = 0x0010
	CGETRSP         CommandKind = 0x8010
	CFINDRQ         CommandKind = 0x0020
	CFINDRSP        CommandKind = 0x8020
	CMOVERQ         CommandKind = 0x0021
	CMOVERSP        CommandKind = 0x8021
	CECHORQ         CommandKind = 0x0030
	CECHORSP        CommandKind = 0x8030
	NEVENTREPORTRQ  CommandKind = 0x0100
	NEVENTREPORTRSP CommandKind = 0x8100
	NGETRQ          CommandKind = 0x0110
	NGETRSP         CommandKind = 0x8110
	NSETRQ          CommandKind = 0x0120
	NSETRSP         CommandKind = 0x8120
	NACTIONRQ       CommandKind = 0x0130
	NACTIONRSP      CommandKind = 0x8130
	NCREATERQ       CommandKind = 0x0140
	NCREATERSP      CommandKind = 0x8140
	NDELETERQ       CommandKind = 0x0150
	NDELETERSP      CommandKind = 0x8150
	CCANCELRQ       CommandKind = 0x0FFF

	CommandDataSetTypeNull    int = 0x101
	CommandDataSetTypeNonNull int = 1
)

func EncodeCmd(cmd *Command) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	w, err := dicom.NewWriter(buf)
	w.SetTransferSyntax(binary.LittleEndian, true)
	if err != nil {
		return nil, err
	} else if err := cmd.encode(w); err != nil {
		return nil, err
	}
	return buf.Bytes(), writeElement(w, tags.CommandGroupLength, []int{buf.Len()})
}

func (c *Command) encode(w *dicom.Writer) error {
	if err := writeElement(w, tags.CommandField, []int{int(c.CommandField)}); err != nil {
		return err
	} else if err := writeElement(w, tags.AffectedSOPClassUID, c.AffectedSOPClassUID); err != nil {
		return err
	} else if err := writeElement(w, tags.MessageID, []int{c.MessageID}); err != nil {
		return err
	} else if err := writeElement(w, tags.CommandDataSetType, []int{c.CommandDataSetType}); err != nil {
		return err
	}

	if c.CommandField == CSTORERQ || c.CommandField == CMOVERQ || c.CommandField == CGETRQ || c.CommandField == CFINDRQ {
		if err := writeElement(w, tags.Priority, []int{c.Priority}); err != nil {
			return err
		}
	}

	if c.CommandField == CMOVERQ {
		if err := writeElement(w, tags.MoveDestination, []string{c.MoveDestination}); err != nil {
			return err
		}
	}

	if c.CommandField == CSTORERQ {
		if err := writeElement(w, tags.AffectedSOPInstanceUID, c.AffectedSOPInstanceUID); err != nil {
			return err
		}
		if c.MoveOriginatorApplicationEntityTitle != "" {
			if err := writeElement(w, tags.MoveOriginatorApplicationEntityTitle, []string{c.MoveOriginatorApplicationEntityTitle}); err != nil {
				return err
			}
		}
		if c.MoveOriginatorMessageID != 0 {
			if err := writeElement(w, tags.MoveOriginatorMessageID, []int{c.MoveOriginatorMessageID}); err != nil {
				return err
			}
		}
	}

	return nil
}

func Decode(data []byte) (*Command, error) {
	d, err := dicom.ParseUntilEOF(bytes.NewBuffer(data), nil)
	if err != nil {
		return nil, err
	}
	cmd := &Command{Dataset: d}
	if err := readElement(d, tags.AffectedSOPClassUID, cmd.AffectedSOPClassUID, true); err != nil {
		return nil, err
	} else if err := readElement(d, tags.MessageIDBeingRespondedTo, cmd.MessageID, true); err != nil {
		return nil, err
	} else if err := readElement(d, tags.CommandDataSetType, cmd.CommandDataSetType, true); err != nil {
		return nil, err
	} else if err := readElement(d, tags.StatusTag, &cmd.Status, true); err != nil {
		return nil, err
	} else if err := readElement(d, tags.ErrorComment, &cmd.ErrorComment, false); err != nil {
		return nil, err
	}

	if cmd.CommandField == CSTORERQ || cmd.CommandField == CMOVERQ || cmd.CommandField == CGETRQ || cmd.CommandField == CFINDRQ {
		if err := readElement(d, tags.Priority, cmd.Priority, true); err != nil {
			return nil, err
		}
	}

	// optional anyways
	if err := readElement(d, tags.NumberOfRemainingSuboperations, cmd.NumberOfRemainingSuboperations, false); err != nil {
		return nil, err
	} else if err := readElement(d, tags.NumberOfCompletedSuboperations, cmd.NumberOfCompletedSuboperations, false); err != nil {
		return nil, err
	} else if err := readElement(d, tags.NumberOfFailedSuboperations, cmd.NumberOfFailedSuboperations, false); err != nil {
		return nil, err
	} else if err := readElement(d, tags.NumberOfWarningSuboperations, cmd.NumberOfWarningSuboperations, false); err != nil {
		return nil, err
	}

	if cmd.CommandField == CSTORERQ {
		if err := readElement(d, tags.AffectedSOPInstanceUID, cmd.AffectedSOPInstanceUID, true); err != nil {
			return nil, err
		}
	}

	return cmd, nil
}

func Find(msgID int, sops []string, dsType int) *Command {
	return &Command{
		CommandField:        CFINDRQ,
		MessageID:           msgID,
		AffectedSOPClassUID: sops,
		CommandDataSetType:  dsType,
	}
}

func Get(msgID int, sops []string, dsType, priority int) *Command {
	return &Command{
		CommandField:        CGETRQ,
		MessageID:           msgID,
		AffectedSOPClassUID: sops,
		CommandDataSetType:  dsType,
		Priority:            priority,
	}
}

func Move(msgID int, sops []string, dsType, priority int, dst string) *Command {
	return &Command{
		CommandField:        CMOVERQ,
		MessageID:           msgID,
		AffectedSOPClassUID: sops,
		CommandDataSetType:  dsType,
		Priority:            priority,
		MoveDestination:     dst,
	}
}

func Store(msgID int, sops, inst []string, dsType, priority, id int, dst, title string) *Command {
	return &Command{
		CommandField:                         CSTORERQ,
		MessageID:                            msgID,
		AffectedSOPClassUID:                  sops,
		CommandDataSetType:                   dsType,
		Priority:                             priority,
		MoveDestination:                      dst,
		AffectedSOPInstanceUID:               inst,
		MoveOriginatorApplicationEntityTitle: title,
		MoveOriginatorMessageID:              id,
	}
}

func writeElement(w *dicom.Writer, t tag.Tag, val any) error {
	elem, err := dicom.NewElement(t, val)
	if err != nil {
		return err
	}
	return w.WriteElement(elem)
}

func readElement[T any](ds dicom.Dataset, t tag.Tag, dst T, required bool) error {
	elem, err := ds.FindElementByTagNested(t)
	if err != nil {
		if err == dicom.ErrorElementNotFound && !required {
			return nil
		}
		return err
	}
	val, ok := elem.Value.GetValue().([]T)
	if !ok {
		return fmt.Errorf("element value is %T and not []%T", elem.Value.GetValue(), dst)
	} else if len(val) == 0 {
		return fmt.Errorf("element value empty")
	}
	dst = val[0]
	return nil
}
