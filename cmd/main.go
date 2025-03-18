package main

import (
	"compilers/phases"
	"compilers/utils"
)

const FIM_ARQUIVO = -2
const ID = -4
const NUMBER = -5

var Ponteiro int = 0
var linhaFonte string = ""
var rdFonte []string = nil
var linhaAtual int = 0
var colunaAtual int = 0
var lookAhead rune = ' '
var token int = -3
var errorMessage string = ""

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
	// sintaxe(fileInput)
}


