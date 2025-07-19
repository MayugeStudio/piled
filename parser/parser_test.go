package parser

import "testing"
import "reflect"
import "piled/expr"
import "piled/token"

func TestParse(t *testing.T) {
	tests := map[string]struct{
		in     []*token.Token
		want   *expr.List
	} {
		"nothing": {
			in: []*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.EOF, Line: 1},
			},
			want: &expr.List{},
		},
		"one value": {
			in: []*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.NUMBER, Value: 69, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.EOF, Line: 1},
			},
			want: &expr.List{
				Elements: []*expr.Literal{{Value: 69}},
			},
		},
		"serveral values": {
			in: []*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.NUMBER, Value: 100, Line: 1},
				{Type: token.NUMBER, Value: 200, Line: 1},
				{Type: token.NUMBER, Value: 300, Line: 1},
				{Type: token.NUMBER, Value: 400, Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.EOF, Line: 1},
			},
			want: &expr.List{
				Elements: []*expr.Literal{
					{Value: 100},
					{Value: 200},
					{Value: 300},
					{Value: 400},
				},
			},
		},
		// TODO: Parse print as a function name
		"print 99": {
			in: []*token.Token{
				{Type: token.LPAREN, Line: 1},
				{Type: token.NUMBER, Value: 69, Line: 1},
				{Type: token.IDENTIFIER, Value: "print", Line: 1},
				{Type: token.RPAREN, Line: 1},
				{Type: token.EOF, Line: 1},
			},
			want: &expr.List{
				Elements: []*expr.Literal{
					{Value: 69},
					{Value: "print"},
				},
			},
		},
	}

	for title, tt := range tests {
		t.Run(title, func(t *testing.T) {
			p := New(tt.in)
			got, err := p.Parse()
			if err != nil {
				t.Errorf("Parser.Parse() err = %v", err)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("want = |%v|", tt.want)
				t.Errorf("got = |%v|", got)
			}
		})
	}
}
