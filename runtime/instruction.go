package runtime

// InstructionKind represent an instruction of piled
// this is used to emulate programming language
type InstructionKind byte

const (
	// Stack
	PUSH InstructionKind = 0x00
	POP  InstructionKind = 0x01
	DUP  InstructionKind = 0x02
	SWAP InstructionKind = 0x03
	ROT  InstructionKind = 0x04
	DROP InstructionKind = 0x05

	// Arithmetic
	ADD InstructionKind = 0x10
	SUB InstructionKind = 0x11
	MUL InstructionKind = 0x12
	DIV InstructionKind = 0x13
	MOD InstructionKind = 0x14
	AND InstructionKind = 0x15
	OR  InstructionKind = 0x16
	SHL InstructionKind = 0x17
	SHR InstructionKind = 0x18

	// Comparison
	GT InstructionKind = 0x20
	LT InstructionKind = 0x21
	EQ InstructionKind = 0x22

	// Program Flow
	JMP   InstructionKind = 0x30
	JMPIF InstructionKind = 0x31

	// I/O
	PRINT InstructionKind = 0xF0

	// NOP
	NOP InstructionKind = 0xFA
)

type Instruction struct {
	Kind InstructionKind
	Args []int
}
