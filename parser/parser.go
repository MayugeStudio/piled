package parser

import "piled/token"
import "piled/expr"
import "fmt"

type Parser struct {
	tokens       []*token.Token
	currentToken *token.Token
	peekToken    *token.Token
	index        int
}

func New(tokens []*token.Token) *Parser{
	return &Parser{
		tokens:       tokens,
		currentToken: tokens[0],
		index:        0,
	}
}

func (p *Parser) Parse() (*expr.List, error) {
	root := expr.List{}
	err := p.parseList(&root)
	if err != nil {
		return nil, err
	}
	return &root, nil
}

func (p *Parser) advance() {
	if p.currentToken.Type == token.EOF {
		return
	}
	p.index += 1
	p.currentToken = p.tokens[p.index]
}

func (p *Parser) peek() {
	if p.currentToken.Type == token.EOF {
		p.peekToken = nil
	}
	p.peekToken = p.tokens[p.index + 1]
}

func (p *Parser) parseList(list *expr.List) error {
	// (0 1 2 3 4 5)
	if p.currentToken.Type != token.LPAREN {
		return fmt.Errorf("Expected '('.")
	}
	for p.peek(); p.peekToken.Type != token.RPAREN;{
		p.advance()
		switch p.currentToken.Type {
		case token.IDENTIFIER, token.NUMBER, token.STRING: {
			e := &expr.Literal{}
			list.Elements = append(list.Elements, e)
		}
		default: {
			return fmt.Errorf("Unexpected token '%s'", p.currentToken.Value)
		}
		}
	}

	p.advance()
	if p.currentToken.Type != token.RPAREN {
		return fmt.Errorf("Expected ')'.")
	}
	p.advance()
	return nil
}

