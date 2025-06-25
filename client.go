package dimse

import (
	"bytes"
	"context"
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"
	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/conn"
	"github.com/tanema/dimse/src/serviceobjectpair"
)

// Client is the object that will interact with the PACs
type Client struct {
	cfg  Config
	pool *conn.Pool
}

// NewClient will create a new client for interacting with a PACs
func NewClient(cfg Config) (*Client, error) {
	return &Client{
		cfg:  cfg,
		pool: conn.NewPool(cfg.AETitle, cfg.Conn),
	}, cfg.Validate()
}

func (c *Client) dispatch(ctx context.Context, entity conn.Entity, cmd *commands.Command, payload []byte) ([]dicom.Dataset, error) {
	// Check if already cancelled
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if cmd.CommandDataSetType != commands.Null && len(payload) == 0 {
		return nil, fmt.Errorf("empty payload provided to a command that requires a payload")
	}
	conn, err := c.pool.Aquire(ctx, entity)
	if err != nil {
		return nil, err
	}
	defer c.pool.Release(conn)

	// Check if cancelled
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	ctxMan, err := conn.Associate(cmd.AffectedSOPClassUID, nil)
	if err != nil {
		return nil, err
	}

	// Check if cancelled
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	respCmd, data, err := conn.Pdata(ctx, ctxMan, cmd, payload)
	if err != nil {
		conn.Abort()
		return nil, err
	} else if err := conn.Realease(); err != nil {
		return nil, err
	} else if cmd.MessageID != respCmd.MessageIDBeingRespondedTo {
		return nil, fmt.Errorf("received %v message id but sent %v", cmd.MessageID, respCmd.MessageID)
	}
	return data, nil
}

// Echo will issue an echo command and will return an error if something went wrong.
// No error will be returned if the error command returned successfully.
func (c *Client) Echo(ctx context.Context, entity conn.Entity) error {
	_, err := c.dispatch(ctx, entity, &commands.Command{
		CommandField:        commands.CECHORQ,
		AffectedSOPClassUID: serviceobjectpair.VerificationClasses,
		CommandDataSetType:  commands.Null,
	}, nil)
	return err
}

func (c *Client) Store(ctx context.Context, entity conn.Entity, ds dicom.Dataset) ([]dicom.Dataset, error) {
	element, err := ds.FindElementByTag(tag.SOPClassUID)
	if err != nil {
		return nil, fmt.Errorf("could not find datasets sop class uid: %v", err)
	}
	val, ok := element.Value.GetValue().([]string)
	if !ok {
		return nil, fmt.Errorf("datasets sop class uid is an invalid value")
	}

	sopUIDs := make([]serviceobjectpair.UID, len(val))
	for i, uid := range val {
		sopUIDs[i] = serviceobjectpair.UID(uid)
	}

	buf := bytes.NewBuffer([]byte{})
	if err := dicom.Write(buf, ds); err != nil {
		return nil, err
	}
	return c.dispatch(ctx, entity, &commands.Command{
		CommandField:        commands.CSTORERQ,
		AffectedSOPClassUID: append(sopUIDs, serviceobjectpair.StorageManagementClasses...),
		CommandDataSetType:  commands.NonNull,
	}, buf.Bytes())
}
