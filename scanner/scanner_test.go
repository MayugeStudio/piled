package scanner

import (
	"reflect"
	"testing"

	"piled/token"
)

func TestScanProgram(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		want    []*token.Token
		wantErr bool
	}{
		{
			"Parentheses", "()",
			[]*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
			}, false,
		},
		{
			"Parentheses with a whitespace", "( )",
			[]*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
			}, false,
		},
		{
			"Parentheses with an identifier", "(hello)",
			[]*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.IDENTIFIER, Value: "hello", Line: 1},
				{Type: token.RPAREN, Line: 1},
			}, false,
		},
		{
			"Two parentheses", "()()",
			[]*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
			}, false,
		},
		{
			"Multiple lines", "()\n()\n()\n",
			[]*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.LPAREN, Line: 2},
				{Type: token.RPAREN, Line: 2},
				{Type: token.LPAREN, Line: 3},
				{Type: token.RPAREN, Line: 3},
			}, false,
		},
		{
			"Multiple lines with varius type of tokens", "(I)\n(like)\n(golang)\n",
			[]*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.IDENTIFIER, Value: "I", Line: 1},
				{Type: token.RPAREN, Line: 1},

				{Type: token.LPAREN, Line: 2},
				{Type: token.IDENTIFIER, Value: "like",Line: 2},
				{Type: token.RPAREN, Line: 2},

				{Type: token.LPAREN, Line: 3},
				{Type: token.IDENTIFIER, Value: "golang", Line: 3},
				{Type: token.RPAREN, Line: 3},
			}, false,
		},
		{
			"DOT", ".",
			[]*token.Token{{Type: token.DOT, Line: 1}}, false,
		},
		{
			"COMMA", ",",
			[]*token.Token{{Type: token.COMMA, Line: 1}}, false,
		},
		{
			"EQUAL", "=",
			[]*token.Token{{Type: token.EQUAL, Line: 1}}, false,
		},
		{
			"PLUS", "+",
			[]*token.Token{{Type: token.PLUS, Line: 1}}, false,
		},
		{
			"MINUS", "-",
			[]*token.Token{{Type: token.MINUS, Line: 1}}, false,
		},
		{
			"ASTERISK", "*",
			[]*token.Token{{Type: token.ASTERISK, Line: 1}}, false,
		},
		{
			"SLASH", "/",
			[]*token.Token{{Type: token.SLASH, Line: 1}}, false,
		},
		{
			"NUMBER", "12345",
			[]*token.Token{{Type: token.NUMBER, Value: 12345, Line: 1}}, false,
		},
		{
			"NUMBER OUT-OF-RANGE", "10101011010101010101010101010101011001001101010",
			nil, true,
		},
		{
			"IDENTIFIER", "HELLO",
			[]*token.Token{{Type: token.IDENTIFIER, Line: 1}}, false,
		},
		{
			"UNEXPECTED-TOKEN", "?",
			nil, true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := ScanProgram(tc.source)
			if err != nil && !tc.wantErr {
				t.Errorf("ScanProgram.err = %s\n", err)
			}
			if err == nil && tc.wantErr {
				t.Errorf("expected error but got nil\n")
			}
			if !reflect.DeepEqual(tokens, tc.want) {
				t.Errorf("ScanProgram returned unexpected result\n")
				for i := range len(tokens) {
					if tokens[i] != tc.want[i] {
						t.Errorf("(index = %d).Value actual %s != want %s\n", i, tokens[i].Value, tc.want[i].Value)
						t.Errorf("(index = %d).Line actual %d != want %d\n", i, tokens[i].Line, tc.want[i].Line)
					}
				}
			}
		})
	}
}
