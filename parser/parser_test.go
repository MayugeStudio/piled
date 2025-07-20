package parser

import (
	"testing"

	"piled/expr"
	"piled/token"
)

func tok(t token.Type, l string) token.Token {
	return token.Token{Type: t, Literal: l}
}
func tokt(t token.Type) token.Token {
	return token.Token{Type: t}
}

func lit(v int) *expr.Literal {
	return &expr.Literal{Value: v}
}

func sym(v string, t expr.SymbolType) *expr.Symbol {
	return &expr.Symbol{Name: v, Type: t}
}

func list(elems ...expr.Expr) *expr.List {
	return &expr.List{Elements: elems}
}

func TestParseExpr(t *testing.T) {
	tests := []struct {
		name string
		in   []token.Token
		want expr.Expr
	}{
		{
			name: "nothing",
			in: []token.Token{
				tokt(token.LPAREN),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &expr.List{},
		},
		{
			name: "one value",
			in: []token.Token{
				tokt(token.LPAREN),
				tok(token.NUMBER, "69"),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &expr.List{Elements: []expr.Expr{lit(69)}},
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
			want: &expr.List{
				Elements: []expr.Expr{
					lit(100), lit(200), lit(300), lit(400),
				},
			},
		},
		{
			name: "print 99",
			in: []token.Token{
				tokt(token.LPAREN),
				tok(token.NUMBER, "69"),
				tok(token.IDENT, "print"),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &expr.List{
				Elements: []expr.Expr{
					lit(69),
					sym("print", expr.SymbolPrint),
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
				tok(token.IDENT, "+"),
				tokt(token.RPAREN),
				tok(token.IDENT, "print"),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &expr.List{
				Elements: []expr.Expr{
					list(
						lit(1),
						lit(1),
						sym("+", expr.SymbolAdd),
					),
					sym("print", expr.SymbolPrint),
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
				tok(token.IDENT, "+"),
				tokt(token.RPAREN),
				tokt(token.LPAREN),
				tok(token.NUMBER, "1"),
				tok(token.NUMBER, "1"),
				tok(token.IDENT, "+"),
				tokt(token.RPAREN),
				tok(token.IDENT, "print"),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &expr.List{
				Elements: []expr.Expr{
					list(
						lit(1),
						lit(1),
						sym("+", expr.SymbolAdd),
					),
					list(
						lit(1),
						lit(1),
						sym("+", expr.SymbolAdd),
					),
					sym("print", expr.SymbolPrint),
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
