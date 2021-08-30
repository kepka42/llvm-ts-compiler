package node

import (
	"fmt"
	"github.com/kepka42/llvm-ts-compiler/src/lib/lexer"
	"strconv"
)

type Number struct {
	value float64
}

func (*Number) Type() Type {
	return TypeNumber
}

func NewNodeNumber(value float64) Number {
	return Number{value: value}
}

func ParseNodeNumber(buffer lexer.TokenBuffer) (Base, error) {
	// todo: add custom error system
	token := buffer.NextToken()
	if token.Type != lexer.TokenTypeNumber {
		return nil, fmt.Errorf("error_token_type")
	}

	val, err := strconv.ParseFloat(token.Data, 64)
	if err != nil {
		return nil, fmt.Errorf("error_value_type")
	}

	node := NewNodeNumber(val)
	return &node, nil
}
