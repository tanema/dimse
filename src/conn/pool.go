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
	timeout  time.Duration
	resource chan int
}

func NewPool(addr string, limit int, timeout time.Duration) *Pool {
	p := &Pool{
		addr:     addr,
		resource: make(chan int, limit),
	}
	for range limit {
		p.resource <- 1
	}
	return p
}

func (p *Pool) Aquire(ctx context.Context) (*Conn, error) {
	<-p.resource
	return Connect(ctx, p.addr, p.timeout)
}

func (p *Pool) Release(con *Conn) error {
	p.resource <- 1
	return con.Close()
}
