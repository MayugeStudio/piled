package cmd

import (
	"piled/compiler"
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

func (*Command_DumpToken) Execute(argv []string) int {
	argv = argv[1:] // skip program name
	argv = argv[1:] // skip command name
	if len(argv) == 0 {
		Log(ERROR, "ERROR: filename is not provided")
		return -1
	}
	inputPath := argv[0]

	Log(INFO, "reading %s ...", inputPath)
	source, err := ReadSourceFromFile(inputPath)
	if err != nil {
		Log(ERROR, "error during reading source file: %s", err)
		return -1
	}

	l := compiler.NewLexer(inputPath, source)
	DumpTokens(l)

	return 0
}
