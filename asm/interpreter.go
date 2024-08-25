package asm

import "fmt"

type Registers struct {
	acc int8
	ret int8
	r0  int8
	r1  int8
	r2  int8
	r3  int8
	pc  int8
}

func (r *Registers) Set(kind OPKind, value int8) {
	switch kind {
	case OP_ACC:
		r.acc = value
	case OP_RET:
		r.ret = value
	case OP_R0:
		r.r0 = value
	case OP_R1:
		r.r1 = value
	case OP_R2:
		r.r2 = value
	case OP_R3:
		r.r3 = value
	case OP_PC:
		r.pc = value
	case OP_IMM:
		panic("Exhaustive handling of `OP_IMM`")
	default:
		panic("Exhaustive handling of `OPKind`")
	}
}

func (r *Registers) Get(operand Operand) int8 {
	if operand.Kind == OP_IMM {
		return operand.Value
	}
	switch operand.Kind {
	case OP_ACC:
		return r.acc
	case OP_RET:
		return r.ret
	case OP_R0:
		return r.r0
	case OP_R1:
		return r.r1
	case OP_R2:
		return r.r2
	case OP_R3:
		return r.r3
	case OP_PC:
		return r.pc
	case OP_IMM:
		panic("Exhaustive handling of `OP_IMM`")
	default:
		panic("Exhaustive handling of `OPKind`")
	}
}

func InterpretInsts(insts []Inst) error {
	registers := &Registers{}
	insts_length := int8(len(insts))

	for registers.pc < insts_length {
		inst := insts[registers.pc]
		op_a := inst.Operand[0]
		op_b := inst.Operand[1]

		switch inst.Kind {
		case INST_MOV:
			val := registers.Get(op_b)
			registers.Set(op_a.Kind, val)
		case INST_ADD:
			val_a := registers.Get(op_a)
			val_b := registers.Get(op_b)
			result := val_a + val_b
			registers.Set(op_a.Kind, result)
		case INST_SUB:
			val_a := registers.Get(op_a)
			val_b := registers.Get(op_b)
			result := val_a - val_b
			registers.Set(op_a.Kind, result)
		case INST_MUL:
			val_a := registers.Get(op_a)
			val_b := registers.Get(op_b)
			result := val_a * val_b
			registers.Set(op_a.Kind, result)
		case INST_DIV:
			val_a := registers.Get(op_a)
			val_b := registers.Get(op_b)
			result := val_a / val_b
			registers.Set(op_a.Kind, result)
		case INST_DUMP:
			n := registers.Get(op_a)
			fmt.Println(n)
		default:
			return fmt.Errorf("unknown opcode kind `%d`", inst.Kind)
		}
		registers.pc += 1
	}
	return nil
}
