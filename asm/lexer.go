package asm

import (
	"fmt"
	"os"
	"strings"
)

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

func lexSourceIntoOPs(source string) ([]*OP, error) {
	ops := make([]*OP, 0)
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
				op, err := newOP(val, newLocation(row+1, start_col+1))
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

func LexProgram(programPath string) ([]*OP, error) {
	source, err := readFile(programPath)
	if err != nil {
		return nil, err
	}

	ops, err := lexSourceIntoOPs(source)
	if err != nil {
		return nil, fmt.Errorf("%s:%s", programPath, err)
	}

	return ops, nil
}
