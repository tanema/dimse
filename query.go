package dimse

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"

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
	level    query.Level
	priority int // CStore CMove CGet CFind
}

// Build the query to be used to run a command. Will return an error if the query
// is empty, or the query elements are invalid.
func (c *Client) Query(level query.Level, q []*dicom.Element) (*Query, error) {
	if len(q) == 0 {
		return nil, fmt.Errorf("Query: empty query")
	}

	buf := bytes.NewBuffer([]byte{})
	w, err := dicom.NewWriter(buf)
	if err != nil {
		return nil, err
	}
	w.SetTransferSyntax(binary.LittleEndian, true)
	if elem, err := dicom.NewElement(tag.QueryRetrieveLevel, []string{string(level)}); err != nil {
		return nil, err
	} else if err := w.WriteElement(elem); err != nil {
		return nil, err
	}
	for _, elem := range q {
		if err := w.WriteElement(elem); err != nil {
			return nil, err
		}
	}

	return &Query{
		client:  c,
		level:   level,
		payload: buf.Bytes(),
	}, nil
}

// SetPriority will set the priority value for the query.
func (q *Query) SetPriority(p int) *Query {
	q.priority = p
	return q
}

// Find will run a C-FIND service command on the built query
func (q *Query) Find(ctx context.Context) ([]dicom.Dataset, error) {
	return q.client.dispatch(ctx, commands.CFINDRSP, &commands.Command{
		CommandField:        commands.CFINDRQ,
		AffectedSOPClassUID: serviceobjectpair.QRFindClasses,
		CommandDataSetType:  commands.NonNull,
		Priority:            commands.Priority(q.priority),
	}, q.payload)
}

// Get will run a C-GET service command on the built query
func (q *Query) Get(ctx context.Context) ([]dicom.Dataset, error) {
	return q.client.dispatch(ctx, commands.CGETRSP, &commands.Command{
		CommandField:        commands.CGETRQ,
		AffectedSOPClassUID: serviceobjectpair.QRGetClasses,
		CommandDataSetType:  commands.NonNull,
		Priority:            commands.Priority(q.priority),
	}, q.payload)
}

// Move will run a C-MOVE service command on the built query
func (q *Query) Move(ctx context.Context, dst string) ([]dicom.Dataset, error) {
	return q.client.dispatch(ctx, commands.CMOVERSP, &commands.Command{
		CommandField:        commands.CMOVERQ,
		AffectedSOPClassUID: serviceobjectpair.QRMoveClasses,
		Priority:            commands.Priority(q.priority),
		MoveDestination:     dst,
		CommandDataSetType:  commands.NonNull,
	}, q.payload)
}
