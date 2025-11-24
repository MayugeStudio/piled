package cmd

import (
	"fmt"
	"os"
	"piled/compiler"
	"piled/runtime"
)

// SubCommand represent subcommand of this application
type SubCommand interface {
	// Name returns name of the command
	Name() string

	// Usage returns usage of the command
	// Usage is string slice and they mean newline
	Usage(programName string) []string

	// Execute execute the command
	// argv startswith program-name
	Execute(argv []string) int
}

type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
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

func Log(level LogLevel, format string, args ...any) {
	switch level {
	case INFO:
		fmt.Fprintf(os.Stderr, "[INFO] ")
	case WARNING:
		fmt.Fprintf(os.Stderr, "[WARNING] ")
	case ERROR:
		fmt.Fprintf(os.Stderr, "[ERROR] ")
	}

	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintf(os.Stderr, "\n")
}
