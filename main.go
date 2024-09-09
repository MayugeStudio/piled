package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const DebugMode = false

type Location struct {
	Row int
	Col int
}

type TokenType int

const (
	Token_PUSH_INT TokenType = iota + 1
	Token_ADD
	Token_SUB
	Token_EQUAL
	Token_PRINT
	Token_INVALID
)

type Token struct {
	Type  TokenType
	Loc   Location
	Value int
}

func nameToTokenType(name string) (TokenType, error) {
	switch name {
	case "+":
		return Token_ADD, nil
	case "-":
		return Token_SUB, nil
	case "=":
		return Token_EQUAL, nil
	case "print":
		return Token_PRINT, nil
	default:
		{
			_, err := strconv.Atoi(name)
			if err != nil {
				return Token_INVALID, fmt.Errorf("unknown word `%s`", name)
			}
			return Token_PUSH_INT, nil
		}
	}
}

type lexerError struct {
	Filename string
	Loc      Location
	Err      error
}

func (l lexerError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s", l.Filename, l.Loc.Row, l.Loc.Col, l.Err)
}

func newLexerError(filename string, loc Location, err error) *lexerError {
	return &lexerError{
		Filename: filename,
		Loc:      loc,
		Err:      err,
	}
}

func readFile(programPath string) (string, error) {
	bytes, err := os.ReadFile(programPath)
	if err != nil {
		return "", fmt.Errorf("could not open file `%s: %w\n", programPath, err)
	}
	return string(bytes), nil
}

func lexWord(filename string, value string, loc Location) (*Token, error) {
	opType, err := nameToTokenType(value)
	if err != nil {
		return nil, newLexerError(filename, loc, err)
	}
	if opType == Token_PUSH_INT {
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		return &Token{
			Type:  opType,
			Loc:   loc,
			Value: v,
		}, nil
	}
	return &Token{
		Type: opType,
		Loc:  loc,
	}, nil
}

func lexSourceIntoTokens(filename string, source string) ([]*Token, error) {
	ops := make([]*Token, 0)
	lines := strings.Split(source, "\n")

	for row, line := range lines {
		val := ""
		start_col := 0
		line_length := len(line)
		for col := 0; col < line_length; col++ {
			char := line[col]
			isEndOfLine := col == line_length-1
			isSpace := char == ' '
			isComma := char == ','

			if !isSpace && !isComma {
				val += string(char)
			}

			if isSpace || isEndOfLine {
				loc := Location{Row: row + 1, Col: start_col + 1}
				op, err := lexWord(filename, val, loc)
				if err != nil {
					return nil, err
				}
				start_col = col + 1
				ops = append(ops, op)
				val = ""
			}
		}
	}
	return ops, nil
}

func LexProgram(programPath string) ([]*Token, error) {
	source, err := readFile(programPath)
	if err != nil {
		return nil, err
	}

	ops, err := lexSourceIntoTokens(programPath, source)
	if err != nil {
		return nil, err
	}

	return ops, nil
}

func GenerateLines(ops []*Token) (string, error) {
	b := strings.Builder{}
	b.WriteString("format ELF64 executable 3\n")
	b.WriteString("\n")
	// Builtin dump
	writeDump(&b)

	// Header
	b.WriteString("entry _start\n")
	// Start Label
	b.WriteString("_start:\n")
	for _, op := range ops {
		switch op.Type {
		case Token_PUSH_INT:
			b.WriteString(fmt.Sprintf("    push %d\n", op.Value))
		case Token_ADD:
			b.WriteString("    pop rax\n")
			b.WriteString("    pop rbx\n")
			b.WriteString("    add rax, rbx\n")
			b.WriteString("    push rax\n")
		case Token_SUB:
			b.WriteString("    pop rax\n")
			b.WriteString("    pop rbx\n")
			b.WriteString("    sub rbx, rax\n")
			b.WriteString("    push rbx\n")
		case Token_EQUAL:
			b.WriteString("    mov rcx, 0\n")
			b.WriteString("    mov rdx, 1\n")
			b.WriteString("    pop rax\n")
			b.WriteString("    pop rbx\n")
			b.WriteString("    cmp rax, rbx\n")
			b.WriteString("    cmove rcx, rdx\n")
			b.WriteString("    push rcx\n")
		case Token_PRINT:
			b.WriteString("    pop rdi\n")
			b.WriteString("    call dump\n")
		}
	}

	b.WriteString("    mov rax, 60\n")
	b.WriteString("    mov rdi, 0\n")
	b.WriteString("    syscall\n")

	return b.String(), nil
}

func WriteFileToString(filePath string, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %s", err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func writeDump(b *strings.Builder) {
	b.WriteString("dump:\n")
	b.WriteString("    mov     r8, -3689348814741910323\n")
	b.WriteString("    sub     rsp, 40\n")
	b.WriteString("    mov     BYTE [rsp+31], 10\n")
	b.WriteString("    lea     rcx, [rsp+30]\n")
	b.WriteString(".L2:\n")
	b.WriteString("    mov     rax, rdi\n")
	b.WriteString("    mul     r8\n")
	b.WriteString("    mov     rax, rdi\n")
	b.WriteString("    shr     rdx, 3\n")
	b.WriteString("    lea     rsi, [rdx+rdx*4]\n")
	b.WriteString("    add     rsi, rsi\n")
	b.WriteString("    sub     rax, rsi\n")
	b.WriteString("    mov     rsi, rcx\n")
	b.WriteString("    sub     rcx, 1\n")
	b.WriteString("    add     eax, 48\n")
	b.WriteString("    mov     BYTE [rcx+1], al\n")
	b.WriteString("    mov     rax, rdi\n")
	b.WriteString("    mov     rdi, rdx\n")
	b.WriteString("    cmp     rax, 9\n")
	b.WriteString("    ja      .L2\n")
	b.WriteString("    lea     rdx, [rsp+32]\n")
	b.WriteString("    mov     edi, 1\n")
	b.WriteString("    sub     rdx, rsi\n")
	b.WriteString("    mov     rax, 1\n")
	b.WriteString("    syscall\n")
	b.WriteString("    add     rsp, 40\n")
	b.WriteString("    ret\n")
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "No input file was provided")
		os.Exit(1)
	}

	filePath := args[1]

	ops, err := LexProgram(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	if DebugMode {
		for _, op := range ops {
			fmt.Printf("%s:%d:%d: -> %d\n", filePath, op.Loc.Row, op.Loc.Col, op.Type)
		}
	}

	fmt.Println("[INFO] Generating assembly ...")

	outContent, err := GenerateLines(ops)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to generate assembly file: %s\n", err)
		os.Exit(1)
	}

	outFilename := strings.TrimSuffix(filePath, filepath.Ext(filePath))
	outAsmFilename := outFilename + ".asm"
	outBinaryFilename := outFilename + ".out"

	err = WriteFileToString(outAsmFilename, outContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to write out-content to file: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("[INFO] Calling flat assembler ...")

	err = exec.Command("fasm", outAsmFilename).Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to run fasm: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("[INFO] Renaming binary-file %s -> %s ...\n", outFilename, outBinaryFilename)

	err = os.Rename(outFilename, outBinaryFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Failed to rename %s to %s\n", outFilename, outFilename+".out")
	}
}
