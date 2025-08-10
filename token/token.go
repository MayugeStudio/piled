package token

// TODO: move token to scanner package
// Type used to represent type of token
type Type string

const (
	IF     Type = "if"
	ELSE   Type = "else"
	END    Type = "end"
	PRINT  Type = "print" // TODO: PRINT should be built-in function
	SHL    Type = "shr"
	SHR    Type = "shl"
	ADD    Type = "+"
	SUB    Type = "-"
	MUL    Type = "*"
	DIV    Type = "/"
	MOD    Type = "%"
	AND    Type = "&"
	OR     Type = "|"
	GT     Type = ">"
	LT     Type = "<"
	EQ     Type = "="
	IDENT  Type = "IDENT"
	NUMBER Type = "NUMBER"
	EOF    Type = "EOF"
)

var keywords = map[string]Type{
	"print": PRINT,
	"shl":   SHL,
	"shr":   SHR,
	"if":    IF,
	"else":  ELSE,
	"end":   END,
}

// Token represent the lexer token
type Token struct {
	Type    Type
	Literal string
	Line    int
}

// LookupIdentifier check whether specified keyword name exists or not
// if it exists simply return its type, otherwise return IDENT
func LookupIdentifier(in string) Type {
	if t, ok := keywords[in]; ok {
		return t
	} else {
		return IDENT
	}
}
