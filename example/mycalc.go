// Copyright 2010 Eric Clark. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"goprotobuf.googlecode.com/hg/proto"
	"os"
	"fmt"
)

type MyCalcService int

func (*MyCalcService) Add(req *CalcRequest, resp *CalcResponse) os.Error {
	resp.Result = proto.Int64(proto.GetInt64(req.A) + proto.GetInt64(req.B))
	return nil
}

func (*MyCalcService) Subtract(req *CalcRequest, resp *CalcResponse) os.Error {
    resp.Result = proto.Int64(proto.GetInt64(req.A) - proto.GetInt64(req.B))
	return nil
}

func (*MyCalcService) Multiply(req *CalcRequest, resp *CalcResponse) os.Error {
    resp.Result = proto.Int64(proto.GetInt64(req.A) * proto.GetInt64(req.B))
	return nil
}

func (*MyCalcService) Divide(req *CalcRequest, resp *CalcResponse) (err os.Error) {
	resp.Result = proto.Int64(0)
	defer func() {
		if x := recover(); x != nil {
			if ex, ok := x.(os.Error); ok {
				err = ex
			} else {
				err = os.ErrorString(fmt.Sprint(x))
			}
		}
	}()
    resp.Result = proto.Int64(proto.GetInt64(req.A) / proto.GetInt64(req.B))
    resp.Remainder = proto.Int64(proto.GetInt64(req.A) % proto.GetInt64(req.B))
	return
}
