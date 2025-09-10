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

	_ = argv[0] // program_name
	argv = argv[1:]

	if len(argv) == 0 {
		os.Exit(runREPL())
	}

	command_name := argv[0]
	argv = argv[1:]

	for _, command := range commands {
		if command.Name() == command_name {
			os.Exit(command.Execute(os.Args))
		}
	}

	fmt.Fprintf(os.Stderr, "Invalid command was provided: %s\n", command_name)

	os.Exit(-1)
}
