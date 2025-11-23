package main

import (
	"bufio"
	"fmt"
	"os"
	"piled/cmd"
)

func runREPL() int {
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
				runningErr := RunSource("REPR", prompt)
				if runningErr != nil {
					fmt.Fprintf(os.Stderr, "Error: %s\n", runningErr)
				}
			}
		}
	}
	return 0
}



var commands []cmd.SubCommand
func main() {
	commands = append(commands, &cmd.Command_DumpToken{})
	commands = append(commands, &cmd.Command_DumpIr{})
	commands = append(commands, &cmd.Command_Run{})
	commands = append(commands, &cmd.Command_Compile{})

	argv := os.Args

	programName := argv[0]
	argv = argv[1:]

	if len(argv) == 0 {
		os.Exit(runREPL())
	}

	for len(argv) > 0 {
		arg := argv[0]
		argv = argv[1:]
		if (arg == "--help" || arg == "-h") {
			fmt.Println("Usage: ")
			for _, command := range commands {
				fmt.Print("  ")
				for _, line := range command.Usage(programName) {
					fmt.Print(line)
				}
				fmt.Print("\n")
			}
			os.Exit(0)
		}
		for _, command := range commands {
			if command.Name() == arg {
				os.Exit(command.Execute(os.Args))
			}
		}
		fmt.Fprintf(os.Stderr, "Invalid command was provided: %s\n", arg)
		os.Exit(1)
	}
	os.Exit(1)
}
