package runtime

type OPCode int

const (
	PUSH OPCode = iota
	ADD
	SUB
	GT
	LT
	EQ
	PRINT
)
