package compiler

import "fmt"
import "os"
import "strconv"
import "piled/token"
import "piled/lexer"
import "piled/runtime"


// Compiler contains opcodes
type Compiler struct {
	l        *lexer.Lexer
	code     []runtime.OPCode
}

// New is constructor for Compiler
func New(l *lexer.Lexer) *Compiler {
	return &Compiler{
		l: l,
		code: make([]runtime.OPCode, 0),
	}
}

// Write output an array of opcode to specified filepath
func (c *Compiler) Write(path string) error {
	out := make([]byte, 0, 1024)
	for _, c := range c.code {
		out = append(out, byte(c))
	}
	return os.WriteFile(path, out, 0644)
}

// Compile generate opcode based on parser.Expr
func (c *Compiler) Compile() []runtime.OPCode {
	for tok := c.l.NextToken(); tok.Type != token.EOF; tok = c.l.NextToken() {
		switch tok.Type {
		case token.LPAREN: // Currently ignored
		case token.RPAREN: // Currently ignored
		case token.ADD:
			c.code = append(c.code, runtime.ADD)
		case token.SUB:
			c.code = append(c.code, runtime.SUB)
		case token.GT:
			c.code = append(c.code, runtime.GT)
		case token.LT:
			c.code = append(c.code, runtime.LT)
		case token.EQ:
			c.code = append(c.code, runtime.EQ)
		case token.IDENT:
			fmt.Printf("currently not supported: %s\n", tok.Literal)
		case token.NUMBER:
			{
				value, _ := strconv.Atoi(tok.Literal)
				c.code = append(c.code, runtime.PUSH)
				c.code = append(c.code, runtime.OPCode(value))
			}
		default:
			fmt.Printf("unhandled token: %v\n", tok)
		}
	}
	return c.code
}

