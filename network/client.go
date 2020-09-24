package network

import (
	"net"
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

func (c *TCPClient) Connect(host, port string) error {
	conn, err := net.Dial("tcp4", net.JoinHostPort(host, port))
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
