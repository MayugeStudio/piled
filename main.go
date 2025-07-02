package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	programName := args[0] // TODO: Introduce some sort of arguments operating function
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input-file>\n", programName)
		fmt.Fprintf(os.Stderr, "ERROR: input file is not provided\n")
		os.Exit(1)
	}

	args = args[1:]
	inputPath := args[0]

	// Reading input file
	source, err := readFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: could not read source file from `%s`: %s\n", inputPath, err)
		os.Exit(1)
	}

	// Lexing
	ops, err := LexProgram(inputPath, source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(ops)
}

