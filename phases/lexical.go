package phases

func Lexical(input *string) *string {

	identifyTokens := ""

	currentRow := 0
	currentColumn := 0
	pointer := 0
	line := ""
	currentToken := ""
	errorMessage := ""
	token := ""

	for {
		if token == "fim" || token == "erro" {
			break
		}
		searchNextToken()

	}

	return &identifyTokens
}

func searchNextToken() *string {
	lexeme := ""
	return &lexeme
}
