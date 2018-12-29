package ping

import (
	"fmt"
	"net"
	"time"
)

type Pinger interface {
	PingOnce(host string, port int, out chan Result)
	Ping(host string, port int, out chan Result, count int)
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

func (p pinger) Ping(host string, port int, out chan Result, count int) {
	panic("implement me")
}

func NewWithDialer(dialer Dialer) Pinger {
	return pinger{
		dialer: dialer,
	}
}
