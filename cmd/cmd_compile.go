package cmd

import (
	"path/filepath"
	"piled/compiler"
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

func (*Command_Compile) Execute(argv []string) int {
	argv = argv[1:] // skip program name
	argv = argv[1:] // skip command name
	if len(argv) == 0 {
		Log(ERROR, "ERROR: filename is not provided")
		return -1
	}
	inputPath := argv[0]

	outpath := strings.TrimSuffix(inputPath, filepath.Ext(inputPath))
	outfile := outpath + ".pdb"

	Log(INFO, "reading %s ...", inputPath)
	source, err := ReadSourceFromFile(inputPath)
	if err != nil {
		Log(ERROR, "Error: %s", err)
		return -1
	}
	Log(INFO, "reading file successfully")

	Log(INFO, "compiling program ...")
	l := compiler.NewLexer(inputPath, source)
	g := compiler.NewIrGen()
	if err := g.CompileProgram(l); err != nil {
		Log(INFO, "Error: %s", err)
		return -1
	}

	program := compiler.GenerateProgram(g.Ops)

	Log(INFO, "generating bytecode to %s...", outfile)
	if err := compiler.Write(outpath+".pdb", program); err != nil {
		Log(ERROR, "Error: %s", err)
		return -1
	}
	return 0
}
