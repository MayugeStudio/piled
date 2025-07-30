package token

// TODO: move token to scanner package
// Type used to represent type of token
type Type string

const (
	LPAREN Type = "("
	RPAREN Type = ")"
	EQ     Type = "="
	PRINT  Type = "print" // TODO: PRINT should be built-in function
	ADD    Type = "+"
	SUB    Type = "-"
	GT     Type = ">"
	LT     Type = "<"
	IDENT  Type = "IDENT"
	NUMBER Type = "NUMBER"
	EOF    Type = "EOF"
)

var keywords = map[string]Type{
}

// Token represent the lexer token
type Token struct {
	Type    Type
	Literal string
	Line    int
}

func LookupIdentifier(in string) Type {
	if t, ok := keywords[in]; ok {
		return t
	} else {
		return IDENT
	}
}
