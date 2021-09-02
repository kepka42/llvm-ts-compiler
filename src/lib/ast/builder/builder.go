package builder

import (
	"fmt"
	"github.com/kepka42/llvm-ts-compiler/src/lib/ast/node"
	"github.com/kepka42/llvm-ts-compiler/src/lib/lexer"
	"strconv"
)

func RootBuild(buffer *lexer.TokenBuffer) (node.Base, error) {
	body := make([]node.Base, 0)
	current := buffer.Current()

	for !buffer.End() {
		current = buffer.Current()
		if current.Type == lexer.TokenTypeEof {
			buffer.Next()
		}

		switch current.Type {
		case lexer.TokenTypeVariable:
			declaration, err := VariableDeclarationBuild(buffer)
			if err != nil {
				return nil, err
			}

			body = append(body, declaration)
			break
		}
	}

	program := node.NewNodeProgram(body)
	return &program, nil
}

func VariableDeclarationBuild(buffer *lexer.TokenBuffer) (node.Base, error) {
	buffer.Next() // var/let

	name := buffer.Current() // name
	if name.Type != lexer.TokenTypeIdentifier {
		return nil, fmt.Errorf("expected variable name")
	}

	buffer.Next() // =

	buffer.Next() // val
	current := buffer.Current()

	var declaration *node.VariableDeclaration
	if current.Type == lexer.TokenTypeNumber {
		number, err := NumberBuild(buffer)
		if err != nil {
			return nil, fmt.Errorf("it is not a number")
		}

		d := node.NewNodeVariableDeclaration(name.Data, number)
		declaration = &d
		buffer.Next() // move because number was saved
	}

	if declaration == nil {
		return nil, fmt.Errorf("empty declaration")
	}

	return declaration, nil
}

func NumberBuild(buffer *lexer.TokenBuffer) (node.Base, error) {
	current := buffer.Current()
	if current.Type != lexer.TokenTypeNumber {
		return nil, fmt.Errorf("not a number")
	}

	val, err := strconv.ParseFloat(current.Data, 64)
	if err != nil {
		return nil, fmt.Errorf("can not parse number")
	}

	number := node.NewNodeNumber(val)
	return &number, nil
}
