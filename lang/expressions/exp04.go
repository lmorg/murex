package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/utils/alter"
)

func expAdd(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	lv, rv, err := validateNumericalDataTypes(tree, left, right, tree.currentSymbol())
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Number,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Number,
			Value:     lv + rv,
		},
	})
}

func expSubtract(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	lv, rv, err := validateNumericalDataTypes(tree, left, right, tree.currentSymbol())
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Number,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Number,
			Value:     lv - rv,
		},
	})
}

func expMergeInto(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	merged, err := alter.Merge(tree.p.Context, right.dt.Value, nil, left.dt.Value)
	if err != nil {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"cannot perform merge '%s' into '%s': %s",
			right.Value(), left.Value(),
			err.Error()))
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Other,
			MxDT:      right.dt.MxDT,
			Value:     merged,
		},
	})
}
