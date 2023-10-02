package expressions

import "github.com/lmorg/murex/lang/expressions/node"

func (tree *ParserT) parseComment() {
	start := tree.charPos

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
	tree.syntaxTree.Add(node.H_COMMENT, tree.expression[start:tree.charPos]...)
	tree.charPos--
}

func (tree *ParserT) parseCommentMultiLine() error {
	start := tree.charPos

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

	tree.syntaxTree.Add(node.H_COMMENT, tree.expression[start:]...)
	return raiseError(tree.expression, nil, tree.charPos, "comment opened with '/#' but no closing token '#/' could be found")

endCommentMultiLine:
	tree.charPos++
	tree.syntaxTree.Add(node.H_COMMENT, tree.expression[start:tree.charPos+1]...)
	return nil
}
