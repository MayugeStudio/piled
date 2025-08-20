package codegen

import (
	"reflect"
	"testing"

	"piled/ir"
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
		in   ir.Op
		want runtime.Instruction
	}{
		{
			name: "single literal",
			in:   &ir.Number{Value: 42},
			want: inst(runtime.PUSH, 42),
		},
		{
			name: "add",
			in:   &ir.Binop{Bkind: ir.Add},
			want: inst(runtime.ADD),
		},
		{
			name: "sub",
			in:   &ir.Binop{Bkind: ir.Sub},
			want: inst(runtime.SUB),
		},
		{
			name: "mul",
			in:   &ir.Binop{Bkind: ir.Mul},
			want: inst(runtime.MUL),
		},
		{
			name: "div",
			in:   &ir.Binop{Bkind: ir.Div},
			want: inst(runtime.DIV),
		},
		{
			name: "mod",
			in:   &ir.Binop{Bkind: ir.Mod},
			want: inst(runtime.MOD),
		},
		{
			name: "and",
			in:   &ir.Binop{Bkind: ir.And},
			want: inst(runtime.AND),
		},
		{
			name: "or",
			in:   &ir.Binop{Bkind: ir.Or},
			want: inst(runtime.OR),
		},
		{
			name: "shift-left",
			in:   &ir.Binop{Bkind: ir.Shl},
			want: inst(runtime.SHL),
		},
		{
			name: "shift-right",
			in:   &ir.Binop{Bkind: ir.Shr},
			want: inst(runtime.SHR),
		},
		{
			name: "gt",
			in:   &ir.Binop{Bkind: ir.Gt},
			want: inst(runtime.GT),
		},
		{
			name: "lt",
			in:   &ir.Binop{Bkind: ir.Lt},
			want: inst(runtime.LT),
		},
		{
			name: "eq",
			in:   &ir.Binop{Bkind: ir.Eq},
			want: inst(runtime.EQ),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateProgram([]ir.Op{tt.in})[0]

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
		in   ir.Op
		want runtime.Instruction
	}{
		{
			name: "dup",
			in:   &ir.Dup{},
			want: inst(runtime.DUP),
		},
		{
			name: "dup2",
			in:   &ir.Dup2{},
			want: inst(runtime.DUP2),
		},
		{
			name: "swap",
			in:   &ir.Swap{},
			want: inst(runtime.SWAP),
		},
		{
			name: "rot",
			in:   &ir.Rot{},
			want: inst(runtime.ROT),
		},
		{
			name: "drop",
			in:   &ir.Drop{},
			want: inst(runtime.DROP),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			program := GenerateProgram([]ir.Op{tt.in})

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
		in   []ir.Op
		want []runtime.Instruction
	}{
		{
			name: "if",
			in: []ir.Op{
				&ir.JmpIfNotLabel{Label: 0},
				&ir.Number{Value: 1},
				&ir.Label{Label: 0},
			},
			want: []runtime.Instruction{
				inst(runtime.JMPIF, 2),
				inst(runtime.PUSH, 1),
				inst(runtime.NOP),
			},
		},
		{
			name: "if-else",
			in: []ir.Op{
				&ir.JmpIfNotLabel{Label: 0},
				&ir.Number{Value: 1},
				&ir.JmpLabel{Label: 1},
				&ir.Label{Label: 0},
				&ir.Number{Value: 0},
				&ir.Label{Label: 1},
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
