package lexer

func getAllSpecialSymbols() []byte {
	return []byte{
		' ', '\n', '(', ')', '[', ']', '{', '}', '+', '-',
		'%', '=', '*', ',', '/', ';', ':', '!', '?', '.',
	}
}

func isSpecialSymbol(symbol byte) bool {
	specialSymbols := getAllSpecialSymbols()
	for _, specialSymbol := range specialSymbols {
		if symbol == specialSymbol {
			return true
		}
	}
	return false
}

func isWhiteSpace(symbol byte) bool {
	return symbol == ' '
}
