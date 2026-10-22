package main

import (
	"io"
	"net"
	"time"
)

const (
	lineDelimiter = '\n'
	writeTimeout  = 5 * time.Second
)

type Client struct {
	conn net.Conn
	addr string
}

func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn, addr: conn.RemoteAddr().String()}
}

func (c *Client) Send(msg string) error {
	if err := c.conn.SetWriteDeadline(time.Now().Add(writeTimeout)); err != nil {
		return err
	}
	_, err := io.WriteString(c.conn, msg+string(lineDelimiter))
	return err
}
