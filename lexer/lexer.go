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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var ret token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			ret = token.Token{Type: token.EQ, Literal: literal}
		} else {
			ret = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		ret = newToken(token.PLUS, l.ch)
	case '-':
		ret = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			ret = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			ret = newToken(token.BANG, l.ch)
		}
	case '*':
		ret = newToken(token.ASTERISK, l.ch)
	case '/':
		ret = newToken(token.SLASH, l.ch)
	case '<':
		ret = newToken(token.LT, l.ch)
	case '>':
		ret = newToken(token.GT, l.ch)
	case ',':
		ret = newToken(token.COMMA, l.ch)
	case ';':
		ret = newToken(token.SEMICOLON, l.ch)
	case '(':
		ret = newToken(token.LPAREN, l.ch)
	case ')':
		ret = newToken(token.RPAREN, l.ch)
	case '{':
		ret = newToken(token.LBRACE, l.ch)
	case '}':
		ret = newToken(token.RBRACE, l.ch)
	case '"':
		ret.Type = token.STRING
		ret.Literal = l.readString()
	case 0:
		ret.Literal = ""
		ret.Type = token.EOF
	default:
		if isLetter(l.ch) {
			ret.Literal = l.readIdentifier()
			ret.Type = token.LookupIdent(ret.Literal)
			return ret
		} else if isDigit(l.ch) {
			ret.Literal = l.readNumber()
			ret.Type = token.INT
			return ret
		} else {
			ret = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return ret
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		//Todo: EOFならエラーを返す。
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readIdentifier() string {
	literalBegin := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[literalBegin:l.position]
}

func (l *Lexer) readNumber() string {
	literalBegin := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[literalBegin:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
