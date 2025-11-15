package p2p

import (
	"fmt"
	"net"
	// "sync"
)

type TCPPeer struct {
	conn net.Conn
	// True if connection is outbound and false if connection is inbound
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandShakeFunc HandShakeFunc
	Decoder       Decoder
	OnPeer        OnPeerFunc
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcch    chan RPC
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

// Consume method implements the Transport interface, which will return a read only channel
// for reading the incomming message revcieved from another peer in the network
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("[ERROR] Unable to accept connection: %s\n", err)
		}
		go t.handleConnection(conn)
	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	var err error

	peer := NewTCPPeer(conn, false)
	fmt.Printf("[DEBUG] New connection peer: %+v\n", peer)

	defer func() {
		fmt.Printf("[ERROR] Closing tcp connection: %v", err)
		conn.Close()
	}()

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}

	if err = t.HandShakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("[ERROR] TCP handshake error: %+v", err)
		return
	}

	// Read the message in loop
	rpc := RPC{}
	rpc.From = conn.RemoteAddr()
	for {
		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("[ERROR] TCP decode error: %+v", err)
			continue
		}
		fmt.Printf("[DEBUG] Message recieved: %+v\n", rpc)
		t.rpcch <- rpc
	}
}
