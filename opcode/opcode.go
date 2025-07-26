package opcode

type Code int

const (
	PUSH Code = iota
	ADD
	SUB
	GT
	LT
	EQ
	PRINT
)
