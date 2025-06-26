package dimse

import (
	"context"
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"

	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/defn/query"
	"github.com/tanema/dimse/src/defn/serviceobjectpair"
)

// Query is a captured, validated query scope for find, get, move, and store
type Query struct {
	client   *Client
	entity   Entity
	payload  *dicom.Dataset
	level    query.Level
	priority int // CStore CMove CGet CFind
}

// Build the query to be used to run a command. Will return an error if the query
// is empty, or the query elements are invalid.
func (c *Client) Query(entity Entity, level query.Level, q []*dicom.Element) (*Query, error) {
	if len(q) == 0 {
		return nil, fmt.Errorf("Query: empty query")
	}

	qr, err := dicom.NewElement(tag.QueryRetrieveLevel, []string{string(level)})
	if err != nil {
		return nil, err
	}

	return &Query{
		entity:  entity,
		client:  c,
		level:   level,
		payload: &dicom.Dataset{Elements: append(q, qr)},
	}, nil
}

// SetPriority will set the priority value for the query.
func (q *Query) SetPriority(p int) *Query {
	q.priority = p
	return q
}

// Find will run a C-FIND service command on the built query
func (q *Query) Find(ctx context.Context) ([]dicom.Dataset, error) {
	return q.client.dispatch(ctx, q.entity, &commands.Command{
		CommandField:        commands.CFINDRQ,
		AffectedSOPClassUID: serviceobjectpair.QRFindClasses,
		Priority:            commands.Priority(q.priority),
	}, q.payload)
}

// Get will run a C-GET service command on the built query
func (q *Query) Get(ctx context.Context) ([]dicom.Dataset, error) {
	return q.client.dispatch(ctx, q.entity, &commands.Command{
		CommandField:        commands.CGETRQ,
		AffectedSOPClassUID: serviceobjectpair.QRGetClasses,
		Priority:            commands.Priority(q.priority),
	}, q.payload)
}

// Move will run a C-MOVE service command on the built query
func (q *Query) Move(ctx context.Context, dst string) ([]dicom.Dataset, error) {
	return q.client.dispatch(ctx, q.entity, &commands.Command{
		MoveDestination:     dst,
		CommandField:        commands.CMOVERQ,
		AffectedSOPClassUID: serviceobjectpair.QRMoveClasses,
		Priority:            commands.Priority(q.priority),
	}, q.payload)
}
