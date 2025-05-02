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

func (c *Compiler) Compile()  {
	tokens := c.lexerValidate()
	if tokens == nil {
		panic("Erro Léxico")
	}
	c.parse(tokens)
}

func (c *Compiler) lexerValidate() *[]token.Token {
	if c.lexer == nil {
		panic("Lexer is null")
	}
	r, err := c.lexer.Validate()
	if err {
		panic("Erro léxico")
	} else {
		fmt.Println("Etapa léxica completada")
	}
	return r
}

func (c *Compiler) parse(tokens *[]token.Token) {
	if c.parser == nil {
		c.parser = NewParser(*tokens)
	}
	c.parser.Parse()
	fmt.Println("Etapa sintática completada")
}
