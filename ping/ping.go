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
	dialer        Dialer
	sleepDuration time.Duration
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
	defer close(out)
	for i := 0; i < count; i++ {
		p.PingOnce(host, port, out)
		time.Sleep(p.sleepDuration)
	}
}

func NewWithDialer(dialer Dialer) Pinger {
	return pinger{
		dialer:        dialer,
		sleepDuration: 0,
	}
}

func New() Pinger {
	return pinger{
		dialer: func(host string, port int) (conn net.Conn, e error) {
			return net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
		},
		sleepDuration: time.Duration(1 * time.Second),
	}
}
