package ws

import "github.com/lxzan/gws"

func NewEventHandle() gws.Event {
	return &Handler{}
}

type msg struct {
	data []byte
	err  error
}

type Handler struct {
	c chan msg
}

// OnClose implements gws.Event.
func (h *Handler) OnClose(socket *gws.Conn, err error) {
	if err != nil {
		socket.WriteClose(1006, []byte(err.Error()))
		return
	}
	socket.WriteClose(1000, nil)
}

// OnMessage implements gws.Event.
func (h *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

}

// OnOpen implements gws.Event.
func (h *Handler) OnOpen(socket *gws.Conn) {
}

// OnPing implements gws.Event.
func (h *Handler) OnPing(socket *gws.Conn, payload []byte) {
}

// OnPong implements gws.Event.
func (h *Handler) OnPong(_ *gws.Conn, _ []byte) {}
