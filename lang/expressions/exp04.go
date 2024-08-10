package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/utils/alter"
)

func expAdd(tree *ParserT) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	lv, rv, err := validateNumericalDataTypes(tree, leftNode, rightNode)
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Number,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Number, lv+rv),
	})
}

func expSubtract(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	lv, rv, err := validateNumericalDataTypes(tree, left, right)
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Number,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Number, lv-rv),
	})
}

const errCannotMergeWith = "cannot merge %s with %s: %s"

func expMerge(tree *ParserT) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	left, err := leftNode.dt.GetValue()
	if err != nil {
		return err
	}
	right, err := rightNode.dt.GetValue()
	if err != nil {
		return err
	}

	/*if left.Primitive != primitives.Array && left.Primitive != primitives.Object {
		return raiseError(tree.expression, leftNode, 0, fmt.Sprintf(
			errCannotMergeWith,
			right.Primitive.String(), left.Primitive.String(),
			"left side needs to be an array or object"))
	}

	if right.Primitive != primitives.Array && right.Primitive != primitives.Object {
		return raiseError(tree.expression, rightNode, 0, fmt.Sprintf(
			errCannotMergeWith,
			right.Primitive.String(), left.Primitive.String(),
			"right side needs to be an array or object"))
	}*/

	merged, err := alter.Merge(tree.p.Context, left.Value, nil, right.Value)
	if err != nil {
		return raiseError(tree.expression, leftNode, 0, fmt.Sprintf(
			errCannotMergeWith,
			right.Value, left.Value,
			err.Error()))
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewScalar(right.DataType, merged),
	})
}
