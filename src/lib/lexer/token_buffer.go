package lexer

type TokenBuffer struct {
	tokens []Token
	cur    int
}

func newTokenBuffer(tokens []Token) TokenBuffer {
	return TokenBuffer{
		tokens: tokens,
		cur:    0,
	}
}

func (t *TokenBuffer) NextToken() *Token {
	if t.cur >= len(t.tokens) {
		return nil
	}

	token := t.tokens[t.cur]
	t.cur = t.cur + 1
	return &token
}
