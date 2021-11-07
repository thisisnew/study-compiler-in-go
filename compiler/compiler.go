package compiler

import (
	"study-compiler-in-go/code"
	"study-compiler-in-go/object"
)

type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}
