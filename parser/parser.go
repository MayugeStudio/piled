package parser

import "piled/token"
import "fmt"
import "strconv"

// Parser contains tokens and position of a current token
type Parser struct {
	tokens []token.Token
	pos    int
}

// New is constructor for parser and return a Parser instance from []token.Token
func New(tokens []token.Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

// ParseExpr generate new Expr from tokens
func (p *Parser) ParseExpr() (Expr, error) {
	tok := p.next()

	switch tok.Type {
	case token.NUMBER:
		{
			val, _ := strconv.Atoi(tok.Literal)
			return &Literal{Value: int(val)}, nil
		}
	// TODO: Too ugly to read
	case token.IDENT, token.ADD, token.SUB, token.EQ, token.GT, token.LT, token.PRINT:
		{
			return p.parseSymbol(tok), nil
		}
	case token.LPAREN:
		{
			elems := []Expr{}
			for p.peek().Type != token.RPAREN && p.peek().Type != token.EOF {
				expr, err := p.ParseExpr()
				if err != nil {
					return nil, err
				}
				elems = append(elems, expr)
			}
			if p.peek().Type != token.RPAREN {
				return nil, fmt.Errorf("Expected ')', got %v", p.peek().Literal)
			}
			p.next() // consume RPAREN
			return &List{Elements: elems}, nil
		}
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.peek().Literal)
	}
}

func (p *Parser) peek() token.Token {
	if p.pos >= len(p.tokens) {
		return token.Token{}
	}

	return p.tokens[p.pos]
}

func (p *Parser) next() token.Token {
	t := p.peek()
	p.pos++
	return t
}

func (p *Parser) parseSymbol(tok token.Token) Expr {
	return &Symbol{Token: tok}
}
