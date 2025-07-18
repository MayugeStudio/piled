package main

import (
	"bufio"
	"fmt"
	"os"
	"piled/scanner"
	"piled/parser"
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
		case "exit": {
			fmt.Println("bye!")
			break cmd
		}
		default: {
			tokens, scanErr := scanner.ScanProgram(prompt)
			if scanErr != nil {
				fmt.Fprintf(os.Stderr, "Scanning Error: %s", scanErr)
				os.Exit(1)
			}

			p := parser.New(tokens)
			exprs, parseErr := p.Parse()
			if parseErr != nil {
				fmt.Fprintf(os.Stderr, "Parsing Error: %s", parseErr)
				os.Exit(1)
			}

			fmt.Println(exprs)
		}
		}
	}
}
