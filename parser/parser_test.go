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
