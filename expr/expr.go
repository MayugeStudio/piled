package expr

import "strconv"
import "strings"

type SymbolType int

const (
	SymbolUnkown SymbolType = iota
	SymbolPrint
	SymbolAdd
	SymbolSub
)

var Keywords = map[string]SymbolType{
	"print": SymbolPrint,
	"+":     SymbolAdd,
	"-":     SymbolSub,
}

type Expr interface {
	String() string
	Equal(Expr)  bool
}

type Literal struct {
	Value int
}

func (l *Literal) String() string {
	return strconv.Itoa(l.Value)
}

func (l *Literal) Equal(other Expr) bool {
	o, ok := other.(*Literal)
	return ok && l.Value == o.Value
}

type Symbol struct {
	Name string
	Type SymbolType
}

func (s *Symbol) String() string {
	return s.Name
}

func (s *Symbol) Equal(other Expr) bool {
	o, ok := other.(*Symbol)
	return ok && s.Name == o.Name && s.Type == o.Type
}

type List struct {
	Elements   []Expr
}

func (l *List) String() string {
	var sb strings.Builder
	sb.WriteByte('(')
	for i, e := range l.Elements {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(e.String())
	}
	sb.WriteByte(')')
	return sb.String()
}

func (l *List) Equal(other Expr) bool {
	o, ok := other.(*List)
	if !ok || len(l.Elements) != len(o.Elements) {
		return false
	}

	for i := range o.Elements {
		if !l.Elements[i].Equal(o.Elements[i]) {
			return false
		}
	}
	return true
}
