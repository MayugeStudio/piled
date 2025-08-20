package runtime

// InstructionKind represent an instruction of piled
// this is used to emulate programming language
type InstructionKind byte

const (
	// Stack
	PUSH InstructionKind = 0x00
	POP                  = 0x01
	DUP                  = 0x02
	SWAP                 = 0x03
	ROT                  = 0x04
	DROP                 = 0x05
	DUP2                 = 0x06

	// Arithmetic
	ADD = 0x10
	SUB = 0x11
	MUL = 0x12
	DIV = 0x13
	MOD = 0x14
	AND = 0x15
	OR  = 0x16
	SHL = 0x17
	SHR = 0x18

	// Comparison
	GT = 0x20
	LT = 0x21
	EQ = 0x22

	// Program Flow
	JMP   = 0x30
	JMPIF = 0x31

	// I/O
	PRINT = 0xF0

	// NOP
	NOP = 0xFA
)

type Instruction struct {
	Kind InstructionKind
	Args []int
}
