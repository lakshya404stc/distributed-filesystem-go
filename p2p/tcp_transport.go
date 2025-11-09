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
		conn : 	  conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener 	  net.Listener
	handShake     HandShakeFunc
	decoder		  Decoder

	mu 			  sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport (listenAddr string) *TCPTransport {
	return &TCPTransport{
		handShake: NOPHandshake,
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept()error { 
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
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

type TempMessage struct {}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, false)
	fmt.Printf("[Debug] New connection peer: %+v\n", peer)
	fmt.Printf("[Debug] Recieved new connection: %+v\n", conn)

	if err := t.handShake(conn); err != nil {
		conn.Close()
		return
	}

	// Read the message loop 
	msg := &TempMessage{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("[Error] TCP error: %+v", err)
			continue
		}
	}
}