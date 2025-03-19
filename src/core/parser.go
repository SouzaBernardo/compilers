package core

import "compilers/src/common"

type Parser struct {
	tokens *[]common.Token
}

func NewParser(tokens *[]common.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() bool {
	return false
}
