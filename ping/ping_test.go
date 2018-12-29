package ping_test

import (
	"euler/ping"
	"net"
	"testing"
)

func TestPing(t *testing.T) {
	channel := make(chan ping.Result)
	sut := ping.NewWithDialer(func(host string, port int) (conn net.Conn, e error) {
		panic("Dont call yet")
	})

	go sut.Ping("google.com", 443, channel)

	result := <-channel

	if !result.Success {
		t.Error("Did not Ping successfully the server.")
	}
}
