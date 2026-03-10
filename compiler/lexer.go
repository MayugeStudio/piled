package compiler

import (
  "fmt"
)

// Loc represents a location in the source code.
// It only contains offsets, and should therefore be used together with Lexer.
type Loc struct {
	// InputPath holds the source file path.
	InputPath string
	// LineNumber holds current line number.
	LineNumber int
	// LineOffset holds the offset from the beginning of the line.
	LineOffset int
}

// ParsePoint represents the current reading position while parsing.
// It differs from Loc, which simply stores a position without current parsing context.
type ParsePoint struct {
	// ch is pointed at a character being parsed.
	ch rune
	// current holds current parsing offset.
	current int
	// lineStart holds the offset of the beginning of current line being parsed.
	lineStart int
	// lineNumber holds the number of the current line (used for error reporting).
	lineNumber int
}

// Lexer is the struct that has fields to lex program.
type Lexer struct {
	InputPath string
	// CurrentPoint holds current parsing position.
	CurrentPoint ParsePoint
	// Source is source program.
	Source []rune
}

// NewLexer is constructor for lexer
func NewLexer(inputPath string, source string) *Lexer {
	point := ParsePoint{
		ch:         rune(0),
		current:    0,
		lineStart:  0,
		lineNumber: 1,
	}

	if len(source) != 0 {
		point.ch = rune(source[0])
	}

	l := &Lexer{
		InputPath:    inputPath,
		CurrentPoint: point,
		Source:       []rune(source),
	}

	// ensure that ch point at a character
	l.nextChar()
	return l
}

// NextToken provide token by reading source code character by character
func (l *Lexer) NextToken() Token {
	var tok Token
	l.skipWhitespace()

	switch l.CurrentPoint.ch {
	case '+':
		tok.Type = ADD
		tok.Literal = "+"
	case '-':
		tok.Type = SUB
		tok.Literal = "-"
	case '*':
		tok.Type = MUL
		tok.Literal = "*"
	case '/':
		tok.Type = DIV
		tok.Literal = "/"
	case '%':
		tok.Type = MOD
		tok.Literal = "%"
	case '&':
		tok.Type = AND
		tok.Literal = "&"
	case '|':
		tok.Type = OR
		tok.Literal = "|"
	case '>':
		tok.Type = GT
		tok.Literal = ">"
	case '<':
		tok.Type = LT
		tok.Literal = "<"
	case '=':
		tok.Type = EQ
		tok.Literal = "="
	case '{':
		tok.Type = OCURLY
		tok.Literal = "{"
	case '}':
		tok.Type = CCURLY
		tok.Literal = "}"
  case ';':
    tok.Type = COMMENT_START
    tok.Literal = ";"
	case rune(0):
		tok.Type = EOF
	default:
		if isDigit(l.CurrentPoint.ch) {
			return l.readNumeric()
		}
		tok.Literal = l.readIdentifier()
		tok.Type = LookupIdentifier(tok.Literal)

    // TODO: Implement IDENT
    if tok.Type == IDENT {
      fmt.Printf("ERROR:%d:%d: unknown token '%s' is found\n", l.CurrentPoint.lineNumber, l.CurrentPoint.lineStart, tok.Literal)
    }
    return tok
	}

	l.nextChar()

	return tok
}

func (l *Lexer) skipUntil(c rune) {
	for l.CurrentPoint.ch != c {
		l.nextChar()
		if l.CurrentPoint.ch == rune(0) {
			break
		}
	}
}

func (l *Lexer) nextChar() {
	if l.CurrentPoint.current >= len(l.Source) {
		l.CurrentPoint.ch = rune(0)
		return
	}
	l.CurrentPoint.ch = l.Source[l.CurrentPoint.current]

	if l.CurrentPoint.ch == '\n' {
		l.CurrentPoint.current += 1
		if l.CurrentPoint.current >= len(l.Source) {
			l.CurrentPoint.ch = rune(0)
			return
		}
		l.CurrentPoint.lineStart = l.CurrentPoint.current
		l.CurrentPoint.lineNumber += 1
	} else {
		l.CurrentPoint.current += 1
	}
}

func (l *Lexer) skipWhitespace() {
	for l.CurrentPoint.ch == rune(' ') || l.CurrentPoint.ch == rune('\n') || l.CurrentPoint.ch == rune('\r') {
		l.nextChar()
	}
}

func (l *Lexer) readNumeric() Token {
	out := ""
	for isDigit(l.CurrentPoint.ch) {
		out += string(l.CurrentPoint.ch)
		l.nextChar()
	}

	return Token{
		Type:    NUMBER,
		Literal: out,
	}

}
func (l *Lexer) readIdentifier() string {
	out := ""

	for isAlpha(l.CurrentPoint.ch) || isDigit(l.CurrentPoint.ch) {
		out += string(l.CurrentPoint.ch)
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
