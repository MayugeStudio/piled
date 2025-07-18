package token

// TODO: move token to scanner package
type TokenType string

const (
	LPAREN TokenType = "("
	RPAREN           = ")"
	DOT              = "."
	COMMA = ","
	EQUAL = "="
	PLUS  = "+"
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"
	IDENTIFIER = "IDENT"
	NUMBER = "NUMBER"
	STRING = "STRING"
	EOF = "EOF"
)

type Token struct {
	Type  TokenType
	Value any
	Line  int
}

func (t Token) String() string {
	return string(t.Type)
}
