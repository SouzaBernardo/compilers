package main

import (
	"compilers/phases"
	"compilers/utils"
	"errors"
	"flag"
	"fmt"
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
	fmt.Println("Output: ", *output)
	fileInput, err := utils.ReadFile(input)
	if err != nil {
		panic(err)
	}
	lexicalResult := phases.Lexical(&fileInput)
	fmt.Println(*lexicalResult)
	// fileOutput, err := utils.OpenFile(output)
	// if err != nil {
	// 	panic(err)
	// }

	// utils.WriteFile(fileOutput, *lexicalResult)
	// fileOutput.Close()
}
