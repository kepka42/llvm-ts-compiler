package node

type Type int32

const (
	TypeProgram Type = iota
	TypeVariable
	TypeNumber
	TypeBinaryOp
)
