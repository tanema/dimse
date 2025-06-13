package dimse

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/tanema/dimse/obj/commands"
	"github.com/tanema/dimse/obj/tags"
)

type (
	Command struct {
		Dataset             dicom.Dataset
		CommandField        commands.Kind
		AffectedSOPClassUID []string
		MessageID           int
		CommandDataSetType  commands.DataSetType
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
	} else if err := writeElement(w, tags.CommandDataSetType, []int{int(c.CommandDataSetType)}); err != nil {
		return err
	}

	if c.CommandField == commands.CSTORERQ || c.CommandField == commands.CMOVERQ || c.CommandField == commands.CGETRQ || c.CommandField == commands.CFINDRQ {
		if err := writeElement(w, tags.Priority, []int{c.Priority}); err != nil {
			return err
		}
	}

	if c.CommandField == commands.CMOVERQ {
		if err := writeElement(w, tags.MoveDestination, []string{c.MoveDestination}); err != nil {
			return err
		}
	}

	if c.CommandField == commands.CSTORERQ {
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

func DecodeCmd(data []byte) (*Command, error) {
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

	if cmd.CommandField == commands.CSTORERQ || cmd.CommandField == commands.CMOVERQ || cmd.CommandField == commands.CGETRQ || cmd.CommandField == commands.CFINDRQ {
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

	if cmd.CommandField == commands.CSTORERQ {
		if err := readElement(d, tags.AffectedSOPInstanceUID, cmd.AffectedSOPInstanceUID, true); err != nil {
			return nil, err
		}
	}

	return cmd, nil
}

func Find(msgID int, sops []string, dsType commands.DataSetType) *Command {
	return &Command{
		CommandField:        commands.CFINDRQ,
		MessageID:           msgID,
		AffectedSOPClassUID: sops,
		CommandDataSetType:  dsType,
	}
}

func Get(msgID int, sops []string, dsType commands.DataSetType, priority int) *Command {
	return &Command{
		CommandField:        commands.CGETRQ,
		MessageID:           msgID,
		AffectedSOPClassUID: sops,
		CommandDataSetType:  dsType,
		Priority:            priority,
	}
}

func Move(msgID int, sops []string, dsType commands.DataSetType, priority int, dst string) *Command {
	return &Command{
		CommandField:        commands.CMOVERQ,
		MessageID:           msgID,
		AffectedSOPClassUID: sops,
		CommandDataSetType:  dsType,
		Priority:            priority,
		MoveDestination:     dst,
	}
}

func Store(msgID int, sops, inst []string, dsType commands.DataSetType, priority, id int, dst, title string) *Command {
	return &Command{
		CommandField:                         commands.CSTORERQ,
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
