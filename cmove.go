package dimse

import (
	"github.com/suyashkumar/dicom"
)

type CMove struct {
	Dataset                        *dicom.Dataset
	AffectedSOPClassUID            string
	MessageID                      int
	Priority                       int
	MoveDestination                string
	CommandDataSetType             int
	NumberOfRemainingSuboperations int
	NumberOfCompletedSuboperations int
	NumberOfFailedSuboperations    int
	NumberOfWarningSuboperations   int
	Status                         Status
}

func (v *CMove) Encode(w *dicom.Writer) error {
	if err := writeElement(w, CommandField, uint16(33)); err != nil {
		return err
	} else if err := writeElement(w, AffectedSOPClassUID, v.AffectedSOPClassUID); err != nil {
		return err
	} else if err := writeElement(w, MessageID, v.MessageID); err != nil {
		return err
	} else if err := writeElement(w, Priority, v.Priority); err != nil {
		return err
	} else if err := writeElement(w, MoveDestination, v.MoveDestination); err != nil {
		return err
	}
	return writeElement(w, CommandDataSetType, v.CommandDataSetType)
}

func decodeCMoveRsp(d *dicom.Dataset) (*CMove, error) {
	move := &CMove{Dataset: d}
	if err := readElement(d, AffectedSOPClassUID, move.AffectedSOPClassUID, true); err != nil {
		return nil, err
	} else if err := readElement(d, MessageIDBeingRespondedTo, move.MessageID, true); err != nil {
		return nil, err
	} else if err := readElement(d, CommandDataSetType, move.CommandDataSetType, true); err != nil {
		return nil, err
	} else if err := readElement(d, NumberOfRemainingSuboperations, move.NumberOfRemainingSuboperations, false); err != nil {
		return nil, err
	} else if err := readElement(d, NumberOfCompletedSuboperations, move.NumberOfCompletedSuboperations, false); err != nil {
		return nil, err
	} else if err := readElement(d, NumberOfFailedSuboperations, move.NumberOfFailedSuboperations, false); err != nil {
		return nil, err
	} else if err := readElement(d, NumberOfWarningSuboperations, move.NumberOfWarningSuboperations, false); err != nil {
		return nil, err
	}
	var err error
	move.Status, err = readStatus(d)
	return move, err
}
