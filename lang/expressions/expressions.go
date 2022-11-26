package expressions

import (
	"github.com/lmorg/murex/lang/expressions/primitives"
)

type getVarCallback func(string) (interface{}, string, error)
type setVarCallback func(string, interface{}, string) error

func Execute(expression []rune, getVar getVarCallback, setVar setVarCallback) (*primitives.DataType, error) {
	tree := newExpTree(expression)
	tree.getVar = getVar
	tree.setVar = setVar
	err := tree.parse(true)
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
	err := tree.parse(false)
	if err != nil {
		return 0, err
	}

	return tree.charPos, nil
}
