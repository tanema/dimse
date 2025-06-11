package dimse

import (
	"github.com/suyashkumar/dicom"
)

type CStore struct {
	Dataset                              *dicom.Dataset
	AffectedSOPClassUID                  string
	MessageID                            int
	Priority                             int
	CommandDataSetType                   int
	AffectedSOPInstanceUID               string
	MoveOriginatorApplicationEntityTitle string
	MoveOriginatorMessageID              int
	Status                               Status
}

func (v *CStore) Encode(w *dicom.Writer) error {
	if err := writeElement(w, CommandField, uint16(1)); err != nil {
		return err
	} else if err := writeElement(w, AffectedSOPClassUID, v.AffectedSOPClassUID); err != nil {
		return err
	} else if err := writeElement(w, MessageID, v.MessageID); err != nil {
		return err
	} else if err := writeElement(w, Priority, v.Priority); err != nil {
		return err
	} else if err := writeElement(w, CommandDataSetType, v.CommandDataSetType); err != nil {
		return err
	} else if err := writeElement(w, AffectedSOPInstanceUID, v.AffectedSOPInstanceUID); err != nil {
		return err
	}

	if v.MoveOriginatorApplicationEntityTitle != "" {
		if err := writeElement(w, MoveOriginatorApplicationEntityTitle, v.MoveOriginatorApplicationEntityTitle); err != nil {
			return err
		}
	}

	if v.MoveOriginatorMessageID != 0 {
		if err := writeElement(w, MoveOriginatorMessageID, v.MoveOriginatorMessageID); err != nil {
			return err
		}
	}

	return nil
}

func decodeCStoreRsp(d *dicom.Dataset) (*CStore, error) {
	store := &CStore{Dataset: d}
	if err := readElement(d, AffectedSOPClassUID, store.AffectedSOPClassUID, true); err != nil {
		return nil, err
	} else if err := readElement(d, MessageIDBeingRespondedTo, store.MessageID, true); err != nil {
		return nil, err
	} else if err := readElement(d, CommandDataSetType, store.CommandDataSetType, true); err != nil {
		return nil, err
	} else if err := readElement(d, AffectedSOPInstanceUID, store.AffectedSOPInstanceUID, true); err != nil {
		return nil, err
	}
	var err error
	store.Status, err = readStatus(d)
	return store, err
}
