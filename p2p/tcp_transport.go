package p2p

import (
	"fmt"
	"net"
	"sync"
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

type TCPTransportOpts struct {
	ListenAddr    string
	HandShakeFunc HandShakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
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
			fmt.Printf("[Error] Unable to accept connection: %s\n", err)
		}
		go t.handleConnection(conn)
	}
}

type TempMessage struct{}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, false)
	fmt.Printf("[Debug] New connection peer: %+v\n", peer)
	fmt.Printf("[Debug] Recieved new connection: %+v\n", conn)

	if err := t.HandShakeFunc(conn); err != nil {
		conn.Close()
		fmt.Printf("[Error] TCP handshake error: %+v", err)
		return
	}

	// Read the message in loop
	msg := &TempMessage{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("[Error] TCP decode error: %+v", err)
			continue
		}
	}
}
