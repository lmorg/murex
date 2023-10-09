package expressions

import (
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func expNullCoalescing(tree *ParserT) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	left, err := leftNode.dt.GetValue()

	switch {
	case err != nil:
		return expElvisRightValue(tree, rightNode)

	case left.DataType == types.Null:
		return expElvisRightValue(tree, rightNode)

	default:
		// valid left operand
		return tree.foldAst(&astNodeT{
			key: symbols.Calculated,
			pos: tree.ast[tree.astPos].pos,
			dt:  primitives.NewScalar(left.DataType, left.Value),
		})
	}
}

func expElvis(tree *ParserT) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	left, err := leftNode.dt.GetValue()

	if err != nil {
		return expElvisRightValue(tree, rightNode)
	}

	v, err := types.ConvertGoType(left.Value, types.Boolean)
	if err != nil {
		return expElvisRightValue(tree, rightNode)
	}

	if !v.(bool) {
		return expElvisRightValue(tree, rightNode)
	}

	// valid left operand
	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewScalar(left.DataType, left.Value),
	})
}

func expElvisRightValue(tree *ParserT, rightNode *astNodeT) error {
	right, err := rightNode.dt.GetValue()
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewScalar(right.DataType, right.Value),
	})
}
