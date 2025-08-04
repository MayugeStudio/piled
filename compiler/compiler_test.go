package compiler

import (
	"reflect"
	"testing"

	"piled/runtime"
	"piled/token"
	"piled/lexer"
)

func tok(t token.Type) token.Token {
	return token.Token{Type: t}
}

func TestCompiler_Compile(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []runtime.OPCode
	}{
		{
			name: "single literal",
			in:   "42",
			want: []runtime.OPCode{runtime.PUSH, 42},
		},
		{
			name: "simple add",
			in: "20 10 +",
			want: []runtime.OPCode{
				runtime.PUSH, 20,
				runtime.PUSH, 10,
				runtime.ADD,
			},
		},
		{
			name: "simple sub",
			in: "20 10 -",
			want: []runtime.OPCode{
				runtime.PUSH, 20,
				runtime.PUSH, 10,
				runtime.SUB,
			},
		},
		{
			name: "simple mul",
			in: "20 10 *",
			want: []runtime.OPCode{
				runtime.PUSH, 20,
				runtime.PUSH, 10,
				runtime.MUL,
			},
		},
		{
			name: "simple div",
			in: "20 10 /",
			want: []runtime.OPCode{
				runtime.PUSH, 20,
				runtime.PUSH, 10,
				runtime.DIV,
			},
		},
		{
			name: "simple mod",
			in: "20 10 %",
			want: []runtime.OPCode{
				runtime.PUSH, 20,
				runtime.PUSH, 10,
				runtime.MOD,
			},
		},
		{
			name: "shift-left",
			in: "10 2 shl",
			want: []runtime.OPCode{
				runtime.PUSH, 10,
				runtime.PUSH, 2,
				runtime.SHL,
			},
		},
		{
			name: "shift-right",
			in: "10 2 shr",
			want: []runtime.OPCode{
				runtime.PUSH, 10,
				runtime.PUSH, 2,
				runtime.SHR,
			},
		},
		{
			name: "gt",
			in: "20 10 >",
			want: []runtime.OPCode{
				runtime.PUSH, 20,
				runtime.PUSH, 10,
				runtime.GT,
			},
		},
		{
			name: "lt",
			in: "20 10 <",
			want: []runtime.OPCode{
				runtime.PUSH, 20,
				runtime.PUSH, 10,
				runtime.LT,
			},
		},
		{
			name: "eq",
			in: "=",
			want: []runtime.OPCode{
				runtime.EQ,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.in)
			c := New(l)
			got := c.Compile()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
			}
		})
	}
}

