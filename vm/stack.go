package vm

type Stack []int

func NewStack() Stack {
	return make(Stack, 0, 0)
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

	// get a top element
	l := len(*s)
	element := (*s)[l-1]

	(*s) = (*s)[:l-1]

	return element, true
}
