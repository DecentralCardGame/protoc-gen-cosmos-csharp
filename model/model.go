package model

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
)

type Model struct {
	NameSpace string
	Clients   []Client
	Source    string
}

type Client struct {
	Name        string
	SendMethods []SendMethod
}

type SendMethod struct {
	OutputType string
	InputType  string
	Name       string
	TypeUrl    string
}

func parsePathName(path string) string {
	paths := strings.Split(path, ".")

	var newPaths []string
	for _, p := range paths[1:] {
		newPaths = append(newPaths, strings.Title(p))
	}

	return strings.Join(newPaths, ".")
}

func NewService(service *descriptorpb.ServiceDescriptorProto) Client {
	client := Client{
		Name: *service.Name,
	}

	for _, msg := range service.Method {
		client.SendMethods = append(client.SendMethods, NewMethod(msg))
	}
	return client
}

func NewMethod(msg *descriptorpb.MethodDescriptorProto) SendMethod {
	return SendMethod{
		OutputType: parsePathName(*msg.OutputType),
		InputType:  parsePathName(*msg.InputType),
		Name:       *msg.Name,
		TypeUrl:    strings.Trim(*msg.InputType, "."),
	}
}

func NewModel(file *protogen.File) *Model {
	m := Model{
		NameSpace: parsePathName("." + *file.Proto.Package),
		Source:    *file.Proto.Name,
	}

	if len(file.Proto.Service) == 0 {
		return nil
	}

	for _, service := range file.Proto.Service {
		m.Clients = append(m.Clients, NewService(service))
	}

	return &m
}
