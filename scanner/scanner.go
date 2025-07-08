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

	for i < len(source) {
		b := source[i]
		switch b {
		case '(':
			{
				token := &token.Token{Type: token.LPAREN}
				result = append(result, token)
			}
		case ')':
			{
				token := &token.Token{Type: token.RPAREN}
				result = append(result, token)
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
							token := &token.Token{Type: token.LITERAL, Value: string(source[start : i+1])}
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
