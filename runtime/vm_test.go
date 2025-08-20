package runtime

import (
	"bytes"
	"testing"
)

func inst(k InstructionKind, args ...int) Instruction {
	if len(args) == 0 {
		return Instruction{Kind: k, Args: nil}
	}
	return Instruction{Kind: k, Args: args}
}

func TestVM_Run(t *testing.T) {
	tests := []struct {
		name   string
		code   []Instruction
		want   []int
		output string
	}{
		{
			"add",
			[]Instruction{
				inst(PUSH, 10),
				inst(PUSH, 20),
				inst(ADD),
			},
			[]int{30}, "",
		},
		{
			"sub",
			[]Instruction{
				inst(PUSH, 50),
				inst(PUSH, 20),
				inst(SUB),
			},
			[]int{30}, "",
		},
		{
			"mul",
			[]Instruction{
				inst(PUSH, 5),
				inst(PUSH, 2),
				inst(MUL),
			},
			[]int{10}, "",
		},
		{
			"div",
			[]Instruction{
				inst(PUSH, 10),
				inst(PUSH, 2),
				inst(DIV),
			},
			[]int{5}, "",
		},
		{
			"mod",
			[]Instruction{
				inst(PUSH, 9),
				inst(PUSH, 2),
				inst(MOD),
				inst(PUSH, 10),
				inst(PUSH, 2),
				inst(MOD),
			},
			[]int{1, 0}, "",
		},
		{
			"and",
			[]Instruction{
				inst(PUSH, 5),
				inst(PUSH, 14),
				inst(AND),
			},
			[]int{4}, "",
		},
		{
			"or",
			[]Instruction{
				inst(PUSH, 6),
				inst(PUSH, 10),
				inst(OR),
			},
			[]int{14}, "",
		},
		{
			"left-shift",
			[]Instruction{
				inst(PUSH, 2),
				inst(PUSH, 1),
				inst(SHL),
			},
			[]int{4}, "",
		},
		{
			"right-shift",
			[]Instruction{
				inst(PUSH, 4),
				inst(PUSH, 1),
				inst(SHR),
			},
			[]int{2}, "",
		},
		{
			"add and print",
			[]Instruction{
				inst(PUSH, 5),
				inst(PUSH, 7),
				inst(ADD),
				inst(PRINT),
			},
			[]int{}, "12\n",
		},
		{
			"jmp",
			[]Instruction{
				inst(PUSH, 0),
				inst(PUSH, 1),
				inst(JMP,  4),
				inst(PUSH, 34),
				inst(PUSH, 2),
			},
			[]int{0, 1, 2}, "",
		},
		{
			"jmpif",
			[]Instruction{
				inst(PUSH,  0),
				inst(JMPIF, 3),
				inst(PUSH,  35),
				inst(PUSH,  34),
			},
			[]int{34}, "",
		},
		{
			"dup",
			[]Instruction{
				inst(PUSH, 69),
				inst(DUP),
			},
			[]int{69, 69}, "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := NewVM(tt.code)

			buf := &bytes.Buffer{}
			vm.stdout = buf
			vm.Run()

			if !equalSlices(vm.stack, tt.want) {
				t.Errorf("stack = %v, want = %v", vm.stack, tt.want)
			}

			if buf.String() != tt.output {
				t.Errorf("output = %q, want = %q", buf.String(), tt.output)
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
