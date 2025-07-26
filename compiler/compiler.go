package compiler

import "fmt"
import "piled/expr"
import "piled/opcode"

type Compiler struct {
	code []int
}

func New() *Compiler {
	return &Compiler{code: make([]int, 0)}
}
func (c *Compiler) Compile(e expr.Expr) ([]int, error) {
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
	c.emit(opcode.PUSH, l.Value)
}

// TODO: Add symbol - opcode relation mapping
func (c *Compiler) compileSymbol(s *expr.Symbol) error {
	switch s.Name {
	case "print":
		c.emit(opcode.PRINT)
	case "+":
		c.emit(opcode.ADD)
	case "-":
		c.emit(opcode.SUB)
	default:
		return fmt.Errorf("unknown symbol: %s", s.Name)
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
func (c *Compiler) emit(op int, args ...int) {
	c.code = append(c.code, op)
	c.code = append(c.code, args...)
}
