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

func (t *TokenBuffer) Current() *Token {
	if t.cur > len(t.tokens) - 1 {
		return nil
	}

	current := t.tokens[t.cur]
	return &current
}

func (t *TokenBuffer) Next() {
	t.cur = t.cur + 1
}

func (t *TokenBuffer) End() bool {
	return t.cur >= len(t.tokens) - 1
}
