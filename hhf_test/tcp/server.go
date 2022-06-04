package main

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
	"log"
	"time"
)

type EchoServer struct {
	*gnet.EventServer
	Pool *goroutine.Pool
}

func (es *EchoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	data := append([]byte{}, frame...)
	fmt.Println(data)

	// Use ants pool to unblock the event-loop.
	_ = es.Pool.Submit(func() {
		time.Sleep(1 * time.Second)
		c.AsyncWrite(data)
	})

	return
}
func main() {
	p := goroutine.Default()
	defer p.Release()

	echo := &EchoServer{Pool: p}
	log.Fatal(gnet.Serve(echo, "tcp://:9000", gnet.WithMulticore(true)))
}
