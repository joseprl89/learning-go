package ping_test

import (
	"euler/ping"
	"testing"
)

func TestPing(t *testing.T) {
	channel := make(chan ping.Result)
	go ping.Server("https://www.google.com/", channel)

	result := <-channel

	if !result.Success {
		t.Error("Did not ping successfully the server.")
	}
}
