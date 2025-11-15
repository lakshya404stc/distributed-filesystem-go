package p2p

// peer is anything that represents a remote node
type Peer interface {
	Close() error
}

// transport is anything that handles the communication between nodes in the network
// using tcp, udp or websockets etc
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}