package node

type Variable struct {
	name string
}

func (*Variable) Type() Type {
	return TypeVariable
}

func NewNodeVariable(name string) Variable {
	return Variable{name: name}
}
