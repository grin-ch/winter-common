package ws

import (
	"github.com/gorilla/websocket"
	"github.com/grin-ch/winter-common/socket"
)

func NewWsConn(c *websocket.Conn, jwt string) socket.Conn {
	return &conn{
		conn: c,
		jwt:  jwt,
	}
}

type conn struct {
	conn *websocket.Conn
	jwt  string
}

func (c *conn) Ip() string {
	return c.conn.RemoteAddr().String()
}

func (c *conn) Protocol() string {
	return "ws"
}

func (c *conn) Close() error {
	return c.conn.Close()
}

func (c *conn) Handshake() ([]byte, error) {
	return []byte(c.jwt), nil
}

func (c *conn) Recv() ([]byte, error) {
	_, msg, err := c.conn.ReadMessage()
	return msg, err
}

func (c *conn) Send(pack []byte) error {
	return c.conn.WriteMessage(websocket.BinaryMessage, pack)
}
