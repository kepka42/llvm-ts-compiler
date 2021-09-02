package node

type Program struct{
	body []Base
}

func (*Program) Type() Type {
	return TypeProgram
}

func NewNodeProgram(body []Base) Program {
	return Program{
		body: body,
	}
}
