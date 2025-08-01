package main

import (
	"fmt"
	"os"
	"piled/compiler"
	"piled/lexer"
	"piled/token"
	"piled/runtime"
)

func DumpTokens(l *lexer.Lexer) {
	i := 0
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("token[%d] = %v\n", i, tok)
		i++
	}
}

func RunSource(source string) error {
	l := lexer.New(source)
	c := compiler.New(l)
	code := c.Compile()
	vm := runtime.NewVM(code)
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
