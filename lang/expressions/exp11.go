package expressions

import (
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func retBooleanFalse(tree *ParserT) error {
	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Boolean, false),
	})
}

func expLogicalAnd(tree *ParserT) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	nv, err := leftNode.dt.GetValue()
	if err != nil {
		return retBooleanFalse(tree)
	}

	v, err := types.ConvertGoType(nv, types.String)
	if err != nil {
		return err
	}

	if !types.IsTrueString(v.(string), nv.ExitNum) {
		return retBooleanFalse(tree)
	}

	nv, err = rightNode.dt.GetValue()
	if err != nil {
		return retBooleanFalse(tree)
	}

	v, err = types.ConvertGoType(nv, types.String)
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Boolean, types.IsTrueString(v.(string), nv.ExitNum)),
	})
}
