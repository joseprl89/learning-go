package ping

type Result struct {
	Success bool
}

func Server(host string, out chan Result) {
	out <- Result{
		Success: true,
	}
}
