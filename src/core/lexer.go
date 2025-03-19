package core

import (
	"compilers/src/common"
)


type Lexer struct {
	phase *Phase
}

func NewLexer(source string) *Lexer {
	return &Lexer{phase: &Phase{source: source, position: 0}}
}

func (l *Lexer) Validate() (*[]common.Token, bool) {
	tokens := []common.Token{}
	for {
		token := l.phase.NextToken()
		if token.Type == TOKEN_UNKNOWN {
			return nil, false
		}
		tokens = append(tokens, *token)
		if token.Type == TOKEN_EOF {
			break
		}
	}
	return &tokens, true
}
