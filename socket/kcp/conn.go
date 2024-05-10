package kcp

import (
	"io"

	"github.com/grin-ch/winter-common/socket"
	kcpgo "github.com/xtaci/kcp-go/v5"
)

const (
	buf_size = 1024
)

func NewConn(sess *kcpgo.UDPSession) socket.Conn {
	return &conn{
		sess: sess,
		rBuf: make([]byte, buf_size),
	}
}

type conn struct {
	sess *kcpgo.UDPSession
	rBuf []byte
}

func (c *conn) Ip() string {
	return c.sess.RemoteAddr().String()
}

func (c *conn) Protocol() string {
	return "kcp"
}

func (c *conn) Handshake() ([]byte, error) {
	return c.Recv()
}

func (c *conn) Close() error {
	return c.sess.Close()
}

func (c *conn) Recv() ([]byte, error) {
	return recv(c.sess, c.rBuf)
}
func recv(r io.Reader, buf []byte) ([]byte, error) {
	data := make([]byte, 0, buf_size)
	for {
		n, err := r.Read(buf)
		if err != nil {
			return nil, err
		}

		if n == 0 {
			break
		}
		data = append(data, buf[:n]...)
		if n < buf_size {
			break
		}
	}
	return data, nil
}

func (c *conn) Send(msg []byte) error {
	_, err := c.sess.Write(msg)
	return err
}
