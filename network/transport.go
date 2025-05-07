package network

type NetAddr string

type RPC struct {
	From NetAddr
	// To string
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
