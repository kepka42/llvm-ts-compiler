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
		// todo: refactor it
		if isSpecialSymbol(symbol) || isWhiteSpace(symbol) {
			data := string(wordBytes)
			if data == "" {
				continue
			}

			tokens = append(tokens, NewToken(data))
			if isSpecialSymbol(symbol) {
				tokens = append(tokens, NewToken(string(symbol)))
			}

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

func isWhiteSpace(symbol byte) bool {
	return symbol == whiteSpace
}

func isSpecialSymbol(symbol byte) bool {
	return symbol == newLine
}
