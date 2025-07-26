package compiler

import (
	"piled/parser"
	"piled/runtime"
	"piled/token"
	"testing"
)

func tok(t token.Type) token.Token {
	return token.Token{Type: t}
}

func TestCompiler_Compile(t *testing.T) {
	tests := []struct {
		name string
		in   parser.Expr
		want []runtime.OPCode
	}{
		{
			name: "single literal",
			in:   &parser.Literal{Value: 42},
			want: []runtime.OPCode{runtime.PUSH, 42},
		},
		{
			name: "simple add",
			in: &parser.List{
				Elements: []parser.Expr{
					&parser.Literal{Value: 20},
					&parser.Literal{Value: 10},
					&parser.Symbol{Token: tok(token.ADD)},
				},
			},
			want: []runtime.OPCode{
				runtime.PUSH, 20,
				runtime.PUSH, 10,
				runtime.ADD,
			},
		},
		{
			name: "simple sub",
			in: &parser.List{
				Elements: []parser.Expr{
					&parser.Literal{Value: 20},
					&parser.Literal{Value: 10},
					&parser.Symbol{Token: tok(token.SUB)},
				},
			},
			want: []runtime.OPCode{
				runtime.PUSH, 20,
				runtime.PUSH, 10,
				runtime.SUB,
			},
		},
		{
			name: "nested expression with print",
			in: &parser.List{
				Elements: []parser.Expr{
					&parser.List{
						Elements: []parser.Expr{
							&parser.Literal{Value: 10},
							&parser.Literal{Value: 20},
							&parser.Symbol{Token: tok(token.ADD)},
						},
					},
					&parser.Symbol{Token: tok(token.PRINT)},
				},
			},
			want: []runtime.OPCode{
				runtime.PUSH, 10,
				runtime.PUSH, 20,
				runtime.ADD,
				runtime.PRINT,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			got, err := c.Compile(tt.in)
			if err != nil {
				t.Fatalf("Compile() error = %v", err)
			}
			if !equalSlices(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func equalSlices(a, b []runtime.OPCode) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
