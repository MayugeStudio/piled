package scanner

func LexProgram(programPath string, source string) ([]*Token, error) {
	ops, err := lexSourceIntoTokens(programPath, source)
	if err != nil {
		return nil, err
	}
	return ops, nil
}

func lexSourceIntoTokens(filepath string, source string) ([]*Token, error) {
	ops := make([]*Token, 0)
	lines := strings.Split(source, "\n")

	for row, line := range lines {
		val := ""
		start_col := 0
		line_length := len(line)
		for col := 0; col < line_length; col++ {
			char := line[col]
			isEndOfLine := col == line_length-1
			isSpace := char == ' '
			isComma := char == ','

			if !isSpace && !isComma {
				val += string(char)
			}

			if isSpace || isEndOfLine {
				loc := Location{Row: row + 1, Col: start_col + 1}
				op, err := lexWord(filepath, val, loc)
				if err != nil {
					return nil, err
				}
				start_col = col + 1
				ops = append(ops, op)
				val = ""
			}
		}
	}
	return ops, nil
}

func lexWord(filepath string, value string, loc Location) (*Token, error) {
	opType, err := nameToTokenType(value)
	switch value {
		case
	}
	if opType == Token_PUSH_INT {
		v, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}

		return &Token{
			Type:  opType,
			Loc:   loc,
			Value: v,
		}, nil
	}
	return &Token{
		Type: opType,
		Loc:  loc,
	}, nil
}
