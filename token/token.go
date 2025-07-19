package token

// TODO: move token to scanner package
type Type string

const (
	LPAREN Type = "("
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
	Type  Type
	Value any
	Line  int
}

func (t Token) String() string {
	return string(t.Type)
}
