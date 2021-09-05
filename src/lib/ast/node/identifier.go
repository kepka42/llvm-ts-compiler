package node

type Identifier struct {
	name string
}

func (*Identifier) Type() Type {
	return TypeIdentifier
}

func NewIdentifier(name string) Identifier {
	return Identifier{
		name: name,
	}
}
