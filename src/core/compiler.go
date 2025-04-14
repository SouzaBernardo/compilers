package core

import (
	"compilers/src/token"
	"fmt"
)

type Compiler struct {
	source string
	lexer  *Lexer
	parser *Parser
}

func NewCompiler(source string) *Compiler {
	return &Compiler{
		source: source,
		lexer:  NewLexer(source),
	}
}

func (c *Compiler) LexerValidate() *[]token.Token {
	if c.lexer == nil {
		panic("Lexer is null")
	}
	r, err := c.lexer.Validate()
	if err {
		panic("error")
	} else {
		fmt.Println("Lexer complete")
	}
	return r
}

func (c *Compiler) Parse(tokens *[]token.Token) {
	if c.parser == nil {
		c.parser = NewParser(*tokens)
	}
	c.parser.Parse()
	fmt.Println("Sintax complete")
}
