package lexer

type Builder struct {
	input []byte
}

func (b *Builder) SetInput(input []byte) {
	b.input = input
}

func (b *Builder) Run() (*TokenBuffer, error) {
	tokens := make([]Token, 0)

	wordBytes := make([]byte, 0)
	for _, symbol := range b.input {
		if !isSpecialSymbol(symbol) {
			wordBytes = append(wordBytes, symbol)
			continue
		}

		data := string(wordBytes)
		if data != "" {
			tokens = append(tokens, NewToken(data))
		}

		if isSpecialSymbol(symbol) && !isWhiteSpace(symbol) {
			tokens = append(tokens, NewToken(string(symbol)))
		}

		wordBytes = make([]byte, 0)
	}

	if len(wordBytes) != 0 {
		tokens = append(tokens, NewToken(string(wordBytes)))
	}

	buffer := newTokenBuffer(tokens)
	return &buffer, nil
}
