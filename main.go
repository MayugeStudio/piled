package main

import (
	"MayugeStudio/piled/asm"

	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const DebugMode = true

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "No input file was provided")
		os.Exit(1)
	}

	filePath := args[1]

	ops, err := asm.LexProgram(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if DebugMode {
		for _, op := range ops {
			fmt.Printf("%s:%d:%d: -> %d\n", filePath, op.Loc.Row, op.Loc.Col, op.Type)
		}
	}

	fmt.Println("[INFO] Generating assembly ...")

	outContent, err := asm.GenerateLines(ops)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to generate assembly file: %s\n", err)
		os.Exit(1)
	}

	outFilename := strings.TrimSuffix(filePath, filepath.Ext(filePath))
	outAsmFilename := outFilename + ".asm"
	outBinaryFilename := outFilename + ".out"

	err = asm.WriteFileToString(outAsmFilename, outContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to write out-content to file: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("[INFO] Calling flat assembler ...")

	err = exec.Command("fasm", outAsmFilename).Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to run fasm: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("[INFO] Renaming binary-file %s -> %s ...\n", outFilename, outBinaryFilename)

	err = os.Rename(outFilename, outBinaryFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to rename %s to %s\n", outFilename, outFilename + ".out")
	}
}
