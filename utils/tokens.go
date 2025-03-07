package utils


const ID = 0
var tokens = map[rune]int{
    '🏃': 1,
    '🔄': 2,
    '👈': 3,
    '👀': 4,
    '👉': 5,
    '🤔': 6,
    '🫣': 7,
}

func GetTokenValue(token rune) int {
    result, exists := tokens[token]
    if exists {
        return result
    } else {
        return ID
    }
}