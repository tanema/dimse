package commands

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"

	"github.com/tanema/dimse/src/chunkreader"
	"github.com/tanema/dimse/src/serviceobjectpair"
	"github.com/tanema/dimse/src/tags"
	"github.com/tanema/dimse/src/transfersyntax"
)

type (
	// Command captures both a request and response of a PDU command
	Command struct {
		Dataset             dicom.Dataset
		CommandField        Kind
		AffectedSOPClassUID []serviceobjectpair.UID
		MessageID           int
		CommandDataSetType  DataSetType
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

func Encode(cmd *Command, ts transfersyntax.UID) ([]byte, error) {
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

func Decode(data []byte, ts transfersyntax.UID) (*Command, error) {
	r := chunkreader.New()
	if err := r.Decode(data, true); err != nil {
		return nil, err
	}
	d := r.Dataset()

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

	cmd.CommandField = Kind(kind)
	cmd.AffectedSOPClassUID = make([]serviceobjectpair.UID, len(asopcuid))
	for i, sop := range asopcuid {
		cmd.AffectedSOPClassUID[i] = serviceobjectpair.UID(sop)
	}
	cmd.CommandDataSetType = DataSetType(cdst)

	if cmd.CommandField == CSTORERQ || cmd.CommandField == CMOVERQ || cmd.CommandField == CGETRQ || cmd.CommandField == CFINDRQ {
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

	if cmd.CommandField == CSTORERQ {
		if err := readElementVals(d, tags.AffectedSOPInstanceUID, &cmd.AffectedSOPInstanceUID, true); err != nil {
			return nil, err
		}
	}

	return cmd, nil
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
