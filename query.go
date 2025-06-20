package dimse

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"slices"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/query"
	"github.com/tanema/dimse/src/serviceobjectpair"
)

// Query is a captured, validated query scope for find, get, move, and store
type Query struct {
	client   *Client
	payload  []byte
	Level    query.Level
	Filter   []*dicom.Element
	Priority int // CStore CMove CGet CFind
}

func (c *Client) Query(level query.Level, q []*dicom.Element) (*Query, error) {
	query := &Query{
		client: c,
		Level:  level,
		Filter: q,
	}

	if len(q) == 0 {
		return nil, fmt.Errorf("Query: empty query")
	} else if err := query.encodePayload(); err != nil {
		return nil, fmt.Errorf("Query: unable to encode query payload %v", err)
	}
	return query, nil
}

func (q *Query) SetPriority(p int) *Query {
	q.Priority = p
	return q
}

func (q *Query) Find(ctx context.Context) (*commands.Command, []dicom.Dataset, error) {
	resp, data, err := q.client.dispatch(ctx, &commands.Command{
		CommandField:        commands.CFINDRQ,
		AffectedSOPClassUID: []serviceobjectpair.UID{q.sopForCmd(commands.CFINDRQ)},
		CommandDataSetType:  commands.NonNull,
		Priority:            q.Priority,
	}, q.payload)
	if err != nil {
		return nil, nil, err
	} else if resp.CommandField != commands.CFINDRSP {
		return nil, nil, fmt.Errorf("received %s in response to find", resp.CommandField)
	}
	return resp, data, nil
}

func (q *Query) Get(ctx context.Context) ([]dicom.Dataset, error) {
	resp, data, err := q.client.dispatch(ctx, &commands.Command{
		CommandField:        commands.CGETRQ,
		AffectedSOPClassUID: []serviceobjectpair.UID{q.sopForCmd(commands.CGETRQ)},
		CommandDataSetType:  commands.NonNull,
		Priority:            q.Priority,
	}, q.payload)
	if err != nil {
		return nil, err
	} else if resp.CommandField != commands.CGETRSP {
		return nil, fmt.Errorf("received %s in response to find", resp.CommandField)
	}
	return data, nil
}

func (q *Query) Move(ctx context.Context, dst string) ([]dicom.Dataset, error) {
	resp, data, err := q.client.dispatch(ctx, &commands.Command{
		CommandField:        commands.CMOVERQ,
		AffectedSOPClassUID: []serviceobjectpair.UID{q.sopForCmd(commands.CMOVERQ)},
		Priority:            q.Priority,
		MoveDestination:     dst,
		CommandDataSetType:  commands.NonNull,
	}, q.payload)
	if err != nil {
		return nil, err
	} else if resp.CommandField != commands.CMOVERSP {
		return nil, fmt.Errorf("received %s in response to find", resp.CommandField)
	}
	return data, nil
}

func (q *Query) Store(ctx context.Context, inst []serviceobjectpair.UID, id int, dst, title string) ([]dicom.Dataset, error) {
	resp, data, err := q.client.dispatch(ctx, &commands.Command{
		CommandField:                         commands.CSTORERQ,
		AffectedSOPClassUID:                  []serviceobjectpair.UID{q.sopForCmd(commands.CMOVERQ)},
		CommandDataSetType:                   commands.NonNull,
		Priority:                             q.Priority,
		MoveDestination:                      dst,
		AffectedSOPInstanceUID:               inst,
		MoveOriginatorApplicationEntityTitle: title,
		MoveOriginatorMessageID:              id,
	}, q.payload)
	if err != nil {
		return nil, err
	} else if resp.CommandField != commands.CMOVERSP {
		return nil, fmt.Errorf("received %s in response to find", resp.CommandField)
	}
	return data, nil
}

// collectSOPs will collect SOPs from all commands to be put into the association
// request, and ensure that they are unique.
func collectSOPs(cmds ...*commands.Command) []serviceobjectpair.UID {
	sops := []serviceobjectpair.UID{}
	for _, cmd := range cmds {
		sops = append(sops, cmd.AffectedSOPClassUID...)
	}
	slices.Sort(sops)
	return slices.Compact(sops)
}

func (q *Query) sopForCmd(kind commands.Kind) serviceobjectpair.UID {
	switch q.Level {
	case query.Patient:
		switch kind {
		case commands.CFINDRQ:
			return serviceobjectpair.PatientRootQueryRetrieveInformationModelFind
		case commands.CGETRQ:
			return serviceobjectpair.PatientRootQueryRetrieveInformationModelGet
		case commands.CMOVERQ:
			return serviceobjectpair.PatientRootQueryRetrieveInformationModelMove
		}
	case query.Study, query.Series:
		switch kind {
		case commands.CFINDRQ:
			return serviceobjectpair.StudyRootQueryRetrieveInformationModelFind
		case commands.CGETRQ:
			return serviceobjectpair.StudyRootQueryRetrieveInformationModelGet
		case commands.CMOVERQ:
			return serviceobjectpair.StudyRootQueryRetrieveInformationModelMove
		}
	}
	return ""
}

func (q *Query) encodePayload() error {
	foundQRLevel := false
	buf := bytes.NewBuffer([]byte{})
	w, err := dicom.NewWriter(buf)
	if err != nil {
		return err
	}
	w.SetTransferSyntax(binary.LittleEndian, true)
	for _, elem := range q.Filter {
		if elem.Tag == tag.QueryRetrieveLevel {
			foundQRLevel = true
		}
		if err := w.WriteElement(elem); err != nil {
			return err
		}
	}
	if !foundQRLevel {
		var qrLevelString string
		switch q.Level {
		case query.Patient:
			qrLevelString = "PATIENT"
		case query.Study:
			qrLevelString = "STUDY"
		case query.Series:
			qrLevelString = "SERIES"
		}
		elem, err := dicom.NewElement(tag.QueryRetrieveLevel, []string{qrLevelString})
		if err != nil {
			return err
		}
		if err := w.WriteElement(elem); err != nil {
			return err
		}
	}
	q.payload = buf.Bytes()
	return nil
}
