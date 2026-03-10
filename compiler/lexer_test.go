package compiler

import (
	"reflect"
	"testing"
)

func tok(t Type, lit string) Token {
	return Token{
		Type:    t,
		Literal: lit,
	}
}

func TestLexerNextToken(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want Token
	}{
		{
			"EOF", "",
			tok(EOF, ""),
		},
		{
			"Comment", "+;this is comment",
			tok(ADD, "+"),
		},
		{
			"binary-operators-add", "+",
			tok(ADD, "+"),
		},
		{
			"binary-operators-sub", "-",
			tok(SUB, "-"),
		},
		{
			"binary-operators-mul", "*",
			tok(MUL, "*"),
		},
		{
			"binary-operators-div", "/",
			tok(DIV, "/"),
		},
		{
			"binary-operators-modulo", "%",
			tok(MOD, "%"),
		},
		{
			"binary-operators-and", "&",
			tok(AND, "&"),
		},
		{
			"binary-operators-or", "|",
			tok(OR, "|"),
		},
		{
			"comparison-operators-gt", ">",
			tok(GT, ">"),
		},
		{
			"comparison-operators-lt", "<",
			tok(LT, "<"),
		},
		{
			"comparison-operators-gt", "=",
			tok(EQ, "="),
		},
		{
			"number", "12",
			tok(NUMBER, "12"),
		},
		{
			"ident-print", "print",
			tok(PRINT, "print"),
		},
		{
			"ident-shl", "shl",
			tok(SHL, "shl"),
		},
		{
			"ident-shr", "shr",
			tok(SHR, "shr"),
		},
		{
			"ident-dup", "dup",
			tok(DUP, "dup"),
		},
		{
			"ident-dup2", "dup2",
			tok(DUP2, "dup2"),
		},
		{
			"ident-swap", "swap",
			tok(SWAP, "swap"),
		},
		{
			"ident-rot", "rot",
			tok(ROT, "rot"),
		},
		{
			"ident-drop", "drop",
			tok(DROP, "drop"),
		},
		{
			"ident-over", "over",
			tok(OVER, "over"),
		},
		{
			"controlflow-if", "if",
			tok(IF, "if"),
		},
		{
			"controlflow-else", "else",
			tok(ELSE, "else"),
		},
		{
			"controlflow-while", "while",
			tok(WHILE, "while"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer("test.piled", tt.in)
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
		want []Token
	}{
		{
			"two-items",
			"1234 print",
			[]Token{
				tok(NUMBER, "1234"),
				tok(PRINT, "print"),
			},
		},
		{
			"three-items",
			"1 1 +",
			[]Token{
				tok(NUMBER, "1"),
				tok(NUMBER, "1"),
				tok(ADD, "+"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer("test.piled", tt.in)
			got := make([]Token, 0)
			for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
				got = append(got, tok)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}

func TestParsePoint(t *testing.T) {
	l := NewLexer("test.piled", "1 2")
	var tok Token

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
