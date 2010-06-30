// Copyright 2010 Eric Clark. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net"
	"github.com/eclark/protorpc"
)

func main() {
	calc := new(MyCalcService)
	RegisterCalcService(calc)

	l, e := net.Listen("tcp",":1234")
	if e != nil {
		log.Exit("listen error:",e)
	}

	protorpc.Serve(l)
}
