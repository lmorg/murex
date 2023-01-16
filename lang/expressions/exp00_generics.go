package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/types"
)

func validateNumericalDataTypes(tree *ParserT, left *astNodeT, right *astNodeT, operation *astNodeT) (float64, float64, error) {
	if tree.StrictTypes() {
		switch {
		case left.dt.Primitive != primitives.Number:
			return 0, 0, raiseError(tree.expression, left, 0, fmt.Sprintf(
				"cannot %s with %s types", tree.currentSymbol().key, left.dt.Primitive,
			))
		case right.dt.Primitive != primitives.Number:
			return 0, 0, raiseError(tree.expression, right, 0, fmt.Sprintf(
				"cannot %s with %s types", tree.currentSymbol().key, right.dt.Primitive,
			))
		default:
			var lv, rv float64
			switch t := left.dt.Value.(type) {
			case float64:
				lv = t
			case int:
				lv = float64(t)
			default:
				return 0, 0, raiseError(tree.expression, left, 0, fmt.Sprintf(
					"value cannot be converted into an integer nor floating point number\nUnderlying data type: %T", t,
				))
			}
			switch t := right.dt.Value.(type) {
			case float64:
				rv = t
			case int:
				rv = float64(t)
			default:
				return 0, 0, raiseError(tree.expression, right, 0, fmt.Sprintf(
					"value cannot be converted into an integer nor floating point number\nUnderlying data type: %T", t,
				))
			}
			return lv, rv, nil
		}
	}

	lv, err := types.ConvertGoType(left.dt.Value, types.Number)
	if err != nil {
		return 0, 0, raiseError(tree.expression, left, 0, err.Error())
	}

	rv, err := types.ConvertGoType(right.dt.Value, types.Number)
	if err != nil {
		return 0, 0, raiseError(tree.expression, right, 0, err.Error())
	}

	return lv.(float64), rv.(float64), nil
}

func compareTypes(tree *ParserT, left *astNodeT, right *astNodeT) (interface{}, interface{}, error) {
	if tree.StrictTypes() {
		if left.dt.Primitive != right.dt.Primitive {
			return nil, nil, raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot compare %s with %s", left.dt.Primitive, right.dt.Primitive,
			))
		}
		return left.dt.Value, right.dt.Value, nil
	}

	if left.dt.Primitive == right.dt.Primitive {
		return left.dt.Value, right.dt.Value, nil
	}

	lv, err := types.ConvertGoType(left.dt.Value, types.String)
	if err != nil {
		return nil, nil, err
	}

	rv, err := types.ConvertGoType(right.dt.Value, types.String)
	if err != nil {
		return nil, nil, err
	}

	return lv, rv, nil
}
