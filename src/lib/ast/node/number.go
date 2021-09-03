package node

type Number struct {
	value float64
}

func (*Number) Type() Type {
	return TypeNumber
}

func (n *Number) Value() float64 {
	return n.value
}

func NewNodeNumber(value float64) Number {
	return Number{value: value}
}
