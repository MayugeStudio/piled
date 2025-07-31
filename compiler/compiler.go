package compiler

import "fmt"
import "os"
import "strconv"
import "piled/token"
import "piled/lexer"
import "piled/parser"
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
func (c *Compiler) Compile(e parser.Expr) ([]runtime.OPCode, error) {
	tok := l.NextToken()
	for tok.Type != token.EOF {
		switch tok.Type {
		case token.LPAREN:
		case token.RPAREN:
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
			fmt.Println("currently not supported")

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
	return c.code, nil
}

func (c *Compiler) compileLiteral(l *parser.Literal) {
	c.emit(runtime.PUSH, runtime.OPCode(l.Value))
}

// TODO: Add symbol - opcode relation mapping
func (c *Compiler) compileSymbol(s *parser.Symbol) error {
	switch s.Token.Type {
	case token.PRINT:
		c.emit(runtime.PRINT)
	case token.ADD:
		c.emit(runtime.ADD)
	case token.SUB:
		c.emit(runtime.SUB)
	case token.GT:
		c.emit(runtime.GT)
	case token.LT:
		c.emit(runtime.LT)
	case token.EQ:
		c.emit(runtime.EQ)
	default:
		return fmt.Errorf("unknown token has been found at compile time: %s", s.Token.Type)
	}
	return nil
}

func (c *Compiler) compileList(lst *parser.List) error {
	for _, elem := range lst.Elements {
		switch e := elem.(type) {
		case *parser.Literal:
			c.compileLiteral(e)
		case *parser.Symbol:
			if err := c.compileSymbol(e); err != nil {
				return err
			}
		case *parser.List:
			if err := c.compileList(e); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported parser type")
		}
	}
	return nil
}

func (c *Compiler) emit(op runtime.OPCode, val ...runtime.OPCode) {
	c.code = append(c.code, op)
	c.code = append(c.code, val...)
}
