package main

import (
	"bufio"
	"fmt"
	"os"
	"piled"
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
				runningErr := piled.RunSource(prompt)
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
		source, readErr := piled.ReadSourceFromFile(filename)
		if readErr != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", readErr)
			os.Exit(1)
		}
		runningErr := piled.RunSource(source)
		if runningErr != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", runningErr)
			os.Exit(1)
		}
	}
}
