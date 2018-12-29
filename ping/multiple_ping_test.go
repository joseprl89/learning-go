package ping_test

import (
	"euler/mocks"
	"euler/ping"
	"net"
	"testing"
)

func TestMultiplePingSuccess(t *testing.T) {
	assertValidResult := func(result ping.Result) {
		if !result.Success {
			t.Error("Did not PingOnce successfully the server when success expected.")
		}
	}

	channel := make(chan ping.Result)

	sut := ping.NewWithDialer(func(host string, port int) (conn net.Conn, e error) {
		return &mocks.Connection{}, nil
	})

	go sut.Ping("google.com", 443, channel, 5)

	count := 0
	for result := range channel {
		count++
		assertValidResult(result)
	}

	if count != 5 {
		t.Errorf("Expected 5 iterations, but received %d", count)
	}
}
