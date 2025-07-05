package token

type TokenType int

const (
	ADD = iota
	SUB
	EQUAL
	PRINT
	LPAREN
	RPAREN
	LITERAL
)

type Token struct {
	Type  TokenType
	Value string
}

