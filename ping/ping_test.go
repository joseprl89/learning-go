package ping_test

import (
	"errors"
	"euler/mocks"
	"euler/ping"
	"net"
	"strings"
	"testing"
)

func TestPingSuccess(t *testing.T) {
	channel := make(chan ping.Result)
	sut := ping.NewWithDialer(func(host string, port int) (conn net.Conn, e error) {
		return &mocks.Connection{}, nil
	})

	go sut.Ping("google.com", 443, channel)

	result := <-channel

	if !result.Success {
		t.Error("Did not Ping successfully the server when success expected.")
	}
}

func TestPingFailure(t *testing.T) {
	channel := make(chan ping.Result)
	sut := ping.NewWithDialer(func(host string, port int) (conn net.Conn, e error) {
		return nil, errors.New("fail")
	})

	go sut.Ping("google.com", 443, channel)

	result := <-channel

	if result.Success {
		t.Error("Did Ping successfully the server when failure expected.")
	}
}

func TestPingFailureWhenErrorAndConnection(t *testing.T) {
	channel := make(chan ping.Result)
	sut := ping.NewWithDialer(func(host string, port int) (conn net.Conn, e error) {
		return &mocks.Connection{}, errors.New("fail")
	})

	go sut.Ping("google.com", 443, channel)

	result := <-channel

	if result.Success {
		t.Error("Did Ping successfully the server when failure expected.")
	}
}

func TestPingFailureWhenBothConnectionAndErrorAreNil(t *testing.T) {
	channel := make(chan ping.Result)
	sut := ping.NewWithDialer(func(host string, port int) (conn net.Conn, e error) {
		return nil, nil
	})

	go sut.Ping("google.com", 443, channel)

	result := <-channel

	if result.Success {
		t.Error("Did Ping successfully the server when failure expected.")
	}
}

func TestPingSuccessDescription(t *testing.T) {
	channel := make(chan ping.Result)
	sut := ping.NewWithDialer(func(host string, port int) (conn net.Conn, e error) {
		return &mocks.Connection{}, nil
	})

	go sut.Ping("google.com", 443, channel)

	result := <-channel

	expectedStartOfDescription := "From google.com: icmp_seq=0 time="

	if strings.HasPrefix(result.Description, expectedStartOfDescription) {
		t.Errorf("Description did not match expectations. Got %s, expected %s.",
			result.Description,
			expectedStartOfDescription)
	}
}
