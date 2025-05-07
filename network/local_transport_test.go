package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tr1 := NewLocalTransport("node1")
	tr2 := NewLocalTransport("node2")

	tr1.Connect(tr2)
	tr2.Connect(tr1)
	assert.Equal(t, tr1.peers[tr2.addr], tr2)
	assert.Equal(t, tr2.peers[tr1.addr], tr1)
}

// assert. Equal(t, 1, 1)

func TestSendMessage(t *testing.T) {
	tr1 := NewLocalTransport("node1")
	tr2 := NewLocalTransport("node2")

	tr1.Connect(tr2)
	tr2.Connect(tr1)

	msg := []byte("hello blockchain world")
	assert.Nil(t, tr1.SendMessage(tr2.addr, msg))

	rpc := <-tr2.Consume()
	assert.Equal(t, rpc.From, tr1.addr)
	assert.Equal(t, rpc.Payload, msg)
}
