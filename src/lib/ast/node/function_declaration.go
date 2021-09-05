package node

type FunctionDeclaration struct {
	params []Identifier
	body   BlockStatement
}

func (*FunctionDeclaration) Type() Type {
	return TypeFunctionDeclaration
}

func NewFunctionDeclaration(params []Identifier, body BlockStatement) FunctionDeclaration {
	return FunctionDeclaration{
		params: params,
		body:   body,
	}
}
