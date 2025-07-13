package stack

import "testing"

func TestStackNew(t *testing.T) {
	in := 64
	want := 64
	stack := New(in)

	length := len(stack)
	if length != 0 {
		t.Errorf("TestNew() wanted 0, but got %d", length)
		return
	}

	capacity := cap(stack)
	if capacity != want {
		t.Errorf("TestNew() wanted %d, but got %d", want, capacity)
		return
	}
}

func TestStackLen(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	want := 6
	stack := New(32)
	for _, v := range in {
		stack.Push(v)
	}

	if stack.Len() != want {
		t.Errorf("TestNew() wanted %d, but got %d", want, stack.Len())
		return
	}
}

func TestStackPush(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{
			"a Single element",
			[]int{1}, []int{1},
		},
		{
			"Two elements",
			[]int{1, 2}, []int{1, 2},
		},
		{
			"Several elements",
			[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stack := New(32)
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
		in        []int
		want      []int
		wantEmpty bool
	}{
		{"123456", []int{1, 2, 3, 4, 5, 6}, []int{6, 5, 4}, false},
		{"empty", []int{}, []int{1}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stack := New(32)
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
