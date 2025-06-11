package dimse

import (
	"github.com/suyashkumar/dicom"
)

type CGet struct {
	Dataset                        *dicom.Dataset
	AffectedSOPClassUID            string
	MessageID                      int
	Priority                       int
	CommandDataSetType             int
	NumberOfRemainingSuboperations int
	NumberOfCompletedSuboperations int
	NumberOfFailedSuboperations    int
	NumberOfWarningSuboperations   int
	Status                         Status
}

func (v *CGet) Encode(w *dicom.Writer) error {
	if err := writeElement(w, CommandField, uint16(16)); err != nil {
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

func decodeCGetRsp(d *dicom.Dataset) (*CGet, error) {
	get := &CGet{Dataset: d}
	if err := readElement(d, AffectedSOPClassUID, get.AffectedSOPClassUID, true); err != nil {
		return nil, err
	} else if err := readElement(d, MessageIDBeingRespondedTo, get.MessageID, true); err != nil {
		return nil, err
	} else if err := readElement(d, CommandDataSetType, get.CommandDataSetType, true); err != nil {
		return nil, err
	} else if err := readElement(d, NumberOfRemainingSuboperations, get.NumberOfRemainingSuboperations, false); err != nil {
		return nil, err
	} else if err := readElement(d, NumberOfCompletedSuboperations, get.NumberOfCompletedSuboperations, false); err != nil {
		return nil, err
	} else if err := readElement(d, NumberOfFailedSuboperations, get.NumberOfFailedSuboperations, false); err != nil {
		return nil, err
	} else if err := readElement(d, NumberOfWarningSuboperations, get.NumberOfWarningSuboperations, false); err != nil {
		return nil, err
	}
	var err error
	get.Status, err = readStatus(d)
	return get, err
}
