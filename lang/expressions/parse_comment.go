package expressions

func (tree *ParserT) parseComment() {
	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {

		if tree.expression[tree.charPos] == '\n' {
			break
		}
	}

	tree.charPos--
}
