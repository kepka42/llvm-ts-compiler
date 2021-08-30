package node

type Program struct{}

func (*Program) Type() Type {
	return TypeProgram
}

func NewNodeProgram() Program {
	return Program{}
}
