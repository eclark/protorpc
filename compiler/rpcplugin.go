package main

import (
//        desc "goprotobuf.googlecode.com/hg/compiler/descriptor"
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
	log.Stderr("generate")
	g.P()
	g.P("// protorpc interface")

	for _, sd := range file.Service {
		log.Stderr(*sd.Name)
		for _, m := range sd.Method {
			name := *m.Name
			input_type := *m.InputType
			output_type := *m.OutputType
			log.Stderr(name, input_type, output_type)
		}
	}
}

func (g *RpcPlugin) GenerateImports(file *FileDescriptor) {
	log.Stderr("generate-imports")

	g.P("// this is where i will put imports")
}
