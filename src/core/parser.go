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

/*
<G> ::= '🛬' '🚧' '(' <VARS> ')' '{' <CMDS> '}'
*/
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

/*
<VARS> ::= <VAR> , <VARS>
<VARS> ::= <VAR>
<VARS> ::= ε
*/
func (p *Parser) vars() {

	if p.getToken().Type != TOKEN_LPAREN {
		panic("ERoooooooo")
	}
	p.var_token()

	if (*p.tokens)[*p.position].Type == TOKEN_ID {
		p.vars()
	}
}

/*
<CMDS> ::= <CMD> <CMDS>
<CMDS> ::= ε
*/
func (p *Parser) cmds() {
	if p.getToken().Type != TOKEN_LBRACE {
		panic("ERoooooooo")
	}

	p.cmd()

	if p.getToken().Type != TOKEN_RBRACE {
		panic("ERoooooooo")
	}
}

/*
<CMD>  ::= <VAR_DECLARATION>
<CMD>  ::= <CMD_IF>
<CMD>  ::= <CMD_LOOP>
<CMD>  ::= <CMD_SHOW>
*/
func (p *Parser) cmd() {

	if p.getToken2().Type == TOKEN_RBRACE {
		return
	}
	tokenType := p.getToken().Type
	if tokenType == TOKEN_ID {
		p.var_token()
	} else if tokenType == TOKEN_SHOW {
		p.show()
	} else if tokenType == TOKEN_IF {
		p.if_cmd()
	} else if tokenType == TOKEN_FOR {
		p.for_cmd()
	} else {
		// panic("ERoooooooo")
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
}

/*
<CMD_LOOP> ::= '🔄' '🫸' <CONDI_LOOP> '🫷' '👇' <CMDS> '👆'
*/
func (p *Parser) for_cmd() {
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
