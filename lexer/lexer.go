package lexer

import "github.com/mbparis/interpreter-book/monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to curent char)
	readPosition int  // cuurent reading positiong (after current character)
	ch           byte //current char under examination

}

// note that the use of the `byte` type for Lexer.ch is restricting
// the Lexer to  ASCII characters
// TO support unicode, the type would need to be `rune`
// and the reading method would need to acount for characters composing multiple bytes
// TODO support unicode !

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()

	return lexer
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	}

	l.readChar()
	return tok
}

func newToken(t token.TokenType, ch byte) token.Token {
	return token.Token{Type: t, Literal: string(ch)}

}

// the purpose of readChar is to give the next  char and to advance position in the input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}
