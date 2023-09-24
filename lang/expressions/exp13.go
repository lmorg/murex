package expressions

import (
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func expElvis(tree *ParserT) error {
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

	/*case left.DataType == types.String:
	s, ok := left.Value.(string)
	if ok && s == "" {
		return expElvisRightValue(tree, rightNode)
	}
	fallthrough*/

	default:
		// valid left operand
		return tree.foldAst(&astNodeT{
			key: symbols.Calculated,
			pos: tree.ast[tree.astPos].pos,
			dt:  primitives.NewPrimitive(left.Primitive, left.Value),
		})
	}
}

func expElvisRightValue(tree *ParserT, rightNode *astNodeT) error {
	right, err := rightNode.dt.GetValue()
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(right.Primitive, right.Value),
	})
}
