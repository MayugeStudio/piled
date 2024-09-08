package asm

import (
	"fmt"
	"strconv"
)

type OPType int

const (
	OP_PUSH_INT OPType = iota + 1
	OP_ADD
	OP_SUB
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
	Type     OPType
	Loc      Location
	Value    int
}

func nameToOPType(name string) (OPType, error) {
	switch name {
	case "+":
		return OP_ADD, nil
	case "-":
		return OP_SUB, nil
	case "print":
		return OP_PRINT, nil
	default:
		{
			_, err := strconv.Atoi(name)
			if err != nil {
				return OP_INVALID, fmt.Errorf("unknown word `%s`", name)
			}
			return OP_PUSH_INT, nil
		}
	}
}
