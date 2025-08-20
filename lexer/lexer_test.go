package lexer

import (
	"reflect"
	"testing"

	"piled/token"
)

func tok(t token.Type, lit string) token.Token {
	return token.Token{
		Type:    t,
		Literal: lit,
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
			tok(token.EOF, ""),
		},
		{
			"binary-operators-add", "+",
			tok(token.ADD, "+"),
		},
		{
			"binary-operators-sub", "-",
			tok(token.SUB, "-"),
		},
		{
			"binary-operators-mul", "*",
			tok(token.MUL, "*"),
		},
		{
			"binary-operators-div", "/",
			tok(token.DIV, "/"),
		},
		{
			"binary-operators-modulo", "%",
			tok(token.MOD, "%"),
		},
		{
			"binary-operators-and", "&",
			tok(token.AND, "&"),
		},
		{
			"binary-operators-or", "|",
			tok(token.OR, "|"),
		},
		{
			"comparison-operators-gt", ">",
			tok(token.GT, ">"),
		},
		{
			"comparison-operators-lt", "<",
			tok(token.LT, "<"),
		},
		{
			"comparison-operators-gt", "=",
			tok(token.EQ, "="),
		},
		{
			"number", "12",
			tok(token.NUMBER, "12"),
		},
		{
			"ident-print", "print",
			tok(token.PRINT, "print"),
		},
		{
			"ident-shl", "shl",
			tok(token.SHL, "shl"),
		},
		{
			"ident-shr", "shr",
			tok(token.SHR, "shr"),
		},
		{
			"ident-dup", "dup",
			tok(token.DUP, "dup"),
		},
		{
			"ident-swap", "swap",
			tok(token.SWAP, "swap"),
		},
		{
			"controlflow-if", "if",
			tok(token.IF, "if"),
		},
		{
			"controlflow-else", "else",
			tok(token.ELSE, "else"),
		},
		{
			"controlflow-while", "while",
			tok(token.WHILE, "while"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New("test.piled", tt.in)
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
				tok(token.NUMBER, "1234"),
				tok(token.PRINT, "print"),
			},
		},
		{
			"three-items",
			"1 1 +",
			[]token.Token{
				tok(token.NUMBER, "1"),
				tok(token.NUMBER, "1"),
				tok(token.ADD, "+"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New("test.piled", tt.in)
			got := make([]token.Token, 0)
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				got = append(got, tok)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}

func TestParsePoint(t *testing.T) {
	l := New("test.piled", "1 2")
	var tok token.Token

	savedPoint := l.CurrentPoint

	tok = l.NextToken()
	if tok.Literal != "1" {
		t.Fatalf("got = %q, want = %q", tok.Literal, "1")
	}

	tok = l.NextToken()
	if tok.Literal != "2" {
		t.Fatalf("got = %q, want = %q", tok.Literal, "2")
	}

	// Restore savedPoint
	l.CurrentPoint = savedPoint

	tok = l.NextToken()
	if tok.Literal != "1" {
		t.Fatalf("got = %q, want = %q", tok.Literal, "1")
	}
}
