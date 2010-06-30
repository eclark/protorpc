// Copyright 2010 Eric Clark. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protorpc

import (
	"rpc"
	"io"
	"net"
	"os"
	"goprotobuf.googlecode.com/hg/proto"
)

type serverCodec struct {
	c    io.ReadWriteCloser
	req  *bufferPair
	resp *bufferPair
}

func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec {
	req := &bufferPair{proto.NewBuffer(nil), proto.NewBuffer(nil)}
	resp := &bufferPair{proto.NewBuffer(nil), proto.NewBuffer(nil)}

	return &serverCodec{conn, req, resp}
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) (err os.Error) {
	c.req.header.Reset()

	lbuf := make([]byte, lenSize)
	_, err = io.ReadFull(c.c, lbuf)
	if err != nil {
		return
	}

	pbuf := make([]byte, decodeLen(lbuf))
	_, err = io.ReadFull(c.c, pbuf)
	if err != nil {
		return
	}

	c.req.header.SetBuf(pbuf)

	h := NewHeader()
	err = c.req.header.Unmarshal(h)
	if err != nil {
		return
	}

	r.Seq = *h.Seq
	r.ServiceMethod = *h.ServiceMethod

	return
}

func (c *serverCodec) ReadRequestBody(message interface{}) (err os.Error) {
	c.req.body.Reset()

	lbuf := make([]byte, lenSize)
	_, err = io.ReadFull(c.c, lbuf)
	if err != nil {
		return
	}

	pbuf := make([]byte, decodeLen(lbuf))
	_, err = io.ReadFull(c.c, pbuf)
	if err != nil {
		return
	}

	c.req.body.SetBuf(pbuf)

	err = c.req.body.Unmarshal(message)
	if err != nil {
		return
	}

	return
}

func (c *serverCodec) WriteResponse(r *rpc.Response, message interface{}) (err os.Error) {
	c.resp.header.Reset()
	c.resp.body.Reset()

	h := NewHeader()
	h.Seq = &r.Seq
	h.ServiceMethod = &r.ServiceMethod
	h.Error = &r.Error

	err = c.resp.header.Marshal(h)
	if err != nil {
		return
	}

	_, err = c.c.Write(encodeLen(len(c.resp.header.Bytes())))
	if err != nil {
		return
	}

	_, err = c.c.Write(c.resp.header.Bytes())
	if err != nil {
		return
	}

	if _, ok := message.(rpc.InvalidRequest); !ok {
		err = c.resp.body.Marshal(message)
		if err != nil {
			return
		}

		_, err = c.c.Write(encodeLen(len(c.resp.body.Bytes())))
		if err != nil {
			return
		}

		_, err = c.c.Write(c.resp.body.Bytes())
		if err != nil {
			return
		}
	}

	return
}

func (c *serverCodec) Close() os.Error {
	return c.c.Close()
}

func ServeConn(conn io.ReadWriteCloser) {
	rpc.ServeCodec(NewServerCodec(conn))
}

func Serve(l net.Listener) os.Error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go ServeConn(conn)
	}
	return nil
}
