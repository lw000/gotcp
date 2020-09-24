package network

import (
	"net"
	"time"
)

type TCPClient struct {
	session *Session
}

func NewTCPClient() *TCPClient {
	return &TCPClient{
		session: newSession(nil, nil, 0, nil),
	}
}

func (c *TCPClient) SetMessageHandler(fn MessageHandler) {
	c.session.onHandler = fn
}

func (c *TCPClient) Connect(host, port string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp4", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return err
	}
	c.session.SetConn(conn)
	return nil
}

func (c *TCPClient) Session() *Session {
	return c.session
}

func (c *TCPClient) Loop() {
	c.session.Run()
}
