package phases

import (
	"compilers/utils"
	"strings"
)

func Lexical(input *string) *string {
	identifyTokens := ""
	rows := strings.Split(*input, "\n")
	for _, currentRow := range rows {
		var acc = ""
		for _, currentColumn := range currentRow {
			
			if currentColumn == ' ' || currentColumn == '\t' || currentColumn == '\n'{
				identifyTokens += acc + string(currentColumn)
				acc = ""
				continue
			}

			if (currentColumn >= 'a' && currentColumn <= 'z') || 
				(currentColumn >= 'A' && currentColumn <= 'Z') {
				acc += string(currentColumn)
				lookAhead := searchNextToken(currentRow, strings.Index(currentRow, string(currentColumn)))
				next := utils.Tokens[lookAhead]
				if next != "" {
					identifyTokens += acc
					acc = ""
				}
				continue
			}

			acc += string(currentColumn)
			r := utils.Tokens[acc]
			if r != "" {
				identifyTokens += r
				acc = ""
				continue
			}
		}
		acc = ""
	}
	return &identifyTokens
}

func searchNextToken(row string, column int) string {
	if column + 1 >= len(row) {
		return ""
	}
	return string(row[column + 1])
}