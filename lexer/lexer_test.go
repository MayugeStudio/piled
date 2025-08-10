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
		name string
		in   string
		want token.Token
	}{
		{
			"EOF", "",
			tok(token.EOF, "", 1),
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
			"binary-operators-mul", "*",
			tok(token.MUL, "*", 1),
		},
		{
			"binary-operators-div", "/",
			tok(token.DIV, "/", 1),
		},
		{
			"binary-operators-modulo", "%",
			tok(token.MOD, "%", 1),
		},
		{
			"binary-operators-and", "&",
			tok(token.AND, "&", 1),
		},
		{
			"binary-operators-or", "|",
			tok(token.OR, "|", 1),
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
			"ident-print", "print",
			tok(token.PRINT, "print", 1),
		},
		{
			"ident-shl", "shl",
			tok(token.SHL, "shl", 1),
		},
		{
			"ident-shr", "shr",
			tok(token.SHR, "shr", 1),
		},
		{
			"controlflow-if", "if",
			tok(token.IF, "if", 1),
		},
		{
			"controlflow-else", "else",
			tok(token.ELSE, "else", 1),
		},
		{
			"controlflow-end", "end",
			tok(token.END, "end", 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.in)
			tok := l.NextToken()
			if !reflect.DeepEqual(tok, tt.want) {
				t.Errorf("got = %v, want = %v\n", tok, tt.want)
			}
		})
	}
}

func TestLexerNextTokenMultiple(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []token.Token
	}{
		{
			"two-items",
			"1234 print",
			[]token.Token{
				tok(token.NUMBER, "1234", 1),
				tok(token.PRINT, "print", 1),
			},
		},
		{
			"three-items",
			"1 1 +",
			[]token.Token{
				tok(token.NUMBER, "1", 1),
				tok(token.NUMBER, "1", 1),
				tok(token.ADD, "+", 1),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.in)
			tokens := make([]token.Token, 0)
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				tokens = append(tokens, tok)
			}

			if !reflect.DeepEqual(tokens, tt.want) {
				t.Errorf("got = %v, want = %v\n", tokens, tt.want)
			}
		})
	}
}
