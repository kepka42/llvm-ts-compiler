package lexer

import "strconv"

type TokenType int32

const (
	TokenTypeEof TokenType = iota
	TokenTypeNumber
	TokenTypeBinaryOp
	TokenTypeVariable
	TokenTypeIdentifier
)

func isBinaryOp(symbol rune) bool {
	binaryOps := []rune{
		'+', '-', '*', '/', '=', '%',
	}

	for _, v := range binaryOps {
		if v == symbol {
			return true
		}
	}

	return false
}

func isVariable(data string) bool {
	return data == "var"
}

func getTokenType(data string) TokenType {
	if data == "\n" {
		return TokenTypeEof
	}

	_, err := strconv.Atoi(data)
	if err == nil {
		return TokenTypeNumber
	}

	_, err = strconv.ParseFloat(data, 64)
	if err == nil {
		return TokenTypeNumber
	}

	if len(data) == 1 {
		symbol := data[0]

		if isBinaryOp(rune(symbol)) {
			return TokenTypeBinaryOp
		}
	}

	if isVariable(data) {
		return TokenTypeVariable
	}

	return TokenTypeIdentifier
}
