package node

type VariableDeclaration struct {
	name  string
	value Base
}

func (*VariableDeclaration) Type() Type {
	return TypeVariable
}

func (v *VariableDeclaration) Name() string {
	return v.name
}

func (v *VariableDeclaration) Value() Base {
	return v.value
}


func NewNodeVariableDeclaration(name string, value Base) VariableDeclaration {
	return VariableDeclaration{
		name:  name,
		value: value,
	}
}
