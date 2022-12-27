package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func expMultiply(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if tree.StrictTypes() {
		if left.dt.Primitive != right.dt.Primitive {
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
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
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s with %s types", tree.currentSymbol().key, left.dt.Primitive,
			))
		}

	} else {
		lv, err := types.ConvertGoType(left.dt.Value, types.Number)
		if err != nil {
			return raiseError(tree.expression, left, 0, err.Error())
		}

		rv, err := types.ConvertGoType(left.dt.Value, types.Number)
		if err != nil {
			return raiseError(tree.expression, right, 0, err.Error())
		}

		return tree.foldAst(&astNodeT{
			key: symbols.Number,
			pos: tree.ast[tree.astPos].pos,
			dt: &primitives.DataType{
				Primitive: primitives.Number,
				Value:     lv.(float64) * rv.(float64),
			},
		})
	}
}

func expDivide(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.dt.Primitive != right.dt.Primitive {
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
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
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s with %s types", tree.currentSymbol().key, left.dt.Primitive,
		))
	}

}
