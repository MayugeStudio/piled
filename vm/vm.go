package vm

import "piled/vm/stack"

type PiledVM struct {
	st     stack.Stack
	logger PiledVMLogger
}

func New(logger PiledVMLogger) *PiledVM {
	return &PiledVM{
		st: stack.New(32),
	}
}
