package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  //current index of input string
	readPosition int  //next index
	ch           byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var ret token.Token
	switch l.ch {
	case '=':
		ret = newToken(token.ASSIGN, l.ch)
	case ';':
		ret = newToken(token.SEMICOLON, l.ch)
	case '(':
		ret = newToken(token.LPAREN, l.ch)
	case ')':
		ret = newToken(token.RPAREN, l.ch)
	case ',':
		ret = newToken(token.COMMA, l.ch)
	case '+':
		ret = newToken(token.PLUS, l.ch)
	case '{':
		ret = newToken(token.LBRACE, l.ch)
	case '}':
		ret = newToken(token.RBRACE, l.ch)
	case 0:
		ret.Literal = ""
		ret.Type = token.EOF
	}

	l.readChar()

	return ret
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
