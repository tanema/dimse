package chunkreader

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/suyashkumar/dicom/pkg/vrraw"
)

type Reader struct {
	buf *bytes.Buffer
	ds  dicom.Dataset
	bo  binary.ByteOrder
}

func New() *Reader {
	return &Reader{
		ds:  dicom.Dataset{Elements: []*dicom.Element{}},
		buf: bytes.NewBuffer(nil),
		bo:  binary.LittleEndian,
	}
}

func (r *Reader) Dataset() dicom.Dataset {
	return r.ds
}

func (r *Reader) Decode(data []byte, implicit bool) error {
	if _, err := r.buf.Write(data); err != nil {
		return err
	}
	for {
		elem, err := r.readElement(implicit)
		if elem != nil {
			r.ds.Elements = append(r.ds.Elements, elem)
		}
		if errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func (r *Reader) read(val any) error {
	return binary.Read(r.buf, r.bo, val)
}

func (r *Reader) skip(l int) error {
	junk := make([]byte, l)
	return binary.Read(r.buf, r.bo, &junk)
}

func (r *Reader) readElement(implicit bool) (*dicom.Element, error) {
	elemTag, err := r.readTag()
	if err != nil {
		return nil, err
	}

	vr, err := r.readVR(implicit, *elemTag)
	if err != nil {
		return nil, err
	}

	vlength, err := r.readVL(implicit, vr)
	if err != nil {
		return nil, fmt.Errorf("readElement: error when reading VL for element %v: %w", elemTag, err)
	}

	val, err := r.readValue(*elemTag, vr, vlength)
	if err != nil {
		return nil, fmt.Errorf("readElement: error when reading Value for element %v: %w", elemTag, err)
	}

	return &dicom.Element{
		Tag:                    *elemTag,
		ValueRepresentation:    tag.GetVRKind(*elemTag, vr),
		RawValueRepresentation: vr,
		ValueLength:            vlength,
		Value:                  val,
	}, nil
}

func (r *Reader) readTag() (*tag.Tag, error) {
	var group, element uint16
	if err := r.read(&group); err != nil {
		return nil, err
	} else if err := r.read(&element); err != nil {
		return nil, err
	}
	return &tag.Tag{Group: group, Element: element}, nil
}

func (r *Reader) readVR(implicit bool, t tag.Tag) (string, error) {
	if implicit {
		if entry, err := tag.Find(t); err == nil {
			switch entry.Tag {
			case tag.PixelData, tag.OverlayData:
				return "OW", nil
			default:
				return entry.VRs[0], nil
			}
		}
		return tag.UnknownVR, nil
	}
	// Explicit Transfer Syntax, read 2 byte VR:
	return r.readRawString(2)
}

func (r *Reader) readVL(implicit bool, vr string) (uint32, error) {
	if implicit {
		var vl uint32
		return vl, r.read(&vl)
	}

	// Explicit Transfer Syntax
	// More details here: https://dicom.nema.org/medical/dicom/current/output/html/part05.html#sect_7.1.2
	switch vr {
	case "NA", vrraw.OtherByte, vrraw.OtherDouble, vrraw.OtherFloat,
		vrraw.OtherLong, vrraw.OtherWord, vrraw.Sequence, vrraw.Unknown,
		vrraw.UnlimitedCharacters, vrraw.UniversalResourceIdentifier,
		vrraw.UnlimitedText:
		r.skip(2) // ignore two reserved bytes (0000H)
		var vl uint32
		err := r.read(&vl)
		if err != nil {
			return 0, err
		}

		if vl == tag.VLUndefinedLength &&
			(vr == vrraw.UnlimitedCharacters ||
				vr == vrraw.UniversalResourceIdentifier ||
				vr == vrraw.UnlimitedText) {
			return 0, errors.New("UC, UR and UT may not have an Undefined Length, i.e.,a Value Length of FFFFFFFFH")
		}
		return vl, nil
	default:
		var vl16 uint16
		err := r.read(&vl16)
		if err != nil {
			return 0, err
		}
		return uint32(vl16), nil
	}
}

func (r *Reader) readValue(t tag.Tag, vr string, vl uint32) (dicom.Value, error) {
	vrkind := tag.GetVRKind(t, vr)
	switch vrkind {
	case tag.VRBytes:
		return nil, fmt.Errorf("cannot read bytes value in command")
		// return r.readBytes(t, vr, vl)
	case tag.VRUInt16List, tag.VRUInt32List, tag.VRInt16List, tag.VRInt32List, tag.VRTagList:
		return r.readInt(t, vr, vl)
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
		return r.readString(t, vl)
	}
}

func (r *Reader) readRawString(l uint32) (string, error) {
	data := make([]byte, l)
	_, err := io.ReadFull(r.buf, data)
	return string(data), err
}

func (r *Reader) readString(t tag.Tag, vl uint32) (dicom.Value, error) {
	str, err := r.readRawString(vl)
	if err != nil {
		return nil, fmt.Errorf("error reading string element value: %w", err)
	}

	info, _ := tag.Find(t)
	if err != nil {
		return nil, fmt.Errorf("error reading string element %v value: %w", info.Name, err)
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

func (r *Reader) readInt(t tag.Tag, vr string, vl uint32) (dicom.Value, error) {
	dataLen := vl
	retVal := []int{}
	for i := 0; dataLen > 0; i++ {
		var err error
		switch vr {
		case vrraw.UnsignedLong:
			var val uint32
			err = r.read(&val)
			dataLen -= 4
			retVal = append(retVal, int(val))
		case vrraw.SignedLong:
			var val int32
			err = r.read(&val)
			dataLen -= 4
			retVal = append(retVal, int(val))
		case vrraw.UnsignedShort, vrraw.AttributeTag:
			var val uint16
			err = r.read(&val)
			dataLen -= 2
			retVal = append(retVal, int(val))
		case vrraw.SignedShort:
			var val int16
			err = r.read(&val)
			dataLen -= 2
			retVal = append(retVal, int(val))
		default:
			return nil, fmt.Errorf("unable to parse integer type due to unknown VR %v", vr)
		}
		if err != nil {
			return nil, fmt.Errorf("error reading int element (%v) value: %w", t, err)
		}
	}
	return dicom.NewValue(retVal)
}
