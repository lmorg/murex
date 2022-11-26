package expressions

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
)

func Execute(p *lang.Process, expression []rune) (*primitives.DataType, error) {
	tree := newExpTree(expression)
	tree.p = p
	err := tree.parse()
	if err != nil {
		return nil, err
	}

	return tree.execute()
}

// ChainParser is intended to be called from other parsers as a way of
// embedding this expressions library into other language syntaxes.
// This function just parses the expression and returns the end of the
// expression.
func ChainParser(expression []rune, offset int) (int, error) {
	tree := newExpTree(expression)
	tree.charOffset = offset
	err := tree.parse()
	if err != nil {
		return 0, err
	}

	return tree.charPos, nil
}
