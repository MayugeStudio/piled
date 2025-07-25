package main

import (
	"bufio"
	"fmt"
	"os"
	"piled/compiler"
	"piled/parser"
	"piled/runtime"
	"piled/scanner"
)

func main() {
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
				tokens, scanErr := scanner.ScanProgram(prompt)
				if scanErr != nil {
					fmt.Fprintf(os.Stderr, "Scanning Error: %s", scanErr)
					os.Exit(1)
				}

				p := parser.New(tokens)
				exprs, parseErr := p.ParseExpr()
				if parseErr != nil {
					fmt.Fprintf(os.Stderr, "Parsing Error: %s", parseErr)
					os.Exit(1)
				}

				c := compiler.NewCompiler() // TODO: change compiler constructor name to .New.
				codes, compileErr := c.Compile(exprs)
				if compileErr != nil {
					fmt.Fprintf(os.Stderr, "Compiling Error: %s", compileErr)
					os.Exit(1)
				}

				vm := runtime.NewVM(codes)

				vm.Run()
			}
		}
	}
}
