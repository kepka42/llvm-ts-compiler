package node

type BlockStatement struct {
	body []Base
}

func (*BlockStatement) Type() Type {
	return TypeBlockStatement
}
