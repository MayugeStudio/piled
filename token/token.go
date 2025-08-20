package token

// Type used to represent type of token
type Type string

const (
	IF    Type = "if"
	ELSE  Type = "else"
	PRINT Type = "print"
	SHL   Type = "shr"
	SHR   Type = "shl"
	WHILE Type = "while"

	// TODO: These builtin-functions should not be here.
	DUP  Type = "dup"
	SWAP Type = "swap"
	ROT  Type = "rot"
	DROP Type = "drop"

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

// TODO: Introduce Punctuations
var keywords = map[string]Type{
	"print": PRINT,
	"dup":   DUP,
	"swap":  SWAP,
	"drop":  DROP,
	"rot":   ROT,
	"shl":   SHL,
	"shr":   SHR,
	"while": WHILE,
	"if":    IF,
	"else":  ELSE,
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
