package ping

import (
	"fmt"
	"net"
)

type Pinger interface {
	Ping(host string, port int, out chan Result)
}

type Result struct {
	Success     bool
	Description string
}

type Dialer func(host string, port int) (net.Conn, error)

type pinger struct {
	dialer Dialer
}

func (p pinger) Ping(host string, port int, out chan Result) {
	conn, err := p.dialer(host, port)

	out <- Result{
		Success:     err == nil && conn != nil,
		Description: fmt.Sprintf("From %s: icmp_seq=%d time=%d", host, 0, 10000000),
	}
}

func NewWithDialer(dialer Dialer) Pinger {
	return pinger{
		dialer: dialer,
	}
}
