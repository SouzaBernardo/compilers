package utils

import (
	"errors"
	"flag"
	"regexp"
)

func isEmojiFile(filename string) bool {
	re := regexp.MustCompile(`\.emoji$`)
	return re.MatchString(filename)
}

func ReadInputs() (*string, *string, error) {

	isValid := isEmojiFile("exemplo.emoji")
	if !isValid {
		return nil, nil, errors.New("the file must have the .emoji extension")
	}

	input := flag.String("input", "exemplo.emoji", "The file to read")
	output := flag.String("output", "output.js", "The file to create")
	flag.Parse()

	if *input == "" || *output == "" {
		return nil, nil, errors.New("you must provide an input and output file")
	}

	return input, output, nil
}
