package ir

import (
	"strconv"
	"fmt"
	"piled/lexer"
	"piled/token"
)

type OpKind string
type BinKind string

const (
	Number OpKind   = "Number"
	Bind            = "Bind"
	Binop           = "Binop"
	Label           = "Label"
	JmpLabel        = "JmpLabel"
	JmpIfNotLabel   = "JmpIfNotLabel"
	Print           = "Print"
)

const (
	// Bin op
	Add BinKind     = "Add"
	Sub             = "Sub"
	Mul             = "Mul"
	Div             = "Div"
	Mod             = "Mod" 
	Gt              = "Gt"
	Lt              = "Lt"
	Eq              = "Eq"
	// TODO: Rename And to BitAnd
	And             = "And"
	// TODO: Rename Or to BitOr
	Or              = "Or"
)

type Op interface {
	Kind() OpKind
	String() string
}

// -------------------- NumberImpl -------------------- 

type NumberImpl struct {
	kind OpKind
	Value int
}

func (p *NumberImpl) Kind() OpKind {
	return p.kind
}

func (p *NumberImpl) String() string {
	return string(p.kind) + "(" + strconv.Itoa(p.Value) + ")"
}

// -------------------- BinopImpl -------------------- 

type BinopImpl struct {
	kind OpKind
	Bkind BinKind
}

func (p *BinopImpl) Kind() OpKind {
	return p.kind
}

func (p *BinopImpl) String() string {
	return string(p.kind) + "(" + string(p.Bkind) + ")"
}

// -------------------- LabelImpl -------------------- 

type LabelImpl struct {
	kind  OpKind
	label int
}

func (p *LabelImpl) Kind() OpKind {
	return p.kind
}

func (p *LabelImpl) String() string {
	return string(p.kind) + "(" + strconv.Itoa(p.label) + ")"
}

// -------------------- JmpLabelImpl -------------------- 

type JmpLabelImpl struct {
	kind OpKind
	label int
}

func (p *JmpLabelImpl) Kind() OpKind {
	return p.kind
}

func (p *JmpLabelImpl) String() string {
	return string(p.kind) + "(" + strconv.Itoa(p.label) + ")"
}

// -------------------- JmpIfNotLabelImpl -------------------- 

type JmpIfNotLabelImpl struct {
	kind  OpKind
	label int
}

func (p *JmpIfNotLabelImpl) Kind() OpKind {
	return p.kind
}

func (p *JmpIfNotLabelImpl) String() string {
	return string(p.kind) + "(" + strconv.Itoa(p.label) + ")"
}

// -------------------- Print -------------------- 

type PrintImpl struct {
	kind OpKind
}

func (p *PrintImpl) Kind() OpKind {
	return p.kind
}

func (p *PrintImpl) String() string {
	return string(p.kind) + "()"
}

// -------------------- IrGen -------------------- 

type IrGen struct {
	ops []Op
	labelCount int
}

func NewIrGen() *IrGen {
	return &IrGen { ops: make([]Op, 0) }
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

func (p *IrGen) compileToken(l *lexer.Lexer, tok token.Token) error {
	switch tok.Type {
		// Literals
		case token.NUMBER:
			value, _ := strconv.Atoi(tok.Literal)
			p.emit(&NumberImpl{ kind: Number, Value: value })
		case token.IDENT:
			// TODO: Report invalid Ident through diagnostics
		// Binops
		case token.ADD: p.emitBin(Add)
		case token.SUB: p.emitBin(Sub)
		case token.MUL: p.emitBin(Mul)
		case token.DIV: p.emitBin(Div)
		case token.MOD: p.emitBin(Mod)
		case token.AND: p.emitBin(And)
		case token.OR:  p.emitBin(Or)
		case token.GT:  p.emitBin(Gt)
		case token.LT:  p.emitBin(Lt)
		case token.EQ:  p.emitBin(Eq)
		case token.IF:
			if err := p.compileIF(l); err != nil {
				return err
			}
		//case token.LET:
		//	p.compileBINDING(l)
		case token.PRINT:
			p.emit(&PrintImpl{kind: Print})
		//case token.SHL:
		//case token.SHR:
		case token.OCURLY:
		case token.CCURLY:
		case token.EOF:
		default:
			return fmt.Errorf("unhandled token: %v %v", tok, p.ops)
	}

	return nil
}

func (p *IrGen) compileIF(l *lexer.Lexer) error {
	// TODO: Introduce allocate label function
	else_label := p.labelCount
	p.emit(&JmpIfNotLabelImpl{ kind: JmpIfNotLabel, label: p.labelCount })
	p.labelCount += 1
	
	// TODO: Introduce block by using curly braces
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
	savePoint := l.ReadPos-1
	tok = l.NextToken() // expect else
	if tok.Type == token.ELSE {
		out_label := p.labelCount
		p.labelCount += 1
		p.emit(&JmpLabelImpl{ kind: JmpLabel, label: out_label })
		p.emit(&LabelImpl{ kind: Label, label: else_label })
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
		p.emit(&LabelImpl{ kind: Label, label: out_label })
	} else {
		l.RestorePos(savePoint)
		p.emit(&LabelImpl{ kind: Label, label: else_label })
	}

	return nil
}

func (p *IrGen) compileBINDING(l *lexer.Lexer) {
}

func (p *IrGen) emit(ir Op) {
	p.ops = append(p.ops, ir)
}

func (p *IrGen) emitBin(b BinKind) {
	p.emit(&BinopImpl{kind: Binop, Bkind: b})
}

