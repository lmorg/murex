package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/json"
)

func Execute(p *lang.Process, expression []rune) (*primitives.DataType, error) {
	tree := newExpTree(p, expression)

	err := tree.parse(true)
	if err != nil {
		return nil, err
	}

	return tree.execute()
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
			"expression failed to execute correctly (AST results > 1).\n%s\n%s",
			json.LazyLoggingPretty(tree.Dump()),
			consts.IssueTrackerURL)
	}

	return tree.ast[0].dt, nil
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
	/*defer func() {
		if err := recover(); err != nil {
			err = fmt.Errorf("panic caught: %v\nExpression: %s\nnode: %d\nAST: %s",
				err,
				string(tree.expression),
				tree.astPos,
				json.LazyLoggingPretty(tree.Dump()))

		}
	}()*/

	for tree.astPos = 0; tree.astPos < len(tree.ast); tree.astPos++ {
		//fmt.Println(tree.astPos, json.LazyLogging(tree.Dump()))
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
			err = expAssignAdd(tree)
		case symbols.AssignAndSubtract:
			err = expAssignSubtract(tree)
		case symbols.AssignAndDivide:
			err = expAssignDivide(tree)
		case symbols.AssignAndMultiply:
			err = expAssignMultiply(tree)

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
			err = raiseError(tree.expression, node, 0, fmt.Sprintf(
				"no code written to handle symbol (%s)",
				consts.IssueTrackerURL))
		}

		if err != nil {
			return err
		}

		tree.astPos = 0
	}

	return nil
}
