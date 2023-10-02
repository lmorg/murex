package expressions

func (tree *ParserT) parseEscape() ([]rune, []rune) {
	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case 's', ' ':
			return []rune{'\\', r}, []rune{' '}
		case 't', '\t':
			return []rune{'\\', r}, []rune{'\t'}
		case 'r':
			return []rune{'\\', r}, []rune{'\r'}
		case 'n', '\n':
			return []rune{'\\', r}, []rune{'\n'}
		case '\r':
			return []rune{'\\', r}, nil
		default:
			return []rune{'\\', r}, []rune{r}
		}
	}

	return []rune{'\\'}, []rune{'\\'}
}
