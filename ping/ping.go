package ping

import (
	"fmt"
	"net"
	"time"
)

type Pinger interface {
	PingOnce(host string, port int, out chan Result)
}

type Result struct {
	Success     bool
	Description string
}

type Dialer func(host string, port int) (net.Conn, error)

type pinger struct {
	dialer Dialer
}

func (p pinger) PingOnce(host string, port int, out chan Result) {
	pingStartTime := time.Now()

	conn, err := p.dialer(host, port)

	pingDuration := time.Since(pingStartTime)

	out <- Result{
		Success:     err == nil && conn != nil,
		Description: fmt.Sprintf("From %s: icmp_seq=%d time=%d", host, 0, pingDuration),
	}
}

func NewWithDialer(dialer Dialer) Pinger {
	return pinger{
		dialer: dialer,
	}
}
