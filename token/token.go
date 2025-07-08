package token

type TokenType int

const (
	LPAREN = iota
	RPAREN
	DOT
	COMMA
	EQUAL
	PLUS
	MINUS
	ASTERISK
	SLASH
	IDENTIFIER
	NUMBER
	EOF
)

type Token struct {
	Type  TokenType
	Value any
	Line  int
}
