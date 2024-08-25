package main

import (
	"MayugeStudio/piled/asm"

	"fmt"
	"os"
)

func main() {
	reset := "MOV: ACC, 0\n"
	dump := "DUMP: ACC\n"
	source := ""
	source += "MOV: ACC, 34\nADD: ACC, 35\n" + dump + reset
	source += "MOV: ACC, 150\nSUB: ACC, 50\n" + dump + reset
	source += "MOV: ACC, 25\nMUL: ACC, 2\n" + dump + reset
	source += "MOV: ACC, 14\nDIV: ACC, 2\n" + dump + reset

	fmt.Printf("raw source:\n%s\n", source)
	tokens, err := asm.LexSource(source)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: In `TokenizeSource`", err)
		os.Exit(1)
	}
	insts, err := asm.LexTokens(tokens)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: In `LexTokensAsInsts`", err)
		os.Exit(1)
	}
	err = asm.InterpretInsts(insts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: In `InterpretInst`", err)
		os.Exit(1)
	}
}
