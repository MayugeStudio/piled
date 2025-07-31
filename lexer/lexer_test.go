package lexer

import (
	"reflect"
	"testing"

	"piled/token"
)


func tok(t token.Type, lit string, l int) token.Token {
	return token.Token{
		Type:    t,
		Literal: lit,
		Line:    l,
	}
}

func TestLexerNextToken(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    token.Token
	}{
		{
			"EOF", "",
			tok(token.EOF, "", 1),
		},
		{
			"left-parenthesis", "(",
			tok(token.LPAREN, "(", 1),
		},
		{
			"right-parenthesis", ")",
			tok(token.RPAREN, ")", 1),
		},
		{
			"binary-operators-add", "+",
			tok(token.ADD, "+", 1),
		},
		{
			"binary-operators-sub", "-",
			tok(token.SUB, "-", 1),
		},
		{
			"comparison-operators-gt", ">",
			tok(token.GT, ">", 1),
		},
		{
			"comparison-operators-lt", "<",
			tok(token.LT, "<", 1),
		},
		{
			"comparison-operators-gt", "=",
			tok(token.EQ, "=", 1),
		},
		{
			"number", "12",
			tok(token.NUMBER, "12", 1),
		},
		{
			"identifier", "print",
			tok(token.IDENT, "print", 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.in)
			//tokens := make([]token.Token, 0, 0)
		    tok := l.NextToken()
		    if !reflect.Equal(tok, tt.want) {
				t.Errorf("got = %v, want = %v\n", tok, tt.want)
			}
		})
	}
}

func TestLexerNextTokenMultiple(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    []token.Token
	}{
		{
			"two-items + EOF",
			"1234 print",
			{
				tok(token.NUMBER, "1234", 1),
				tok(token.IDENT, "print", 1),
				tok(token.EOF, "", 1),
			},
		},
		{
			"three-items + EOF",
			"1 1 +",
			{
				tok(token.NUMBER, "1", 1),
				tok(token.NUMBER, "1", 1),
				tok(token.ADD, "+", 1),
				tok(token.EOF, "", 1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.in)
			tokens := make([]token.Token, 0, 0)
		    for  {
				tok := l.NextToken()
				tokens = append(tokens, tok)
				if tok.Type == token.EOF {
					break
				}
			}

		    if !reflect.Equal(tokens, tt.want) {
				t.Errorf("got = %v, want = %v\n", tokens, tt.want)
			}
		})
	}
}

