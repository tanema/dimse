package conn

import (
	"context"
	"time"
)

// A pool will create a limited amount of connections to aquire. It does not keep
// them alive however because dimse connections are closed after the association
// is released.
type Pool struct {
	aetitle  string
	cfg      *Config
	timeout  time.Duration
	resource chan int
}

func NewPool(aetitle string, cfg Config) *Pool {
	cfg.Validate()
	p := &Pool{
		aetitle:  aetitle,
		cfg:      &cfg,
		resource: make(chan int, cfg.MaxConnections),
	}
	for range cfg.MaxConnections {
		p.resource <- 1
	}
	return p
}

func (p *Pool) Aquire(ctx context.Context, ent Entity) (*Conn, error) {
	<-p.resource
	return Connect(ctx, p.aetitle, ent, p.cfg)
}

func (p *Pool) Release(con *Conn) error {
	p.resource <- 1
	return con.Close()
}
