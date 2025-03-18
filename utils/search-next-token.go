package utils

import (
	"errors"
	"strings"
)

func SearchNextToken(pointer *int, currentLine *string, lineCounter int, programSize int, lookAhead *rune) error {
	token := ' '
	lexeme := ""
	for ((*lookAhead) == 9) ||
		((*lookAhead) == '\n') ||
		((*lookAhead) == 8) ||
		((*lookAhead) == 11) ||
		((*lookAhead) == 12) ||
		((*lookAhead) == '\r') ||
		((*lookAhead) == 32) {
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
	}
	if (*lookAhead >= 'A') && (*lookAhead <= 'Z') {
		lexeme += string(*lookAhead)
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
		for ((*lookAhead >= 'A') && (*lookAhead <= 'Z')) ||
			((*lookAhead >= '0') && (*lookAhead <= '9')) ||
			(*lookAhead == '_') {
			lexeme += string(*lookAhead)
			MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
		}

		_, exists := Tokens[lexeme]

		if !exists {
			token = ID
		}
	} else if (*lookAhead >= '0') && (*lookAhead <= '9') {
		lexeme += string(*lookAhead)
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
		for (*lookAhead >= '0') && (*lookAhead <= '9') {
			lexeme += string(*lookAhead)
			MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
		}
		token = ' '
	} else if *lookAhead == '(' {
		lexeme += string(*lookAhead)
		token = '('
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
	} else if *lookAhead == ')' {
		lexeme += string(*lookAhead)
		token = '('
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
	} else if *lookAhead == ',' {
		lexeme += string(*lookAhead)
		token = ','
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
	} else if *lookAhead == '+' {
		lexeme += string(*lookAhead)
		token = '+'
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
	} else if *lookAhead == '-' {
		lexeme += string(*lookAhead)
		token = '-'
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
	} else if *lookAhead == '*' {
		lexeme += string(*lookAhead)
		token = '*'
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
	} else if *lookAhead == '/' {
		lexeme += string(*lookAhead)
		token = '/'
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
	} else if *lookAhead == '%' {
		lexeme += string(*lookAhead)
		token = '%'
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
	} else if *lookAhead == '<' {
		lexeme += string(*lookAhead)
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
		if *lookAhead == '=' {
			lexeme += string(*lookAhead)
			token = '<'
		} else {
			token = '<'
		}
	} else if *lookAhead == '>' {
		lexeme += string(*lookAhead)
		MoveLookAhead(pointer, currentLine, lineCounter, programSize, lookAhead)
		if *lookAhead == '=' {
			lexeme += string(*lookAhead)
			token = '>'
		} else {
			token = '>'
		}
	} else if *lookAhead == -10 {
		token = -10
	} else {
		return errors.New(strings.NewReplacer("{lexeme}", lexeme, "{token}", string(token)).Replace("erro: Caractere inválido {lexeme} {token}"))
	}

	return nil
}
