package scanner

import (
	"strings"
	"fmt"
	"piled/utils"
	"piled/token"
)

func LexProgram(filepath string, source string) ([]*token.Token, error) {
	ops := make([]*token.Token, 0)
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
				loc := utils.Location{Row: row + 1, Col: start_col + 1}
				op, err := lexLiteralIntoToken(filepath, val, loc)
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

func lexLiteralIntoToken(filepath string, literal string, loc utils.Location) (*token.Token, error) {
	switch literal {
		case "(": {
			return &token.Token{ Type: token.T_LPAREN, Loc: loc }, nil
		}
		case ")": {
			return &token.Token{ Type: token.T_RPAREN, Loc: loc }, nil
		}
		default: {
			return nil, fmt.Errorf("invalid token: %s", literal)
		}
	}
}
