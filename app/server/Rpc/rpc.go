package Rpc

import (
	App "gin-api/app/rpc"
	"github.com/spiral/goridge"
	"net"
	"net/rpc"
)

func Run(port string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	rpc.Register(new(App.App))

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeCodec(goridge.NewCodec(conn))
	}
}
