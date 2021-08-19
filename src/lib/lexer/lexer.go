package lexer

type Builder struct {
	input []byte
	tokens []Token
}

const (
	whiteSpace = ' '
	newLine    = '\n'
)

func (b* Builder) SetInput(input []byte) {
	b.input = input
}

func (b* Builder) Run() ([]Token, error) {
	tokens := make([]Token, 0)

	wordBytes := make([]byte, 0)
	for _, symbol := range b.input {
		if symbol == whiteSpace || symbol == newLine {
			tokens = append(tokens, NewToken(string(symbol)))
			tokens = append(tokens, NewToken(string(wordBytes)))

			wordBytes = make([]byte, 0)
			continue
		}
		wordBytes = append(wordBytes, symbol)
	}

	if len(wordBytes) != 0 {
		tokens = append(tokens, NewToken(string(wordBytes)))
	}

	b.tokens = tokens
	return tokens, nil
}
