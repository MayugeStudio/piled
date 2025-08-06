package compiler

import "fmt"
import "os"
import "strconv"
import "piled/token"
import "piled/lexer"
import "piled/runtime"

// Compiler contains opcodes
type Compiler struct {
	l    *lexer.Lexer
	code []runtime.OPCode
}

// New is constructor for Compiler
func New(l *lexer.Lexer) *Compiler {
	return &Compiler{
		l:    l,
		code: make([]runtime.OPCode, 0),
	}
}

// Compile generate opcode from source-code
func (c *Compiler) Compile() []runtime.OPCode {
	backpatch_stack := make([]runtime.OPCode, 0)
	for tok := c.l.NextToken(); tok.Type != token.EOF; tok = c.l.NextToken() {
		switch tok.Type {
		case token.ADD:
			c.code = append(c.code, runtime.ADD)
		case token.SUB:
			c.code = append(c.code, runtime.SUB)
		case token.MUL:
			c.code = append(c.code, runtime.MUL)
		case token.DIV:
			c.code = append(c.code, runtime.DIV)
		case token.MOD:
			c.code = append(c.code, runtime.MOD)
		case token.AND:
			c.code = append(c.code, runtime.AND)
		case token.OR:
			c.code = append(c.code, runtime.OR)
		case token.SHL:
			c.code = append(c.code, runtime.SHL)
		case token.SHR:
			c.code = append(c.code, runtime.SHR)
		case token.GT:
			c.code = append(c.code, runtime.GT)
		case token.LT:
			c.code = append(c.code, runtime.LT)
		case token.EQ:
			c.code = append(c.code, runtime.EQ)
		case token.LPAREN: // Currently ignored
			fmt.Printf("lparen is currently not supported: %v\n", tok)
		case token.RPAREN: // Currently ignored
			fmt.Printf("rparen is currently not supported: %v\n", tok)
		case token.IF:
			c.code = append(c.code, runtime.JMPIF)
			c.code = append(c.code, runtime.OPCode(0)) // for backpatching
			// save current ip onto the stack to backpatch it
			backpatch_stack = append(backpatch_stack, runtime.OPCode(len(c.code)-1))
		case token.ELSE:
			// append JMP opcode
			c.code = append(c.code, runtime.JMP)
			c.code = append(c.code, runtime.OPCode(0)) // for backpatching
			// save current ip onto the stack to backpatch it
			backpatch_stack = append(backpatch_stack, runtime.OPCode(len(c.code)-1))

			// jump destination of 'if'
			c.code = append(c.code, runtime.NOP)
			else_addr := len(c.code) - 1

			// backpatching if-block
			if_addr := backpatch_stack[0]
			backpatch_stack = backpatch_stack[1:]
			c.code[if_addr] = runtime.OPCode(else_addr)
		case token.END:
			c.code = append(c.code, runtime.NOP)
			end_addr := len(c.code) - 1
			// backpatching block
			block_addr := backpatch_stack[0]
			backpatch_stack = backpatch_stack[1:]
			c.code[block_addr] = runtime.OPCode(end_addr)
		case token.PRINT:
			c.code = append(c.code, runtime.PRINT)
		case token.NUMBER:
			{
				value, _ := strconv.Atoi(tok.Literal)
				c.code = append(c.code, runtime.PUSH)
				c.code = append(c.code, runtime.OPCode(value))
			}
		case token.IDENT:
			fmt.Printf("ident is currently not supported: %v\n", tok)
		default:
			fmt.Printf("unhandled token: %v\n", tok)
		}
	}
	return c.code
}

// Write output an array of opcode to specified filepath
func (c *Compiler) Write(path string) error {
	out := make([]byte, 0, 1024)
	for _, c := range c.code {
		out = append(out, byte(c))
	}
	return os.WriteFile(path, out, 0644)
}
