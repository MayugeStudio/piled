package asm

import (
	"fmt"
	"strconv"
	"strings"
)

type Token struct {
	Value string
	Loc   Location
}

func (t Token) String() string {
	return fmt.Sprintf("%s: `%s`", t.Loc, t.Value)
}

type LexerError struct {
	Filename string
	Loc      Location
	Err      error
}

func (l LexerError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", l.Filename, l.Loc.Row, l.Loc.Col, l.Err)
}

func NewLexerError(filename string, loc Location, err error) *LexerError {
	return &LexerError{
		Filename: filename,
		Loc:      loc,
		Err:      err,
	}
}

func ParseRawInst(s string) (InstKind, error) {
	switch s {
	case "mov":
		return INST_MOV, nil
	case "add":
		return INST_ADD, nil
	case "sub":
		return INST_SUB, nil
	case "mul":
		return INST_MUL, nil
	case "div":
		return INST_DIV, nil
	case "dump":
		return INST_DUMP, nil
	default:
		return INST_INVALID, fmt.Errorf("unknown opcode name `%s`", s)
	}
}

func ParseRawOPKind(s string) (OPKind, error) {
	if _, err := strconv.Atoi(s); err == nil {
		return OP_IMM, nil
	}
	switch s {
	case "acc":
		return OP_ACC, nil
	case "ret":
		return OP_RET, nil
	case "r0":
		return OP_R0, nil
	case "r1":
		return OP_R1, nil
	case "r2":
		return OP_R2, nil
	case "r3":
		return OP_R3, nil
	case "pc":
		return OP_PC, nil
	default:
		return OP_INVALID, fmt.Errorf("unknown operand `%s`", s)
	}
}

func OperandCount(kind InstKind) int {
	switch kind {
	case INST_MOV, INST_ADD, INST_SUB, INST_MUL, INST_DIV:
		return 2
	case INST_DUMP:
		return 1
	case INST_INVALID:
		panic("unreachable INST_INVALID")
	default:
		panic("unreachable DEFAULT")
	}
}

func LexProgram(programPath string, source string) ([]Inst, error) {
	if len(source) == 0 {
		return nil, nil
	}

	tokens := make([]Token, 0, 0)
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
			isComma := char == ','

			isEndOfLine := col == line_length-1

			if !isSpace && !isComma {
				val += string(char)
			}

			if char == ' ' || isEndOfLine {
				token := Token{
					Value: val,
					Loc:   Location{Row: row + 1, Col: start_col + 1},
				}
				start_col = col + 1
				tokens = append(tokens, token)
				val = ""
			}
		}
	}

	insts := make([]Inst, 0, 0)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		inst_kind, err := ParseRawInst(token.Value)
		if err != nil {
			return nil, NewLexerError(programPath, token.Loc, err)
		}

		inst := NewInst(token.Loc, inst_kind)

		operandCound := OperandCount(inst_kind)

		for n := 0; n < operandCound; n++ {
			i++
			operand := tokens[i]
			operand_kind, err := ParseRawOPKind(operand.Value)
			if err != nil {
				return nil, NewLexerError(programPath, operand.Loc, err)
			}

			var value int
			if operand_kind == OP_IMM {
				value, err = strconv.Atoi(operand.Value)
				if err != nil {
					return nil, NewLexerError(programPath, operand.Loc, err)
				}
			}
			inst.Operand[n] = NewOperand(operand.Loc, operand_kind, int8(value))
		}

		insts = append(insts, inst)
	}
	return insts, nil
}
