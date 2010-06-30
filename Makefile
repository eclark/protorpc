# Copyright 2010 Eric Clark. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=github.com/eclark/protorpc

GOFILES=\
    client.go\
    server.go\
    common.go\
    header.pb.go\

CLEANFILES += header.pb.go
INSTALLFILES += compiler.make

include $(GOROOT)/src/Make.pkg
include $(GOROOT)/src/pkg/goprotobuf.googlecode.com/hg/Make.protobuf

dist: header.pb.go
	cp header.pb.go header.dist.go

