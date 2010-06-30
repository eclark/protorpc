// Copyright 2010 Eric Clark. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"log"
	"net"
	"local/protorpc"
)

type MyChatService int

func (e *MyChatService) Chat(req *ChatRequest, resp *ChatResponse) os.Error {
	println(*req.Line);

	v := int32(1);
	resp.Ok = &v

	return nil
}

func main() {
	echo := new(MyChatService)
	RegisterChatService(echo)

	l, e := net.Listen("tcp",":1234")
	if e != nil {
		log.Exit("listen error:",e)
	}

	protorpc.Serve(l)
}
