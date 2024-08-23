package asm

import "testing"
import "reflect"

const wantErr = true
const noErr = false

type wantToken struct {
	value string
	row   int
	col   int
}

func ToToken(t *testing.T, want wantToken) AsmToken {
	t.Helper()
	return AsmToken{
		value: want.value,
		loc:   Location{want.row, want.col},
	}
}

func ToTokens(t *testing.T, wants []wantToken) (atokens []AsmToken) {
	for _, wt := range wants {
		atokens = append(atokens, ToToken(t, wt))
	}
	return
}

func Test_TokenizeSource(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		want    []wantToken
	}{
		{"single-line", "MOV: ACC, 420", noErr, []wantToken{{"MOV", 0, 0}, {"ACC", 0, 5}, {"420", 0, 10}}},
		{"multiple-line", "MOV: R1, 69\nMOV: R0, R1", noErr, []wantToken{{"MOV", 0, 0}, {"R1", 0, 5}, {"69", 0, 9}, {"MOV", 1, 0}, {"R0", 1, 5}, {"R1", 1, 9}}},
		{"allow-empty-line", "MOV: R1, 69\n\nMOV: R0, R1", noErr, []wantToken{{"MOV", 0, 0}, {"R1", 0, 5}, {"69", 0, 9}, {"MOV", 2, 0}, {"R0", 2, 5}, {"R1", 2, 9}}},
		{"allow-empty-source", "", noErr, []wantToken{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := TokenizeSource(tt.input)
			if err != nil && !tt.wantErr {
				t.Fatalf("This test wants error but no error : actual: `%v`, want `%v`", actual, tt.want)
			}
			tokens := ToTokens(t, tt.want)
			if !reflect.DeepEqual(actual, tokens) {
				t.Errorf("source:\n%s", tt.input)
				t.Fatalf("actual = `%v`, want `%v`", actual, tokens)
			}
		})
	}
}
