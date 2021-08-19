package lexer

import "strconv"

type Token struct {
	Data string
	Type TokenType
}

func NewToken(data string) Token {
	token := Token{
		Data: data,
		Type: getTokenType(data),
	}

	return token
}

func getTokenType(data string) TokenType {
	if data == " " {
		return TokenTypeEof
	}

	_, err := strconv.Atoi(data)
	if err == nil {
		return TokenTypeNumber
	}

	return TokenTypeExtern
}