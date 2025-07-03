package token

import "piled/utils"

type TokenType int

const (
	T_ADD = iota
	T_SUB
	T_EQUAL
	T_PRINT
	T_LPAREN
	T_RPAREN
)

type Token struct {
	Type  TokenType
	Loc   utils.Location
	Value int
}

