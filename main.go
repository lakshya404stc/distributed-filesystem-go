package main

import (
	"fmt"
	"log"

	"github.com/lakshya404stc/distributed-filesystem-go/p2p"
)

func main() {
	fmt.Println("Hello!! from go main package")

	listenAddr := ":4000"
	tr := p2p.NewTCPTransport(listenAddr)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select{}
}