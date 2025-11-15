package p2p

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	opts := TCPTransportOpts{
		ListenAddr: ":4000",
		HandShakeFunc: NOPHandshake,
		Decoder: DefaultDecoder{},
		OnPeer: NOPOnPeer,
	}
	tr := NewTCPTransport(opts)
 
	assert.Equal(t, tr.ListenAddr, opts.ListenAddr)
	assert.Nil(t, tr.ListenAndAccept())
	select{}
}