package main

import (
	"fmt"
	"log"

	"github.com/lakshya404stc/distributed-filesystem-go/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":4000",
		HandShakeFunc: p2p.NOPHandshake,
		Decoder: p2p.DefaultDecoder{},
		OnPeer: p2p.NOPOnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func (){
		for {
			msg := <-tr.Consume()
			fmt.Printf("[DEBUG] New message recieved from %v: %v", msg.From, msg.Payload)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
