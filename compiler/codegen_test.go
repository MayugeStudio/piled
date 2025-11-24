package compiler

import (
	"reflect"
	"testing"

	"piled/runtime"
)

func inst(k runtime.InstructionKind, args ...int) runtime.Instruction {
	if len(args) == 0 {
		return runtime.Instruction{Kind: k, Args: nil}
	}
	return runtime.Instruction{Kind: k, Args: args}
}

func TestCodegen_Compile(t *testing.T) {
	// TODO: Separate TestCodegen_Compile from Binop tests
	tests := []struct {
		name string
		in   Op
		want runtime.Instruction
	}{
		{
			name: "single literal",
			in:   &Number{Value: 42},
			want: inst(runtime.PUSH, 42),
		},
		{
			name: "add",
			in:   &Binop{Bkind: Add},
			want: inst(runtime.ADD),
		},
		{
			name: "sub",
			in:   &Binop{Bkind: Sub},
			want: inst(runtime.SUB),
		},
		{
			name: "mul",
			in:   &Binop{Bkind: Mul},
			want: inst(runtime.MUL),
		},
		{
			name: "div",
			in:   &Binop{Bkind: Div},
			want: inst(runtime.DIV),
		},
		{
			name: "mod",
			in:   &Binop{Bkind: Mod},
			want: inst(runtime.MOD),
		},
		{
			name: "and",
			in:   &Binop{Bkind: And},
			want: inst(runtime.AND),
		},
		{
			name: "or",
			in:   &Binop{Bkind: Or},
			want: inst(runtime.OR),
		},
		{
			name: "shift-left",
			in:   &Binop{Bkind: Shl},
			want: inst(runtime.SHL),
		},
		{
			name: "shift-right",
			in:   &Binop{Bkind: Shr},
			want: inst(runtime.SHR),
		},
		{
			name: "gt",
			in:   &Binop{Bkind: Gt},
			want: inst(runtime.GT),
		},
		{
			name: "lt",
			in:   &Binop{Bkind: Lt},
			want: inst(runtime.LT),
		},
		{
			name: "eq",
			in:   &Binop{Bkind: Eq},
			want: inst(runtime.EQ),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateProgram([]Op{tt.in})[0]

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
				return
			}
		})
	}
}

func TestCodegen_Compile_Stack(t *testing.T) {
	tests := []struct {
		name string
		in   Op
		want runtime.Instruction
	}{
		{
			name: "dup",
			in:   &Dup{},
			want: inst(runtime.DUP),
		},
		{
			name: "swap",
			in:   &Swap{},
			want: inst(runtime.SWAP),
		},
		{
			name: "rot",
			in:   &Rot{},
			want: inst(runtime.ROT),
		},
		{
			name: "drop",
			in:   &Drop{},
			want: inst(runtime.DROP),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			program := GenerateProgram([]Op{tt.in})

			if len(program) == 0 {
				t.Errorf("got = %v, want = %v", program, tt.want)
			}

			got := program[0]

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
				return
			}
		})
	}
}

func TestCodegen_Compile_ControlFlow(t *testing.T) {
	tests := []struct {
		name string
		in   []Op
		want []runtime.Instruction
	}{
		{
			name: "if",
			in: []Op{
				&JmpIfNotLabel{Label: 0},
				&Number{Value: 1},
				&Label{Label: 0},
			},
			want: []runtime.Instruction{
				inst(runtime.JMPIF, 2),
				inst(runtime.PUSH, 1),
				inst(runtime.NOP),
			},
		},
		{
			name: "if-else",
			in: []Op{
				&JmpIfNotLabel{Label: 0},
				&Number{Value: 1},
				&JmpLabel{Label: 1},
				&Label{Label: 0},
				&Number{Value: 0},
				&Label{Label: 1},
			},
			want: []runtime.Instruction{
				inst(runtime.JMPIF, 3),
				inst(runtime.PUSH, 1),
				inst(runtime.JMP, 5),
				inst(runtime.NOP),
				inst(runtime.PUSH, 0),
				inst(runtime.NOP),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateProgram(tt.in)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want = %v", got, tt.want)
				return
			}
		})
	}
}
