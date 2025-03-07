package main

import (
	"compilers/phases"
	"compilers/utils"
	"errors"
	"flag"
)

func readInputs() (*string, *string, error) {
	input := flag.String("input", "exemplo.emo", "The file to read")
	output := flag.String("output", "output.js", "The file to create")
	flag.Parse()

	if *input == "" || *output == "" {
		return nil, nil, errors.New("you must provide an input and output file")
	}

	return input, output, nil
}

func main() {
	input, output, err := readInputs()

	if err != nil {
		panic(err)
	}

	fileInput, err := utils.ReadFile(input)
	if err != nil {
		panic(err)
	}

	lexicalResult := phases.Lexical(&fileInput)
	fileOutput, err := utils.OpenFile(output)
	if err != nil {
		panic(err)
	}

	utils.WriteFile(fileOutput, lexicalResult)
}
