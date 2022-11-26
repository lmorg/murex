package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
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

	if right.key <= symbols.Bareword {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"right side should not be a %s", right.key))
	}

	if right.key == symbols.Bareword {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"right side should not be a %s", right.key))
	}

	err = tree.setVar(left.Value(), right.dt.Value, right.dt.DataType())
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
