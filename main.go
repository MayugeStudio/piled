package main

import (
	"bufio"
	"fmt"
	"os"
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
				runningErr := RunSource(prompt)
				if runningErr != nil {
					fmt.Fprintf(os.Stderr, "Error: %s\n", runningErr)
				}
			}
		}
	}
	return 0
}

var commands []SubCommand

func main() {
	pl := PiledLogger{w: os.Stdout}
	commands = append(commands, &Command_DumpToken{})
	commands = append(commands, &Command_Run{})
	commands = append(commands, &Command_Compile{})

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
			os.Exit(command.Execute(os.Args, pl))
		}
	}

	fmt.Fprintf(os.Stderr, "Invalid command was provided: %s\n", command_name)

	os.Exit(CommandError)
}
