package dimse

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"

	"github.com/tanema/dimse/src/chunkreader"
	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/query"
	"github.com/tanema/dimse/src/serviceobjectpair"
	"github.com/tanema/dimse/src/tags"
	"github.com/tanema/dimse/src/transfersyntax"
)

type (
	// Query is a captured, validated query scope for find, get, move, and store
	Query struct {
		client   *Client
		payload  []byte
		Level    query.Level
		Filter   []*dicom.Element
		Priority int // CStore CMove CGet CFind
	}
	// Command captures both a request and response of a PDU command
	Command struct {
		Dataset             dicom.Dataset
		CommandField        commands.Kind
		AffectedSOPClassUID []serviceobjectpair.UID
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
		AffectedSOPInstanceUID               []serviceobjectpair.UID
		MoveOriginatorApplicationEntityTitle string
		MoveOriginatorMessageID              int
	}
)

func EncodeCmd(cmd *Command, ts transfersyntax.UID) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	w, err := dicom.NewWriter(buf)
	if err != nil {
		return nil, err
	}
	w.SetTransferSyntax(transfersyntax.Info(ts))
	if err := cmd.encode(w); err != nil {
		return nil, err
	}
	return buf.Bytes(), writeElement(w, tags.CommandGroupLength, []int{buf.Len()})
}

func (c *Command) encode(w *dicom.Writer) error {
	sops := []string{}
	for _, s := range c.AffectedSOPClassUID {
		sops = append(sops, string(s))
	}
	if err := writeElement(w, tags.CommandField, []int{int(c.CommandField)}); err != nil {
		return err
	} else if err := writeElement(w, tags.AffectedSOPClassUID, sops); err != nil {
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
		sops := []string{}
		for _, s := range c.AffectedSOPInstanceUID {
			sops = append(sops, string(s))
		}
		if err := writeElement(w, tags.AffectedSOPInstanceUID, sops); err != nil {
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

func DecodeCmd(data []byte, ts transfersyntax.UID) (*Command, error) {
	d, err := decode(data, true)
	if err != nil {
		return nil, err
	}
	cmd := &Command{Dataset: d}
	var cdst, kind int
	var asopcuid []string
	if err := readElementVal(d, tags.CommandField, &kind, true); err != nil {
		return nil, fmt.Errorf("issues reading CommandField %v", err)
	} else if err := readElementVals(d, tags.AffectedSOPClassUID, &asopcuid, true); err != nil {
		return nil, fmt.Errorf("issues reading AffectedSOPClassUID %v", err)
	} else if err := readElementVal(d, tags.MessageIDBeingRespondedTo, &cmd.MessageID, true); err != nil {
		return nil, fmt.Errorf("issues reading MessageIDBeingRespondedTo %v", err)
	} else if err := readElementVal(d, tags.CommandDataSetType, &cdst, true); err != nil {
		return nil, fmt.Errorf("issues reading CommandDataSetType %v", err)
	} else if err := readElementVal(d, tags.StatusTag, &cmd.Status, true); err != nil {
		return nil, fmt.Errorf("issues reading StatusTag %v", err)
	} else if err := readElementVal(d, tags.ErrorComment, &cmd.ErrorComment, false); err != nil {
		return nil, fmt.Errorf("issues reading ErrorComment %v", err)
	}

	cmd.CommandField = commands.Kind(kind)
	cmd.AffectedSOPClassUID = make([]serviceobjectpair.UID, len(asopcuid))
	for i, sop := range asopcuid {
		cmd.AffectedSOPClassUID[i] = serviceobjectpair.UID(sop)
	}
	cmd.CommandDataSetType = commands.DataSetType(cdst)

	if cmd.CommandField == commands.CSTORERQ || cmd.CommandField == commands.CMOVERQ || cmd.CommandField == commands.CGETRQ || cmd.CommandField == commands.CFINDRQ {
		if err := readElementVal(d, tags.Priority, &cmd.Priority, true); err != nil {
			return nil, err
		}
	}

	// optional anyways
	if err := readElementVal(d, tags.NumberOfRemainingSuboperations, &cmd.NumberOfRemainingSuboperations, false); err != nil {
		return nil, err
	} else if err := readElementVal(d, tags.NumberOfCompletedSuboperations, &cmd.NumberOfCompletedSuboperations, false); err != nil {
		return nil, err
	} else if err := readElementVal(d, tags.NumberOfFailedSuboperations, &cmd.NumberOfFailedSuboperations, false); err != nil {
		return nil, err
	} else if err := readElementVal(d, tags.NumberOfWarningSuboperations, &cmd.NumberOfWarningSuboperations, false); err != nil {
		return nil, err
	}

	if cmd.CommandField == commands.CSTORERQ {
		if err := readElementVals(d, tags.AffectedSOPInstanceUID, &cmd.AffectedSOPInstanceUID, true); err != nil {
			return nil, err
		}
	}

	return cmd, nil
}

func decode(data []byte, implicit bool) (dicom.Dataset, error) {
	r := chunkreader.New()
	if err := r.Decode(data, implicit); err != nil {
		return dicom.Dataset{}, err
	}
	return r.Dataset(), nil
}

func writeElement(w *dicom.Writer, t tag.Tag, val any) error {
	elem, err := dicom.NewElement(t, val)
	if err != nil {
		return err
	}
	return w.WriteElement(elem)
}

func findValue(ds dicom.Dataset, t tag.Tag, required bool) (any, error) {
	elem, err := ds.FindElementByTagNested(t)
	if err != nil {
		if errors.Is(err, dicom.ErrorElementNotFound) && !required {
			return nil, nil
		}
		return nil, err
	}
	return elem.Value.GetValue(), nil
}

func readElementVal[T any](ds dicom.Dataset, t tag.Tag, dst *T, required bool) error {
	eval, err := findValue(ds, t, required)
	if err != nil || eval == nil {
		return err
	}
	val, ok := eval.([]T)
	if !ok {
		return fmt.Errorf("element value is %T and not []%T", eval, dst)
	} else if len(val) == 0 {
		return fmt.Errorf("element value empty")
	}
	*dst = val[0]
	return nil
}

func readElementVals[T any](ds dicom.Dataset, t tag.Tag, dst *T, required bool) error {
	eval, err := findValue(ds, t, required)
	if err != nil {
		return err
	}
	val, ok := eval.(T)
	if !ok {
		return fmt.Errorf("element value is %T and not []%T", eval, dst)
	}
	*dst = val
	return nil
}
