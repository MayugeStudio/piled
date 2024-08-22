package asm

import "fmt"
import "strings"

type Register int8

const (
	R_ACC Register = iota << 1
	R_IDX
	R_TMP
	R_RET
	R_0
	R_1
	R_2
	R_3
	R_4
	R_5
	R_6
	R_7
	R_PC
	R_SP
	R_BP
	R_FL
)

// Location
type Location struct {
	row int
	col int
}

func (l Location) String() string {
	return fmt.Sprintf("%d:%d", l.row, l.col)
}

// AsmToken
type AsmToken struct {
	value string
	loc   Location
}

func (t AsmToken) String() string {
	return fmt.Sprintf("%s: `%s`", t.loc, t.value)
}

type OpCode struct{}

func TokenizeSource(source string) (tokens []AsmToken, err error) {
	if len(source) == 0 {
		return
	}

	lines := strings.Split(source, "\n")

	for row, line := range lines {
		if len(line) == 0 {
			continue
		}
		val := ""
		start_col := 0
		line_length := len(line)
		for col := 0; col < line_length; col++ {
			char := line[col]
			isSpace := char == ' '
			isColon := char == ':'
			isComma := char == ','

			isEndOfLine := col == line_length-1

			if !isSpace && !isColon && !isComma {
				val += string(char)
			}

			if char == ' ' || isEndOfLine {
				token := AsmToken{
					value: val,
					loc:   Location{row: row, col: start_col},
				}
				start_col = col+1
				tokens = append(tokens, token)
				val = ""
			}
		}
	}
	return
}

