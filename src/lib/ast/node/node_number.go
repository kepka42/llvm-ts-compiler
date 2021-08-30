package node

type Number struct {
	value float64
}

func (*Number) Type() Type {
	return TypeNumber
}

func NewNodeNumber(value float64) Number {
	return Number{value: value}
}
