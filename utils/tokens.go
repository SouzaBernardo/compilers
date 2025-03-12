package utils

const (
	ID = 32
)

var tokens = map[string]int{
	"🏃": 1,
	"🔄": 2,
	"👈": 3,
	"👀": 4,
	"👉": 5,
	"👇": 6,
	"🤔": 7,
	"🫣": 8,
	"🤩": 9,
	"😇": 10,
	"🫤": 11,
}

func GetLexemes(token string) int {
	result, exists := tokens[token]
	if exists {
		return result
	}
	return ID
}
