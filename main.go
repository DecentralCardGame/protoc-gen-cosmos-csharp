package main

import (
	"bytes"
	_ "embed"
	"strings"
	"text/template"

	"github.com/DecentralCardgame/protoc-gen-cosmosCsharp/model"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

//go:embed templates/Client.pb.cs.tmpl
var clientTmpl string

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

func parsePathName(path string) string {
	paths := strings.Split(path, ".")

	var newPaths []string

	for _, p := range paths[1:] {
		newPaths = append(newPaths, strings.Title(p))
	}

	return strings.Join(newPaths, ".")
}

func generateMethod(msg *descriptorpb.MethodDescriptorProto) model.SendMethod {
	return model.SendMethod{
		OutputType: parsePathName(*msg.OutputType),
		InputType:  parsePathName(*msg.InputType),
		Name:       *msg.Name,
		TypeUrl:    strings.Trim(*msg.InputType, "."),
	}
}

func generateService(service *descriptorpb.ServiceDescriptorProto) model.Client {
	client := model.Client{
		Name: *service.Name,
	}

	for _, msg := range service.Method {
		client.SendMethods = append(client.SendMethods, generateMethod(msg))
	}
	return client
}

func generateFile(gen *protogen.Plugin, file *protogen.File) {
	m := model.Model{
		NameSpace: parsePathName("." + *file.Proto.Package),
	}
	path := strings.Replace(m.NameSpace, ".", "/", -1)
	filename := path + "/TxClient.pb.cs"
	g := gen.NewGeneratedFile(filename, protogen.GoImportPath(path))
	if !strings.Contains(file.GeneratedFilenamePrefix, "tx") {
		g.Skip()
		return
	}

	for _, service := range file.Proto.Service {
		m.Clients = append(m.Clients, generateService(service))
	}

	if len(m.Clients) == 0 {
		g.Skip()
		return
	}

	tmpl, err := template.New("client").Parse(clientTmpl)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBufferString("")

	err = tmpl.Execute(buf, m)
	if err != nil {
		panic(err)
	}

	g.P(buf.String())
}
