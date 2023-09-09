package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/types"
)

func validateNumericalDataTypes(tree *ParserT, leftNode *astNodeT, rightNode *astNodeT, operation *astNodeT) (float64, float64, error) {
	left, err := leftNode.dt.GetValue()
	if err != nil {
		return 0, 0, err
	}
	right, err := rightNode.dt.GetValue()
	if err != nil {
		return 0, 0, err
	}

	if tree.StrictTypes() {
		switch {
		case left.Primitive != primitives.Number:
			return 0, 0, raiseError(tree.expression, leftNode, 0, fmt.Sprintf(
				"cannot %s with %s types", tree.currentSymbol().key, left.Primitive,
			))
		case right.Primitive != primitives.Number:
			return 0, 0, raiseError(tree.expression, rightNode, 0, fmt.Sprintf(
				"cannot %s with %s types", tree.currentSymbol().key, right.Primitive,
			))
		default:
			var lv, rv float64
			switch t := left.Value.(type) {
			case float64:
				lv = t
			case int:
				lv = float64(t)
			default:
				return 0, 0, raiseError(tree.expression, leftNode, 0, fmt.Sprintf(
					"value cannot be converted into an integer nor floating point number\nUnderlying data type: %T", t,
				))
			}

			switch t := right.Value.(type) {
			case float64:
				rv = t
			case int:
				rv = float64(t)
			default:
				return 0, 0, raiseError(tree.expression, rightNode, 0, fmt.Sprintf(
					"value cannot be converted into an integer nor floating point number\nUnderlying data type: %T", t,
				))
			}
			return lv, rv, nil
		}
	}

	lv, err := types.ConvertGoType(left.Value, types.Number)
	if err != nil {
		return 0, 0, raiseError(tree.expression, leftNode, 0, err.Error())
	}

	rv, err := types.ConvertGoType(right.Value, types.Number)
	if err != nil {
		return 0, 0, raiseError(tree.expression, rightNode, 0, err.Error())
	}

	return lv.(float64), rv.(float64), nil
}

func compareTypes(tree *ParserT, leftNode *astNodeT, rightNode *astNodeT) (interface{}, interface{}, error) {
	left, err := leftNode.dt.GetValue()
	if err != nil {
		return nil, nil, err
	}
	right, err := rightNode.dt.GetValue()
	if err != nil {
		return nil, nil, err
	}

	if tree.StrictTypes() {
		if left.Primitive != right.Primitive {
			return nil, nil, raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot compare %s with %s", left.Primitive, right.Primitive,
			))
		}
		return left.Value, right.Value, nil
	}

	if left.Primitive == right.Primitive {
		return left.Value, right.Value, nil
	}

	if left.Primitive == primitives.Number || right.Primitive == primitives.Number {
		lv, lErr := types.ConvertGoType(left.Value, types.Number)
		rv, rErr := types.ConvertGoType(right.Value, types.Number)
		if lErr == nil && rErr == nil {
			return lv, rv, nil
		}
	}

	lv, err := types.ConvertGoType(left.Value, types.String)
	if err != nil {
		return nil, nil, err
	}

	rv, err := types.ConvertGoType(right.Value, types.String)
	if err != nil {
		return nil, nil, err
	}

	return lv, rv, nil
}
