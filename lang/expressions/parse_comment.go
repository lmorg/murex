package expressions

func (tree *ParserT) parseComment() {
	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]
		switch r {

		case '\n':
			goto endComment

		case '\\':
			next := tree.nextChar()
			if next == '\r' || next == '\n' {
				tree.statement.ignoreCrLf = true
			}

		}
	}

endComment:
	tree.charPos--
}
