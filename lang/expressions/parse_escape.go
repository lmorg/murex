package expressions

import "github.com/lmorg/murex/lang/expressions/node"

func (tree *ParserT) parseEscape() []rune {
	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		tree.syntaxTree.Add(node.H_ESCAPE, '\\', r)

		switch r {
		case 's', ' ':
			return []rune{' '}
		case 't', '\t':
			return []rune{'\t'}
		case 'r':
			return []rune{'\r'}
		case 'n', '\n':
			return []rune{'\n'}
		case '\r':
			return nil
		default:
			return []rune{r}
		}
	}

	tree.charPos--
	tree.syntaxTree.Append('\\')
	return []rune{'\\'}
}
