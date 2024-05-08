package main

import (
	"bytes"
	_ "embed"
	"flag"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/DecentralCardGame/protoc-gen-cosmos-csharp/model"
	"google.golang.org/protobuf/compiler/protogen"
)

//go:embed templates/Client.pb.cs.tmpl
var clientTmpl string

var suffixFlag string

func main() {
	log.SetOutput(os.Stderr)
	flag.StringVar(&suffixFlag, "suffix", "Client.pb.cs", "file suffixFlag")

	protogen.Options{ParamFunc: flag.Set}.Run(func(gen *protogen.Plugin) error {
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

	path := strings.Replace(m.NameSpace.String(), ".", "/", -1)
	filename := path + "/Tx" + suffixFlag
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
