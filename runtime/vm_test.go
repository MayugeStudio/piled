package runtime

import (
	"bytes"
	"testing"
)

func TestVM_Run(t *testing.T) {
	tests := []struct {
		name   string
		code   []OPCode
		want   []int
		output string
	}{
		{
			"add",
			[]OPCode{
				PUSH, 10,
				PUSH, 20,
				ADD,
			},
			[]int{30}, "",
		},
		{
			"sub",
			[]OPCode{
				PUSH, 50,
				PUSH, 20,
				SUB,
			},
			[]int{30}, "",
		},
		{
			"mul",
			[]OPCode{
				PUSH, 5,
				PUSH, 2,
				MUL,
			},
			[]int{10}, "",
		},
		{
			"div",
			[]OPCode{
				PUSH, 10,
				PUSH, 2,
				DIV,
			},
			[]int{5}, "",
		},
		{
			"mod",
			[]OPCode{
				PUSH, 9,
				PUSH, 2,
				MOD,
				PUSH, 10,
				PUSH, 2,
				MOD,
			},
			[]int{1, 0}, "",
		},
		{
			"and",
			[]OPCode{
				PUSH, 5, // 0101
				PUSH, 14, // 1110
				AND,
			},
			[]int{4}, "",
		},
		{
			"or",
			[]OPCode{
				PUSH, 6, // 0110
				PUSH, 10, // 1010
				OR,
			},
			[]int{14}, "",
		},
		{
			"left-shift",
			[]OPCode{
				PUSH, 2,
				PUSH, 1,
				SHL,
			},
			[]int{4}, "",
		},
		{
			"right-shift",
			[]OPCode{
				PUSH, 4,
				PUSH, 1,
				SHR,
			},
			[]int{2}, "",
		},
		{
			"add and print",
			[]OPCode{
				PUSH, 5,
				PUSH, 7,
				ADD,
				PRINT,
			},
			[]int{}, "12\n",
		},
		{
			"jmp",
			[]OPCode{
				PUSH, 0,
				PUSH, 1,
				JMP,  8,
				PUSH, 34,
				PUSH, 2,
			},
			[]int{0, 1, 2}, "",
		},
		{
			"jmpif",
			[]OPCode{
				PUSH,  0,
				JMPIF, 5,
				PUSH,  35,
				NOP,
				PUSH,  34,
			},
			[]int{34}, "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: Introduce vm constructor
			vm := &VM{code: tt.code}

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
