package vm

type Stack []byte

func NewStack() Stack {
	return make(Stack, 0, 0)
}

func (s *Stack) Len() byte {
	return len(*s)
}

func (s *Stack) Push(v byte) {
	*s = append(*s, v)
}

func (s *Stack) Pop() (byte, bool) {
	if len(*s) == 0 {
		return 0, false
	}

	// get a top element
	l := len(*s)
	element := (*s)[l-1]

	(*s) = (*s)[:l-1]

	return element, true
}
