package scanner

import (
	"fmt"
	"piled/token"
	"strconv"
)

type ScanContext struct {
	line    int
	index   int
	source  string
	current rune
}

func isAlpha(c rune) bool {
	return (c <= 'z' && c >= 'a') || (c <= 'Z' && c >= 'A')
}

func isNumeric(c rune) bool {
	return (c <= '9' && c >= '0')
}

func ScanProgram(source string) ([]*token.Token, error) {
	result := make([]*token.Token, 0)

	i := 0
	line := 1

	for i < len(source) {
		switch source[i] {
		case '(':
			{
				token := &token.Token{Type: token.LPAREN, Line: line}
				result = append(result, token)
			}
		case ')':
			{
				token := &token.Token{Type: token.RPAREN, Line: line}
				result = append(result, token)
			}
		case '.':
			{
				token := &token.Token{Type: token.DOT, Line: line}
				result = append(result, token)
			}
		case ',':
			{
				token := &token.Token{Type: token.COMMA, Line: line}
				result = append(result, token)
			}
		case '=':
			{
				token := &token.Token{Type: token.EQUAL, Line: line}
				result = append(result, token)
			}
		case '+':
			{
				token := &token.Token{Type: token.PLUS, Line: line}
				result = append(result, token)
			}
		case '-':
			{
				token := &token.Token{Type: token.MINUS, Line: line}
				result = append(result, token)
			}
		case '*':
			{
				token := &token.Token{Type: token.ASTERISK, Line: line}
				result = append(result, token)
			}
		case '/':
			{
				token := &token.Token{Type: token.SLASH, Line: line}
				result = append(result, token)
			}
		case '\n':
			{
				line += 1
			}
		case ' ':
			{
			} // ignore
		default:
			{
				if isAlpha(rune(source[i])) {
					start := i
					for i+1 < len(source) && isAlpha(rune(source[i+1])) {
						i += 1
					}
					token := &token.Token{Type: token.IDENTIFIER, Value: source[start : i+1], Line: line}
					result = append(result, token)
				} else if isNumeric(rune(source[i])) {
					start := i
					for i+1 < len(source) && isNumeric(rune(source[i+1])) {
						i += 1
					}
					value, err := strconv.Atoi(source[start : i+1])
					if err != nil {
						return nil, fmt.Errorf("got error while parsing number literal: %s", err)
					}
					token := &token.Token{Type: token.NUMBER, Value: value, Line: line}
					result = append(result, token)
				} else {
					return nil, fmt.Errorf("got unknown literal: %c", rune(source[i]))
				}
			}
		}
		i += 1
	}

	return result, nil
}
