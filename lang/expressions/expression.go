package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/utils/consts"
)

func (tree *ParserT) executeExpr() (*primitives.DataType, error) {
	err := tree.validateExpression()
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
			"expression failed to execute correctly (AST results > 1).\n%s",
			consts.IssueTrackerURL)
	}

	return tree.ast[0].dt, nil
}

// To allow for extendability and developer expectations, the order of operations
// will follow what is defined by (for example) C, as outlined in the following:
// https://en.wikipedia.org/wiki/Order_of_operations#Programming_languages
// Not all operations will be available in murex and some are likely to be added
// in future versions of this package.
//
// Please also note that the slice below is just defining the groupings. Each
// operator within the _same_ group will then be processed from left to right.
// Read the `executeExpression` function further down this source file to view
// every supported operator
var orderOfOperations = []symbols.Exp{
	// 01. Function call, scope, array/member access
	// 02. (most) unary operators, sizeof and type casts (right to left)
	// 03. Multiplication, division, modulo
	symbols.Multiply,

	// 04. Addition and subtraction
	symbols.Add,

	// 04.1 Merge
	symbols.Merge,

	// 05. Bitwise shift left and right
	// 06. Comparisons: less-than and greater-than
	symbols.GreaterThan,

	// 07. Comparisons: equal and not equal
	symbols.EqualTo,

	// 08. Bitwise AND
	// 09. Bitwise exclusive OR (XOR)
	// 10. Bitwise inclusive (normal) OR
	// 11. Logical AND
	symbols.LogicalAnd,

	// 12. Logical OR
	symbols.LogicalOr,

	// 13. Conditional expression (ternary)
	symbols.Elvis,

	// 14. Assignment operators (right to left)
	symbols.Assign,

	// 15. Comma operator
}

func executeExpression(tree *ParserT, order symbols.Exp) (err error) {
	for tree.astPos = 0; tree.astPos < len(tree.ast); tree.astPos++ {
		node := tree.ast[tree.astPos]

		if node.key < order {
			continue
		}

		switch node.key {

		// 15. Comma operator
		// 14. Assignment operators (right to left)
		case symbols.Assign:
			err = expAssign(tree, true)
		case symbols.AssignUpdate:
			err = expAssign(tree, false)
		case symbols.AssignAndAdd:
			//err = expAssignAdd(tree)
			err = expAssignAndOperate(tree, _assAdd)
		case symbols.AssignAndSubtract:
			err = expAssignAndOperate(tree, _assSub)
		case symbols.AssignAndDivide:
			err = expAssignAndOperate(tree, _assDiv)
		case symbols.AssignAndMultiply:
			err = expAssignAndOperate(tree, _assMulti)
		case symbols.AssignOrMerge:
			err = expAssignMerge(tree)

		// 13. Conditional expression (ternary)
		case symbols.NullCoalescing:
			err = expNullCoalescing(tree)
		case symbols.Elvis:
			err = expElvis(tree)

		// 12. Logical OR
		case symbols.LogicalOr:
			err = expLogicalOr(tree)

		// 11. Logical AND
		case symbols.LogicalAnd:
			err = expLogicalAnd(tree)

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
			err = expGtLt(tree, _gtF, _gtS)
		case symbols.GreaterThanOrEqual:
			err = expGtLt(tree, _gtEqF, _gtEqS)
		case symbols.LessThan:
			err = expGtLt(tree, _ltF, _ltS)
		case symbols.LessThanOrEqual:
			err = expGtLt(tree, _ltEqF, _ltEqS)

		// 05. Bitwise shift left and right

		// 04.1 Merge
		case symbols.Merge:
			err = expMerge(tree)

			// 04. Addition and subtraction
		case symbols.PlusPlus:
			err = expPlusPlus(tree, 1)
		case symbols.MinusMinus:
			err = expPlusPlus(tree, -1)
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
