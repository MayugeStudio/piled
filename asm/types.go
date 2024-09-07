package asm

import (
	"strconv"
	"fmt"
)

type OPType int

const (
	OP_PUSH_INT OPType = iota + 1
	OP_ADD
	OP_SUB
	OP_MUL
	OP_DIV
	OP_PRINT
	OP_INVALID
)

type Location struct {
	Row int
	Col int
}

func newLocation(row, col int) Location {
	return Location{row, col}
}

type OP struct {
	Type  OPType
	Loc   Location
	Value int
}

func nameToOPType(name string) (OPType, error) {
	switch name {
	case "+":
		return OP_ADD, nil
	case "-":
		return OP_SUB, nil
	case "*":
		return OP_MUL, nil
	case "/":
		return OP_DIV, nil
	case "print":
		return OP_PRINT, nil
	default:
		{
			_, err := strconv.Atoi(name)
			if err != nil {
				return OP_INVALID, fmt.Errorf("unknown builtin word `%s`", name)
			}
			return OP_PUSH_INT, nil
		}
	}
}

func newOP(value string, loc Location) (*OP, error) {
	opType, err := nameToOPType(value)
	if err != nil {
		return nil, err
	}
	if opType == OP_PUSH_INT {
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("unknown builtin word `%s`", value)
		}

		return &OP{
			Type: opType,
			Loc:  loc,
			Value: v,
		}, nil
	}
	return &OP{
		Type: opType,
		Loc:  loc,
	}, nil
}

