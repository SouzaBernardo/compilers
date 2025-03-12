package phases

import (
	"compilers/utils"
	"strings"
)

var currentLine int = 0
var pointer int = 0
var lookAhead rune = -1
var fontLine string = ""
var token int = NULL
var errorMessage string = ""
var currentColumn int = 0
var font []string = []string{}

const (
	NULL = 99
	NUMBER = 31
	FIM  = 26
)

func Lexical(input *string) {
	font = strings.Split(*input, "\n")
}

func searchNextToken() {

	lexeme := ""

	// Salto espaçoes enters e tabs até o inicio do proximo token
	for (lookAhead == 9) ||
		(lookAhead == '\n') ||
		(lookAhead == 8) ||
		(lookAhead == 11) ||
		(lookAhead == 12) ||
		(lookAhead == '\r') ||
		(lookAhead == 32) {
		moveLookAhead()
	}

	// tokens
	if (lookAhead >= 'A') && (lookAhead <= 'Z') {
		lexeme += string(lookAhead)
		moveLookAhead()

		for (lookAhead >= 'A' && lookAhead <= 'Z') ||
			(lookAhead >= '0' && lookAhead <= '9') {
			lexeme += string(lookAhead)
			moveLookAhead()
		}

		token = utils.GetLexemes(lexeme)
	} else if (lookAhead >= '0') && (lookAhead <= '9') {
		lexeme += string(lookAhead)
		moveLookAhead()
		for (lookAhead >= '0' && lookAhead <= '9') {
			lexeme += string(lookAhead)
			moveLookAhead()
		}
		token = NUMBER
	} else if lookAhead == '(' {
		lexeme += string(lookAhead)
		token = 16 // T_ABRE_PAR
		moveLookAhead()
	} else if lookAhead == ')' {
	} else if lookAhead == ';' {
	} else if lookAhead == ',' {
	} else if lookAhead == '+' {
	} else if lookAhead == '-' {
	} else if lookAhead == '*' {
	} else if lookAhead == '/' {
	} else if lookAhead == '%' {
	} else if lookAhead == '<' {
	} else if lookAhead == '>' {
	} else if lookAhead == FIM {
	} else {
		errorMessage = "Caractere inválido"
	}

}

func moveLookAhead() {
	if (pointer + 1) > len(fontLine) {
		currentLine++
		pointer = 0
		if currentLine < len(font) {
			lookAhead = FIM
		} else {
			fontLine = fontLine + "\n"
			lookAhead = rune(fontLine[pointer])
		}
	} else {
		lookAhead = rune(fontLine[pointer])
	}
	if lookAhead >= 'a' && lookAhead <= 'z' {
		lookAhead = lookAhead - 'a' + 'A'
	}
	pointer++
	currentColumn = pointer + 1
}
