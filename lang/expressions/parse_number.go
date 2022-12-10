package expressions

func (tree *ParserT) parseNumber(first rune) []rune {
	// TODO: don't append each time, just return a range
	value := []rune{first}

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case (r >= '0' && '9' >= r) || r == '.':
			value = append(value, r)

		default:
			// not a number
			goto endNumber
		}
	}

endNumber:
	return value
}
