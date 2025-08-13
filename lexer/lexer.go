package lexer

import (
	"piled/token"
)

// Lexer provide methods to lex source code
type Lexer struct {
	ch         rune
	characters []rune
	pos        int
	ReadPos    int

	Line int
	Col  int
}

// New is constructor for lexer
func New(source string) *Lexer {
	l := &Lexer{
		characters: []rune(source),
		pos:        0,
		ReadPos:    0,
		Line:       1,
		Col:        0,
	}
	// ensure that ch point at a character
	l.nextChar()
	return l
}

// TODO: Introduce ParsePoint
func (l *Lexer) RestorePos(pos int) {
	l.ReadPos = pos

	// Ensure that lexer.ch point at a valid character
	l.nextChar()
}

// NextToken provide token by reading source-code character by character
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '+':
		tok.Type = token.ADD
		tok.Literal = "+"
	case '-':
		tok.Type = token.SUB
		tok.Literal = "-"
	case '*':
		tok.Type = token.MUL
		tok.Literal = "*"
	case '/':
		tok.Type = token.DIV
		tok.Literal = "/"
	case '%':
		tok.Type = token.MOD
		tok.Literal = "%"
	case '&':
		tok.Type = token.AND
		tok.Literal = "&"
	case '|':
		tok.Type = token.OR
		tok.Literal = "|"
	case '>':
		tok.Type = token.GT
		tok.Literal = ">"
	case '<':
		tok.Type = token.LT
		tok.Literal = "<"
	case '=':
		tok.Type = token.EQ
		tok.Literal = "="
	case '{':
		tok.Type = token.OCURLY
		tok.Literal = "{"
	case '}':
		tok.Type = token.CCURLY
		tok.Literal = "}"
	case rune(0):
		tok.Type = token.EOF
	default:
		if isDigit(l.ch) {
			return l.readNumeric()
		}
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdentifier(tok.Literal)
	}

	tok.Line = l.Line
	l.nextChar()

	return tok
}

func (l *Lexer) nextChar() {
	if l.ReadPos >= len(l.characters) {
		l.ch = rune(0)
	} else {
		l.ch = l.characters[l.ReadPos]
	}

	if l.ch == ('\n') {
		l.Line++
		l.Col = 0
	}

	l.pos = l.ReadPos
	l.ReadPos++
	l.Col++
}

// TODO: Delete peekChar and lexer.pos
func (l *Lexer) peekChar() rune {
	if l.pos >= len(l.characters) {
		return rune(0)
	} else {
		return l.characters[l.ReadPos]
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == rune(' ') || l.ch == rune('\n') || l.ch == rune('\r') {
		l.nextChar()
	}
}

func (l *Lexer) readNumeric() token.Token {
	out := ""
	for isDigit(l.ch) {
		out += string(l.ch)
		l.nextChar()
	}

	return token.Token{
		Type:    token.NUMBER,
		Literal: out,
		Line:    l.Line,
	}

}
func (l *Lexer) readIdentifier() string {
	out := ""

	for isAlpha(l.ch) {
		out += string(l.ch)
		l.nextChar()
	}
	return out
}

func isAlpha(c rune) bool {
	return (c <= 'z' && c >= 'a') || (c <= 'Z' && c >= 'A')
}

func isDigit(c rune) bool {
	return (c <= '9' && c >= '0')
}
