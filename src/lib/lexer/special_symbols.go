package lexer

const (
	symbolWhiteSpace         = ' '
	symbolNewLine            = '\n'
	symbolBracketOpen        = '('
	symbolBracketClose       = ')'
	symbolBracketSquareOpen  = '['
	symbolBracketSquareClose = ']'
	symbolBraceOpen          = '{'
	symbolBraceClose         = '}'
	symbolPlus               = '+'
	symbolMinus              = '-'
	symbolEqual              = '='
	symbolStar               = '*'
	symbolComma              = ','
	symbolDevide             = '/'
	symbolSemicolon          = ';'
	symbolColon              = ':'
	symbolExclamationMark    = '!'
	symbolQuestionMark       = '?'
)

func getAllSpecialSymbols() []byte {
	return []byte {
		symbolWhiteSpace,
		symbolNewLine,
		symbolBracketOpen,
		symbolBracketClose,
		symbolBracketSquareOpen,
		symbolBracketSquareClose,
		symbolBraceOpen,
		symbolBraceClose,
		symbolPlus,
		symbolMinus,
		symbolEqual,
		symbolStar,
		symbolComma,
		symbolDevide,
		symbolSemicolon,
		symbolColon,
		symbolExclamationMark,
		symbolQuestionMark,
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
	return symbol == symbolWhiteSpace
}