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
	LITERAL
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
}
