package main

import (
	"fmt"
	"os"
	"piled/compiler/codegen"
	"piled/compiler/ir"
	"piled/compiler/lexer"
	"piled/runtime"
	"piled/compiler/token"
)

func DumpTokens(l *lexer.Lexer) {
	i := 0
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("token[%d] = %v\n", i, tok)
		i++
	}
}

func RunSource(inputPath string, source string) error {
	l := lexer.New(inputPath, source)
	g := ir.NewIrGen()
	if err := g.CompileProgram(l); err != nil {
		return err
	}

	program := codegen.GenerateProgram(g.Ops)
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
