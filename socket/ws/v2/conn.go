package ws

import (
	"github.com/grin-ch/winter-common/socket"
	"github.com/lxzan/gws"
)

func NewWsConn(c *gws.Conn, jwt string) socket.Conn {
	return &conn{
		conn: c,
		jwt:  jwt,
	}
}

type conn struct {
	conn *gws.Conn
	jwt  string
}

// Close implements socket.Conn.
func (c *conn) Close() error {
	c.conn.WriteClose(1000, nil)
	return nil
}

// Ip implements socket.Conn.
func (c *conn) Ip() string {
	return c.conn.RemoteAddr().String()
}

// Protocol implements socket.Conn.
func (c *conn) Protocol() string {
	return "ws"
}

// Recv implements socket.Conn.
func (c *conn) Recv() ([]byte, error) {
	c.conn.ReadLoop()
	return nil, nil
}

// Send implements socket.Conn.
func (c *conn) Send(pack []byte) error {
	return c.conn.WriteMessage(gws.OpcodeBinary, pack)
}
