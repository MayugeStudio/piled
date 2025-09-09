package cmd

import (
	"fmt"
	"os"
	"piled/compiler/ir"
	"piled/compiler/lexer"
)

type Command_DumpIr struct {
	SubCommand
}

func (*Command_DumpIr) Name() string {
	return "dumpir"
}

func (*Command_DumpIr) Usage(programName string) []string {
	return []string{
		programName + " <filename>",
		"    This command prints irs",
	}
}

func (*Command_DumpIr) Execute(argv []string, pl PiledLogger) int {
	argv = argv[1:] // skip program name
	argv = argv[1:] // skip command name
	if len(argv) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: filename is not provided\n")
		return CommandError
	}
	inputPath := argv[0]

	pl.Info("reading %s ...", inputPath)
	source, err := ReadSourceFromFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during reading source file: %s", err)
		return CommandError
	}

	l := lexer.New(inputPath, source)

	g := ir.NewIrGen()
	g.CompileProgram(l)

	irs := g.Ops

	for _, op := range irs {
		fmt.Println(op)
	}

	return CommandSuccess
}
