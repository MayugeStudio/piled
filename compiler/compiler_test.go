package compiler

import (
	"reflect"
	"testing"

	"piled/lexer"
	"piled/runtime"
)

func inst(k runtime.InstructionKind, args ...int) runtime.Instruction {
	if len(args) == 0 {
		return runtime.Instruction{Kind: k, Args: nil}
	}
	return runtime.Instruction{Kind: k, Args: args}
}

func TestCompiler_Compile(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want runtime.Instruction
	}{
		{
			name: "single literal",
			in:   "42",
			want: inst(runtime.PUSH, 42),
		},
		{
			name: "add",
			in:   "+",
			want: inst(runtime.ADD),
		},
		{
			name: "sub",
			in:   "-",
			want: inst(runtime.SUB),
		},
		{
			name: "mul",
			in:   "*",
			want: inst(runtime.MUL),
		},
		{
			name: "div",
			in:   "/",
			want: inst(runtime.DIV),
		},
		{
			name: "mod",
			in:   "%",
			want: inst(runtime.MOD),
		},
		{
			name: "and",
			in:   "&",
			want: inst(runtime.AND),
		},
		{
			name: "or",
			in:   "|",
			want: inst(runtime.OR),
		},
		{
			name: "shift-left",
			in:   "shl",
			want: inst(runtime.SHL),
		},
		{
			name: "shift-right",
			in:   "shr",
			want: inst(runtime.SHR),
		},
		{
			name: "gt",
			in:   ">",
			want: inst(runtime.GT),
		},
		{
			name: "lt",
			in:   "<",
			want: inst(runtime.LT),
		},
		{
			name: "eq",
			in:   "=",
			want: inst(runtime.EQ),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.in)
			c := New(l)
			got := c.Compile()[0]

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
				return
			}
		})
	}
}

func TestCompiler_Compile_ControlFlow(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []runtime.Instruction
	}{
		{
			name: "if",
			in:   "if 1 end",
			want: []runtime.Instruction{
				inst(runtime.JMPIF, 2),    // IF
				inst(runtime.PUSH, 1),     // PUSH 1
				inst(runtime.NOP),         // END
			},
		},
		{
			name: "if-else",
			in:   "if 1 else 0 end",
			want: []runtime.Instruction{
				inst(runtime.JMPIF, 3),    // IF  ELSE_ADDR <<
				inst(runtime.PUSH, 1),     // PUSH 1
				inst(runtime.JMP, 5),      // JMP END_ADDR  <<
				inst(runtime.NOP),         // ELSE
				inst(runtime.PUSH, 0),     // PUSH 0
				inst(runtime.NOP),         // END
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
				return
			}
		})
	}
}
