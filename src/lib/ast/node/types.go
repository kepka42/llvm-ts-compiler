package node

type Type int32

const (
	TypeProgram Type = iota
	TypeVariableDeclaration
	TypeFunctionDeclaration
	TypeNumber
	TypeBinaryOp
	TypeIdentifier
	TypeBlockStatement
)
