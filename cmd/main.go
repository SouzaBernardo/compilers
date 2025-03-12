package main

import (
	"compilers/phases"
	"compilers/utils"
)

func main() {
	input, _, err := utils.ReadInputs()

	if err != nil {
		panic(err)
	}

	fileInput, err := utils.ReadFile(input)
	if err != nil {
		panic(err)
	}

	phases.Lexical(&fileInput)
}
