package main

import (
//        desc "goprotobuf.googlecode.com/hg/compiler/descriptor"
	//"goprotobuf.googlecode.com/hg/proto"
	. "goprotobuf.googlecode.com/hg/compiler/generator"
	"log"
)

func init() {
	rpc := new(RpcPlugin)
	RegisterPlugin(rpc)
}

type RpcPlugin struct {
	*Generator
}

func (*RpcPlugin) Name() string {
	return "protorpc"
}

func (g *RpcPlugin) Init(ng *Generator) {
	g.Generator = ng
}

func (g *RpcPlugin) Generate(file *FileDescriptor) {
	g.P()
	g.P("// protorpc code")

	for _, sd := range file.Service {
		serviceName := *sd.Name
		if serviceName == "" {
			log.Stderr("no service name")
			continue
		}

		// build the interface
		g.P("type ", serviceName, " interface {")
		g.In()
		for _, m := range sd.Method {
			name := *m.Name
			input_type := CamelCaseSlice(g.ObjectNamed(*m.InputType).TypeName())
			output_type := CamelCaseSlice(g.ObjectNamed(*m.OutputType).TypeName())

			g.P(name, "(*", input_type, ", *", output_type, ") os.Error")

		}
		g.Out()
		g.P("}")

		// build server registration helper
		g.P("func Register", serviceName, "(s ", serviceName, ") os.Error {")
		g.In()
		g.P("return rpc.Register(s)")
		g.Out()
		g.P("}")

		// build the concrete client
		g.P("type ", serviceName, "Client struct {")
		g.In()
		g.P("*rpc.Client")
		g.P("remoteName string")
		g.Out()
		g.P("}")

		// client constructor
		g.P("func New", serviceName, "Client(rname, net, laddr, raddr string) (csc *", serviceName, "Client, err os.Error) {")
		g.In()
		g.P("client, err := protorpc.Dial(net, laddr, raddr)")
		g.P("if err != nil {")
		g.In()
		g.P("return")
		g.Out()
		g.P("}")
		g.P("csc = new(", serviceName, "Client)")
		g.P("csc.Client = client")
		g.P("csc.remoteName = rname")
		g.P("return")
		g.Out()
		g.P("}")

		// build methods on client
                for _, m := range sd.Method {
                        name := *m.Name
                        input_type := CamelCaseSlice(g.ObjectNamed(*m.InputType).TypeName())
                        output_type := CamelCaseSlice(g.ObjectNamed(*m.OutputType).TypeName())

			g.P("func (self *", serviceName, "Client) ", name, "(request *", input_type, ", response *", output_type, ") os.Error {")
			g.In()
			g.P("return self.Call(self.remoteName + ",Quote("." + name),", request, response)")
			g.Out()
			g.P("}")
                }

	}
}

func (g *RpcPlugin) GenerateImports(file *FileDescriptor) {
	g.P()
	g.P("// protorpc imports")
	g.P("import ", Quote("os"))
	g.P("import ", Quote("rpc"))
	g.P("import ", Quote("local/protorpc"))
}


