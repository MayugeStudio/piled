package ir

import (
	"fmt"
	"piled/lexer"
	"piled/token"
	"strconv"
)

type BinKind string

const (
	// Bin op
	Add BinKind = "Add"
	Sub         = "Sub"
	Mul         = "Mul"
	Div         = "Div"
	Mod         = "Mod"
	Gt          = "Gt"
	Lt          = "Lt"
	Eq          = "Eq"
	// TODO: Rename And to BitAnd
	// TODO: Rename Or to BitOr
	And = "And"
	Or  = "Or"
	Shl = "BitShl"
	Shr = "BitShr"
)

type Op interface {
	String() string
}

// -------------------- Number --------------------

type Number struct {
	Value int
}

func (p *Number) String() string {
	return "Number(" + strconv.Itoa(p.Value) + ")"
}

// -------------------- Binop --------------------

type Binop struct {
	Bkind BinKind
}

func (p *Binop) String() string {
	return "Binop(" + string(p.Bkind) + ")"
}

// -------------------- Bind --------------------

type Bind struct{}

func (p *Bind) String() string {
	return "Bind()"
}

// -------------------- Label --------------------

type Label struct {
	Label int
}

func (p *Label) String() string {
	return "Label(" + strconv.Itoa(p.Label) + ")"
}

// -------------------- JmpLabel --------------------

type JmpLabel struct {
	Label int
}

func (p *JmpLabel) String() string {
	return "JmpLabel(" + strconv.Itoa(p.Label) + ")"
}

// -------------------- JmpIfNotLabel --------------------

type JmpIfNotLabel struct {
	Label int
}

func (p *JmpIfNotLabel) String() string {
	return "JmpIfNotLabel(" + strconv.Itoa(p.Label) + ")"
}

// -------------------- Print --------------------

type Print struct{}

func (p *Print) String() string {
	return "Print()"
}

// -------------------- Dup --------------------

type Dup struct{}

func (p *Dup) String() string {
	return "Dup()"
}

// -------------------- Swap --------------------

type Swap struct{}

func (p *Swap) String() string {
	return "Swap()"
}

// -------------------- Rot --------------------

type Rot struct{}

func (p *Rot) String() string {
	return "Rot()"
}

// -------------------- IrGen --------------------

type IrGen struct {
	Ops        []Op
	labelCount int
}

func NewIrGen() *IrGen {
	return &IrGen{Ops: make([]Op, 0)}
}

func (p *IrGen) CompileProgram(l *lexer.Lexer) error {
	var tok token.Token
	var err error
	for tok.Type != token.EOF {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}

		err = p.compileToken(l, tok)
		if err != nil {
			return err
		}
	}
	return nil
}

// TODO: Report invalid Ident through diagnostics
func (p *IrGen) compileToken(l *lexer.Lexer, tok token.Token) error {
	switch tok.Type {
	// Literals
	case token.NUMBER:
		value, _ := strconv.Atoi(tok.Literal)
		p.emit(&Number{Value: value})
	case token.IDENT:
	// Binops
	case token.ADD:
		p.emitBin(Add)
	case token.SUB:
		p.emitBin(Sub)
	case token.MUL:
		p.emitBin(Mul)
	case token.DIV:
		p.emitBin(Div)
	case token.MOD:
		p.emitBin(Mod)
	case token.AND:
		p.emitBin(And)
	case token.OR:
		p.emitBin(Or)
	case token.GT:
		p.emitBin(Gt)
	case token.LT:
		p.emitBin(Lt)
	case token.EQ:
		p.emitBin(Eq)
	case token.IF:
		if err := p.compileIF(l); err != nil {
			return err
		}
	case token.WHILE:
		if err := p.compileWHILE(l); err != nil {
			return err
		}
	//case token.LET:
	//	p.compileBINDING(l)
	case token.PRINT:
		p.emit(&Print{})
	case token.DUP:
		p.emit(&Dup{})
	case token.SWAP:
		p.emit(&Swap{})
	case token.ROT:
		p.emit(&Rot{})
	//case token.SHL:
	//case token.SHR:
	case token.OCURLY:
	case token.CCURLY:
	case token.EOF:
	default:
		return fmt.Errorf("IR-GEN: unhandled token: %v %v", tok, p.Ops)
	}

	return nil
}

func (p *IrGen) allocateLabel() int {
	result := p.labelCount
	p.labelCount += 1
	return result

}

func (p *IrGen) compileIF(l *lexer.Lexer) error {
	else_label := p.allocateLabel()
	p.emit(&JmpIfNotLabel{Label: else_label})

	// Parse if block
	var err error
	tok := l.NextToken() // expect OCurly
	if tok.Type == token.OCURLY {
		for {
			// TODO: Introduce an Expect(TokenType) helper method in the lexer to easily return errors.
			tok = l.NextToken() // expect OCurly
			if tok.Type == token.CCURLY {
				break
			}
			err = p.compileToken(l, tok)
			if err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("expected } but got %s", tok.Type)
	}

	// Parse else block (if it exists)
	savePoint := l.CurrentPoint
	tok = l.NextToken() // expect else
	if tok.Type == token.ELSE {
		out_label := p.allocateLabel()
		p.emit(&JmpLabel{Label: out_label})
		p.emit(&Label{Label: else_label})
		tok := l.NextToken() // expect OCurly
		if tok.Type == token.OCURLY {
			for {
				tok = l.NextToken()
				if tok.Type == token.CCURLY {
					break
				}
				err = p.compileToken(l, tok)
				if err != nil {
					return err
				}
			}
		} else {
			return fmt.Errorf("expected } but got %s", tok.Type)
		}
		p.emit(&Label{Label: out_label})
	} else {
		l.CurrentPoint = savePoint
		p.emit(&Label{Label: else_label})
	}

	return nil
}

func (p *IrGen) compileWHILE(l *lexer.Lexer) error {
	top_label := p.allocateLabel()
	out_label := p.allocateLabel()

	p.emit(&Label{Label: top_label})

	var err error
	var tok token.Token

	// condition
	for {
		tok = l.NextToken() // expect OCurly
		if tok.Type == token.OCURLY {
			break
		}
		err = p.compileToken(l, tok)
		if err != nil {
			return err
		}
	}
	p.emit(&JmpIfNotLabel{Label: out_label})

	// body
	for {
		tok = l.NextToken() // expect OCurly
		if tok.Type == token.CCURLY {
			break
		}
		err = p.compileToken(l, tok)
		if err != nil {
			return err
		}

	}

	p.emit(&JmpLabel{Label: top_label})
	p.emit(&Label{Label: out_label})

	return nil
}

func (p *IrGen) compileBINDING(l *lexer.Lexer) {
}

func (p *IrGen) emit(ir Op) {
	p.Ops = append(p.Ops, ir)
}

func (p *IrGen) emitBin(b BinKind) {
	p.emit(&Binop{Bkind: b})
}
