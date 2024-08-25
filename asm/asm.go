package asm

import "fmt"

type Operand struct {
	Loc   Location
	Kind  OPKind
	Value int8
}

func (o Operand) String() string {
	return fmt.Sprintf("[kind=`%d`, value=`%d`]", o.Kind, o.Value)
}

func (o Operand) IsImmediate() bool {
	return o.Kind == OP_IMM
}

type Inst struct {
	Loc     Location
	Kind    InstKind
	Operand [2]Operand
}

func (i Inst) String() string {
	return fmt.Sprintf("kind=%d, operand=`%s`", i.Kind, i.Operand)
}

func NewInst(loc Location, kind InstKind) Inst {
	return Inst{
		Loc:  loc,
		Kind: kind,
	}
}

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

