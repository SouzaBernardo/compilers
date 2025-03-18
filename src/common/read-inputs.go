package common

import (
	"errors"
	"flag"
	"regexp"
)

func isValidFile(filename string) bool {
	re := regexp.MustCompile(`\.be$`)
	return re.MatchString(filename)
}

func ReadInputs() (*string, *string, error) {

	
	input := flag.String("input", "docs/exemplo.be", "The file to read")
	output := flag.String("output", "output.js", "The file to create")
	flag.Parse()

	isValid := isValidFile(*input)
	if !isValid {
		return nil, nil, errors.New("the file must have the .emoji extension")
	}


	if *input == "" || *output == "" {
		return nil, nil, errors.New("you must provide an input and output file")
	}

	return input, output, nil
}
