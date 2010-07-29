// Copyright 2010 Eric Clark. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protorpc

import (
	"rpc"
	"io"
	"os"
	"net"
	"goprotobuf.googlecode.com/hg/proto"
)

type clientCodec struct {
	c    io.ReadWriteCloser
	req  *bufferPair
	resp *bufferPair
}

func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec {
	req := &bufferPair{proto.NewBuffer(nil), proto.NewBuffer(nil)}
	resp := &bufferPair{proto.NewBuffer(nil), proto.NewBuffer(nil)}

	return &clientCodec{conn, req, resp}
}

func (c *clientCodec) WriteRequest(r *rpc.Request, message interface{}) (err os.Error) {
	c.req.header.Reset()
	c.req.body.Reset()

	h := new(Header)
	h.Seq = &r.Seq
	h.ServiceMethod = &r.ServiceMethod

	err = c.req.header.Marshal(h)
	if err != nil {
		return
	}

	err = c.req.body.Marshal(message)
	if err != nil {
		return
	}

	_, err = c.c.Write(encodeLen(len(c.req.header.Bytes())))
	if err != nil {
		return
	}

	_, err = c.c.Write(c.req.header.Bytes())
	if err != nil {
		return
	}

	_, err = c.c.Write(encodeLen(len(c.req.body.Bytes())))
	if err != nil {
		return
	}

	_, err = c.c.Write(c.req.body.Bytes())
	if err != nil {
		return
	}

	return
}

func (c *clientCodec) ReadResponseHeader(r *rpc.Response) (err os.Error) {
	c.resp.header.Reset()

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

	c.resp.header.SetBuf(pbuf)

	h := new(Header)
	err = c.resp.header.Unmarshal(h)
	if err != nil {
		return
	}

	r.Seq = *h.Seq
	r.ServiceMethod = *h.ServiceMethod
	r.Error = *h.Error

	return nil
}

func (c *clientCodec) ReadResponseBody(message interface{}) (err os.Error) {
	c.resp.body.Reset()

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

	c.resp.body.SetBuf(pbuf)

	err = c.resp.body.Unmarshal(message)
	if err != nil {
		return
	}

	return
}

func (c *clientCodec) Close() os.Error {
	return c.c.Close()
}

func Dial(netw, laddr, raddr string) (*rpc.Client, os.Error) {
	conn, err := net.Dial(netw, laddr, raddr)
	if err != nil {
		return nil, err
	}

	codec := NewClientCodec(conn)
	client := rpc.NewClientWithCodec(codec)

	return client, nil
}
