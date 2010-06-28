package main

import (
	"rpc"
	"net"
	"log"
	"local/protorpc"
)

func main() {

	conn, err := net.Dial("tcp", "", "127.0.0.1:1234")
	if err != nil {
		log.Exit("dialing:", err)
	}

	codec := protorpc.NewClientCodec(conn)
	client := rpc.NewClientWithCodec(codec)

	line := "hi martha focker";

	crq := NewChatRequest()
	crq.Line = &line
	crs := NewChatResponse()

	err = client.Call("Echo.Echo", crq, crs)
	if err != nil {
		log.Exit("echo error:", err)
	}

	log.Stderr("resp:", *crs.Ok)
}
