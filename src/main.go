package main

import (
	"compilers/src/common"
	"compilers/src/core"
)

func main() {
	input, _, err := common.ReadInputs()

	if err != nil {
		panic(err)
	}

	source, err := common.ReadFile(input)
	if err != nil {
		panic(err)
	}
	core.NewCompiler(source).Compile()
}
