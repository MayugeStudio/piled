package main

import "io"
import "fmt"

// LogLevel represent log level of PiledLogger
type LogLevel string

const (
	Info  LogLevel = "INFO"
	Error          = "ERROR"
)

// PiledLogger provide logging functionalities to out compiler
type PiledLogger struct {
	w io.Writer
}

// Info
func (p *PiledLogger) Info(msg string, args ...any) {
	p.Log(Info, msg, args...)
}

// Error
func (p *PiledLogger) Error(msg string, args ...any) {
	p.Log(Error, msg, args...)
}

// Log provide basic functionality of logging
func (p *PiledLogger) Log(level LogLevel, msg string, args ...any) {
	prefix := level

	formatted := fmt.Sprintf(msg, args...)

	out := fmt.Sprintf("[%s] %s\n", prefix, formatted)

	p.write(out)
}

func (p *PiledLogger) write(msg string) {
	p.w.Write([]byte(msg))
}
