package asm

import (
	"fmt"
	"os"
	"strings"
)

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
