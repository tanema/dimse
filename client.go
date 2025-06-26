package dimse

import (
	"context"
	"fmt"
	"net"

	"github.com/suyashkumar/dicom"
	"github.com/suyashkumar/dicom/pkg/tag"

	"github.com/tanema/dimse/src/commands"
	"github.com/tanema/dimse/src/defn/serviceobjectpair"
	"github.com/tanema/dimse/src/defn/transfersyntax"
	"github.com/tanema/dimse/src/pdu"
)

// Client is the object that will interact with the PACs
type Client struct {
	cfg      Config
	listener net.Listener
	pool     chan int
}

// NewClient will create a new client for interacting with a PACs
func NewClient(cfg Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		return nil, err
	}
	client := &Client{
		cfg:      cfg,
		pool:     make(chan int, cfg.Conn.MaxConnections),
		listener: l,
	}
	for range cfg.Conn.MaxConnections {
		client.pool <- 1
	}
	go client.listen()
	return client, nil
}

func (c *Client) listen() {
	for {
		conn, err := c.listener.Accept()
		if err != nil {
			return
		}
		fmt.Println(pdu.NewReader(conn).Next())
	}
}

func (c *Client) Close() error {
	return c.listener.Close()
}

func (c *Client) aquireConn(ctx context.Context, entity Entity) (*Conn, error) {
	<-c.pool
	return Connect(ctx, c.cfg.AETitle, entity, &c.cfg.Conn)
}

func (c *Client) releaseConn(con *Conn) error {
	c.pool <- 1
	return con.Close()
}

func (c *Client) dispatch(ctx context.Context, entity Entity, cmd *commands.Command, ds *dicom.Dataset) ([]dicom.Dataset, error) {
	// Check if already cancelled
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	conn, err := c.aquireConn(ctx, entity)
	if err != nil {
		return nil, err
	}
	defer c.releaseConn(conn)

	// Check if cancelled
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if err := conn.Associate(cmd.AffectedSOPClassUID, readTransferSyntax(ds)); err != nil {
		return nil, err
	}

	// Check if cancelled
	if err := ctx.Err(); err != nil {
		conn.Abort()
		return nil, err
	}

	respCmd, data, err := conn.Pdata(cmd, ds)
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
func (c *Client) Echo(ctx context.Context, entity Entity) error {
	_, err := c.dispatch(ctx, entity, &commands.Command{
		CommandField:        commands.CECHORQ,
		AffectedSOPClassUID: serviceobjectpair.VerificationClasses,
	}, nil)
	return err
}

func (c *Client) Store(ctx context.Context, entity Entity, ds dicom.Dataset) error {
	sopClassUIDs, err := getSOPUIDs(ds, tag.SOPClassUID)
	if err != nil {
		return err
	}

	sopInstanceUIDs, err := getSOPUIDs(ds, tag.SOPInstanceUID)
	if err != nil {
		return err
	}

	_, err = c.dispatch(ctx, entity, &commands.Command{
		CommandField:           commands.CSTORERQ,
		AffectedSOPClassUID:    append(sopClassUIDs, serviceobjectpair.StorageManagementClasses...),
		AffectedSOPInstanceUID: sopInstanceUIDs,
	}, &ds)
	return err
}

func getSOPUIDs(ds dicom.Dataset, t tag.Tag) ([]serviceobjectpair.UID, error) {
	val, err := readElementVal(ds, t)
	if err != nil {
		return nil, err
	}
	sopUIDs := make([]serviceobjectpair.UID, len(val))
	for i, uid := range val {
		sopUIDs[i] = serviceobjectpair.UID(uid)
	}
	return sopUIDs, nil
}

func readTransferSyntax(ds *dicom.Dataset) []transfersyntax.UID {
	if ds == nil {
		return transfersyntax.StandardSyntaxes
	}

	val, err := readElementVal(*ds, tag.TransferSyntaxUID)
	if err != nil {
		return transfersyntax.StandardSyntaxes
	}
	ts := make([]transfersyntax.UID, len(val))
	for i, uid := range val {
		ts[i] = transfersyntax.UID(uid)
	}
	return ts
}

func readElementVal(ds dicom.Dataset, t tag.Tag) ([]string, error) {
	info, _ := tag.Find(t)

	element, err := ds.FindElementByTag(t)
	if err != nil {
		return nil, err
	}
	val, ok := element.Value.GetValue().([]string)
	if !ok {
		return nil, fmt.Errorf("%v is wrong format", info.Name)
	}
	return val, nil
}
