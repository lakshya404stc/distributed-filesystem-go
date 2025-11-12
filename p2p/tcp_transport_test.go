package p2p

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	ListenAddr := ":4000"
	tr := NewTCPTransport(TCPTransportOpts{ListenAddr})

	assert.Equal(t, tr.listenAddr, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())

	select{}
}