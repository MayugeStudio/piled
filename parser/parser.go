package parser

import "piled/token"
import "piled/expr"
import "fmt"
import "strconv"

type Parser struct {
	tokens []token.Token
	pos    int
}

func New(tokens []token.Token) *Parser {
	return &Parser{tokens: tokens, pos: 0}
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

func (p *Parser) ParseExpr() (expr.Expr, error) {
	tok := p.next()

	switch tok.Type {
	case token.NUMBER:
		{
			val, _ := strconv.Atoi(tok.Literal)
			return &expr.Literal{Value: int(val)}, nil
		}
	// TODO: Too ugly to read
	case token.IDENT, token.ADD, token.SUB, token.EQ, token.GT, token.LT, token.PRINT:
		{
			return p.parseSymbol(tok), nil
		}
	case token.LPAREN:
		{
			elems := []expr.Expr{}
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
			return &expr.List{Elements: elems}, nil
		}
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.peek().Literal)
	}
}

func (p *Parser) parseSymbol(tok token.Token) expr.Expr {
	return &expr.Symbol{Token: tok}
}
