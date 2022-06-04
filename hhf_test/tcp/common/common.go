package common

import (
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
	"time"

	"github.com/panjf2000/gnet"
)

type EchoServer struct {
	*gnet.EventServer
	Pool *goroutine.Pool
}

func (es *EchoServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	data := append([]byte{}, frame...)

	// Use ants pool to unblock the event-loop.
	_ = es.Pool.Submit(func() {
		time.Sleep(1 * time.Second)
		c.AsyncWrite(data)
	})

	return
}
