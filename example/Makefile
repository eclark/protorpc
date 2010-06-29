# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

all: server client

include $(GOROOT)/src/Make.$(GOARCH)

CLEANFILES += server client message.pb.go

#include $(GOROOT)/src/pkg/goprotobuf.googlecode.com/hg/Make.protobuf

# ugly hack to deal with whitespaces in $GOROOT
nullstring :=
space := $(nullstring) # a space at the end
QUOTED_GOROOT:=$(subst $(space),\ ,$(GOROOT))

include $(QUOTED_GOROOT)/src/Make.common

%.pb.go:    %.proto
	protoc --protorpc_out=. $<

server: server.$O
	$(QUOTED_GOBIN)/$(LD) -o $@ $@.$O

server.$O: server.go message.pb.go
	$(QUOTED_GOBIN)/$(GC) -o $@ $^

client: client.$O
	$(QUOTED_GOBIN)/$(LD) -o $@ $@.$O

client.$O: client.go message.pb.go
	$(QUOTED_GOBIN)/$(GC) -o $@ $^

