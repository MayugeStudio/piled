package asm

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Location struct {
	Row int
	Col int
}

type TokenType int

const (
	Token_PUSH_INT TokenType = iota + 1
	Token_ADD
	Token_SUB
	Token_EQUAL
	Token_PRINT
	Token_INVALID
)

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
	case "=":
		return Token_EQUAL, nil
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
type lexerError struct {
	Filename string
	Loc      Location
	Err      error
}

func (l lexerError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", l.Filename, l.Loc.Row, l.Loc.Col, l.Err)
}

func newLexerError(filename string, loc Location, err error) *lexerError {
	return &lexerError{
		Filename: filename,
		Loc:      loc,
		Err:      err,
	}
}

func readFile(programPath string) (string, error) {
	bytes, err := os.ReadFile(programPath)
	if err != nil {
		return "", fmt.Errorf("could not open file `%s: %w\n", programPath, err)
	}
	return string(bytes), nil
}

func lexWord(filename string, value string, loc Location) (*Token, error) {
	opType, err := nameToTokenType(value)
	if err != nil {
		return nil, newLexerError(filename, loc, err)
	}
	if opType == Token_PUSH_INT {
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		return &Token{
			Type:  opType,
			Loc:   loc,
			Value: v,
		}, nil
	}
	return &Token{
		Type: opType,
		Loc:  loc,
	}, nil
}

func lexSourceIntoTokens(filename string, source string) ([]*Token, error) {
	ops := make([]*Token, 0)
	lines := strings.Split(source, "\n")

	for row, line := range lines {
		val := ""
		start_col := 0
		line_length := len(line)
		for col := 0; col < line_length; col++ {
			char := line[col]
			isEndOfLine := col == line_length-1
			isSpace := char == ' '
			isComma := char == ','

			if !isSpace && !isComma {
				val += string(char)
			}

			if isSpace || isEndOfLine {
				loc := Location{Row: row+1, Col: start_col+1}
				op, err := lexWord(filename, val, loc)
				if err != nil {
					return nil, err
				}
				start_col = col + 1
				ops = append(ops, op)
				val = ""
			}
		}
	}
	return ops, nil
}

func LexProgram(programPath string) ([]*Token, error) {
	source, err := readFile(programPath)
	if err != nil {
		return nil, err
	}

	ops, err := lexSourceIntoTokens(programPath, source)
	if err != nil {
		return nil, err
	}

	return ops, nil
}
