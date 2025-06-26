package dimse

import (
	"time"

	"github.com/tanema/dimse/src/pdu"
)

type (
	// Config allows for configuring how the client connects to PACS
	Config struct {
		AETitle string           `json:"aetitle" yaml:"aetitle"`
		Port    int              `json:"port" yaml:"port"`
		Conn    ConnectionConfig `json:"conn" yaml:"conn"`
	}
	ConnectionConfig struct {
		MaxConnections int           `json:"max_connections" yaml:"maxConnections"`
		Timeout        time.Duration `json:"timeout" yaml:"timeout"`
		ChunkSize      uint32        `json:"chunk_size" yaml:"chunkSize"`
	}
	Entity struct {
		Title string `json:"title" yaml:"title"`
		Host  string `json:"host" yaml:"host"`
		Port  int    `json:"port" yaml:"port"`
	}
)

func (c *Config) Validate() error {
	if c.AETitle == "" {
		c.AETitle = "anon-user"
	}
	if c.Port <= 0 {
		c.Port = 104
	}
	if c.Conn.MaxConnections <= 0 {
		c.Conn.MaxConnections = 10
	}
	if c.Conn.Timeout <= 0 {
		c.Conn.Timeout = time.Second
	}
	if c.Conn.ChunkSize <= 0 {
		c.Conn.ChunkSize = pdu.DefaultMaxPDUSize
	}
	return nil
}
