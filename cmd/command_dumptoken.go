package cmd

import (
	"fmt"
	"os"
	"piled/compiler/lexer"
)

type Command_DumpToken struct {
	SubCommand
}

func (*Command_DumpToken) Name() string {
	return "dumptoken"
}

func (*Command_DumpToken) Usage(programName string) []string {
	return []string{
		programName + " <filename>",
		"    This command print tokens that is created for specified filename.",
	}
}

func (*Command_DumpToken) Execute(argv []string, pl PiledLogger) int {
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
	DumpTokens(l)

	return CommandSuccess
}
