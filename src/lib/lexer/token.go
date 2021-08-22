package lexer

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
