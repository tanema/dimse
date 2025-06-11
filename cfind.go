package dimse

import (
	"github.com/suyashkumar/dicom"
)

type CFind struct {
	AffectedSOPClassUID string
	MessageID           int
	Priority            int
	CommandDataSetType  int
	Status              Status
	DataSet             *dicom.Dataset
}

func (v *CFind) Encode(w *dicom.Writer) error {
	if err := writeElement(w, CommandField, uint16(32)); err != nil {
		return err
	} else if err := writeElement(w, AffectedSOPClassUID, v.AffectedSOPClassUID); err != nil {
		return err
	} else if err := writeElement(w, MessageID, v.MessageID); err != nil {
		return err
	} else if err := writeElement(w, Priority, v.Priority); err != nil {
		return err
	}
	return writeElement(w, CommandDataSetType, v.CommandDataSetType)
}

func decodeCFindRsp(d *dicom.Dataset) (*CFind, error) {
	find := &CFind{DataSet: d}
	if err := readElement(d, AffectedSOPClassUID, find.AffectedSOPClassUID, true); err != nil {
		return nil, err
	} else if err := readElement(d, MessageIDBeingRespondedTo, find.MessageID, true); err != nil {
		return nil, err
	} else if err := readElement(d, CommandDataSetType, find.CommandDataSetType, true); err != nil {
		return nil, err
	} else if err := readElement(d, Priority, find.Priority, true); err != nil {
		return nil, err
	}
	var err error
	find.Status, err = readStatus(d)
	return find, err
}
