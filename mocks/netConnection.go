package mocks

import (
	"errors"
	"net"
	"time"
)

type Connection struct {
	IsClosed      bool
	WrittenBytes  []byte
	LocalAddress  net.Addr
	RemoteAddress net.Addr
}

func (Connection) Read(b []byte) (n int, err error) {
	return 0, errors.New("not implemented")
}

func (conn Connection) Write(b []byte) (n int, err error) {
	conn.WrittenBytes = append(conn.WrittenBytes, b...)
	return len(b), nil
}

func (conn Connection) Close() error {
	conn.IsClosed = true
	return nil
}

func (conn Connection) LocalAddr() net.Addr {
	return conn.LocalAddress
}

func (conn Connection) RemoteAddr() net.Addr {
	return conn.RemoteAddress
}

func (Connection) SetDeadline(t time.Time) error {
	panic("implement me")
}

func (Connection) SetReadDeadline(t time.Time) error {
	panic("implement me")
}

func (Connection) SetWriteDeadline(t time.Time) error {
	panic("implement me")
}
