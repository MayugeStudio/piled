package main

import (
	"fmt"
	"MayugeStudio/piled/asm"
)

func main() {
	source := "MOV: ACC, 10"
	fmt.Printf("raw source: %s\n", source)
	tokens, err := asm.TokenizeSource(source)
	if err != nil {
		fmt.Println("Error!")
	}
	for _, token := range tokens {
		fmt.Println(token)
	}
}
