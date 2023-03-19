package websocket

type Handler interface {
	// NOTE: it's blocking method
	Connect() error

	// Close connection
	Close() error
}
