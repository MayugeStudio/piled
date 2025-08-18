package main

import (
	"fmt"
	"os"
	"path/filepath"
	"piled/ir"
	"piled/codegen"
	"piled/lexer"
	"strings"
)

type Command_Compile struct {
	SubCommand
}

func (*Command_Compile) Name() string {
	return "compile"
}

func (*Command_Compile) Usage(programName string) []string {
	return []string{
		programName + " <filename>",
	}
}

func (*Command_Compile) Execute(argv []string, pl PiledLogger) int {
	argv = argv[1:] // skip program name
	argv = argv[1:] // skip command name
	if len(argv) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: filename is not provided\n")
		return CommandError
	}
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
	// TODO: Rename filename to inputPath
	l := lexer.New(filename, source)
	g := ir.NewIrGen()
	if err := g.CompileProgram(l); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return CommandError
	}
	
	program := codegen.GenerateProgram(g.Ops)


	pl.Info("generating bytecode to %s...", outfile)
	if err := codegen.Write(outpath + ".pdb", program); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return CommandError
	}
	return CommandSuccess
}
