// Copyright 2010 Eric Clark. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

func main() {

	chatter, err := NewChatServiceClient("MyChatService", "tcp", "", "127.0.0.1:1234")
	if err != nil {
		log.Exit("making chat service:", err)
	}

	line := "hi martha focker";

	crq := NewChatRequest()
	crq.Line = &line
	crs := NewChatResponse()

	err = chatter.Chat(crq,crs)
	if err != nil {
		log.Exit("chat error:", err)
	}

	log.Stderr("resp:", *crs.Ok)
}
