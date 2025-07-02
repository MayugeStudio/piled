package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Location struct {
	Row int
	Col int
}

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
	FilePath string
	Loc      Location
	Err      error
}

func (l lexerError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", l.FilePath, l.Loc.Row, l.Loc.Col, l.Err)
}

func newLexerError(filepath string, loc Location, err error) *lexerError {
	return &lexerError{
		FilePath: filepath,
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

func lexWord(filepath string, value string, loc Location) (*Token, error) {
	opType, err := nameToTokenType(value)
	if err != nil {
		return nil, newLexerError(filepath, loc, err)
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

func lexSourceIntoTokens(filepath string, source string) ([]*Token, error) {
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
				loc := Location{Row: row + 1, Col: start_col + 1}
				op, err := lexWord(filepath, val, loc)
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

func LexProgram(programPath string, source string) ([]*Token, error) {
	ops, err := lexSourceIntoTokens(programPath, source)
	if err != nil {
		return nil, err
	}

	return ops, nil
}

func main() {
	args := os.Args
	programName := args[0] // TODO: Introduce some sort of arguments operating function
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input-file>\n", programName)
		fmt.Fprintf(os.Stderr, "ERROR: input file was not provided\n")
		os.Exit(1)
	}

	args = args[1:]
	inputPath := args[0]

	// Reading input file
	source, err := readFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not read source file from `%s`: %s\n", inputPath, err)
		os.Exit(1)
	}

	// Lexing
	ops, err := LexProgram(inputPath, source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}

