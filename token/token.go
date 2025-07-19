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
	IDENT = "IDENT"
	NUMBER = "NUMBER"
	STRING = "STRING"
	EOF = "EOF"
)

type Token struct {
	Type    Type
	Literal string
	Line    int
}

func (t Token) String() string {
	return string(t.Type) + " " + t.Literal
}
