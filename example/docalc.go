// Copyright 2010 Eric Clark. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"goprotobuf.googlecode.com/hg/proto"
	"log"
)

func doCalc(calc CalcService) {
	crq := new(CalcRequest)
	crs := new(CalcResponse)

	// add
	crq.A = proto.Int64(61)
	crq.B = proto.Int64(35)

	err := calc.Add(crq, crs)
	if err != nil {
		log.Println("add error:", err)
	}

	log.Println("add result:", proto.GetInt64(crs.Result))

	crq.Reset()
	crs.Reset()

	// subtract
	crq.A = proto.Int64(61)
	crq.B = proto.Int64(35)

	err = calc.Subtract(crq, crs)
	if err != nil {
		log.Println("subtract error:", err)
	}

	log.Println("subtract result:", proto.GetInt64(crs.Result))

	crq.Reset()
	crs.Reset()

	// multiply
	crq.A = proto.Int64(9)
	crq.B = proto.Int64(11)

	err = calc.Multiply(crq, crs)
	if err != nil {
		log.Println("multiply error:", err)
	}

	log.Println("multiply result:", proto.GetInt64(crs.Result))

	crq.Reset()
	crs.Reset()

	// divide
	crq.A = proto.Int64(20)
	crq.B = proto.Int64(8)

	err = calc.Divide(crq, crs)
	if err != nil {
		log.Println("divide error:", err)
	}

	log.Println("divide result:", proto.GetInt64(crs.Result))
	log.Println("divide remainder:", proto.GetInt64(crs.Remainder))

	crq.Reset()
	crs.Reset()

}
