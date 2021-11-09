package compiler

import (
	"study-compiler-in-go/ast"
	"study-compiler-in-go/code"
	"study-compiler-in-go/object"
)

type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instruction: c.instructions,
		Constants:   c.constants,
	}
}

type Bytecode struct {
	Instruction code.Instructions
	Constants   []object.Object
}
