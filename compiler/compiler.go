package compiler

import "fmt"
import "os"
import "strconv"
import "piled/token"
import "piled/lexer"
import "piled/runtime"

// Compiler contains instructions
type Compiler struct {
	l    *lexer.Lexer
	code []runtime.Instruction
}

// New is constructor for Compiler
func New(l *lexer.Lexer) *Compiler {
	return &Compiler{
		l:    l,
		code: make([]runtime.Instruction, 0),
	}
}

// Compile generate instruction from source-code
func (c *Compiler) Compile() []runtime.Instruction {
	backpatch_stack := make([]int, 0)
	for tok := c.l.NextToken(); tok.Type != token.EOF; tok = c.l.NextToken() {
		switch tok.Type {
		case token.ADD:
			c.emit(runtime.Instruction{Kind: runtime.ADD, Args: nil})
		case token.SUB:
			c.emit(runtime.Instruction{Kind: runtime.SUB, Args: nil})
		case token.MUL:
			c.emit(runtime.Instruction{Kind: runtime.MUL, Args: nil})
		case token.DIV:
			c.emit(runtime.Instruction{Kind: runtime.DIV, Args: nil})
		case token.MOD:
			c.emit(runtime.Instruction{Kind: runtime.MOD, Args: nil})
		case token.AND:
			c.emit(runtime.Instruction{Kind: runtime.AND, Args: nil})
		case token.OR:
			c.emit(runtime.Instruction{Kind: runtime.OR, Args: nil})
		case token.SHL:
			c.emit(runtime.Instruction{Kind: runtime.SHL, Args: nil})
		case token.SHR:
			c.emit(runtime.Instruction{Kind: runtime.SHR, Args: nil})
		case token.GT:
			c.emit(runtime.Instruction{Kind: runtime.GT, Args: nil})
		case token.LT:
			c.emit(runtime.Instruction{Kind: runtime.LT, Args: nil})
		case token.EQ:
			c.emit(runtime.Instruction{Kind: runtime.EQ, Args: nil})
		case token.LPAREN: // Currently ignored
			fmt.Printf("lparen is currently not supported: %v\n", tok)
		case token.RPAREN: // Currently ignored
			fmt.Printf("rparen is currently not supported: %v\n", tok)
		case token.IF:
			c.emit(runtime.Instruction{Kind: runtime.JMPIF, Args: []int{}})
			// save current ip onto the stack to backpatch it
			backpatch_stack = append(backpatch_stack, len(c.code)-1)
		case token.ELSE:
			// append JMP instruction
			c.emit(runtime.Instruction{Kind: runtime.JMP, Args: []int{}})
			// save current ip onto the stack to backpatch it
			backpatch_stack = append(backpatch_stack, len(c.code)-1)

			// jump destination of 'if'
			c.emit(runtime.Instruction{Kind: runtime.NOP, Args: nil}) // for backpatching
			else_addr := len(c.code) - 1

			// backpatching if-block
			if_addr := backpatch_stack[0]
			backpatch_stack = backpatch_stack[1:]
			c.code[if_addr].Args = append(c.code[if_addr].Args, else_addr)
		case token.END:
			c.emit(runtime.Instruction{Kind: runtime.NOP, Args: nil})
			end_addr := len(c.code) - 1
			// backpatching block
			block_addr := backpatch_stack[0]
			backpatch_stack = backpatch_stack[1:]
			c.code[block_addr].Args = append(c.code[end_addr].Args, end_addr)
		case token.PRINT:
			c.emit(runtime.Instruction{Kind: runtime.PRINT, Args: nil})
		case token.NUMBER:
			{
				value, _ := strconv.Atoi(tok.Literal)
				c.emit(runtime.Instruction{Kind: runtime.PUSH, Args: []int{value}})
			}
		case token.IDENT:
			fmt.Printf("ident is currently not supported: %v\n", tok)
		default:
			fmt.Printf("unhandled token: %v\n", tok)
		}
	}
	return c.code
}

func (c *Compiler) emit(i runtime.Instruction) {
	c.code = append(c.code, i)
}

// Write output an array of instruction to specified filepath
func (p *Compiler) Write(path string) error {
	out := make([]byte, 0, 1024)
	for _, inst := range p.code {
		out = append(out, byte(inst.Kind))
	}
	return os.WriteFile(path, out, 0644)
}
