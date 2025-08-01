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

