package expressions

import (
	"github.com/lmorg/murex/lang"
)

func init() {
	lang.ChainParser = ChainParser
}

// ChainParser is intended to be called from other parsers as a way of
// embedding this expressions library into other language syntaxes.
// This function just parses the expression and returns the end of the
// expression.
func ChainParser(expression []rune, offset int) (int, error) {
	tree := newExpTree(nil, expression)
	tree.charOffset = offset

	err := tree.parse(false)
	if err != nil {
		return 0, err
	}

	err = validateExpression(tree)
	if err != nil {
		return 0, err
	}

	return tree.charPos, nil
}
