package dimse

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/suyashkumar/dicom"
	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/conn"
	"github.com/tanema/dimse/src/serviceobjectpair"
)

type (
	Client struct {
		msgID int32
		cfg   ClientConfig
		pool  *conn.Pool
	}
	ClientConfig struct {
		Conn conn.Config
	}
)

var DefaultConfig = &ClientConfig{
	Conn: *conn.DefaultConfig,
}

func NewClient(addr string, cfg *ClientConfig) *Client {
	if cfg == nil {
		cfg = DefaultConfig
	}
	return &Client{
		cfg:  *cfg,
		pool: conn.NewPool(addr, cfg.Conn),
	}
}

func (c *Client) dispatch(ctx context.Context, cmd *commands.Command, payload []byte) (*commands.Command, []dicom.Dataset, error) {
	if cmd.CommandDataSetType != commands.Null && len(payload) == 0 {
		return nil, nil, fmt.Errorf("empty payload provided to a command that requires a payload")
	}
	conn, err := c.pool.Aquire(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer c.pool.Release(conn)

	cmd.MessageID = int(atomic.AddInt32(&c.msgID, 1))
	ctxMan, err := conn.Associate(collectSOPs(cmd), nil)
	if err != nil {
		return nil, nil, err
	}
	respCmd, data, err := conn.Pdata(ctxMan, cmd, payload)
	if err != nil {
		conn.Abort()
		return nil, nil, err
	} else if err := conn.Realease(); err != nil {
		return nil, nil, err
	} else if cmd.MessageID != respCmd.MessageID {
		return nil, nil, fmt.Errorf("received %v message id but sent %v", cmd.MessageID, respCmd.MessageID)
	}
	return respCmd, data, nil
}

// Echo will issue an echo command and will return an error if something went wrong.
// No error will be returned if the error command returned successfully.
func (c *Client) Echo(ctx context.Context) error {
	resp, _, err := c.dispatch(ctx, &commands.Command{
		CommandField:        commands.CECHORQ,
		AffectedSOPClassUID: serviceobjectpair.VerificationClasses,
		CommandDataSetType:  commands.Null,
	}, nil)
	if err != nil {
		return err
	} else if resp.CommandField != commands.CECHORSP {
		return fmt.Errorf("received %s in response to echo", resp.CommandField)
	}
	return nil
}
