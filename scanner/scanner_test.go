package scanner

import (
	"reflect"
	"testing"

	"piled/token"
)

func TestScanProgram(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   []*token.Token
	}{
		{
			"Parentheses", "()",
			[]*token.Token{
				{Type: token.LPAREN},
				{Type: token.RPAREN},
			},
		},
		{
			"Parentheses with a whitespace", "( )",
			[]*token.Token{
				{Type: token.LPAREN},
				{Type: token.RPAREN},
			},
		},
		{
			"Parentheses with a literal", "(hello)",
			[]*token.Token{
				{Type: token.LPAREN},
				{Type: token.LITERAL, Value: "hello"},
				{Type: token.RPAREN},
			},
		},
		{
			"Two parentheses", "()()",
			[]*token.Token{
				{Type: token.LPAREN},
				{Type: token.RPAREN},
				{Type: token.LPAREN},
				{Type: token.RPAREN},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := ScanProgram(tc.source)
			if err != nil {
				t.Errorf("ScanProgram.err = %s\n", err)
			}
			if !reflect.DeepEqual(tokens, tc.want) {
				t.Errorf("ScanProgram returned unexpected result\n")
				for i := range len(tokens) {
					if tokens[i] != tc.want[i] {
						t.Errorf("(index = %d) actual %s != want %s\n", i, tokens[i].Value, tc.want[i].Value)
					}
				}
			}
		})
	}
}
