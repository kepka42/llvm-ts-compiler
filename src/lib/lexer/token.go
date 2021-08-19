package lexer

type Token struct {
	Data string
}

func NewToken(data string) Token {
	token := Token{
		Data: data,
	}

	return token
}