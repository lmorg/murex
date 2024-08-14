package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func (tree *ParserT) validateExpression() error {
	// compile data types and check for errors in the AST
	//
	// first walk to ensure we have a:
	// value, expression, value, expression, value....etc

	if tree.charPos < 0 {
		tree.charPos = 0
	}

	if len(tree.ast) == 0 {
		return fmt.Errorf("missing expression: '%s'", string(tree.expression[:tree.charPos]))
	}

	if len(tree.ast) == 1 &&
		(tree.ast[0].key == symbols.Bareword || tree.ast[0].key == symbols.SubExpressionBegin ||
			tree.ast[0].key == symbols.QuoteSingle || tree.ast[0].key == symbols.QuoteDouble) {
		if tree.charPos+1 > len(tree.expression) {
			tree.charPos--
		}
		return fmt.Errorf("not an expression: '%s'", string(tree.expression[:tree.charPos+1]))
	}

	var expectValue bool

	for tree.astPos = 0; tree.astPos < len(tree.ast); tree.astPos++ {
		prev := tree.prevSymbol()
		node := tree.ast[tree.astPos]
		next := tree.nextSymbol()

		expectValue = !expectValue

		// check for errors raised by the parser
		if node.key < symbols.DataValues {
			return raiseError(tree.expression, node, 0, errMessage[node.key])
		}

		// check each operation has a left side and right side data value
		if expectValue {
			if (node.key < symbols.DataValues) ||
				node.key > symbols.Operations {
				return raiseError(tree.expression, node, 0, "expecting a data value")
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

			switch {
			case prev == nil:
				return raiseError(tree.expression, node, 0, fmt.Sprintf("nil symbol preceding %s", node.key))

			case node.key == symbols.PlusPlus, node.key == symbols.MinusMinus:
				if prev.key != symbols.Scalar {
					return raiseError(tree.expression, node, 0, fmt.Sprintf("%s can only follow a %s, instead got %s", node.key, symbols.Scalar, prev.key))
				}
				expectValue = !expectValue

			case next == nil:
				return raiseError(tree.expression, node, 0, fmt.Sprintf("nil symbol following %s", node.key))

			case node.key < symbols.Operations:
				return raiseError(tree.expression, node, 0, fmt.Sprintf("expecting an operation, instead got %s", node.key))

			case node.key >= symbols.Add:
				if !isSymbolNumeric(prev.key) {
					return raiseError(tree.expression, node, 0, fmt.Sprintf("cannot %s non-numeric data types, left is %s", node.key, prev.key))
				}
				if !isSymbolNumeric(next.key) {
					return raiseError(tree.expression, node, 0, fmt.Sprintf("cannot %s non-numeric data types, right is %s", node.key, next.key))
				}

			case node.key >= symbols.Elvis:
				if prev.key == symbols.Bareword {
					return raiseError(tree.expression, node, 0, fmt.Sprintf("cannot %s left %s", node.key, prev.key))
				}
				if next.key == symbols.Bareword {
					return raiseError(tree.expression, node, 0, fmt.Sprintf("cannot %s right %s", node.key, next.key))
				}

			default:
				if !isSymbolAssignable(prev.key) {
					return raiseError(tree.expression, node, 0, fmt.Sprintf("cannot %s to %s", node.key, prev.key))
				}
			}
		}

	}

	if !expectValue {
		return raiseError(tree.expression, tree.ast[len(tree.ast)-1], 0, "unexpected end of expression")
	}

	return nil
}

func isSymbolNumeric(sym symbols.Exp) bool {
	return sym == symbols.Number ||
		sym == symbols.Calculated ||
		sym == symbols.SubExpressionBegin ||
		sym == symbols.Scalar
}

func isSymbolAssignable(sym symbols.Exp) bool {
	return sym == symbols.Bareword ||
		sym == symbols.Scalar ||
		sym == symbols.SubExpressionBegin
}
