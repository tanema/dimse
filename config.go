package dimse

import "github.com/tanema/dimse/src/conn"

// Config allows for configuring how the client connects to PACS
type Config struct {
	AETitle string
	Conn    conn.Config
}

func (c *Config) Validate() error {
	if c.AETitle == "" {
		c.AETitle = "anon-user"
	}
	return c.Conn.Validate()
}
