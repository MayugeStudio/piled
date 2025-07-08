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
			"Multiple lines with various type of tokens", "(I)\n(like)\n(golang)\n",
			[]*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.IDENTIFIER, Value: "I", Line: 1},
				{Type: token.RPAREN, Line: 1},

				{Type: token.LPAREN, Line: 2},
				{Type: token.IDENTIFIER, Value: "like", Line: 2},
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
			"IDENTIFIER", "HELLO",
			[]*token.Token{{Type: token.IDENTIFIER, Value: "HELLO", Line: 1}}, false,
		},
		{
			"STRING", "\"MYSTRING\"",
			[]*token.Token{{Type: token.STRING, Value: "MYSTRING", Line: 1}}, false,
		},
		{
			"UNTERMINATED-STRING", "\"A+B-C*D/E=F.G,H",
			nil, true,
		},
		{
			"UNEXPECTED-TOKEN", "?",
			nil, true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := ScanProgram(tc.source)
			if tc.wantErr {
				if err == nil {
					t.Errorf("ScanProgram() error = %v, wantErr = %v", err, tc.wantErr)
					return
				}
			} else {
				if err != nil {
					t.Errorf("ScanProgram() error = %v, wantErr = %v", err, tc.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(tokens, tc.want) {
				t.Errorf("ScanProgram doesn't returned a result which we expect")
				if len(tokens) != len(tc.want) {
					t.Errorf("length of tokens = %d, but want = %d", len(tokens), len(tc.want))
				}
				for i := range len(tokens) {
					t.Errorf("(index = %d).Type actual %d != want %d", i, tokens[i].Type, tc.want[i].Type)
					t.Errorf("(index = %d).Value actual %s != want %s", i, tokens[i].Value, tc.want[i].Value)
					t.Errorf("(index = %d).Line actual %d != want %d", i, tokens[i].Line, tc.want[i].Line)
				}
			}
		})
	}
}
