package scanner

import (
	"fmt"
	"piled/token"
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

func ScanProgram(source string) ([]*token.Token, error) {
	result := make([]*token.Token, 0)

	i := 0
	line := 1

	for i < len(source) {
		b := source[i]
		switch b {
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
				if isAlpha(rune(b)) {
					start := i
					for i+1 < len(source) {
						if isAlpha(rune(source[i+1])) {
							i += 1
						} else {
							token := &token.Token{Type: token.LITERAL, Value: string(source[start : i+1]), Line: 1}
							result = append(result, token)
							break
						}
					}
				} else {
					return nil, fmt.Errorf("got unknown literal: %c", rune(b))
				}
			}
		}
		i += 1
	}

	return result, nil
}
