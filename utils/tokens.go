package utils

const (
	ID = ""
)

var tokens = map[rune]string{
	'🏃': "func",
	'🔄': "for",
	'👈': ":=",
	'👀': "println",
	'👉': "range make([]int,",
    '👇': ")",
	'🤔': "if",
	'🫣': "else if",
    '🤩': "bool",
    '😇': "int",
    '🫤': "string",
}

func GetTokenValue(token rune) (string, bool) {
	result, exists := tokens[token]
	if exists {
		return result, true
	} else {
		return ID, false
	}
}
