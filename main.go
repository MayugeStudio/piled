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
	argv := os.Args
	if len(argv) == 1 {
		runREPL()
	} else {
		_ = argv[0]
		argv = argv[1:]
		filename := argv[0]
		source, err := ReadSourceFromFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		l := lexer.New(source)
		c := compiler.New(l)
		c.Compile()
		outpath := strings.TrimSuffix(filename, filepath.Ext(filename))
		
		if err := c.Write(outpath + ".pdb"); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		code, err := runtime.ReadBytecodeFile(outpath + ".pdb")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		}
		vm := runtime.NewVM(code)
		vm.Run()
	}
}
