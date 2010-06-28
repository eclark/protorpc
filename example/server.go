package main

import (
	"os"
	"rpc"
	"log"
	"net"
	"local/protorpc"
)

type Echo int

func (e *Echo) Echo(req *ChatRequest, resp *ChatResponse) os.Error {
	println(*req.Line);

	v := int32(1);
	resp.Ok = &v

	return nil
}


func main() {
	echo := new(Echo)
	rpc.Register(echo)


	l, e := net.Listen("tcp",":1234")
	if e != nil {
		log.Exit("listen error:",e)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Stderr("accept error:",err)
		}

		go protorpc.ServeConn(conn)
	}
}
