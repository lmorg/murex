package expressions

func (tree *ParserT) parseNumber() []rune {
	start := tree.charPos

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case (r >= '0' && '9' >= r) || r == '.':
			// valid numeric character

		default:
			// not a number
			goto endNumber
		}
	}

endNumber:
	return tree.expression[start:tree.charPos]
}
