package expressions

import (
	"fmt"

	"github.com/lmorg/murex/builtins/core/expressions/primitives"
	"github.com/lmorg/murex/builtins/core/expressions/symbols"
)

func expAssign(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.key != symbols.Bareword {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"left side should be a bareword, instead got %s", left.key))
	}

	err = tree.p.Variables.Set(tree.p, left.Value(), right.dt.Value, right.dt.DataType())
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.DataValues,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Null,
			Value:     nil,
		},
	})
}
