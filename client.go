package dimse

import (
	"context"
	"fmt"

	"github.com/suyashkumar/dicom"
	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/conn"
	"github.com/tanema/dimse/src/serviceobjectpair"
)

type (
	// Client is the object that will interact with the PACs
	Client struct {
		cfg  ClientConfig
		pool *conn.Pool
	}
	// ClientConfig allows for configuring how the client connects to PACS
	ClientConfig struct {
		Conn conn.Config
	}
)

// DefaultConfig is the fallback config for new Clients
var DefaultConfig = &ClientConfig{
	Conn: *conn.DefaultConfig,
}

// NewClient will create a new client for interacting with a PACs
func NewClient(addr string, cfg *ClientConfig) *Client {
	if cfg == nil {
		cfg = DefaultConfig
	}
	return &Client{
		cfg:  *cfg,
		pool: conn.NewPool(addr, cfg.Conn),
	}
}

func (c *Client) dispatch(ctx context.Context, expected commands.Kind, cmd *commands.Command, payload []byte) ([]dicom.Dataset, error) {
	if cmd.CommandDataSetType != commands.Null && len(payload) == 0 {
		return nil, fmt.Errorf("empty payload provided to a command that requires a payload")
	}
	conn, err := c.pool.Aquire(ctx)
	if err != nil {
		return nil, err
	}
	defer c.pool.Release(conn)

	ctxMan, err := conn.Associate(cmd.AffectedSOPClassUID, nil)
	if err != nil {
		return nil, err
	}

	respCmd, data, err := conn.Pdata(ctxMan, cmd, payload)
	if err != nil {
		conn.Abort()
		return nil, err
	} else if err := conn.Realease(); err != nil {
		return nil, err
	} else if cmd.MessageID != respCmd.MessageIDBeingRespondedTo {
		return nil, fmt.Errorf("received %v message id but sent %v", cmd.MessageID, respCmd.MessageID)
	} else if respCmd.CommandField != expected {
		return nil, fmt.Errorf("received %s in response to %s", expected, cmd.CommandField)
	}
	return data, nil
}

// Echo will issue an echo command and will return an error if something went wrong.
// No error will be returned if the error command returned successfully.
func (c *Client) Echo(ctx context.Context) error {
	_, err := c.dispatch(ctx, commands.CECHORSP, &commands.Command{
		CommandField:        commands.CECHORQ,
		AffectedSOPClassUID: serviceobjectpair.VerificationClasses,
		CommandDataSetType:  commands.Null,
	}, nil)
	return err
}
