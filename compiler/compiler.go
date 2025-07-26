package compiler

import "fmt"
import "piled/token"
import "piled/expr"
import "piled/opcode"

type Compiler struct {
	code []opcode.Code
}

func New() *Compiler {
	return &Compiler{code: make([]opcode.Code, 0)}
}
func (c *Compiler) Compile(e expr.Expr) ([]opcode.Code, error) {
	switch v := e.(type) {
	case *expr.Literal:
		c.compileLiteral(v)
	case *expr.Symbol:
		if err := c.compileSymbol(v); err != nil {
			return nil, err
		}
	case *expr.List:
		if err := c.compileList(v); err != nil {
			return nil, err
		}
	}
	return c.code, nil
}
func (c *Compiler) compileLiteral(l *expr.Literal) {
	c.emit(opcode.PUSH, opcode.Code(l.Value))
}

// TODO: Add symbol - opcode relation mapping
func (c *Compiler) compileSymbol(s *expr.Symbol) error {
	switch s.Token.Type {
	case token.PRINT:
		c.emit(opcode.PRINT)
	case token.ADD:
		c.emit(opcode.ADD)
	case token.SUB:
		c.emit(opcode.SUB)
	case token.GT:
		c.emit(opcode.GT)
	case token.LT:
		c.emit(opcode.LT)
	case token.EQ:
		c.emit(opcode.EQ)
	default:
		return fmt.Errorf("unknown token has been found at compile time: %s", s.Token.Type)
	}
	return nil
}

func (c *Compiler) compileList(lst *expr.List) error {
	for _, elem := range lst.Elements {
		switch e := elem.(type) {
		case *expr.Literal:
			c.compileLiteral(e)
		case *expr.Symbol:
			if err := c.compileSymbol(e); err != nil {
				return err
			}
		case *expr.List:
			if err := c.compileList(e); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported expr type")
		}
	}
	return nil
}
func (c *Compiler) emit(op opcode.Code, val ...opcode.Code) {
	c.code = append(c.code, op)
	c.code = append(c.code, val...)
}
