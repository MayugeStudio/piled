package compiler

import (
	"piled/parser"
	"piled/opcode"
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
		want []opcode.Code
	}{
		{
			name: "single literal",
			in:   &parser.Literal{Value: 42},
			want: []opcode.Code{opcode.PUSH, 42},
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
			want: []opcode.Code{
				opcode.PUSH, 20,
				opcode.PUSH, 10,
				opcode.ADD,
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
			want: []opcode.Code{
				opcode.PUSH, 20,
				opcode.PUSH, 10,
				opcode.SUB,
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
			want: []opcode.Code{
				opcode.PUSH, 10,
				opcode.PUSH, 20,
				opcode.ADD,
				opcode.PRINT,
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

func equalSlices(a, b []opcode.Code) bool {
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
