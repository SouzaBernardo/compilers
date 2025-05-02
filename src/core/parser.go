package core

import (
	"compilers/src/token"
	"fmt"
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

func (p *Parser) getTokenPlus() token.Token {
	currentPosition := *p.position
	*p.position++
	return (*p.tokens)[currentPosition]
}

func (p *Parser) getToken() token.Token {
	return (*p.tokens)[*p.position]
}

/*
<G> ::= '🛬' '🚧' '(' <VARS> ')' '{' <CMDS> '}'
*/
func (p *Parser) g() {
	p.validTokenPlus(TOKEN_FUNC)
	p.validTokenPlus(TOKEN_FUNC_MAIN)
	p.vars()
	p.cmds()
	p.validTokenPlus(TOKEN_EOF)
}

/*
<VARS> ::= <VAR> , <VARS>
<VARS> ::= <VAR>
<VARS> ::= ε
*/
func (p *Parser) vars() {
	p.validTokenPlus(TOKEN_LPAREN)
	p.var_token()
	if p.getToken().Type == TOKEN_ID {
		p.vars()
	}
}

/*
<CMDS> ::= <CMD> <CMDS>
<CMDS> ::= ε
*/
func (p *Parser) cmds() {
	p.validTokenPlus(TOKEN_LBRACE)
	p.cmd()
	p.validTokenPlus(TOKEN_RBRACE)
}

/*
<CMD>  ::= <VAR_DECLARATION>
<CMD>  ::= <CMD_IF>
<CMD>  ::= <CMD_LOOP>
<CMD>  ::= <CMD_SHOW>
*/
func (p *Parser) cmd() {

	if p.getToken().Type == TOKEN_RBRACE {
		return
	}
	tokenType := p.getTokenPlus().Type
	if tokenType == TOKEN_ID {
		p.var_token()
	} else if tokenType == TOKEN_IF {
		p.if_cmd()
	} else if tokenType == TOKEN_FOR {
		p.for_cmd()
	} else if tokenType == TOKEN_SHOW {
		p.show()
	} else {
		p.showError()
	}
}

/*
<CMD_SHOW> ::= '👀' '🫸' <E> '🫷'
*/
func (p *Parser) show() {

}

/*
<CMD_IF> ::= '🤨' '🫸' <CONDI> '🫷' '👇' <CMDS> '👆'
*/
func (p *Parser) if_cmd() {
	p.validTokenPlus(TOKEN_IF)
	p.validTokenPlus(TOKEN_LPAREN)
	p.if_condi()
	p.validTokenPlus(TOKEN_RPAREN)
	p.validTokenPlus(TOKEN_LBRACE)
	p.cmds()
	p.validToken(TOKEN_RBRACE)
}

/*
<CONDI>  ::= <E> '>' <E>
<CONDI> ::= <E> '>=' <E>
<CONDI> ::= <E> '<>' <E>
<CONDI> ::= <E> '<=' <E>
<CONDI> ::= <E> '<' <E>
<CONDI> ::= <E> '==' <E>
*/
func (p *Parser) if_condi() {

}

/*
<CMD_LOOP> ::= '🔄' '🫸' <CONDI_LOOP> '🫷' '👇' <CMDS> '👆'
*/
func (p *Parser) for_cmd() {
}


func (p *Parser) var_token() {
	p.validToken(TOKEN_ID)
	p.id()
	p.idType()
	p.comma()
}

func (p *Parser) id() {
	p.validTokenPlus(TOKEN_ID)
}

func (p *Parser) idType() {
	p.validTokenPlus(TOKEN_TYPE)
}

func (p *Parser) comma() {
	token := p.getTokenPlus()

	if token.Type == TOKEN_COMMA || token.Type == TOKEN_RPAREN {
		return
	}
	p.showError()
}

func (p *Parser) validTokenPlus(tokenType token.TokenType) {
	if p.getTokenPlus().Type != tokenType {
		p.showError()
	}
}

func (p *Parser) validToken(tokenType token.TokenType) {
	if p.getToken().Type != tokenType {
		p.showError()
	}
}

func (p *Parser) showError() {	
	message, _ := fmt.Printf("Erro Sintático encontrado com o caractere: %s \n", p.getToken().Content)
	panic(message)
}
