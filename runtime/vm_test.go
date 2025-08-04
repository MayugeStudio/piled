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
