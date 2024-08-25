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

	bytes, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not open file `%s: %w\n", filepath, err)
	}

	source := string(bytes)

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
