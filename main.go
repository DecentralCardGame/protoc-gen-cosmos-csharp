package main

import (
	"bytes"
	_ "embed"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/DecentralCardgame/protoc-gen-cosmosCsharp/model"
	"google.golang.org/protobuf/compiler/protogen"
)

//go:embed templates/Client.pb.cs.tmpl
var clientTmpl string

func main() {
	log.SetOutput(os.Stderr)

	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			if err := generateFile(gen, f); err != nil {
				return err
			}
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File) error {
	m := model.NewModel(file)
	if m == nil {
		return nil
	}

	path := strings.Replace(m.NameSpace, ".", "/", -1)
	filename := path + "/TxClient.pb.cs"
	g := gen.NewGeneratedFile(filename, protogen.GoImportPath(path))
	if !strings.Contains(file.GeneratedFilenamePrefix, "tx") {
		g.Skip()
		return nil
	}

	tmpl, err := template.New("client").Parse(clientTmpl)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, m); err != nil {
		return err
	}

	g.P(buf.String())

	return nil
}
