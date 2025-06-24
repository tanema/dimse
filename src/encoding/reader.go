package encoding

import (
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

type (
	Reader struct {
		bo       binary.ByteOrder
		implicit bool
		in       []io.Reader
	}
	Skip int64
)

func NewReader(in io.Reader, bo binary.ByteOrder, implicit bool) *Reader {
	return &Reader{
		bo:       bo,
		implicit: implicit,
		in:       []io.Reader{in},
	}
}

func (r *Reader) Decode() (dicom.Dataset, error) {
	ds := dicom.Dataset{}
	for {
		elem, err := r.readElement()
		if elem != nil {
			ds.Elements = append(ds.Elements, elem)
		}
		if errors.Is(err, io.EOF) {
			return ds, nil
		} else if err != nil {
			return ds, err
		}
	}
}

func (r *Reader) reader() io.Reader {
	return r.in[len(r.in)-1]
}

func (r *Reader) PushLimit(l int) {
	r.in = append(r.in, &io.LimitedReader{R: r.reader(), N: int64(l)})
}

func (r *Reader) PopLimit() {
	if len(r.in) <= 1 {
		return
	}
	x := len(r.in) - 1
	r.in = r.in[:x:x]
}

func (r *Reader) Read(parts ...any) error {
	for _, data := range parts {
		if skip, isSkip := data.(Skip); isSkip {
			zeros := make([]byte, int(skip))
			if err := binary.Read(r.reader(), r.bo, zeros); err != nil {
				return err
			}
		} else if err := binary.Read(r.reader(), r.bo, data); err != nil {
			return err
		}
	}
	return nil
}

func (r *Reader) String(length int) (string, error) {
	data := make([]byte, length)
	return string(data), r.Read(&data)
}

func (r *Reader) readElement() (*dicom.Element, error) {
	if t, err := r.readTag(); err != nil {
		return nil, err
	} else if vr, err := r.readRepresentation(t); err != nil {
		return nil, err
	} else if vlength, err := r.readLength(vr); err != nil {
		return nil, fmt.Errorf("readElement: error when reading VL for element %v: %w", t, err)
	} else if val, err := r.readValue(t, vr, vlength); err != nil {
		return nil, fmt.Errorf("readElement: error when reading Value for element %v: %w", t, err)
	} else {
		return &dicom.Element{
			Tag:                    t,
			ValueRepresentation:    tag.GetVRKind(t, vr),
			RawValueRepresentation: vr,
			ValueLength:            vlength,
			Value:                  val,
		}, nil
	}
}

func (r *Reader) readTag() (tag.Tag, error) {
	var group, element uint16
	if err := r.Read(&group, &element); err != nil {
		return tag.Tag{}, err
	}
	return tag.Tag{Group: group, Element: element}, nil
}

func (r *Reader) readRepresentation(t tag.Tag) (string, error) {
	if r.implicit {
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
	return r.readRawString(2)
}

func (r *Reader) readLength(vr string) (uint32, error) {
	if r.implicit {
		var vl uint32
		return vl, r.Read(&vl)
	}
	switch vr {
	case "NA", vrraw.OtherByte, vrraw.OtherDouble, vrraw.OtherFloat,
		vrraw.OtherLong, vrraw.OtherWord, vrraw.Sequence, vrraw.Unknown,
		vrraw.UnlimitedCharacters, vrraw.UniversalResourceIdentifier,
		vrraw.UnlimitedText:
		var vl uint32
		if err := r.Read(Skip(2), &vl); err != nil {
			return 0, err
		} else if vl == tag.VLUndefinedLength && (vr == vrraw.UnlimitedCharacters || vr == vrraw.UniversalResourceIdentifier || vr == vrraw.UnlimitedText) {
			return 0, errors.New("UC, UR and UT may not have an Undefined Length, i.e.,a Value Length of FFFFFFFFH")
		}
		return vl, nil
	default:
		var vl16 uint16
		if err := r.Read(&vl16); err != nil {
			return 0, err
		}
		return uint32(vl16), nil
	}
}

// limited value read support because the protocol data contains very few data types.
func (r *Reader) readValue(t tag.Tag, vr string, vl uint32) (dicom.Value, error) {
	vrkind := tag.GetVRKind(t, vr)
	switch vrkind {
	case tag.VRBytes, tag.VRFloat32List, tag.VRFloat64List, tag.VRSequence, tag.VRItem,
		tag.VRUnknown, tag.VRPixelData:
		return nil, fmt.Errorf("read %v value unsupported right now", vr)
	case tag.VRUInt16List, tag.VRUInt32List, tag.VRInt16List, tag.VRInt32List, tag.VRTagList:
		return r.readInt(t, vr, vl)
	default:
		return r.readString(vl)
	}
}

func (r *Reader) readRawString(l uint32) (string, error) {
	data := make([]byte, l)
	_, err := io.ReadFull(r.reader(), data)
	return string(data), err
}

func (r *Reader) readString(vl uint32) (dicom.Value, error) {
	str, err := r.readRawString(vl)
	if err != nil {
		return nil, fmt.Errorf("error reading string element value: %w", err)
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
			err = r.Read(&val)
			dataLen -= 4
			retVal = append(retVal, int(val))
		case vrraw.SignedLong:
			var val int32
			err = r.Read(&val)
			dataLen -= 4
			retVal = append(retVal, int(val))
		case vrraw.UnsignedShort, vrraw.AttributeTag:
			var val uint16
			err = r.Read(&val)
			dataLen -= 2
			retVal = append(retVal, int(val))
		case vrraw.SignedShort:
			var val int16
			err = r.Read(&val)
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
