package piled

import (
	"os"
	"piled/compiler"
	"piled/parser"
	"piled/runtime"
	"piled/scanner"
)

func RunSource(source string) error {
	tokens, scanErr := scanner.ScanProgram(source)
	if scanErr != nil {
		return scanErr
	}

	p := parser.New(tokens)
	exprs, parseErr := p.ParseExpr()
	if parseErr != nil {
		return parseErr
	}

	c := compiler.New()
	codes, compileErr := c.Compile(exprs)
	if compileErr != nil {
		return compileErr
	}

	vm := runtime.NewVM(codes)

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
