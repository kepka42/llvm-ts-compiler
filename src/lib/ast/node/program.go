package node

type Program struct {
	body []Base
}

func (*Program) Type() Type {
	return TypeProgram
}

func (p *Program) Body() []Base {
	return p.body
}

func NewNodeProgram(body []Base) Program {
	return Program{
		body: body,
	}
}
