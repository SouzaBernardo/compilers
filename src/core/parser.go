package core

import (
	"compilers/src/common"
)

type Parser struct {
	tokens *[]common.Token
}

func NewParser(tokens *[]common.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() bool {

	position := 0

	for {
		currentToken := (*p.tokens)[position]
		a := handleFunctions(p.tokens, &position, currentToken.Type)
		err := a()
		position++
		if currentToken.Type == TOKEN_EOF || err {
			break
		}
	}

	return false
}

func handleFunctions(tokens *[]common.Token, position *int, tokenType common.TokenType) func() bool {
	switch tokenType {
	case TOKEN_FUNC:
		return func() bool {
			return tokenMain(tokens, position)
		}
	default:
		return func() bool {
			println("Não foi identificado o token ", (*tokens)[*position].Content, " ", (*tokens)[*position].Type)
			return true
		}
	}

}

func tokenMain(tokens *[]common.Token, position *int) bool {

	token := (*tokens)[*position]
	if token.Type != TOKEN_FUNC {
		return true
	}

	*position += 1
	token = (*tokens)[*position]
	if token.Type != TOKEN_FUNC_MAIN {
		return true
	}

	*position += 1
	token = (*tokens)[*position]
	if token.Type != TOKEN_LPAREN {
		return true
	}

	return false
}
