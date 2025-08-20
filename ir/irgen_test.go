package ir

import (
	"piled/lexer"
	"reflect"
	"testing"
)

func num(v int) *Number {
	return &Number{Value: v}
}

func bin(bkind BinKind) *Binop {
	return &Binop{Bkind: bkind}
}

func Test_CompileProgram(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []Op
	}{
		{
			"Number",
			"1 print",
			[]Op{
				num(1),
				&Print{},
			},
		},
		{
			"Add",
			"1 1 + print",
			[]Op{
				num(1),
				num(1),
				&Binop{Bkind: Add},
				&Print{},
			},
		},
		{
			"Sub",
			"1 1 - print",
			[]Op{
				num(1),
				num(1),
				&Binop{Bkind: Sub},
				&Print{},
			},
		},
		{
			"Mul",
			"1 1 * print",
			[]Op{
				num(1),
				num(1),
				&Binop{Bkind: Mul},
				&Print{},
			},
		},
		{
			"Div",
			"1 1 / print",
			[]Op{
				num(1),
				num(1),
				&Binop{Bkind: Div},
				&Print{},
			},
		},
		{
			"if",
			"1 if {34} print",
			[]Op{
				num(1),
				&JmpIfNotLabel{Label: 0},
				num(34),
				&Label{Label: 0},
				&Print{},
			},
		},
		{
			"if-else",
			"1 if {34} else {35} print",
			[]Op{
				num(1),
				&JmpIfNotLabel{Label: 0},
				num(34),
				&JmpLabel{Label: 1},
				&Label{Label: 0},
				num(35),
				&Label{Label: 1},
				&Print{},
			},
		},
		{
			"dup",
			"69 dup print print",
			[]Op{
				num(69),
				&Dup{},
				&Print{},
				&Print{},
			},
		},
		{
			"swap",
			"69 420 swap",
			[]Op{
				num(69),
				num(420),
				&Swap{},
			},
		},
		{
			"rot",
			"1 2 3 rot",
			[]Op{
				num(1),
				num(2),
				num(3),
				&Rot{},
			},
		},
		{
			"drop",
			"69 drop",
			[]Op{
				num(69),
				&Drop{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New("test.piled", tt.in)
			g := NewIrGen()
			err := g.CompileProgram(l)
			if err != nil {
				t.Errorf("g.CompileProgram() error = %v", err)
				return
			}

			got := g.Ops

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want = %v, got = %v", tt.want, got)
				return
			}
		})
	}
}

func Test_compileIF_SingleIF(t *testing.T) {
	// if is already found by compileProgram
	in := "{1 1 +}"
	want := []Op{
		&JmpIfNotLabel{Label: 0},
		num(1),
		num(1),
		bin(Add),
		&Label{Label: 0},
	}
	l := lexer.New("test.piled", in)
	g := NewIrGen()

	err := g.compileIF(l)
	if err != nil {
		t.Errorf("compileIF() error = %v", err)
		return
	}

	got := g.Ops

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want = %v, got = %v", want, got)
		return
	}
}

func Test_compileBINDING(t *testing.T) {}
