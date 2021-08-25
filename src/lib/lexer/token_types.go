package lexer

import "strconv"

type TokenType int32

const (
	TokenTypeEof TokenType = iota
	TokenTypeNumber
	TokenTypeMul
	TokenTypeAdd
	TokenTypeSub
	TokenTypeDiv
	TokenTypeMod
	TokenTypeEqual
	TokenTypeComma
	TokenTypeBracketOpen
	TokenTypeBracketClose
	TokenTypeBracketSquareOpen
	TokenTypeBracketSquareClose
	TokenTypeBraceOpen
	TokenTypeBraceClose
	TokenTypeSemicolon
	TokenTypeColon
	TokenTypeExclamation
	TokenTypeQuestion
	TokenTypeUndefined
)

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
		switch symbol {
		case symbolPlus:
			return TokenTypeAdd
		case symbolMinus:
			return TokenTypeSub
		case symbolStar:
			return TokenTypeMul
		case symbolDevide:
			return TokenTypeDiv
		case symbolMod:
			return TokenTypeMod
		case symbolEqual:
			return TokenTypeEqual
		case symbolBracketOpen:
			return TokenTypeBracketOpen
		case symbolBracketClose:
			return TokenTypeBracketClose
		case symbolComma:
			return TokenTypeComma
		case symbolBraceOpen:
			return TokenTypeBraceOpen
		case symbolBraceClose:
			return TokenTypeBraceClose
		case symbolBracketSquareOpen:
			return TokenTypeBracketSquareOpen
		case symbolBracketSquareClose:
			return TokenTypeBracketSquareClose
		case symbolSemicolon:
			return TokenTypeSemicolon
		case symbolColon:
			return TokenTypeColon
		case symbolExclamationMark:
			return TokenTypeExclamation
		case symbolQuestionMark:
			return TokenTypeQuestion
		}
	}

	return TokenTypeUndefined
}
