package expr

// (0 print)
// ((1 1 +) print)
// (((1 1 +) (1 1 +) +) print)

type Literal struct {
	Value any
}

type List struct {
	Elements   []*Literal
}

