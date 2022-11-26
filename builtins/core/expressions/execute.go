package expressions

import (
	"fmt"

	"github.com/lmorg/murex/builtins/core/expressions/primitives"
	"github.com/lmorg/murex/builtins/core/expressions/symbols"
	"github.com/lmorg/murex/utils/consts"
)

func raiseError(node *astNodeT, message string) error {
	if node == nil {
		return fmt.Errorf("nil ast (%s)", consts.IssueTrackerURL)
	}

	return fmt.Errorf("%s at char %d (expression: %s, value '%s')",
		message, node.pos+1, node.key.String(), node.Value())
}

var errMessage = map[symbols.Exp]string{
	symbols.Unexpected:       "unexpected symbol",
	symbols.SubExpressionEnd: "more closing parenthesis then opening parenthesis",
	symbols.ObjectEnd:        "more closing curly braces then opening braces",
	symbols.ArrayEnd:         "more closing square brackets then opening brackets",
	symbols.InvalidHyphen:    "unexpected hyphen",
}

func (tree *expTreeT) execute() (*primitives.DataType, error) {
	err := validateExpression(tree)
	if err != nil {
		return nil, err
	}

	for i := range orderOfOperations {
		err = executeExpression(tree, orderOfOperations[i])
		if err != nil {
			return nil, err
		}
	}

	if len(tree.ast) > 1 {
		return nil, fmt.Errorf(
			"expression failed to execute correctly. %s",
			consts.IssueTrackerURL)
	}

	return tree.ast[0].dt, nil
}

func validateExpression(tree *expTreeT) error {
	// compile data types and check for errors in the AST
	//
	// first walk to ensure we have a:
	// value, expression, value, expression, value....etc

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
		}
	}

	if !expectValue {
		return raiseError(tree.ast[len(tree.ast)-1], "unexpected end of expression")
	}

	return nil
}

// To allow for extendability and developer expectations, the order of operations
// will follow what is defined by (for example) C, as outlined in the following:
// https://en.wikipedia.org/wiki/Order_of_operations#Programming_languages
// Not all operations will be available in murex and some are likely to be added
// in future versions of this package.
var orderOfOperations = []symbols.Exp{
	// 01. Function call, scope, array/member access
	// 02. (most) unary operators, sizeof and type casts (right to left)
	// 03. Multiplication, division, modulo
	symbols.Multiply,

	// 04. Addition and subtraction
	symbols.Add,

	// 05. Bitwise shift left and right
	// 06. Comparisons: less-than and greater-than
	symbols.GreaterThan,

	// 07. Comparisons: equal and not equal
	symbols.EqualTo,

	// 08. Bitwise AND
	// 09. Bitwise exclusive OR (XOR)
	// 10. Bitwise inclusive (normal) OR
	// 11. Logical AND
	// 12. Logical OR
	// 13. Conditional expression (ternary)
	// 14. Assignment operators (right to left)
	symbols.Assign,

	// 15. Comma operator
}

func executeExpression(tree *expTreeT, order symbols.Exp) (err error) {
	for /*i := 0;*/ tree.astPos = 0; tree.astPos < len(tree.ast); tree.astPos++ {
		node := tree.ast[tree.astPos]

		if node.key < order {
			continue
		}

		switch node.key {

		// 15. Comma operator
		// 14. Assignment operators (right to left)
		case symbols.Assign:
			err = expAssign(tree)
		case symbols.AssignAndAdd:
		case symbols.AssignAndSubtract:
		case symbols.AssignAndDivide:
		case symbols.AssignAndMultiply:

		// 13. Conditional expression (ternary)
		// 12. Logical OR
		// 11. Logical AND
		// 10. Bitwise inclusive (normal) OR
		// 09. Bitwise exclusive OR (XOR)
		// 08. Bitwise AND
		// 07. Comparisons: equal and not equal
		case symbols.EqualTo:
			err = expEqualTo(tree)
		case symbols.NotEqualTo:
			err = expNotEqualTo(tree)
		case symbols.Like:
			err = expLike(tree, true)
		case symbols.NotLike:
			err = expLike(tree, false)
		case symbols.Regexp:
			err = expRegexp(tree, true)
		case symbols.NotRegexp:
			err = expRegexp(tree, false)
		// 06. Comparisons: less-than and greater-than
		case symbols.GreaterThan:
			err = expGreaterThan(tree)
		case symbols.GreaterThanOrEqual:
			err = expGreaterThanOrEqual(tree)
		case symbols.LessThan:
			err = expLessThan(tree)
		case symbols.LessThanOrEqual:
			err = expLessThanOrEqual(tree)

		// 05. Bitwise shift left and right
		// 04. Addition and subtraction
		case symbols.Add:
			err = expAdd(tree)
		case symbols.Subtract:
			err = expSubtract(tree)

		// 03. Multiplication, division, modulo
		case symbols.Multiply:
			err = expMultiply(tree)
		case symbols.Divide:
			err = expDivide(tree)

		// 02. (most) unary operators, sizeof and type casts (right to left)
		// 01. Function call, scope, array/member access

		default:
			err = raiseError(node, fmt.Sprintf(
				"no code written to handle symbol (%s)",
				consts.IssueTrackerURL,
			))
		}

		if err != nil {
			return err
		}
	}

	return nil
}
