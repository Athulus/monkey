package lexer

import (
	"github.com/Athulus/monkey/token"
)

//Lexer is the tyoe that will parse a string and return tokens
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

//New creates a new Lexer
func New(input string) *Lexer {
	return &Lexer{"", 0, 0, 0}
}

//NextToken does something
func (l *Lexer) NextToken() token.Token {
	return token.Token{Type: token.ILLEGAL, Literal: "ILLEGAL"}
}
