package ping_test

import (
	"euler/mocks"
	"euler/ping"
	"net"
	"testing"
)

func TestPing(t *testing.T) {
	channel := make(chan ping.Result)
	sut := ping.NewWithDialer(func(host string, port int) (conn net.Conn, e error) {
		return &mocks.Connection{}, nil
	})

	go sut.Ping("google.com", 443, channel)

	result := <-channel

	if !result.Success {
		t.Error("Did not Ping successfully the server.")
	}
}
