package ast

import (
	"github.com/kepka42/llvm-ts-compiler/src/lib/ast/builder"
	"github.com/kepka42/llvm-ts-compiler/src/lib/ast/node"
	"github.com/kepka42/llvm-ts-compiler/src/lib/lexer"
)

type Tree struct {
	root node.Base
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Build(buffer *lexer.TokenBuffer) error {
	root, err := builder.RootBuild(buffer)
	if err != nil {
		return err
	}

	t.root = root
	return nil
}
