package runtime

import (
	"bytes"
	"testing"
)

func TestVM_Run(t *testing.T) {
	tests := []struct {
		name   string
		code   []int
		want   []int
		output string
	}{
		{
			name: "simple add",
			code: []int{
				OP_PUSH, 10,
				OP_PUSH, 20,
				OP_ADD,
			},
			want:   []int{30},
			output: "",
		},
		{
			name: "simple sub",
			code: []int{
				OP_PUSH, 50,
				OP_PUSH, 20,
				OP_SUB,
			},
			want:   []int{30},
			output: "",
		},
		{
			name: "add and print",
			code: []int{
				OP_PUSH, 5,
				OP_PUSH, 7,
				OP_ADD,
				OP_PRINT,
			},
			want:   []int{},
			output: "12\n",
		},
		{
			name: "push multiple values",
			code: []int{
				OP_PUSH, 1,
				OP_PUSH, 2,
				OP_PUSH, 3,
			},
			want:   []int{1, 2, 3},
			output: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vm := &VM{code: tt.code}

			buf := &bytes.Buffer{}
			writer = buf
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
