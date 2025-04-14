package core

import (
	"compilers/src/token"
	"regexp"
	"unicode"
)

const (
	TOKEN_FUNC       token.TokenType = "TOKEN_FUNC"
	TOKEN_ID         token.TokenType = "TOKEN_ID"
	TOKEN_FUNC_MAIN  token.TokenType = "TOKEN_FUNC_MAIN"
	TOKEN_TYPE       token.TokenType = "TOKEN_TYPE"
	TOKEN_SHOW       token.TokenType = "TOKEN_SHOW"
	TOKEN_IF         token.TokenType = "TOKEN_IF"
	TOKEN_FOR        token.TokenType = "TOKEN_FOR"
	TOKEN_ASSINGMENT token.TokenType = "TOKEN_ASSINGMENT"
	TOKEN_OP         token.TokenType = "TOKEN_OP"
	TOKEN_NUMBER     token.TokenType = "TOKEN_NUMBER"
	TOKEN_LOOP_SPLIT token.TokenType = "TOKEN_LOOP_SPLIT"
	TOKEN_STRING     token.TokenType = "TOKEN_STRING"
	TOKEN_LPAREN     token.TokenType = "TOKEN_LPAREN"
	TOKEN_RPAREN     token.TokenType = "TOKEN_RPAREN"
	TOKEN_LBRACE     token.TokenType = "TOKEN_LBRACE"
	TOKEN_RBRACE     token.TokenType = "TOKEN_RBRACE"
	TOKEN_COMMA      token.TokenType = "TOKEN_COMMA"
	TOKEN_NEWLINE    token.TokenType = "TOKEN_NEWLINE"
	TOKEN_UNKNOWN    token.TokenType = "TOKEN_UNKNOWN"
	TOKEN_EOF        token.TokenType = "TOKEN_EOF"
)

var patterns = map[token.TokenType]*regexp.Regexp{
	"TOKEN_FUNC":       regexp.MustCompile(`^🛬`),
	"TOKEN_FUNC_MAIN":  regexp.MustCompile(`^🚧`),
	"TOKEN_SHOW":       regexp.MustCompile(`^👀`),
	"TOKEN_IF":         regexp.MustCompile(`^🤨`),
	"TOKEN_BOOL_TRUE":  regexp.MustCompile(`^👍`),
	"TOKEN_BOOL_FALSE": regexp.MustCompile(`^👎`),
	"TOKEN_FOR":        regexp.MustCompile(`^🔄`),
	"TOKEN_ASSINGMENT": regexp.MustCompile(`^👈`),
	"TOKEN_LOOP_SPLIT": regexp.MustCompile(`^;`),
	"TOKEN_ID":         regexp.MustCompile(`^[A-Z]+[0-9]*`),
	"TOKEN_TYPE":       regexp.MustCompile(`^(string|int|bool)`),
	"TOKEN_OP":         regexp.MustCompile(`^(==|>=|<=|<>|>|<|➕|➖|✖️|➗|%)`),
	"TOKEN_NUMBER":     regexp.MustCompile(`^[0-9]+(\.[0-9]+)?`),
	"TOKEN_STRING":     regexp.MustCompile(`^"([^"]*)"`),
	"TOKEN_LPAREN":     regexp.MustCompile(`^🫸`),
	"TOKEN_RPAREN":     regexp.MustCompile(`^🫷`),
	"TOKEN_LBRACE":     regexp.MustCompile(`^👇`),
	"TOKEN_RBRACE":     regexp.MustCompile(`^👆`),
	"TOKEN_COMMA":      regexp.MustCompile(`^,`),
	"TOKEN_NEWLINE":    regexp.MustCompile(`^\n`),
}

type Phase struct {
	source   string
	position int
}

type IPhase interface {
	Validate() bool
}

func (p *Phase) match(pattern *regexp.Regexp) (string, bool) {
	if match := pattern.FindString(p.source[p.position:]); match != "" {
		return match, true
	}
	return "", false
}

func (p *Phase) NextToken() *token.Token {

	for p.position < len(p.source) && unicode.IsSpace(rune(p.source[p.position])) {
		p.position++
	}

	if p.position >= len(p.source) {
		return &token.Token{Type: TOKEN_EOF, Content: ""}
	}

	for tokenType, pattern := range patterns {
		if match, found := p.match(pattern); found {
			p.position += len(match)
			return &token.Token{Type: tokenType, Content: match}
		}
	}
	ch := string(p.source[p.position])
	p.position++
	return &token.Token{Type: TOKEN_UNKNOWN, Content: ch}
}
