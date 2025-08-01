package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"path/filepath"

	"piled/lexer"
	"piled/compiler"
	"piled/runtime"
)

func runREPL() {
	s := bufio.NewScanner(os.Stdin)
cmd:
	for {
		fmt.Print("piled> ")
		if !s.Scan() {
			break
		}
		prompt := s.Text()
		switch prompt {
		case "exit":
			{
				fmt.Println("bye!")
				break cmd
			}
		default:
			{
				runningErr := RunSource(prompt)
				if runningErr != nil {
					fmt.Fprintf(os.Stderr, "Error: %s\n", runningErr)
				}
			}
		}
	}
}

func main() {
	pl := PiledLogger{ w: os.Stdout }

	argv := os.Args
	
	if len(argv) == 1 {
		runREPL()
	} else {
		_ = argv[0]
		argv = argv[1:]
		if argv[0] == "dumptoken" {
			_ = argv[0]
			argv = argv[1:]
			filename := argv[0]
			source, err := ReadSourceFromFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error during reading source file: %s", err)
				os.Exit(1)
			}
			l := lexer.New(source)
			DumpTokens(l)
		} else {
			filename := argv[0]
			outpath := strings.TrimSuffix(filename, filepath.Ext(filename))
			outfile := outpath + ".pdb"
			
			pl.Info("reading %s ...", filename)
			source, err := ReadSourceFromFile(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				os.Exit(1)
			}
			pl.Info("reading file successfully")

			pl.Info("compiling program ...")
			l := lexer.New(source)
			c := compiler.New(l)
			c.Compile()

			pl.Info("generating bytecode to %s...", outfile)
			if err := c.Write(outpath + ".pdb"); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
				os.Exit(1)
			}

			pl.Info("reading bytecode from %s ...", outfile)
			code, err := runtime.ReadBytecodeFile(outfile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			}

			vm := runtime.NewVM(code)
			vm.Run()
		}
	}
}
