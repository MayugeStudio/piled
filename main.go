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
					fmt.Fprintf(os.Stderr, "Scanning Error: %s\n", scanErr)
					continue
				}

				p := parser.New(tokens)
				exprs, parseErr := p.ParseExpr()
				if parseErr != nil {
					fmt.Fprintf(os.Stderr, "Parsing Error: %s\n", parseErr)
					continue
				}

				c := compiler.New()
				codes, compileErr := c.Compile(exprs)
				if compileErr != nil {
					fmt.Fprintf(os.Stderr, "Compiling Error: %s\n", compileErr)
					continue
				}

				vm := runtime.NewVM(codes)

				vm.Run()
			}
		}
	}
}
