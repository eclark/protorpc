// Copyright 2010 Eric Clark. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

func main() {
	calc, err := NewCalcServiceClient("MyCalcService", "tcp", "", "127.0.0.1:1234")
	if err != nil {
		log.Exit("cant setup calc service:", err)
	}

	doCalc(calc)
}
