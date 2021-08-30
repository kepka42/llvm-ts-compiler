package node

type BinaryOp struct {
	op rune
	lhs Base
	rhs Base
}

func (*BinaryOp) Type() Type {
	return TypeBinaryOp
}

func NewNodeBinaryOp(op rune, lhs, rhs Base) BinaryOp {
	return BinaryOp{
		op: op,
		lhs: lhs,
		rhs: rhs,
	}
}
