package asm

import (
	"strings"
	"fmt"
	"strconv"
)

type AsmToken struct {
	Value string
	Loc   Location
}

func (t AsmToken) String() string {
	return fmt.Sprintf("%s: `%s`", t.Loc, t.Value)
}

func ParseRawOPCode(s string) (InstKind, error) {
	switch s {
	case "MOV":
		return INST_MOV, nil
	case "ADD":
		return INST_ADD, nil
	case "SUB":
		return INST_SUB, nil
	case "MUL":
		return INST_MUL, nil
	case "DIV":
		return INST_DIV, nil
	case "DUMP":
		return INST_DUMP, nil
	default:
		return INST_INVALID, fmt.Errorf("unknown opcode name `%s`", s)
	}
}


func ParseRawOperandKind(s string) (OPKind, error) {
	if _, err := strconv.Atoi(s); err == nil {
		return OP_IMM, nil
	}
	switch s {
	case "ACC":
		return OP_ACC, nil
	case "RET":
		return OP_RET, nil
	case "R_0":
		return OP_R0, nil
	case "R_1":
		return OP_R1, nil
	case "R_2":
		return OP_R2, nil
	case "R_3":
		return OP_R3, nil
	case "R_PC":
		return OP_PC, nil
	default:
		return OP_INVALID, fmt.Errorf("unknown operand `%s`", s)
	}
}
func LexSource(source string) ([]AsmToken, error) {
	tokens := make([]AsmToken, 0, 0)
	if len(source) == 0 {
		return nil, nil
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
					Value: val,
					Loc:   Location{Row: row, Col: start_col},
				}
				start_col = col + 1
				tokens = append(tokens, token)
				val = ""
			}
		}
	}
	return tokens, nil
}

func LexTokens(tokens []AsmToken) (ops []Inst, err error) {
	for i := 0; i < len(tokens); i++ {
		var inst Inst
		token := tokens[i]
		opcode_kind, err := ParseRawOPCode(token.Value)
		inst.Kind = opcode_kind
		if err != nil {
			return nil, err
		}

		operand_num := 0
		switch opcode_kind {
		case INST_MOV, INST_ADD, INST_SUB, INST_MUL, INST_DIV:
			operand_num = 2
		case INST_DUMP:
			operand_num = 1
		case INST_INVALID:
			panic("unreachable INST_INVALID")
		default:
			panic("unreachable DEFAULT")
		}

		for n := 0; n < operand_num; n++ {
			i++
			operand := tokens[i]
			operand_kind, err := ParseRawOperandKind(operand.Value)
			if err != nil {
				return nil, err
			}

			var value int
			if operand_kind == OP_IMM {
				value, err = strconv.Atoi(operand.Value)
				if err != nil {
					return nil, err
				}
			}
			inst.Operand[n] = Operand{
				Loc:   operand.Loc,
				Kind:  operand_kind,
				Value: int8(value),
			}
		}

		ops = append(ops, inst)
	}
	return
}

