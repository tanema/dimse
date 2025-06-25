package conn

import (
	"time"

	"github.com/tanema/dimse/src/pdu"
)

type (
	Config struct {
		MaxConnections    int
		ConnectionTimeout time.Duration
		ChunkSize         uint32
	}
	Entity struct {
		Title string
		Host  string
		Port  int
	}
)

func (c *Config) Validate() error {
	if c.MaxConnections <= 0 {
		c.MaxConnections = 10
	}
	if c.ConnectionTimeout <= 0 {
		c.ConnectionTimeout = time.Second
	}
	if c.ChunkSize <= 0 {
		c.ChunkSize = pdu.DefaultMaxPDUSize
	}
	return nil
}
