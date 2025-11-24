package main

import (
	"fmt"
	"os"
	"piled/compiler"
	"piled/runtime"
)

func DumpTokens(l *compiler.Lexer) {
	i := 0
	for tok := l.NextToken(); tok.Type != compiler.EOF; tok = l.NextToken() {
		fmt.Printf("token[%d] = %v\n", i, tok)
		i++
	}
}

func RunSource(inputPath string, source string) error {
	l := compiler.NewLexer(inputPath, source)
	g := compiler.NewIrGen()
	if err := g.CompileProgram(l); err != nil {
		return err
	}

	program := compiler.GenerateProgram(g.Ops)
	vm := runtime.NewVM(program)
	vm.Run()
	return nil
}

func ReadSourceFromFile(filename string) (string, error) {
	source, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(source), nil
}
