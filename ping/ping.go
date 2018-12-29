package ping

import "net"

type Pinger interface {
	Ping(host string, port int, out chan Result)
}

type Result struct {
	Success bool
}

type Dialer func(host string, port int) (net.Conn, error)

type pinger struct {
	dialer Dialer
}

func (p pinger) Ping(host string, port int, out chan Result) {
	conn, err := p.dialer(host, port)

	out <- Result{
		Success: err == nil && conn != nil,
	}
}

func NewWithDialer(dialer Dialer) Pinger {
	return pinger{
		dialer: dialer,
	}
}
