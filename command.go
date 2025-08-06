package main

const (
	CommandSuccess = 0
	CommandError   = 1
)

// SubCommand represent subcommand of this application
type SubCommand interface {
	// Name returns name of the command
	Name() string

	// Usage returns usage of the command
	// Usage is string slice and they mean newline
	Usage(programName string) []string

	// Execute execute the command
	// argv startswith program-name
	Execute(argv []string, pl PiledLogger) int
}
