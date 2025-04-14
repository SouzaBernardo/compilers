package core

import (
	"compilers/src/token"
)

type Parser struct {
	position *int
	tokens   *[]token.Token
	err      bool
}

func NewParser(tokens []token.Token) *Parser {
	counter := 0
	return &Parser{position: &counter, tokens: &tokens, err: false}
}

func (p *Parser) Parse() bool {
	p.g()
	return p.err
}

func (p *Parser) getToken() token.Token {
	currentPosition := *p.position
	*p.position++
	return (*p.tokens)[currentPosition]
}

func (p *Parser) getToken2() token.Token {
	return (*p.tokens)[*p.position]
}

// <G> ::= '🛬' '🚧' '(' <VARS> ')' '{' <CMDS> '}'
func (p *Parser) g() {
	if p.getToken().Type != TOKEN_FUNC {
		panic("ERoooooooo")
	}

	if p.getToken().Type != TOKEN_FUNC_MAIN {
		panic("ERoooooooo")
	}

	p.vars()
	p.cmds()

	if p.getToken().Type != TOKEN_EOF {
		panic("ERoooooooo")
	}
}

// <VARS> ::= <VAR> , <VARS>
// <VARS> ::= <VAR>
// <VARS> ::= ε
func (p *Parser) vars() {

	if p.getToken().Type != TOKEN_LPAREN {
		panic("ERoooooooo")
	}
	p.var_token()

	if (*p.tokens)[*p.position].Type == TOKEN_ID {
		p.vars()
	}
}

func (p *Parser) cmds() {
	if p.getToken().Type != TOKEN_LBRACE {
		panic("ERoooooooo")
	}

	// others commands

	if p.getToken().Type != TOKEN_RBRACE {
		panic("ERoooooooo")
	}
}

func (p *Parser) var_token() {
	if (*p.tokens)[*p.position].Type != TOKEN_ID {
		return
	}
	p.id()
	p.idType()
	p.comma()
}

func (p *Parser) id() {
	if p.getToken().Type != TOKEN_ID {
		panic("error")
	}
}

func (p *Parser) idType() {
	if p.getToken().Type != TOKEN_TYPE {
		panic("Erro aqui")
	}
}

func (p *Parser) comma() {
	token := p.getToken()

	if token.Type == TOKEN_COMMA || token.Type == TOKEN_RPAREN {
		return
	}
	panic("err")
}
