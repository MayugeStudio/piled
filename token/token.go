package token

// TODO: move token to scanner package
type Type string

const (
	LPAREN Type = "("
	RPAREN      = ")"
	IDENT       = "IDENT"
	NUMBER      = "NUMBER"
	EOF         = "EOF"
)

type Token struct {
	Type    Type
	Literal string
	Line    int
}

func (t Token) String() string {
	return string(t.Type) + " " + t.Literal
}
