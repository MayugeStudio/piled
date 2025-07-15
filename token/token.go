package token

// TODO: move token to scanner package
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
	STRING
	EOF
)

type Token struct {
	Type  TokenType
	Value any
	Line  int
}
