package expressions

import (
	"errors"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func validateExpression(tree *expTreeT) error {
	// compile data types and check for errors in the AST
	//
	// first walk to ensure we have a:
	// value, expression, value, expression, value....etc

	if len(tree.ast) == 0 {
		return errors.New("missing expression")
	}
	if len(tree.ast) == 1 &&
		tree.ast[0].key != symbols.ArrayBegin && tree.ast[0].key != symbols.ObjectBegin {
		return errors.New("not an expression")
	}

	var expectValue bool

	for tree.astPos = 0; tree.astPos < len(tree.ast); tree.astPos++ {
		node := tree.ast[tree.astPos]
		expectValue = !expectValue

		// check for errors raised by the parser
		if node.key < symbols.DataValues {
			return raiseError(node, errMessage[node.key])
		}

		// check each operation has a left side and right side data value
		if expectValue {
			if node.key < symbols.DataValues || node.key > symbols.Operations {
				return raiseError(node, "expecting a data value")
			}

			if node.dt != nil {
				continue
			}
			primitive, err := node2primitive(node)
			if err != nil {
				return err
			}
			node.dt = primitive

		} else {
			if node.key < symbols.Operations {
				return raiseError(node, "expecting an operation")
			}

			switch node.key {
			case symbols.Add:
				if tree.prevSymbol().key == symbols.Bareword || tree.nextSymbol().key == symbols.Bareword {
					return raiseError(node, "cannot add barewords")
				}
			case symbols.Subtract, symbols.Divide, symbols.Multiply:
				if tree.prevSymbol().key != symbols.Number || tree.nextSymbol().key != symbols.Number {
					return raiseError(node, "cannot subtract non-numeric data types")
				}
			}
		}
	}

	if !expectValue {
		return raiseError(tree.ast[len(tree.ast)-1], "unexpected end of expression")
	}

	return nil
}
