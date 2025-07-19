package vm

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack()
	if len(stack) != 0 {
		t.Errorf("NewStack() len has to be 0, but got %d", len(stack))
	}
	if cap(stack) != 0 {
		t.Errorf("NewStack() cap has to be 0, but got %d", cap(stack))

	}
}

func TestStackLen(t *testing.T) {
	in := []byte{1, 2, 3, 4, 5, 6}
	want := 6
	stack := NewStack()
	for _, v := range in {
		stack.Push(v)
	}

	if stack.Len() != want {
		t.Errorf("TestNewStack() wanted %d, but got %d", want, stack.Len())
		return
	}
}

func TestStackPush(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want []byte
	}{
		{
			"a Single element",
			[]byte{1}, []byte{1},
		},
		{
			"Two elements",
			[]byte{1, 2}, []byte{1, 2},
		},
		{
			"Several elements",
			[]byte{1, 2, 3, 4, 5, 6}, []byte{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stack := NewStack()
			for _, v := range tc.in {
				stack.Push(v)
			}

			for i := range len(tc.in) {
				if tc.in[i] != tc.want[i] {
					t.Errorf("input[%d] = %d, want[%d] = %d", i, tc.in[i], i, tc.want[i])
					return
				}
			}
		})
	}
}

func TestStackPop(t *testing.T) {
	tests := []struct {
		name      string
		in        []byte
		want      []byte
		wantEmpty bool
	}{
		{"123456", []byte{1, 2, 3, 4, 5, 6}, []byte{6, 5, 4}, false},
		{"empty", []byte{}, []byte{1}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stack := NewStack()
			for _, v := range tc.in {
				stack.Push(v)
			}

			for _, wantValue := range tc.want {
				value, exists := stack.Pop()

				if tc.wantEmpty && exists {
					t.Errorf("stack.Pop() expected to be empty but got %d", value)
					return
				}

				if tc.wantEmpty && !exists {
					return
				}

				if !tc.wantEmpty && !exists {
					t.Errorf("")
					return
				}

				if !exists {
					t.Errorf("stack.Pop() couldn't Pop a element from a stack.")
					return
				}

				if value != wantValue {
					t.Errorf("stack.Pop() = %d, but want %d", value, wantValue)
					return
				}
			}
		})
	}
}
