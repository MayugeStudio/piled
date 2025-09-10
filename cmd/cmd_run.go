package cmd

import (
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

func (*Command_Run) Execute(argv []string) int {
	argv = argv[1:] // skip program name
	argv = argv[1:] // skip command name
	if len(argv) == 0 {
		Log(ERROR, "ERROR: filename is not provided")
		return -1
	}
	filename := argv[0]

	code, err := runtime.ReadBytecodeFile(filename)
	if err != nil {
		Log(ERROR, "ERROR: %s", err)
		return -1
	}

	vm := runtime.NewVM(code)
	vm.Run()

	return 0
}
