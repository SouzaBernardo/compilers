package core

import (
	"compilers/src/common"
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

func (c *Compiler) LexerValidate() (*[]common.Token, bool) {
	if c.lexer == nil {
		panic("Lexer is null")
	}
	return c.lexer.Validate()
}

func (c *Compiler) Parse(tokens *[]common.Token) {
	if c.parser == nil {
		c.parser = NewParser(tokens)
	}
	c.parser.Parse()
}
