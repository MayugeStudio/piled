package main

import (
	"MayugeStudio/piled/asm"

	"fmt"
	"os"
)

func main() {
	reset := "mov acc, 0\n"
	dump := "dump acc\n"
	source := ""
	source += "mov acc, 34\nadd acc, 35\n" + dump + reset
	source += "mov acc, 150\nsub acc, 50\n" + dump + reset
	source += "mov acc, 25\nmul acc, 2\n" + dump + reset
	source += "mov acc, 14\ndiv acc, 2\n" + dump + reset

	fmt.Printf("raw source:\n%s\n", source)
	tokens, err := asm.LexProgram(source)
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
