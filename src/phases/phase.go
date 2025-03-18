package phases

import (
	"compilers/src/common"
	"regexp"
	"unicode"
)

const (
	TOKEN_FUNC    common.TokenType = "FUNC"
	TOKEN_ID      common.TokenType = "ID"
	TOKEN_TYPE    common.TokenType = "TYPE"
	TOKEN_SHOW    common.TokenType = "SHOW"
	TOKEN_IF      common.TokenType = "IF"
	TOKEN_OP      common.TokenType = "OP"
	TOKEN_NUMBER  common.TokenType = "NUMBER"
	TOKEN_STRING  common.TokenType = "STRING"
	TOKEN_LPAREN  common.TokenType = "LPAREN"
	TOKEN_RPAREN  common.TokenType = "RPAREN"
	TOKEN_LBRACE  common.TokenType = "LBRACE"
	TOKEN_RBRACE  common.TokenType = "RBRACE"
	TOKEN_COMMA   common.TokenType = "COMMA"
	TOKEN_NEWLINE common.TokenType = "NEWLINE"
	TOKEN_UNKNOWN common.TokenType = "UNKNOWN"
	TOKEN_EOF     common.TokenType = "EOF"
)

var patterns = map[common.TokenType]*regexp.Regexp{
	"TOKEN_FUNC":       regexp.MustCompile(`^funcao`),
	"TOKEN_MAIN_FUNC":  regexp.MustCompile(`^principal`),
	"TOKEN_SHOW":       regexp.MustCompile(`^mostrar`),
	"TOKEN_IF":         regexp.MustCompile(`^se`),
	"TOKEN_BOOL_TRUE":  regexp.MustCompile(`^verdadeiro`),
	"TOKEN_BOOL_FALSE": regexp.MustCompile(`^falso`),
	"TOKEN_ID":         regexp.MustCompile(`^[A-Z]+[0-9]*`),
	"TOKEN_TYPE":       regexp.MustCompile(`^(string|int|bool)`),
	"TOKEN_OP":         regexp.MustCompile(`^(==|>=|<=|<>|>|<|\+|-|\*|/|%)`),
	"TOKEN_NUMBER":     regexp.MustCompile(`^[0-9]+(\.[0-9]+)?`),
	"TOKEN_STRING":     regexp.MustCompile(`^"([^"]*)"`),
	"TOKEN_LPAREN":     regexp.MustCompile(`^\(`),
	"TOKEN_RPAREN":     regexp.MustCompile(`^\)`),
	"TOKEN_LBRACE":     regexp.MustCompile(`^\{`),
	"TOKEN_RBRACE":     regexp.MustCompile(`^\}`),
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

func (p *Phase) NextToken() *common.Token {

	for p.position < len(p.source) && unicode.IsSpace(rune(p.source[p.position])) {
		p.position++
	}

	if p.position >= len(p.source) {
		return &common.Token{Type: TOKEN_EOF, Content: ""}
	}

	for tokenType, pattern := range patterns {
		if match, found := p.match(pattern); found {
			p.position += len(match)
			return &common.Token{Type: tokenType, Content: match}
		}
	}
	ch := string(p.source[p.position])
	p.position++
	return &common.Token{Type: TOKEN_UNKNOWN, Content: ch}
}
