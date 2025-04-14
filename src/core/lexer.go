package core

import "compilers/src/token"

type Lexer struct {
	phase *Phase
}

func NewLexer(source string) *Lexer {
	return &Lexer{phase: &Phase{source: source, position: 0}}
}

func (l *Lexer) Validate() (*[]token.Token, bool) {
	tokens := []token.Token{}
	for {
		token := l.phase.NextToken()
		if token.Type == TOKEN_UNKNOWN {
			return nil, true
		}
		tokens = append(tokens, *token)
		if token.Type == TOKEN_EOF {
			break
		}
	}
	return &tokens, false
}
