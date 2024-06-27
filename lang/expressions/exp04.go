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

func expMergeInto(tree *ParserT) error {
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

	merged, err := alter.Merge(tree.p.Context, right.Value, nil, left.Value)
	if err != nil {
		return raiseError(tree.expression, leftNode, 0, fmt.Sprintf(
			"cannot perform merge '%s' into '%s': %s",
			right.Value, left.Value,
			err.Error()))
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewScalar(right.DataType, merged),
	})
}
