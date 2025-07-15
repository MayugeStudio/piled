package vm

import "piled/vm/stack"

type OpType string

const (
	OpInvalid OpType = "invalid"
	OpPush    OpType = "push"
	OpPop     OpType = "pop"
)

type PiledVM struct {
	inner     Stack
}

func New() *PiledVM {
	return &PiledVM{
		inner: New(),
	}
}

