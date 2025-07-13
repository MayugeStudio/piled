package stack

type Stack []int

func New(cap int) Stack {
	return make(Stack, 0, 64)
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	index := len(*s) - 1
	element := (*s)[index] // Get the index of the top element

	(*s) = (*s)[:index]

	return element, true
}
