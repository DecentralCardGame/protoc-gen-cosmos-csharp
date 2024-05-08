package model

import (
	"github.com/DecentralCardGame/protoc-gen-cosmos-csharp/descriptor"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
)

type Model struct {
	NameSpace descriptor.Desriptor
	Clients   []Client
	Source    string
}

type Client struct {
	Name        string
	SendMethods []SendMethod
}

type SendMethod struct {
	OutputType descriptor.Desriptor
	InputType  descriptor.Desriptor
	Name       string
	TypeUrl    string
}

func NewService(service *descriptorpb.ServiceDescriptorProto, nameSpace descriptor.Desriptor) Client {
	client := Client{
		Name: *service.Name,
	}

	for _, msg := range service.Method {
		client.SendMethods = append(client.SendMethods, NewMethod(msg, nameSpace))
	}
	return client
}

func NewMethod(msg *descriptorpb.MethodDescriptorProto, nameSpace descriptor.Desriptor) SendMethod {
	return SendMethod{
		OutputType: descriptor.FromTypeUrl(*msg.OutputType).CutNameSpace(nameSpace),
		InputType:  descriptor.FromTypeUrl(*msg.InputType).CutNameSpace(nameSpace),
		Name:       *msg.Name,
		TypeUrl:    strings.Trim(*msg.InputType, "."),
	}
}

func NewModel(file *protogen.File) *Model {
	nameSpace := descriptor.FromTypeUrl("." + *file.Proto.Package)

	m := Model{
		NameSpace: nameSpace,
		Source:    *file.Proto.Name,
	}

	if len(file.Proto.Service) == 0 {
		return nil
	}

	for _, service := range file.Proto.Service {
		m.Clients = append(m.Clients, NewService(service, nameSpace))
	}

	return &m
}
