package main

import (
	"strings"
	"path/filepath"
	"os"
	"fmt"
	"piled/lexer"
	"piled/compiler"
	"piled/runtime"
)

type Command_Run struct {
	SubCommand
}

func (*Command_Run) Name() string {
	return "run"
}

func (*Command_Run) Usage(programName string) []string {
	return []string{
		programName + " <filename>",
	}
}

func (*Command_Run) Execute(argv []string, pl PiledLogger) int {
	argv = argv[1:] // skip program name
	argv = argv[1:] // skip command name
	filename := argv[0]
	outpath := strings.TrimSuffix(filename, filepath.Ext(filename))
	outfile := outpath + ".pdb"
	
	pl.Info("reading %s ...", filename)
	source, err := ReadSourceFromFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return CommandError
	}
	pl.Info("reading file successfully")

	pl.Info("compiling program ...")
	l := lexer.New(source)
	c := compiler.New(l)
	c.Compile()

	pl.Info("generating bytecode to %s...", outfile)
	if err := c.Write(outpath + ".pdb"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return CommandError
	}

	pl.Info("reading bytecode from %s ...", outfile)
	code, err := runtime.ReadBytecodeFile(outfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return CommandError
	}

	vm := runtime.NewVM(code)
	vm.Run()

	return CommandSuccess
}
