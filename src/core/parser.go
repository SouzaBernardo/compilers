package core

import "compilers/src/common"

type Parser struct {
	tokens *[]common.Token
}

func NewParser(tokens *[]common.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() bool {

	for i, token := range *(p.tokens) {
		println(i, token.Content)
	}

	return false
}

// func handleSyntax(tokens *common.Token, position int) {
// 	switch tokens.Type {
// 	case TOKEN_FUNC:
// 		tokenFunc

// 	}

// }
// func tokenFunc() {

// }
