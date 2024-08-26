package main

import (
	"MayugeStudio/piled/asm"

	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "No input file was provided")
		os.Exit(1)
	}

	filepath := args[1]

	insts, err := asm.LexProgram(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	err = asm.InterpretInsts(insts)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: In `InterpretInst`", err)
		os.Exit(1)
	}
}
