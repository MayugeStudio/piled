package main

import (
	"os"
	"piled/compiler"
	"piled/lexer"
	"piled/runtime"
)

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
