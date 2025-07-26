package token

// TODO: move token to scanner package
type Type string

const (
	LPAREN Type = "("
	RPAREN Type     = ")"
	EQ     Type     = "="
	PRINT  Type     = "print" // TODO: PRINT should be built-in function
	ADD    Type     = "+"
	SUB    Type     = "-"
	GT     Type     = ">"
	LT     Type     = "<"
	IDENT  Type     = "IDENT"
	NUMBER Type     = "NUMBER"
	EOF    Type     = "EOF"
)

type Token struct {
	Type    Type
	Literal string
	Line    int
}

func (t Token) String() string {
	return string(t.Type) + " " + t.Literal
}
