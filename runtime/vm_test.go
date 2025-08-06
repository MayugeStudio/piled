package runtime

import (
	"bytes"
	"testing"
)

// TODO: TestVM_Run should be only care about OPCode but not token such as if or else.
func TestVM_Run(t *testing.T) {
	tests := []struct {
		name   string
		code   []OPCode
		want   []int
		output string
	}{
		{
			name: "simple add",
			code: []OPCode{
				PUSH, 10,
				PUSH, 20,
				ADD,
			},
			want:   []int{30},
			output: "",
		},
		{
			name: "simple sub",
			code: []OPCode{
				PUSH, 50,
				PUSH, 20,
				SUB,
			},
			want:   []int{30},
			output: "",
		},
		{
			name: "simple mul",
			code: []OPCode{
				PUSH, 5,
				PUSH, 2,
				MUL,
			},
			want:   []int{10},
			output: "",
		},
		{
			name: "simple DIV",
			code: []OPCode{
				PUSH, 10,
				PUSH, 2,
				DIV,
			},
			want:   []int{5},
			output: "",
		},
		{
			name: "simple MOD",
			code: []OPCode{
				PUSH, 9,
				PUSH, 2,
				MOD,
				PUSH, 10,
				PUSH, 2,
				MOD,
			},
			want:   []int{1, 0},
			output: "",
		},
		{
			name: "simple AND",
			code: []OPCode{
				PUSH, 5,  // 0101
				PUSH, 14, // 1110
				AND,
			},
			want:   []int{4},
			output: "",
		},
		{
			name: "simple OR",
			code: []OPCode{
				PUSH, 6,  // 0110
				PUSH, 10, // 1010
				OR,
			},
			want:   []int{14},
			output: "",
		},
		{
			name: "simple left-shift operator",
			code: []OPCode{
				PUSH, 2,
				PUSH, 1,
				SHL,
			},
			want:   []int{4},
			output: "",
		},
		{
			name: "simple right-shift operator",
			code: []OPCode{
				PUSH, 4,
				PUSH, 1,
				SHR,
			},
			want:   []int{2},
			output: "",
		},
		{
			name: "add and print",
			code: []OPCode{
				PUSH, 5,
				PUSH, 7,
				ADD,
				PRINT,
			},
			want:   []int{},
			output: "12\n",
		},
		{
			name: "push multiple values",
			code: []OPCode{
				PUSH, 1,
				PUSH, 2,
				PUSH, 3,
			},
			want:   []int{1, 2, 3},
			output: "",
		},
		{
			name: "if-end",
			code: []OPCode{// stack trace
				PUSH, 0,   // 0: 0
				PUSH, 1,   // 1: 0 1
				LT,        // 2: 1
				JMPIF,     // 3: 
				6,         // 4
				PUSH, 35,  // 5: 35
				NOP,       // 6:
				PUSH, 34,  // 7: 35 34
			},
			want:   []int{35, 34},
			output: "",
		},
		{
			name: "jmp",
			code: []OPCode{
				PUSH, 0,
				PUSH, 1,
				JMP,  8,
				PUSH, 34,
				PUSH, 2,
			},
			want:   []int{0, 1, 2},
			output: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
