package phases

import (
	"compilers/src/common"
	"strings"
)

type ILexer interface {
	skipWhitespace()
	match() (string, bool)
	NextToken() *common.Token
	Tokenize() *[]common.Token
}

type Lexer struct {
	phase *Phase
}

func NewLexer(source string) *Lexer {
	return &Lexer{phase: &Phase{source: source, position: 0}}
}

func (l *Lexer) Validate() bool {
	lines := strings.Split(l.phase.source, "\n")
	println(lines[0])
	for {
		token := l.phase.NextToken()
		println(token.Content, token.Type)
		if token.Type == TOKEN_UNKNOWN {
			return false
		}
		if token.Type == TOKEN_EOF {
			break
		}
	}
	return true
}
