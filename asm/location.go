package asm

import "fmt"

type Location struct {
	Row int
	Col int
}

func (l Location) String() string {
	return fmt.Sprintf("%d:%d", l.Row, l.Col)
}
