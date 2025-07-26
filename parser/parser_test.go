package parser

import (
	"testing"

	"piled/token"
)

func tok(t token.Type, l string) token.Token {
	return token.Token{Type: t, Literal: l}
}
func tokt(t token.Type) token.Token {
	return token.Token{Type: t}
}

func lit(v int) *Literal {
	return &Literal{Value: v}
}

func sym(t token.Token) *Symbol {
	return &Symbol{Token: t}
}

func list(elems ...Expr) *List {
	return &List{Elements: elems}
}

func TestParseExpr(t *testing.T) {
	tests := []struct {
		name string
		in   []token.Token
		want Expr
	}{
		{
			name: "nothing",
			in: []token.Token{
				tokt(token.LPAREN),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &List{},
		},
		{
			name: "one value",
			in: []token.Token{
				tokt(token.LPAREN),
				tok(token.NUMBER, "69"),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &List{Elements: []Expr{lit(69)}},
		},
		{
			name: "several values",
			in: []token.Token{
				tokt(token.LPAREN),
				tok(token.NUMBER, "100"),
				tok(token.NUMBER, "200"),
				tok(token.NUMBER, "300"),
				tok(token.NUMBER, "400"),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &List{
				Elements: []Expr{
					lit(100), lit(200), lit(300), lit(400),
				},
			},
		},
		{
			name: "print 99",
			in: []token.Token{
				tokt(token.LPAREN),
				tok(token.NUMBER, "69"),
				tokt(token.PRINT),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &List{
				Elements: []Expr{
					lit(69),
					sym(tokt(token.PRINT)),
				},
			},
		},
		{
			name: "((1 1 +) print)",
			in: []token.Token{
				tokt(token.LPAREN),
				tokt(token.LPAREN),
				tok(token.NUMBER, "1"),
				tok(token.NUMBER, "1"),
				tokt(token.ADD),
				tokt(token.RPAREN),
				tokt(token.PRINT),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &List{
				Elements: []Expr{
					list(
						lit(1),
						lit(1),
						sym(tokt(token.ADD)),
					),
					sym(tokt(token.PRINT)),
				},
			},
		},
		{
			name: "((1 1 +) (1 1 +) print)",
			in: []token.Token{
				tokt(token.LPAREN),
				tokt(token.LPAREN),
				tok(token.NUMBER, "1"),
				tok(token.NUMBER, "1"),
				tokt(token.ADD),
				tokt(token.RPAREN),
				tokt(token.LPAREN),
				tok(token.NUMBER, "1"),
				tok(token.NUMBER, "1"),
				tokt(token.ADD),
				tokt(token.RPAREN),
				tokt(token.PRINT),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &List{
				Elements: []Expr{
					list(
						lit(1),
						lit(1),
						sym(tokt(token.ADD)),
					),
					list(
						lit(1),
						lit(1),
						sym(tokt(token.ADD)),
					),
					sym(tokt(token.PRINT)),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.in)
			got, err := p.ParseExpr()
			if err != nil {
				t.Errorf("ParseExpr() error = %v", err)
				return
			}
			if !tt.want.Equal(got) {
				t.Errorf("want = %v", tt.want)
				t.Errorf(" got = %v", got)
			}
		})
	}
}
