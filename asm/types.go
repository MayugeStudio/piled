package asm

import (
	"fmt"
	"strconv"
)

type TokenType int

const (
	Token_PUSH_INT TokenType = iota + 1
	Token_ADD
	Token_SUB
	Token_PRINT
	Token_INVALID
)

type Location struct {
	Row int
	Col int
}

type Token struct {
	Type     TokenType
	Loc      Location
	Value    int
}

func nameToTokenType(name string) (TokenType, error) {
	switch name {
	case "+":
		return Token_ADD, nil
	case "-":
		return Token_SUB, nil
	case "print":
		return Token_PRINT, nil
	default:
		{
			_, err := strconv.Atoi(name)
			if err != nil {
				return Token_INVALID, fmt.Errorf("unknown word `%s`", name)
			}
			return Token_PUSH_INT, nil
		}
	}
}
