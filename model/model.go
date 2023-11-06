package model

type Model struct {
	NameSpace string
	Clients []Client
}

type Client struct {
	Name string
	SendMethods []SendMethod
}

type SendMethod struct {
	OutputType string
	InputType string
	Name string
	TypeUrl string
}