package runtime

// OPCode represent an instruction of piled
// this is used to emulate programming language
type OPCode byte

const (
	// Stack
	PUSH OPCode = 0x00
	POP  OPCode = 0x01

	// Arithmetic
	ADD OPCode = 0x10
	SUB OPCode = 0x11
	MUL OPCode = 0x12
	DIV OPCode = 0x13
	MOD OPCode = 0x14
	AND OPCode = 0x15
	OR  OPCode = 0x16
	SHL OPCode = 0x17
	SHR OPCode = 0x18

	// Comparison
	GT  OPCode = 0x20
	LT  OPCode = 0x21
	EQ  OPCode = 0x22

	// Program Flow
	JMPIF OPCode = 0x30

	// I/O
	PRINT OPCode = 0xF0

	// NOP
	NOP OPCode = 0xFA
)
