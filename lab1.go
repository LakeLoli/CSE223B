package triblab

import (
	"trib"
	"net/rpc"
	"net/http"
	"log"
	"net"	
)

// Creates an RPC client that connects to addr.
func NewClient(addr string) trib.Storage {
	return &client{addr: addr}
}

// Serve as a backend based on the given configuration
func ServeBack(b *trib.BackConfig) error {
	laddr := b.Addr
	//storage := b.Store
	
	rpc.Register(b.Store)		
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", laddr)
	log.Printf("Listen port has been set up!\n")
	if e!= nil {
		if b.Ready != nil {
			b.Ready <- false
		}
		log.Fatal("listen error: ", e)
		return e;
	}
	if b.Ready != nil {
		log.Printf("")
		b.Ready <- true
	}
	
	return http.Serve(l, nil)
}
