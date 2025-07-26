package lexer

import (
	"reflect"
	"testing"

	"piled/token"
)

func TestLexProgram(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		want    []token.Token
		wantErr bool
	}{
		{
			"Parentheses", "()",
			[]token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.EOF, Line: 1},
			}, false,
		},
		{
			"Parentheses with a whitespace", "( )",
			[]token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.EOF, Line: 1},
			}, false,
		},
		{
			"Parentheses with an identifier", "(hello)",
			[]token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.IDENT, Literal: "hello", Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.EOF, Line: 1},
			}, false,
		},
		{
			"Two parentheses", "()()",
			[]token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.EOF, Line: 1},
			}, false,
		},
		{
			"Multiple lines", "()\n()\n()\n",
			[]token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.LPAREN, Line: 2},
				{Type: token.RPAREN, Line: 2},
				{Type: token.LPAREN, Line: 3},
				{Type: token.RPAREN, Line: 3},
				{Type: token.EOF, Line: 4},
			}, false,
		},
		{
			"Multiple lines with various type of tokens", "(I)\n(like)\n(golang)\n",
			[]token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.IDENT, Literal: "I", Line: 1},
				{Type: token.RPAREN, Line: 1},

				{Type: token.LPAREN, Line: 2},
				{Type: token.IDENT, Literal: "like", Line: 2},
				{Type: token.RPAREN, Line: 2},

				{Type: token.LPAREN, Line: 3},
				{Type: token.IDENT, Literal: "golang", Line: 3},
				{Type: token.RPAREN, Line: 3},
				{Type: token.EOF, Line: 4},
			}, false,
		},
		{
			"NUMBER", "12345",
			[]token.Token{{Type: token.NUMBER, Literal: "12345", Line: 1}, {Type: token.EOF, Line: 1}}, false,
		},
		{
			"TWO NUMBERS", "69 420",
			[]token.Token{
				{Type: token.NUMBER, Literal: "69", Line: 1},
				{Type: token.NUMBER, Literal: "420", Line: 1},
				{Type: token.EOF, Line: 1},
			}, false,
		},
		{
			"IDENT", "HELLO",
			[]token.Token{
				{Type: token.IDENT, Literal: "HELLO", Line: 1},
				{Type: token.EOF, Line: 1},
			}, false,
		},
		{
			"DOUBLE IDENT", "HELLO HELLO",
			[]token.Token{
				{Type: token.IDENT, Literal: "HELLO", Line: 1},
				{Type: token.IDENT, Literal: "HELLO", Line: 1},
				{Type: token.EOF, Line: 1},
			}, false,
		},
		{
			"PRINT NUMBER", "(69 print)",
			[]token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.NUMBER, Literal: "69", Line: 1},
				{Type: token.PRINT, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.EOF, Line: 1},
			}, false,
		},
		{
			"PRINT NESTED", "((3 5 +) print)",
			[]token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.LPAREN, Line: 1},
				{Type: token.NUMBER, Literal: "3", Line: 1},
				{Type: token.NUMBER, Literal: "5", Line: 1},
				{Type: token.ADD, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.PRINT, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.EOF, Line: 1},
			}, false,
		},
		{
			"UNEXPECTED-TOKEN", "?",
			nil, true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tokens, err := LexProgram(tc.source)
			if tc.wantErr {
				if err == nil {
					t.Errorf("LexProgram() error = %v, wantErr = %v", err, tc.wantErr)
					return
				}
			} else {
				if err != nil {
					t.Errorf("LexProgram() error = %v, wantErr = %v", err, tc.wantErr)
					return
				}
			}
			if !reflect.DeepEqual(tokens, tc.want) {
				t.Errorf("LexProgram doesn't returned a result which we expect")
				if len(tokens) != len(tc.want) {
					t.Errorf("length of tokens = %d, but want = %d", len(tokens), len(tc.want))
				}
				for i := range len(tokens) {
					t.Errorf("(index = %d).Type actual %s != want %s", i, tokens[i].Type, tc.want[i].Type)
					t.Errorf("(index = %d).Literal actual %s != want %s", i, tokens[i].Literal, tc.want[i].Literal)
					t.Errorf("(index = %d).Line actual %d != want %d", i, tokens[i].Line, tc.want[i].Line)
				}
			}
		})
	}
}
