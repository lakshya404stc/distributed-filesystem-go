package p2p

import "fmt"

type HandShakeFunc func(Peer) error
func NOPHandshake (Peer) error {return nil}

type OnPeerFunc func(Peer) error
func NOPOnPeer (Peer) error {
	fmt.Printf("[DEBUG] A new peer is available")
	return nil
}