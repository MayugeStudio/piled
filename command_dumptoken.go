package main

import (
	"os"
	"fmt"
	"piled/lexer"
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
	filename := argv[0]

	pl.Info("reading %s ...", filename)
	source, err := ReadSourceFromFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during reading source file: %s", err)
		return CommandError
	}

	l := lexer.New(source)
	DumpTokens(l)

	return CommandSuccess
}
