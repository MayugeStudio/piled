package scanner

import (
	"fmt"
	"piled/token"
	"strconv"
)

func isAlpha(c rune) bool {
	return (c <= 'z' && c >= 'a') || (c <= 'Z' && c >= 'A')
}

// TODO: Symbol has to be handled more politely
func isValidSymbol(c rune) bool {
	switch c {
	case '+', '-', '>', '<', '=':
		return true
	default:
		return false
	}
}

func isNumeric(c rune) bool {
	return (c <= '9' && c >= '0')
}


// TODO: I probably have to change the strategy of scanning
func ScanProgram(source string) ([]token.Token, error) {
	result := make([]token.Token, 0)

	i := 0
	line := 1

	for i < len(source) {
		switch source[i] {
		case '(':
			{
				token := token.Token{Type: token.LPAREN, Line: line}
				result = append(result, token)
			}
		case ')':
			{
				token := token.Token{Type: token.RPAREN, Line: line}
				result = append(result, token)
			}
		case '\n', '\r':
			{
				if source[i] == '\r' {
					if i+1 >= len(source) {
						return nil, fmt.Errorf("single \\r is used")
					}
					if source[i+1] != '\n' {
						return nil, fmt.Errorf("single \\r is used")
					}
					i += 1
				}
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
					token := token.Token{Type: token.IDENT, Literal: source[start : i+1], Line: line}
					result = append(result, token)
				} else if isValidSymbol(rune(source[i])) {
					token := token.Token{Type: token.IDENT, Literal: string(source[i]), Line: line}
					result = append(result, token)
				} else if isNumeric(rune(source[i])) {
					start := i
					for i+1 < len(source) && isNumeric(rune(source[i+1])) {
						i += 1
					}
					_, err := strconv.Atoi(source[start : i+1])
					if err == nil {
						token := token.Token{Type: token.NUMBER, Literal: source[start : i+1], Line: line}
						result = append(result, token)
					} else {
						return nil, fmt.Errorf("got unknown literal: %c", source[start:i+1])
					}
				} else {
					return nil, fmt.Errorf("got unknown literal: %c", rune(source[i]))
				}
			}
		}
		i += 1
	}

	eof := token.Token{Type: token.EOF, Line: line}
	result = append(result, eof)

	return result, nil
}
