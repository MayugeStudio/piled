package ir

import (
	"testing"
	"reflect"
	"piled/lexer"
)

func num(v int) *NumberImpl {
	return &NumberImpl{ kind: Number, Value: v }
}

func bin(bkind BinKind) *BinopImpl {
	return &BinopImpl{ kind: Binop, Bkind: bkind }
}

func Test_CompileProgram(t *testing.T) {
	tests := []struct{
		name string
		in   string
		want  []Op
	}{
		{
			"Number",
			"1 print",
			[]Op{
				num(1),
				&PrintImpl{ kind: Print },
			},
		},
		{
			"Add",
			"1 1 + print",
			[]Op{
				num(1),
				num(1),
				&BinopImpl{ kind: Binop, Bkind: Add },
				&PrintImpl{ kind: Print },
			},
		},
		{
			"Sub",
			"1 1 - print",
			[]Op{
				num(1),
				num(1),
				&BinopImpl{ kind: Binop, Bkind: Sub },
				&PrintImpl{ kind: Print },
			},
		},
		{
			"Mul",
			"1 1 * print",
			[]Op{
				num(1),
				num(1),
				&BinopImpl{ kind: Binop, Bkind: Mul },
				&PrintImpl{ kind: Print },
			},
		},
		{
			"Div",
			"1 1 / print",
			[]Op{
				num(1),
				num(1),
				&BinopImpl{ kind: Binop, Bkind: Div },
				&PrintImpl{ kind: Print },
			},
		},
		{
			"if",
			"1 if {34} print",
			[]Op{
				num(1),
				&JmpIfNotLabelImpl{ kind: JmpIfNotLabel, label: 0 },
				num(34),
				&LabelImpl{ kind: Label, label: 0 },
				&PrintImpl{ kind: Print },
			},
		},
		{
			"if-else",
			"1 if {34} else {35} print",
			[]Op{
				num(1),
				&JmpIfNotLabelImpl{ kind: JmpIfNotLabel, label: 0 },
				num(34),
				&JmpLabelImpl{ kind: JmpLabel, label: 1 },
				&LabelImpl{ kind: Label, label: 0 },
				num(35),
				&LabelImpl{ kind: Label, label: 1 },
				&PrintImpl{ kind: Print },
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.New(tt.in)
			g := NewIrGen()
			err := g.CompileProgram(l)
			if err != nil {
				t.Errorf("g.CompileProgram() error = %v", err)
				return
			}

			got := g.ops

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
		&JmpIfNotLabelImpl{ kind: JmpIfNotLabel, label: 0 },
		num(1),
		num(1),
		bin(Add),
		&LabelImpl{ kind: Label, label: 0 },
	}
	l := lexer.New(in)
	g := NewIrGen()

	err := g.compileIF(l)
	if err != nil {
		t.Errorf("compileIF() error = %v", err)
		return
	}

	got := g.ops

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want = %v, got = %v", want, got)
		return
	}
}

func Test_compileBINDING(t *testing.T) {}


