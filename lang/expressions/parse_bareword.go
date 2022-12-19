package expressions

func isBareChar(r rune) bool {
	return r == '_' ||
		(r >= 'a' && 'z' >= r) ||
		(r >= 'A' && 'Z' >= r) ||
		(r >= '0' && '9' >= r)
}

func (tree *ParserT) parseBareword() []rune {
	i := tree.charPos + 1

	for ; i < len(tree.expression); i++ {
		switch {
		case isBareChar(tree.expression[i]):
			// valid bareword character

		default:
			// not a valid bareword character
			goto endBareword
		}
	}

endBareword:
	value := tree.expression[tree.charPos:i]
	tree.charPos = i
	return value
}
