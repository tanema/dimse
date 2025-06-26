package commands

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"

	"github.com/tanema/dimse/src/defn/serviceobjectpair"
	"github.com/tanema/dimse/src/defn/status"
	"github.com/tanema/dimse/src/defn/tags"
	"github.com/tanema/dimse/src/defn/transfersyntax"
	"github.com/tanema/dimse/src/encoding"
)

//go:generate stringer -type Kind
//go:generate stringer -type DataSetType
//go:generate stringer -type Priority
type (
	// Command captures both a request and response of a PDU command
	Command struct {
		CommandField                         Kind
		AffectedSOPClassUID                  []serviceobjectpair.UID
		MessageID                            int
		MessageIDBeingRespondedTo            int
		HasData                              bool
		Status                               status.Status
		ErrorComment                         string
		Priority                             Priority
		MoveDestination                      string
		NumberOfRemainingSuboperations       int
		NumberOfCompletedSuboperations       int
		NumberOfFailedSuboperations          int
		NumberOfWarningSuboperations         int
		AffectedSOPInstanceUID               []serviceobjectpair.UID
		MoveOriginatorApplicationEntityTitle string
		MoveOriginatorMessageID              int
		CommandGroupLength                   int
	}

	Kind        int
	DataSetType int
	Priority    int
)

const (
	CSTORERQ        Kind = 0x0001
	CSTORERSP       Kind = 0x8001
	CGETRQ          Kind = 0x0010
	CGETRSP         Kind = 0x8010
	CFINDRQ         Kind = 0x0020
	CFINDRSP        Kind = 0x8020
	CMOVERQ         Kind = 0x0021
	CMOVERSP        Kind = 0x8021
	CECHORQ         Kind = 0x0030
	CECHORSP        Kind = 0x8030
	NEVENTREPORTRQ  Kind = 0x0100
	NEVENTREPORTRSP Kind = 0x8100
	NGETRQ          Kind = 0x0110
	NGETRSP         Kind = 0x8110
	NSETRQ          Kind = 0x0120
	NSETRSP         Kind = 0x8120
	NACTIONRQ       Kind = 0x0130
	NACTIONRSP      Kind = 0x8130
	NCREATERQ       Kind = 0x0140
	NCREATERSP      Kind = 0x8140
	NDELETERQ       Kind = 0x0150
	NDELETERSP      Kind = 0x8150
	CCANCELRQ       Kind = 0x0FFF

	Null    DataSetType = 0x101
	NonNull DataSetType = 1

	Low    Priority = 2
	Medium Priority = 0
	High   Priority = 1
)

func (c *Command) String() string {
	return fmt.Sprintf("[%s]:%v %s HasData: %v ErrComment: %s Priority: %v",
		c.CommandField,
		c.MessageID,
		c.Status,
		c.HasData,
		c.ErrorComment,
		c.Priority,
	)
}

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

	dst := Null
	if c.HasData {
		dst = NonNull
	}

	if err := writeElement(w, tags.CommandField, []int{int(c.CommandField)}); err != nil {
		return err
	} else if err := writeElement(w, tags.StatusTag, []int{int(c.Status)}); err != nil {
		return err
	} else if err := writeElement(w, tags.AffectedSOPClassUID, sops); err != nil {
		return err
	} else if err := writeElement(w, tags.MessageID, []int{c.MessageID}); err != nil {
		return err
	} else if err := writeElement(w, tags.MessageIDBeingRespondedTo, []int{c.MessageIDBeingRespondedTo}); err != nil {
		return err
	} else if err := writeElement(w, tags.CommandDataSetType, []int{int(dst)}); err != nil {
		return err
	}

	if c.CommandField == CSTORERQ || c.CommandField == CMOVERQ || c.CommandField == CGETRQ || c.CommandField == CFINDRQ {
		if err := writeElement(w, tags.Priority, []int{int(c.Priority)}); err != nil {
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
	d, err := encoding.NewReader(bytes.NewBuffer(data), binary.LittleEndian, true).Decode()
	if err != nil {
		return nil, err
	}
	cmd := &Command{}
	var cdst, kind, stat, priority int
	var asopcuid, asopiuid []string
	if err := readElementVal(d, tags.CommandField, &kind, true); err != nil {
		return nil, fmt.Errorf("issues reading CommandField %v", err)
	} else if err := readElementVals(d, tags.AffectedSOPClassUID, &asopcuid, true); err != nil {
		return nil, fmt.Errorf("issues reading AffectedSOPClassUID %v", err)
	} else if err := readElementVal(d, tags.CommandDataSetType, &cdst, true); err != nil {
		return nil, fmt.Errorf("issues reading CommandDataSetType %v", err)
	}

	readElementVal(d, tags.StatusTag, &stat, false)
	readElementVal(d, tags.ErrorComment, &cmd.ErrorComment, false)
	readElementVal(d, tags.Priority, &priority, false)
	readElementVals(d, tags.AffectedSOPInstanceUID, &asopiuid, false)
	readElementVals(d, tags.MoveOriginatorApplicationEntityTitle, &cmd.MoveOriginatorApplicationEntityTitle, false)
	readElementVals(d, tags.MoveOriginatorMessageID, &cmd.MoveOriginatorMessageID, false)
	readElementVal(d, tags.NumberOfRemainingSuboperations, &cmd.NumberOfRemainingSuboperations, false)
	readElementVal(d, tags.NumberOfCompletedSuboperations, &cmd.NumberOfCompletedSuboperations, false)
	readElementVal(d, tags.NumberOfFailedSuboperations, &cmd.NumberOfFailedSuboperations, false)
	readElementVal(d, tags.NumberOfWarningSuboperations, &cmd.NumberOfWarningSuboperations, false)
	readElementVal(d, tags.MessageIDBeingRespondedTo, &cmd.MessageIDBeingRespondedTo, false)
	readElementVal(d, tags.MessageID, &cmd.MessageID, false)
	readElementVal(d, tags.CommandGroupLength, &cmd.CommandGroupLength, false)

	cmd.CommandField = Kind(kind)
	cmd.Status = status.Status(stat)
	cmd.Priority = Priority(priority)
	cmd.AffectedSOPClassUID = make([]serviceobjectpair.UID, len(asopcuid))
	for i, sop := range asopcuid {
		cmd.AffectedSOPClassUID[i] = serviceobjectpair.UID(sop)
	}
	cmd.AffectedSOPInstanceUID = make([]serviceobjectpair.UID, len(asopiuid))
	for i, sop := range asopiuid {
		cmd.AffectedSOPInstanceUID[i] = serviceobjectpair.UID(sop)
	}
	if DataSetType(cdst) != Null {
		cmd.HasData = true
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
		if !required {
			return nil
		}
		return fmt.Errorf("element value is %T and not []%T", eval, dst)
	} else if len(val) == 0 {
		if !required {
			return nil
		}
		return fmt.Errorf("element value empty")
	}
	*dst = val[0]
	return nil
}

func readElementVals[T any](ds dicom.Dataset, t tag.Tag, dst *T, required bool) error {
	eval, err := findValue(ds, t, required)
	if err != nil {
		info, _ := tag.Find(t)
		return fmt.Errorf("tag: %v, %v", info.Name, err)
	}
	val, ok := eval.(T)
	if !ok {
		if !required {
			return nil
		}
		return fmt.Errorf("element value is %T and not []%T", eval, dst)
	}
	*dst = val
	return nil
}
