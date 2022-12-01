package expressions

import (
	"errors"
	"fmt"

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
		return fmt.Errorf("not an expression: '%s'", string(tree.expression))
	}

	var expectValue bool

	for tree.astPos = 0; tree.astPos < len(tree.ast); tree.astPos++ {
		node := tree.ast[tree.astPos]
		prev := tree.prevSymbol()
		next := tree.nextSymbol()
		expectValue = !expectValue

		// check for errors raised by the parser
		if node.key < symbols.DataValues {
			return raiseError(tree.expression, node, errMessage[node.key])
		}

		// check each operation has a left side and right side data value
		if expectValue {
			if node.key < symbols.DataValues || node.key > symbols.Operations {
				return raiseError(tree.expression, node, "expecting a data value")
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
				return raiseError(tree.expression, node, "expecting an operation")
			}

			switch node.key {
			case symbols.Add, symbols.GreaterThan, symbols.LessThan:
				if prev == nil || prev.key == symbols.Bareword ||
					next == nil || next.key == symbols.Bareword {
					return raiseError(tree.expression, node, fmt.Sprintf("cannot %s barewords", node.key))
				}
			case symbols.Subtract, symbols.Divide, symbols.Multiply:
				if prev == nil || prev.key != symbols.Number ||
					next == nil || next.key != symbols.Number {
					return raiseError(tree.expression, node, fmt.Sprintf("cannot %s non-numeric data types", node.key))
				}
			}
		}
	}

	if !expectValue {
		return raiseError(tree.expression, tree.ast[len(tree.ast)-1], "unexpected end of expression")
	}

	return nil
}
