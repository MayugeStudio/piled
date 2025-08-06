package compiler

import (
	"reflect"
	"testing"

	"piled/runtime"
	"piled/lexer"
)

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
			in: "+",
			want: []runtime.OPCode{runtime.ADD},
		},
		{
			name: "simple sub",
			in: "-",
			want: []runtime.OPCode{runtime.SUB},
		},
		{
			name: "simple mul",
			in: "*",
			want: []runtime.OPCode{runtime.MUL},
		},
		{
			name: "simple div",
			in: "/",
			want: []runtime.OPCode{runtime.DIV},
		},
		{
			name: "simple mod",
			in: "%",
			want: []runtime.OPCode{runtime.MOD},
		},
		{
			name: "simple and",
			in: "&",
			want: []runtime.OPCode{runtime.AND},
		},
		{
			name: "simple or",
			in: "|",
			want: []runtime.OPCode{runtime.OR},
		},
		{
			name: "shift-left",
			in: "shl",
			want: []runtime.OPCode{runtime.SHL},
		},
		{
			name: "shift-right",
			in: "shr",
			want: []runtime.OPCode{runtime.SHR},
		},
		{
			name: "gt",
			in: ">",
			want: []runtime.OPCode{runtime.GT},
		},
		{
			name: "lt",
			in: "<",
			want: []runtime.OPCode{runtime.LT},
		},
		{
			name: "eq",
			in: "=",
			want: []runtime.OPCode{runtime.EQ},
		},
		// Control flow
		{
			name: "simple if",
			in: "if 1 end",
			want: []runtime.OPCode{
				runtime.JMPIF, runtime.OPCode(4), // IF 
				runtime.PUSH, runtime.OPCode(1),  // PUSH 1 
				runtime.NOP,                      // END
			},
		},
		{
			name: "simple if-else",
			in: "if 1 else 0 end",
			want: []runtime.OPCode{
				runtime.JMPIF, runtime.OPCode(6), // IF  ELSE_ADDR <<
				runtime.PUSH,  runtime.OPCode(1), // PUSH 1 
				runtime.JMP,   runtime.OPCode(9), // JMP END_ADDR  <<
				runtime.NOP,                      // ELSE
				runtime.PUSH,  runtime.OPCode(0), // PUSH 0
				runtime.NOP,                      // END
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

