package phases

import "compilers/src/common"

type ILexer interface {
	skipWhitespace()
	match() (string, bool)
	NextToken() *common.Token
	Tokenize() *[]common.Token
}

type Lexer struct {
	source   string
	position int
}

func NewLexer(source string) *Lexer {
	return &Lexer{source: source, position: 0}
}

func (l *Lexer) skipWhitespace() {

}
func (l *Lexer) match() {

}
func (l *Lexer) NextToken() *common.Token {
	return &common.Token{}
}

func (l *Lexer) Tokenize() *[]common.Token {
	return &[]common.Token{}
}

func (l *Lexer) Validate() bool {



	return false
}
