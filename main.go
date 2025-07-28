package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"path/filepath"

	"piled/lexer"
	"piled/parser"
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

		tokens, err := lexer.LexProgram(source)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		p := parser.New(tokens)
		exprs, err := p.ParseExpr()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		c := compiler.New()
		if _, err := c.Compile(exprs); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

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

		//runningErr := RunSource(source)
		//if runningErr != nil {
		//	fmt.Fprintf(os.Stderr, "Error: %s\n", runningErr)
		//	os.Exit(1)
		//}
	}
}
