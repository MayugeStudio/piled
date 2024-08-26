package asm

type InstKind int8
type OPKind int8

const (
	OP_INVALID OPKind = iota
	OP_IMM
	OP_ACC
	OP_RET
	OP_R0
	OP_R1
	OP_R2
	OP_R3
	OP_PC

	INST_INVALID InstKind = iota
	INST_MOV
	INST_ADD
	INST_SUB
	INST_MUL
	INST_DIV
	INST_DUMP
)

type Operand struct {
	Loc   Location
	Kind  OPKind
	Value int8
}

func (o Operand) IsImmediate() bool {
	return o.Kind == OP_IMM
}

func NewOperand(loc Location, kind OPKind, value int8) Operand {
	return Operand{
		Loc:   loc,
		Kind:  kind,
		Value: value,
	}
}

type Inst struct {
	Loc     Location
	Kind    InstKind
	Operand [2]Operand
}

func NewInst(loc Location, kind InstKind) Inst {
	return Inst{
		Loc:  loc,
		Kind: kind,
	}
}
