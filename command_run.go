package main

import (
	"fmt"
	"os"
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
	if len(argv) == 0 {
		fmt.Fprintf(os.Stderr, "ERROR: filename is not provided\n")
		return CommandError
	}
	filename := argv[0]

	code, err := runtime.ReadBytecodeFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		return CommandError
	}

	vm := runtime.NewVM(code)
	vm.Run()

	return CommandSuccess
}
