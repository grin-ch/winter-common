package socket

type Sender interface {
	Send(pack []byte) error
}

type Receiver interface {
	Recv() ([]byte, error)
}

type Conn interface {
	Protocol() string
	Sender
	Receiver
	Close() error
	Ip() string
}
