package node

type VariableDeclaration struct {
	name  string
	value Base
}

func (*VariableDeclaration) Type() Type {
	return TypeVariable
}

func NewNodeVariableDeclaration(name string, value Base) VariableDeclaration {
	return VariableDeclaration{
		name:  name,
		value: value,
	}
}
