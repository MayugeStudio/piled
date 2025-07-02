package token


type TokenType int

const (
	Token_INVALID TokenType = iota
	Token_PUSH_INT
	Token_ADD
	Token_SUB
	Token_EQUAL
	Token_PRINT
)

type Token struct {
	Type  TokenType
	Loc   Location
	Value int
}

