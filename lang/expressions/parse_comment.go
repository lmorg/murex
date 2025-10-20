package expressions

func (tree *ParserT) parseComment() {
	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]
		switch r {

		case '\n':
			goto endComment

		case '\\':
			next := tree.nextChar()
			if tree.statement != nil && (next == '\r' || next == '\n') {
				tree.statement.escapeLf = true
			}

		}
	}

endComment:
	tree.charPos--
}

func (tree *ParserT) parseCommentMultiLine() error {
	for tree.charPos += 2; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '\n':
			tree.crLf()

		case '#':
			if tree.nextChar() == '/' {
				goto endCommentMultiLine
			}

		}
	}

	return raiseError(tree.expression, nil, tree.charPos, "comment opened with '/#' but no closing token '#/' could be found")

endCommentMultiLine:
	tree.charPos++

	return nil
}
