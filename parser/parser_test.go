package parser

import (
	"reflect"
	"testing"

	"piled/expr"
	"piled/token"
)

func tok(t token.Type, v any) *token.Token {
	return &token.Token{Type: t, Value: v}
}
func tokt(t token.Type) *token.Token {
	return &token.Token{Type: t}
}

func lit(v any) *expr.Literal {
	return &expr.Literal{Value: v}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		in   []*token.Token
		want *expr.List
	}{
		{
			name: "nothing",
			in: []*token.Token{
				tokt(token.LPAREN),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &expr.List{},
		},
		{
			name: "one value",
			in: []*token.Token{
				tokt(token.LPAREN),
				tok(token.NUMBER, 69),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &expr.List{Elements: []*expr.Literal{lit(69)}},
		},
		{
			name: "several values",
			in: []*token.Token{
				tokt(token.LPAREN),
				tok(token.NUMBER, 100),
				tok(token.NUMBER, 200),
				tok(token.NUMBER, 300),
				tok(token.NUMBER, 400),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &expr.List{
				Elements: []*expr.Literal{
					lit(100), lit(200), lit(300), lit(400),
				},
			},
		},
		{
			name: "print 99",
			in: []*token.Token{
				tokt(token.LPAREN),
				tok(token.NUMBER, 69),
				tok(token.IDENTIFIER, "print"),
				tokt(token.RPAREN),
				tokt(token.EOF),
			},
			want: &expr.List{
				Elements: []*expr.Literal{
					lit(69),
					lit("print"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.in)
			got, err := p.Parse()
			if err != nil {
				t.Errorf("Parse() error = %v", err)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want = %v", tt.want)
				t.Errorf(" got = %v", got)
			}
		})
	}
}
