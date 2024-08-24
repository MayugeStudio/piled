package asm

import "fmt"
import "strings"
import "strconv"

// Location
type Location struct {
	row int
	col int
}

func (l Location) String() string {
	return fmt.Sprintf("%d:%d", l.row, l.col)
}

// AsmToken
type AsmToken struct {
	value string
	loc   Location
}

func (t AsmToken) String() string {
	return fmt.Sprintf("%s: `%s`", t.loc, t.value)
}

////// Instruction //////

// OPCode //

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
	OperandKind_MOV
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
	case "MOV":
		return OperandKind_MOV, nil
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

// Operand
type Operand struct {
	loc   Location
	kind  OperandKind
	value int8
}

// OP
type Inst struct {
	loc     Location
	kind    OPCodeKind
	operand [2]Operand
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
