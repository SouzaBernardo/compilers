package main

import (
	"compilers/src/common"
	"compilers/src/phases"
)

type Compiller struct {
	source string
	lexer  *phases.Lexer
}

func NewCompiler(source string) *Compiller {
	return &Compiller{
		source: source,
		lexer:  phases.NewLexer(source),
	}
}

func main() {
	input, _, err := common.ReadInputs()

	if err != nil {
		panic(err)
	}

	source, err := common.ReadFile(input)
	if err != nil {
		panic(err)
	}

	compiler := NewCompiler(source)
	complete := compiler.lexer.Validate()
	if !complete {
		panic("error")
	}
}
