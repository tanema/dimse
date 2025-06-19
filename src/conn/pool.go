package conn

import (
	"context"
	"time"
)

// A pool will create a limited amount of connections to aquire. It does not keep
// them alive however because dimse connections are closed after the association
// is released.
type Pool struct {
	addr     string
	cfg      Config
	timeout  time.Duration
	resource chan int
}

func NewPool(addr string, cfg Config) *Pool {
	p := &Pool{
		addr:     addr,
		cfg:      cfg,
		resource: make(chan int, cfg.MaxConnections),
	}
	for range cfg.MaxConnections {
		p.resource <- 1
	}
	return p
}

func (p *Pool) Aquire(ctx context.Context) (*Conn, error) {
	<-p.resource
	return Connect(ctx, p.addr, p.cfg)
}

func (p *Pool) Release(con *Conn) error {
	p.resource <- 1
	return con.Close()
}
