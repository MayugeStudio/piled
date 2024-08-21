package main

import "fmt"
import "strings"

type Register int8

// Location
type Location struct {
	row int
	col int
}

func (l Location) String() string {
	return fmt.Sprintf("%d:%d", l.row, l.col)
}

// AsmToken
type AsmToken struct {
	value string
	loc Location
}

func (t AsmToken) String() string {
	return fmt.Sprintf("%s: `%s`", t.loc, t.value)
}

type OpCode struct{}

func tokenizeSource(source string) (tokens []AsmToken, err error) {
	if len(source) == 0 {
		return
	}

	lines := strings.Split(source, "\n")

	for row, line := range lines {
		if len(line) == 0 {
			continue
		}
		val := ""
		line_length := len(line)
		for col := 0; col < line_length; col++ {
			char := line[col]
			isEndOfLine := col == line_length - 1

			if char != ' ' {
				val += string(char)
			}

			if char == ' ' || isEndOfLine {
				token := AsmToken{
					value: val,
					loc: Location{row: row, col: col},
				}
				tokens = append(tokens, token)
				val = ""
			}
		}
	}
	return
}

func main() {
	source := "MOV: ACC $10"
	fmt.Printf("raw source: %s\n", source)
	tokens, err := tokenizeSource(source)
	if err != nil {
		fmt.Println("Error!")
	}
	for _, token := range tokens {
		fmt.Println(token)
	}
}

