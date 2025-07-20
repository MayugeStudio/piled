package compiler

import (
	"piled/expr"
	"piled/opcode"
	"testing"
)

func TestCompiler_Compile(t *testing.T) {
	tests := []struct {
		name string
		in   expr.Expr
		want []int
	}{
		{
			name: "single literal",
			in:   &expr.Literal{Value: 42},
			want: []int{opcode.PUSH, 42},
		},
		{
			name: "simple add",
			in: &expr.List{
				Elements: []expr.Expr{
					&expr.Literal{Value: 20},
					&expr.Literal{Value: 10},
					&expr.Symbol{Name: "+"},
				},
			},
			want: []int{
				opcode.PUSH, 20,
				opcode.PUSH, 10,
				opcode.ADD,
			},
		},
		{
			name: "simple sub",
			in: &expr.List{
				Elements: []expr.Expr{
					&expr.Literal{Value: 20},
					&expr.Literal{Value: 10},
					&expr.Symbol{Name: "-"},
				},
			},
			want: []int{
				opcode.PUSH, 20,
				opcode.PUSH, 10,
				opcode.SUB,
			},
		},
		{
			name: "nested expression with print",
			in: &expr.List{
				Elements: []expr.Expr{
					&expr.List{
						Elements: []expr.Expr{
							&expr.Literal{Value: 10},
							&expr.Literal{Value: 20},
							&expr.Symbol{Name: "+"},
						},
					},
					&expr.Symbol{Name: "print"},
				},
			},
			want: []int{
				opcode.PUSH, 10,
				opcode.PUSH, 20,
				opcode.ADD,
				opcode.PRINT,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCompiler()
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

func equalSlices(a, b []int) bool {
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
