package dimse

import (
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
)

type (
	// Message defines the common interface for all DIMSE message types.
	Message interface{ Encode(d *dicom.Writer) error }
	Status  struct {
		Status       int
		ErrorComment string
	}
)

const (
	CommandDataSetTypeNull    int = 0x101
	CommandDataSetTypeNonNull int = 1
)

func writeElement(w *dicom.Writer, t tag.Tag, val any) error {
	elem, err := dicom.NewElement(t, val)
	if err != nil {
		return err
	}
	return w.WriteElement(elem)
}

func readElement[T any](ds *dicom.Dataset, t tag.Tag, dst T, required bool) error {
	elem, err := ds.FindElementByTagNested(t)
	if err != nil {
		if err == dicom.ErrorElementNotFound && !required {
			return nil
		}
		return err
	}
	val, ok := elem.Value.GetValue().(T)
	if !ok {
		return fmt.Errorf("element value is %T", elem.Value.GetValue())
	}
	dst = val
	return nil
}

func readStatus(ds *dicom.Dataset) (Status, error) {
	s := Status{}
	if err := readElement(ds, StatusTag, &s.Status, true); err != nil {
		return s, err
	}
	return s, readElement(ds, ErrorComment, &s.ErrorComment, false)
}

func decodeMessage(d *dicom.Dataset) (Message, error) {
	var commandField uint16
	for _, elem := range d.Elements {
		fmt.Println(tag.MustFind(elem.Tag).Name, elem.Value.String())
	}
	if err := readElement(d, CommandField, commandField, true); err != nil {
		return nil, err
	}
	switch commandField {
	case 0x8020:
		return decodeCFindRsp(d)
	case 0x8030:
		return decodeCEchoRsp(d)
	case 0x8001:
		return decodeCStoreRsp(d)
	case 0x8010:
		return decodeCGetRsp(d)
	case 0x8021:
		return decodeCMoveRsp(d)
	case 0x1:
		return nil, fmt.Errorf("client received cstore request instead of response.")
	case 0x21:
		return nil, fmt.Errorf("client received cmove request instead of response.")
	case 0x10:
		return nil, fmt.Errorf("client received cget request instead of response.")
	case 0x20:
		return nil, fmt.Errorf("client received cfind request instead of response.")
	case 0x30:
		return nil, fmt.Errorf("client received echo request instead of response.")
	default:
		return nil, fmt.Errorf("Unknown DIMSE command 0x%x", commandField)
	}
}
