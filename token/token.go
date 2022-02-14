package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifier + literal
	IDENT = "IDENT" //ex.) add, foobar, x, y ...
	INT   = "INT"   //ex.) 123, 456 ...

	ASSIGN = "="
	PLUS   = "+"

	// delimeter
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
