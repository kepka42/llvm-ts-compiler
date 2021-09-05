package compiler

import (
	"fmt"
	"github.com/kepka42/llvm-ts-compiler/src/lib/ast"
	"github.com/kepka42/llvm-ts-compiler/src/lib/ast/node"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func Run(tree *ast.Tree) (string, error) {
	// Just try to generate llvm IR
	root := tree.Root()

	if root.Type() != node.TypeProgram {
		return "", fmt.Errorf("root type error")
	}

	module := ir.NewModule()
	program := root.(*node.Program)

	for _, val := range program.Body() {
		if val.Type() == node.TypeVariableDeclaration {
			variable := val.(*node.VariableDeclaration)
			variableValue := variable.Value()

			if variableValue.Type() == node.TypeNumber {
				variableValueNumber := variableValue.(*node.Number)
				module.NewGlobalDef(variable.Name(), constant.NewFloat(types.Double, variableValueNumber.Value()))
			}
		}
	}

	return module.String(), nil
}
