package main

import (
	"fmt"
	"github.com/kepka42/llvm-ts-compiler/src/lib/ast"
	"github.com/kepka42/llvm-ts-compiler/src/lib/lexer"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("file not input")
		return
	}

	filename := os.Args[1]
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	builder := lexer.Builder{}
	builder.SetInput(data)

	tokens, err := builder.Run()
	if err != nil {
		panic(err)
	}

	tree := ast.NewTree()

	err = tree.Build(tokens)
	if err != nil {
		panic(err)
	}

	fmt.Println("ready")
}
