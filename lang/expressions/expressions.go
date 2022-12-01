package expressions

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
)

func Execute(p *lang.Process, expression []rune) (*primitives.DataType, error) {
	tree := newExpTree(p, expression)

	err := tree.parse(true)
	if err != nil {
		return nil, err
	}

	return tree.execute()
}

func init() {
	lang.ChainParser = ChainParser
}

// ChainParser is intended to be called from other parsers as a way of
// embedding this expressions library into other language syntaxes.
// This function just parses the expression and returns the end of the
// expression.
func ChainParser(expression []rune, offset int) (int, error) {
	var err error
	tree := newExpTree(nil, expression)

	/*defer func() {
		fmt.Printf("%s: %v %s",
			string(expression),
			err,
			json.LazyLoggingPretty(tree.Dump()))
	}()*/

	tree.charOffset = offset
	err = tree.parse(false)
	if err != nil {
		return 0, err
	}

	err = validateExpression(tree)
	if err != nil {
		return 0, err
	}

	return tree.charPos, nil
}
