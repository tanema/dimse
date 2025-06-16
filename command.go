package dimse

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/dicomio"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/suyashkumar/dicom/pkg/vrraw"
	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/tags"
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
	d, err := decodeCommand(data)
	if err != nil {
		return nil, err
	}
	cmd := &Command{Dataset: d}
	var cdst, kind int
	if err := readElementVal(d, tags.CommandField, &kind, true); err != nil {
		return nil, fmt.Errorf("issues reading CommandField %v", err)
	} else if err := readElementVals(d, tags.AffectedSOPClassUID, &cmd.AffectedSOPClassUID, true); err != nil {
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

func decodeCommand(data []byte) (dicom.Dataset, error) {
	d := dicom.Dataset{Elements: []*dicom.Element{}}
	r := dicomio.NewReader(bufio.NewReader(bytes.NewBuffer(data)), binary.LittleEndian, int64(len(data)))
	r.SetTransferSyntax(binary.LittleEndian, true)
	for {
		group, gerr := r.ReadUInt16()
		if errors.Is(gerr, io.EOF) {
			break
		}
		element, eerr := r.ReadUInt16()
		if err := errors.Join(gerr, eerr); err != nil {
			return d, fmt.Errorf("error reading tag: %w", errors.Join(gerr, eerr))
		}
		elemTag := tag.Tag{Group: group, Element: element}

		valueRepresentation := tag.UnknownVR
		if entry, err := tag.Find(elemTag); err == nil {
			valueRepresentation = entry.VRs[0]
		}

		vlength, err := r.ReadUInt32()
		if err != nil {
			return d, fmt.Errorf("readElement: error when reading VL for element %v: %w", elemTag, err)
		}

		val, err := readValue(r, elemTag, valueRepresentation, vlength)
		if err != nil {
			return d, fmt.Errorf("readElement: error when reading Value for element %v: %w", elemTag, err)
		}

		d.Elements = append(d.Elements, &dicom.Element{
			Tag:                    elemTag,
			ValueRepresentation:    tag.GetVRKind(elemTag, valueRepresentation),
			RawValueRepresentation: valueRepresentation,
			ValueLength:            vlength,
			Value:                  val,
		})
	}
	return d, nil
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

func readValue(r *dicomio.Reader, t tag.Tag, vr string, vl uint32) (dicom.Value, error) {
	vrkind := tag.GetVRKind(t, vr)
	switch vrkind {
	case tag.VRBytes:
		return nil, fmt.Errorf("cannot read bytes value in command")
		// return r.readBytes(t, vr, vl)
	case tag.VRUInt16List, tag.VRUInt32List, tag.VRInt16List, tag.VRInt32List, tag.VRTagList:
		return readInt(r, t, vr, vl)
	case tag.VRFloat32List, tag.VRFloat64List:
		// return r.readFloat(t, vr, vl)
		return nil, fmt.Errorf("cannot read float value in command")
	case tag.VRSequence, tag.VRItem:
		return nil, fmt.Errorf("cannot read sequence value in command")
	case tag.VRUnknown, tag.VRPixelData:
		// if vl == tag.VLUndefinedLength { return r.readSequence(t, vr, vl) }
		// return r.readBytes(t, vr, vl)
		return nil, fmt.Errorf("cannot read pixel value in command")
	default:
		return readString(r, t, vl)
	}
}

func readString(r *dicomio.Reader, t tag.Tag, vl uint32) (dicom.Value, error) {
	str, err := r.ReadString(vl)
	if err != nil {
		return nil, fmt.Errorf("error reading string element (%v) value: %w", t, err)
	}
	onlySpaces := true
	for _, char := range str {
		if !unicode.IsSpace(char) {
			onlySpaces = false
			break
		}
	}
	if !onlySpaces {
		str = strings.Trim(str, " \000")
	}
	return dicom.NewValue(strings.Split(str, "\\"))
}

func readInt(r *dicomio.Reader, t tag.Tag, vr string, vl uint32) (dicom.Value, error) {
	err := r.PushLimit(int64(vl))
	if err != nil {
		return nil, err
	}
	retVal := make([]int, 0, vl/2)
	for !r.IsLimitExhausted() {
		switch vr {
		case vrraw.UnsignedShort, vrraw.AttributeTag:
			val, err := r.ReadUInt16()
			if err != nil {
				return nil, fmt.Errorf("error reading int element (%v) value (ReadUInt16): %w", t, err)
			}
			retVal = append(retVal, int(val))
		case vrraw.UnsignedLong:
			val, err := r.ReadUInt32()
			if err != nil {
				return nil, fmt.Errorf("error reading int element (%v) value (ReadUInt32): %w", t, err)
			}
			retVal = append(retVal, int(val))
		case vrraw.SignedLong:
			val, err := r.ReadInt32()
			if err != nil {
				return nil, fmt.Errorf("error reading int element (%v) value (ReadInt32): %w", t, err)
			}
			retVal = append(retVal, int(val))
		case vrraw.SignedShort:
			val, err := r.ReadInt16()
			if err != nil {
				return nil, fmt.Errorf("error reading int element (%v) value (ReadInt16): %w", t, err)
			}
			retVal = append(retVal, int(val))
		default:
			return nil, fmt.Errorf("unable to parse integer type due to unknown VR %v", vr)
		}
	}
	r.PopLimit()
	return dicom.NewValue(retVal)
}

// func (r *reader) readBytes(t tag.Tag, vr string, vl uint32) (Value, error) {
//	if vr == vrraw.OtherByte || vr == vrraw.Unknown {
//		data := make([]byte, vl)
//		_, err := io.ReadFull(r.rawReader, data)
//		return &bytesValue{value: data}, err
//	} else if vr == vrraw.OtherWord {
//		if vl%2 != 0 {
//			return nil, fmt.Errorf("error reading bytes element (%v) value: %w", t, ErrorOWRequiresEvenVL)
//		}

//		buf := bytes.NewBuffer(make([]byte, 0, vl))
//		numWords := int(vl / 2)
//		for i := 0; i < numWords; i++ {
//			word, err := r.rawReader.ReadUInt16()
//			if err != nil {
//				return nil, fmt.Errorf("error reading bytes element (%v) value: %w", t, err)
//			}
//			err = binary.Write(buf, binary.LittleEndian, word)
//			if err != nil {
//				return nil, err
//			}
//		}
//		return &bytesValue{value: buf.Bytes()}, nil
//	}
//	return nil, fmt.Errorf("error reading bytes element (%v): %w", t, ErrorUnsupportedVR)
// }

// func (r *reader) readFloat(t tag.Tag, vr string, vl uint32) (Value, error) {
//	err := r.rawReader.PushLimit(int64(vl))
//	if err != nil {
//		return nil, err
//	}
//	retVal := &floatsValue{value: make([]float64, 0, vl/2)}
//	for !r.rawReader.IsLimitExhausted() {
//		switch vr {
//		case vrraw.FloatingPointSingle:
//			val, err := r.rawReader.ReadFloat32()
//			if err != nil {
//				return nil, fmt.Errorf("error reading floating point element (%v) value: %w", t, err)
//			}
//			pval, err := strconv.ParseFloat(fmt.Sprint(val), 64)
//			if err != nil {
//				return nil, fmt.Errorf("error reading floating point element (%v) value during strconv.ParseFloat: %w", t, err)
//			}
//			retVal.value = append(retVal.value, pval)
//			break
//		case vrraw.FloatingPointDouble:
//			val, err := r.rawReader.ReadFloat64()
//			if err != nil {
//				return nil, fmt.Errorf("error reading floating point element (%v) value: %w", t, err)
//			}
//			retVal.value = append(retVal.value, val)
//			break
//		default:
//			return nil, fmt.Errorf("error reading floating point element(%v) value: unsupported VR: %w", t, errorUnableToParseFloat)
//		}
//	}
//	r.rawReader.PopLimit()
//	return retVal, nil
// }

// func (r *reader) readSequence(t tag.Tag, vr string, vl uint32) (Value, error) {
//	var sequences sequencesValue

//	seqElements := &dicom.Dataset{}
//	if vl == tag.VLUndefinedLength {
//		for {
//			subElement, err := r.readElement(seqElements, nil)
//			if err != nil {
//				return nil, fmt.Errorf("readSequence: error reading subitem in a sequence: %w", err)
//			}
//			if subElement.Tag == tag.SequenceDelimitationItem {
//				break
//			}
//			if subElement.Tag != tag.Item || subElement.Value.ValueType() != dicom.SequenceItem {
//				return nil, fmt.Errorf("readSequence: error, non item element found in sequence. got: %v", subElement)
//			}
//			sequences.value = append(sequences.value, subElement.Value.(*SequenceItemValue))
//		}
//	} else {
//		err := r.rawReader.PushLimit(int64(vl))
//		if err != nil {
//			return nil, err
//		}
//		for !r.rawReader.IsLimitExhausted() {
//			subElement, err := r.readElement(seqElements, nil)
//			if err != nil {
//				return nil, fmt.Errorf("readSequence: error reading subitem in a sequence: %w", err)
//			}
//			sequences.value = append(sequences.value, subElement.Value.(*dicom.SequenceItemValue))
//		}
//		r.rawReader.PopLimit()
//	}

//	return &sequences, nil
// }

// func (r *reader) readSequenceItem(t tag.Tag, vr string, vl uint32) (Value, error) {
//	var sequenceItem dicom.SequenceItemValue
//	seqElements := Dataset{}
//	if vl == tag.VLUndefinedLength {
//		for {
//			subElem, err := r.readElement(&seqElements, nil)
//			if err != nil {
//				return nil, fmt.Errorf("readSequenceItem: error reading subitem in a sequence item: %w", err)
//			}
//			if subElem.Tag == tag.ItemDelimitationItem {
//				break
//			}
//			sequenceItem.elements = append(sequenceItem.elements, subElem)
//			seqElements.Elements = append(seqElements.Elements, subElem)
//		}
//	} else {
//		if err := r.rawReader.PushLimit(int64(vl)); err != nil {
//			return nil, err
//		}
//		for !r.rawReader.IsLimitExhausted() {
//			subElem, err := r.readElement(&seqElements, nil)
//			if err != nil {
//				return nil, fmt.Errorf("readSequenceItem: error reading subitem in a sequence item: %w", err)
//			}
//			sequenceItem.elements = append(sequenceItem.elements, subElem)
//			seqElements.Elements = append(seqElements.Elements, subElem)
//		}
//		r.rawReader.PopLimit()
//	}
//	return &sequenceItem, nil
// }
