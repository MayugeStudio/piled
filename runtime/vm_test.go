package runtime

import (
	"bytes"
	"piled/opcode"
	"testing"
)

func TestVM_Run(t *testing.T) {
	tests := []struct {
		name   string
		code   []opcode.Code
		want   []int
		output string
	}{
		{
			name: "simple add",
			code: []opcode.Code{
				opcode.PUSH, 10,
				opcode.PUSH, 20,
				opcode.ADD,
			},
			want:   []int{30},
			output: "",
		},
		{
			name: "simple sub",
			code: []opcode.Code{
				opcode.PUSH, 50,
				opcode.PUSH, 20,
				opcode.SUB,
			},
			want:   []int{30},
			output: "",
		},
		{
			name: "add and print",
			code: []opcode.Code{
				opcode.PUSH, 5,
				opcode.PUSH, 7,
				opcode.ADD,
				opcode.PRINT,
			},
			want:   []int{},
			output: "12\n",
		},
		{
			name: "push multiple values",
			code: []opcode.Code{
				opcode.PUSH, 1,
				opcode.PUSH, 2,
				opcode.PUSH, 3,
			},
			want:   []int{1, 2, 3},
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
