package ping

type Result struct {
	Success bool
}

func Server(host string, port int, out chan Result) {
	out <- Result{
		Success: true,
	}
}
