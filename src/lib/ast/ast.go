package ast

import (
	"github.com/kepka42/llvm-ts-compiler/src/lib/ast/node"
	"github.com/kepka42/llvm-ts-compiler/src/lib/lexer"
)

type Tree struct {
	root node.Base
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Build(tokens []lexer.Token) error {
	t.root = new(node.Program)
	return nil
}
