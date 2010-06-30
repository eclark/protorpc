// Copyright 2010 Eric Clark. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protorpc

import "goprotobuf.googlecode.com/hg/proto"

type bufferPair struct {
	header *proto.Buffer
	body   *proto.Buffer
}

const lenSize = 4

func encodeLen(length int) []byte {
	b := make([]byte, 4)
	b[0] = byte(length >> 24)
	b[1] = byte(length >> 16)
	b[2] = byte(length >> 8)
	b[3] = byte(length)

	return b
}

func decodeLen(b []byte) int {
	return int(int32(b[0])<<24 + int32(b[1])<<16 + int32(b[2])<<8 + int32(b[3]))
}
