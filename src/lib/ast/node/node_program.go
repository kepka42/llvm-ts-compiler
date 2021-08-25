package node

type Program struct{}

func (*Program) Type() Type {
	return TypeProgram
}
