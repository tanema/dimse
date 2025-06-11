package dimse

import (
	"github.com/suyashkumar/dicom"
)

type CEcho struct {
	DataSet            *dicom.Dataset
	MessageID          int
	CommandDataSetType int
	Status             Status
}

func (v *CEcho) Encode(w *dicom.Writer) error {
	if err := writeElement(w, CommandField, []int{48}); err != nil {
		return err
	} else if err := writeElement(w, MessageID, []int{v.MessageID}); err != nil {
		return err
	}
	return writeElement(w, CommandDataSetType, []int{v.CommandDataSetType})
}

func decodeCEchoRsp(d *dicom.Dataset) (*CEcho, error) {
	echo := &CEcho{DataSet: d}
	if err := readElement(d, MessageIDBeingRespondedTo, &echo.MessageID, true); err != nil {
		return nil, err
	} else if err := readElement(d, CommandDataSetType, &echo.CommandDataSetType, true); err != nil {
		return nil, err
	}
	var err error
	echo.Status, err = readStatus(d)
	return echo, err
}
