package lexer

type TokenType int32

const (
	TokenTypeEof = iota
	TokenTypeNumber
	TokenTypeExtern
)
