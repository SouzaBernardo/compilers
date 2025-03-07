package phases

import (
	"compilers/utils"
)

func Lexical(input *string) *string {
	identifyTokens := "package main\n"
	for _, currentLexeme := range *input {
		value := searchNextToken(currentLexeme)
		identifyTokens += value
	}
	return &identifyTokens
}

func searchNextToken(currentLexeme rune) string {

	token, exists := utils.GetTokenValue(currentLexeme)

	if exists {
		return string(token)
	}

	return string(currentLexeme)
}
