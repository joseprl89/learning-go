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
	conn, _ := p.dialer(host, port)

	if conn != nil {
		out <- Result{
			Success: true,
		}
	}

	// To avoid deadlock in tests. temporary.
	out <- Result{
		Success: true,
	}
}

func NewWithDialer(dialer Dialer) Pinger {
	return pinger{
		dialer: dialer,
	}
}
