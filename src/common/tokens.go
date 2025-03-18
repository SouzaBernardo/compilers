package common

type TokenType string

type Token struct {
	Type  TokenType
	Content string
}