package asm

import "fmt"
import "strings"
import "strconv"

type Location struct {
	row int
	col int
}

func (l Location) String() string {
	return fmt.Sprintf("%d:%d", l.row, l.col)
}

type AsmToken struct {
	value string
	loc   Location
}

func (t AsmToken) String() string {
	return fmt.Sprintf("%s: `%s`", t.loc, t.value)
}

type OPCodeKind int8

const (
	OP_INVALID OPCodeKind = iota
	OP_MOV
	OP_ADD
	OP_SUB
	OP_MUL
	OP_DIV
	OP_DUMP
)

func ParseRawOPCode(s string) (OPCodeKind, error) {
	switch s {
	case "MOV":
		return OP_MOV, nil
	case "ADD":
		return OP_ADD, nil
	case "SUB":
		return OP_SUB, nil
	case "MUL":
		return OP_MUL, nil
	case "DIV":
		return OP_DIV, nil
	case "DUMP":
		return OP_DUMP, nil
	default:
		return OP_INVALID, fmt.Errorf("unknown opcode name `%s`", s)
	}
}

type OperandKind int8

const (
	OperandKind_INVALID OperandKind = iota
	OperandKind_IMM
	OperandKind_ACC
	OperandKind_RET
	OperandKind_0
	OperandKind_1
	OperandKind_2
	OperandKind_3
	OperandKind_PC
)

func ParseRawOperandKind(s string) (OperandKind, error) {
	if _, err := strconv.Atoi(s); err == nil {
		return OperandKind_IMM, nil
	}
	switch s {
	case "ACC":
		return OperandKind_ACC, nil
	case "RET":
		return OperandKind_RET, nil
	case "R_0":
		return OperandKind_0, nil
	case "R_1":
		return OperandKind_1, nil
	case "R_2":
		return OperandKind_2, nil
	case "R_3":
		return OperandKind_3, nil
	case "R_PC":
		return OperandKind_PC, nil
	default:
		return OperandKind_INVALID, fmt.Errorf("unknown operand `%s`", s)
	}
}

type Operand struct {
	loc   Location
	kind  OperandKind
	value int8
}

func (o Operand) String() string {
	return fmt.Sprintf("[kind=`%d`, value=`%d`]", o.kind, o.value)
}

type Inst struct {
	loc     Location
	kind    OPCodeKind
	operand [2]Operand
}

func (i Inst) String() string {
	return fmt.Sprintf("kind=%d, operand=`%s`", i.kind, i.operand)
}

////// Tokenizer //////

func TokenizeSource(source string) (tokens []AsmToken, err error) {
	if len(source) == 0 {
		return
	}

	lines := strings.Split(source, "\n")

	for row, line := range lines {
		if len(line) == 0 {
			continue
		}
		val := ""
		start_col := 0
		line_length := len(line)
		for col := 0; col < line_length; col++ {
			char := line[col]
			isSpace := char == ' '
			isColon := char == ':'
			isComma := char == ','

			isEndOfLine := col == line_length-1

			if !isSpace && !isColon && !isComma {
				val += string(char)
			}

			if char == ' ' || isEndOfLine {
				token := AsmToken{
					value: val,
					loc:   Location{row: row, col: start_col},
				}
				start_col = col + 1
				tokens = append(tokens, token)
				val = ""
			}
		}
	}
	return
}

func LexTokensAsInsts(tokens []AsmToken) (ops []Inst, err error) {
	for i := 0; i < len(tokens); i++ {
		var inst Inst
		token := tokens[i]
		opcode_kind, err := ParseRawOPCode(token.value)
		inst.kind = opcode_kind
		if err != nil {
			return nil, err
		}

		operand_num := 0
		switch opcode_kind {
		case OP_MOV, OP_ADD, OP_SUB, OP_MUL, OP_DIV:
			operand_num = 2
		case OP_DUMP:
			operand_num = 1
		case OP_INVALID:
			panic("unreachable OP_INVALID")
		default:
			panic("unreachable DEFAULT")
		}

		for n := 0; n < operand_num; n++ {
			i++
			operand := tokens[i]
			operand_kind, err := ParseRawOperandKind(operand.value)
			if err != nil {
				return nil, err
			}

			var value int
			if operand_kind == OperandKind_IMM {
				value, err = strconv.Atoi(operand.value)
				if err != nil {
					return nil, err
				}
			}
			inst.operand[n] = Operand{
				loc:   operand.loc,
				kind:  operand_kind,
				value: int8(value),
			}
		}

		ops = append(ops, inst)
	}
	return
}

type Registers map[OperandKind]int8


func (r *Registers) SetByOperand(a Operand, value int8) {
	r.SetByKind(a.kind, value)
}

func (r *Registers) SetByKind(kind OperandKind, value int8) {
	(*r)[kind] = value
}

func (r *Registers) GetByOperand(operand Operand) int8 {
	if operand.kind == OperandKind_IMM {
		return operand.value
	}
	return (*r)[operand.kind]
}

func (r *Registers) GetByKind(kind OperandKind) int8 {
	return (*r)[kind]
}

func InterpretInsts(insts []Inst) error {
	var registers Registers = Registers{
		OperandKind_ACC: 0,
		OperandKind_RET: 0,
		OperandKind_0:   0,
		OperandKind_1:   0,
		OperandKind_2:   0,
		OperandKind_3:   0,
		OperandKind_PC:  0,
	}

	insts_length := int8(len(insts))

	for registers[OperandKind_PC] < insts_length {
		inst := insts[registers[OperandKind_PC]]
		op_a := inst.operand[0]
		op_b := inst.operand[1]

		switch inst.kind {
		case OP_MOV:
			val := registers.GetByOperand(op_b)
			registers.SetByOperand(op_a, val)
		case OP_ADD:
			val_a := registers.GetByOperand(op_a)
			val_b := registers.GetByOperand(op_b)
			result := val_a + val_b
			registers.SetByOperand(op_a, result)
		case OP_SUB:
			val_a := registers.GetByOperand(op_a)
			val_b := registers.GetByOperand(op_b)
			result := val_a - val_b
			registers.SetByOperand(op_a, result)
		case OP_MUL:
			val_a := registers.GetByOperand(op_a)
			val_b := registers.GetByOperand(op_b)
			result := val_a * val_b
			registers.SetByOperand(op_a, result)
		case OP_DIV:
			val_a := registers.GetByOperand(op_a)
			val_b := registers.GetByOperand(op_b)
			result := val_a / val_b
			registers.SetByOperand(op_a, result)
		case OP_DUMP:
			n := registers.GetByOperand(op_a)
			fmt.Println(n)
		}
		pc := registers.GetByKind(OperandKind_PC)
		registers.SetByKind(OperandKind_PC, pc+1)
	}
	return nil
}

