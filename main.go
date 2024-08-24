package main

import (
	"MayugeStudio/piled/asm"

	"fmt"
	"os"
)

func main() {
	source := "MOV: ACC, 34\nADD: ACC, 35\nDUMP: ACC"
	fmt.Printf("raw source:\n%s\n", source)
	tokens, err := asm.TokenizeSource(source)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error in TokenizeSource!", err)
		os.Exit(1)
	}
	insts, err := asm.LexTokensAsInsts(tokens)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error in LexTokensAsInsts!", err)
		os.Exit(1)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
	for _, inst := range insts {
		fmt.Println(inst)
	}
}
