package token

// Type used to represent type of token
type Type string

const (
	IF     Type = "if"
	ELSE   Type = "else"
	END    Type = "end"
	PRINT  Type = "print"
	SHL    Type = "shr"
	SHR    Type = "shl"
	DUP    Type = "dup"
	// TODO: Token should represent meaning of token.
	//       For example, ADD has to be PLUS, SUB has to be MINUS, and so on.
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
	OCURLY Type = "{"
	CCURLY Type = "}"
	IDENT  Type = "IDENT"
	NUMBER Type = "NUMBER"
	EOF    Type = "EOF"
)

var keywords = map[string]Type{
	"print": PRINT,
	"dup":   DUP,
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
}

// LookupIdentifier check whether specified keyword name exists or not.
// If it exists simply return its type, otherwise return IDENT
func LookupIdentifier(in string) Type {
	if t, ok := keywords[in]; ok {
		return t
	} else {
		return IDENT
	}
}
