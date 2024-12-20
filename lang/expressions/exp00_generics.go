package expressions

import (
	"encoding/json"
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/types"
)

const errCannotConvertToFloat = "value cannot be converted into an integer nor floating point number"

func validateNumericalDataTypes(tree *ParserT, leftNode *astNodeT, rightNode *astNodeT) (float64, float64, error) {
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
				return 0, 0, raiseError(tree.expression, leftNode, 0, fmt.Sprintf("%s\nUnderlying data type: %T", errCannotConvertToFloat, t))
			}

			switch t := right.Value.(type) {
			case float64:
				rv = t
			case int:
				rv = float64(t)
			default:
				return 0, 0, raiseError(tree.expression, rightNode, 0, fmt.Sprintf("%s\nUnderlying data type: %T", errCannotConvertToFloat, t))
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

const errCannotCompareBareword = "cannot compare with bareword '%s': strings should be quoted, variables prefixed with `$`"

func compareTypes(tree *ParserT, leftNode *astNodeT, rightNode *astNodeT) (interface{}, interface{}, error) {
	left, err := leftNode.dt.GetValue()
	if err != nil {
		return nil, nil, err
	}
	right, err := rightNode.dt.GetValue()
	if err != nil {
		return nil, nil, err
	}

	if left.Primitive == primitives.Bareword {
		return nil, nil, raiseError(tree.expression, leftNode, 0, fmt.Sprintf(errCannotCompareBareword, leftNode.Value()))
	}
	if right.Primitive == primitives.Bareword {
		return nil, nil, raiseError(tree.expression, rightNode, 0, fmt.Sprintf(errCannotCompareBareword, rightNode.Value()))
	}

	if tree.StrictTypes() {
		if left.Primitive != right.Primitive ||
			!left.Primitive.IsComparable() || !right.Primitive.IsComparable() {
			return nil, nil, raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot compare %s with %s", left.Primitive, right.Primitive,
			))
		}
		return left.Value, right.Value, nil
	}

	var (
		lv, rv        any
		notCompatible bool
	)

	if !left.Primitive.IsComparable() {
		var ok bool
		lv, ok = left.Value.(string)
		if !ok {
			b, err := json.Marshal(left.Value)
			if err != nil {
				return nil, nil, err
			}
			lv = string(b)
		}
		notCompatible = true
	} else {
		lv = left.Value
	}

	if !right.Primitive.IsComparable() {
		var ok bool
		rv, ok = right.Value.(string)
		if !ok {
			b, err := json.Marshal(right.Value)
			if err != nil {
				return nil, nil, err
			}
			rv = string(b)
		}
		notCompatible = true
	} else {
		rv = right.Value
	}

	if left.DataType == right.DataType || notCompatible {
		return lv, rv, nil
	}

	if left.Primitive == primitives.Number || right.Primitive == primitives.Number {
		lv, err = types.ConvertGoType(left.Value, types.Number)
		if err != nil {
			goto compareAsString
		}
		rv, err = types.ConvertGoType(right.Value, types.Number)
		if err != nil {
			goto compareAsString
		}

		return lv, rv, nil
	}

compareAsString:
	lv, err = types.ConvertGoType(left.Value, types.String)
	if err != nil {
		return nil, nil, err
	}
	rv, err = types.ConvertGoType(right.Value, types.String)
	if err != nil {
		return nil, nil, err
	}

	return lv, rv, nil
}
