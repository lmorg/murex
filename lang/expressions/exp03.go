package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
)

func expMultiply(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.dt.Primitive != right.dt.Primitive {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s %s with %s",
			tree.currentSymbol().key, left.dt.Primitive, right.dt.Primitive,
		))
	}

	switch left.dt.Primitive {
	case primitives.Number:
		return tree.foldAst(&astNodeT{
			key: symbols.Number,
			pos: tree.ast[tree.astPos].pos,
			dt: &primitives.DataType{
				Primitive: primitives.Number,
				Value:     left.dt.Value.(float64) * right.dt.Value.(float64),
			},
		})

	default:
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s with %s types", tree.currentSymbol().key, left.dt.Primitive,
		))
	}

}

func expDivide(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.dt.Primitive != right.dt.Primitive {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s %s with %s",
			tree.currentSymbol().key, left.dt.Primitive, right.dt.Primitive,
		))
	}

	switch left.dt.Primitive {
	case primitives.Number:
		return tree.foldAst(&astNodeT{
			key: symbols.Number,
			pos: tree.ast[tree.astPos].pos,
			dt: &primitives.DataType{
				Primitive: primitives.Number,
				Value:     left.dt.Value.(float64) / right.dt.Value.(float64),
			},
		})

	default:
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s with %s types", tree.currentSymbol().key, left.dt.Primitive,
		))
	}

}
