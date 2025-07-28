package runtime

type OPCode byte

const (
	// Stack
	PUSH OPCode = 0x00
	POP  OPCode = 0x01

	// Arithmetic
	ADD OPCode = 0x10
	SUB OPCode = 0x11

	// Comparison
	GT  OPCode = 0x20
	LT  OPCode = 0x21
	EQ  OPCode = 0x22

	// I/O
	PRINT OPCode = 0xF0
)
