package expressions

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/alter"
)

func scalarNameDetokenised(r []rune) []rune {
	switch len(r) {
	case 0:
		return r

	case 1:
		return r[1:]

	case 2, 3:
		if r[1] == '(' {
			return []rune{}
		} else {
			return r[1:]
		}

	default:
		if r[1] == '(' {
			return r[2 : len(r)-1]
		} else {
			return r[1:]
		}
	}
}

func convertAssigneeToBareword(tree *ParserT, node *astNodeT) error {
	switch {
	case node.key == symbols.Bareword:
		return nil

	case node.key == symbols.Scalar && len(node.value) > 1 &&
		node.value[0] == '$' && node.value[1] != '{':

		node.key = symbols.Bareword
		node.value = scalarNameDetokenised(node.value)
		return nil

	case node.key == symbols.QuoteSingle && len(node.value) > 0 && node.value[0] == '(':
		val, err := node.dt.GetValue()
		if err != nil {
			return raiseError(tree.expression, node, 0, fmt.Sprintf(
				"cannot assign a value to %s: %s", node.key.String(), err.Error()))
		}
		s, err := types.ConvertGoType(val.Value, types.String)
		if err != nil {
			return raiseError(tree.expression, node, 0, fmt.Sprintf(
				"cannot assign a value to %s: %s", node.key.String(), err.Error()))
		}
		node.key = symbols.Bareword
		node.value = []rune(s.(string))
		return nil

	default:
		return raiseError(tree.expression, node, 0, fmt.Sprintf(
			"cannot assign a value to %s", node.key.String()))
	}
}

func expAssign(tree *ParserT, overwriteType bool) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if err = convertAssigneeToBareword(tree, leftNode); err != nil {
		return err
	}

	if rightNode.key <= symbols.Bareword {
		return raiseError(tree.expression, rightNode, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, rightNode.key))
	}

	var (
		v  interface{}
		dt string
	)

	right, err := rightNode.dt.GetValue()
	if err != nil {
		return err
	}

	switch right.Primitive {
	case primitives.Array, primitives.Object:
		if overwriteType {
			dt = types.Json
		} else {
			dt = tree.p.Variables.GetDataType(leftNode.Value())
			if dt == "" {
				dt = types.Json
			}
		}

		// this is ugly but Go's JSON marshaller is better behaved than Murexes on with empty values
		if dt == types.Json {
			b, err := right.Marshal()
			if err != nil {
				raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
			}
			v = string(b)
		} else {
			b, err := lang.MarshalData(tree.p, dt, right.Value)
			if err != nil {
				raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
			}
			v = string(b)
		}

	default:
		if overwriteType {
			dt = right.DataType
			v = right.Value

		} else {
			dt = tree.p.Variables.GetDataType(leftNode.Value())
			if dt == "" || dt == types.Null {
				dt = right.DataType
				v = right.Value

			} else {

				v, err = types.ConvertGoType(right.Value, dt)
				if err != nil {
					raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
				}
			}
		}
	}

	err = tree.setVar(leftNode.value, v, dt)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Null, nil),
	})
}

/*func expAssignAdd(tree *ParserT) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if err = convertAssigneeToBareword(tree, leftNode); err != nil {
		return err
	}

	right, err := rightNode.dt.GetValue()
	if err != nil {
		return err
	}

	if leftNode.key != symbols.Bareword {
		return raiseError(tree.expression, leftNode, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, leftNode.key))
	}

	v, dt, err := tree.getVar(leftNode.value, varAsValue)
	if err != nil {
		if !tree.StrictTypes() && strings.Contains(err.Error(), lang.ErrDoesNotExist) {
			// var doesn't exist and we have strict types disabled so lets create var
			v, dt, err = float64(0), types.Number, nil
		} else {
			return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
		}
	}

	var result interface{}

	switch dt {
	case types.Number, types.Float:
		if right.Primitive != primitives.Number {
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.Primitive, dt))
		}
		result = v.(float64) + right.Value.(float64)

	case types.Integer:
		if right.Primitive != primitives.Number {
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.Primitive, dt))
		}
		result = float64(v.(int)) + right.Value.(float64)

	case types.Boolean:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))

	case types.Null:
		switch right.Primitive {
		case primitives.String:
			result = right.Value.(string)
		case primitives.Number:
			result = right.Value.(float64)
		default:
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.Primitive, dt))
		}

	default:
		if right.Primitive != primitives.String {
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.Primitive, dt))
		}
		result = v.(string) + right.Value.(string)
	}

	err = tree.setVar(leftNode.value, result, right.DataType)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Null, nil),
	})
}*/

type assFnT func(float64, float64) float64

func _assAdd(lv float64, rv float64) float64   { return lv + rv }
func _assSub(lv float64, rv float64) float64   { return lv - rv }
func _assMulti(lv float64, rv float64) float64 { return lv * rv }
func _assDiv(lv float64, rv float64) float64   { return lv / rv }

func expAssignAndOperate(tree *ParserT, operation assFnT) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if err = convertAssigneeToBareword(tree, leftNode); err != nil {
		return err
	}

	right, err := rightNode.dt.GetValue()
	if err != nil {
		return err
	}

	if leftNode.key != symbols.Bareword {
		return raiseError(tree.expression, leftNode, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, leftNode.key))
	}

	if rightNode.key != symbols.Number {
		return raiseError(tree.expression, rightNode, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, rightNode.key))
	}

	v, dt, err := tree.getVar(leftNode.value, varAsValue)
	if err != nil {
		if !tree.StrictTypes() && strings.Contains(err.Error(), lang.ErrDoesNotExist) {
			// var doesn't exist and we have strict types disabled so lets create var
			v, dt, err = float64(0), types.Number, nil
		} else {
			return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
		}
	}

	var f float64

	switch dt {
	case types.Number, types.Float:
		f = operation(v.(float64), right.Value.(float64))
	case types.Integer:
		f = operation(float64(v.(int)), right.Value.(float64))
	case types.Null:
		f = operation(0, right.Value.(float64))
	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))
	}

	err = tree.setVar(leftNode.value, f, right.DataType)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Null, nil),
	})
}

func expAssignMerge(tree *ParserT) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	right, err := rightNode.dt.GetValue()
	if err != nil {
		return err
	}

	if err = convertAssigneeToBareword(tree, leftNode); err != nil {
		return err
	}

	if leftNode.key != symbols.Bareword {
		return raiseError(tree.expression, leftNode, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, leftNode.key))
	}

	rightVal := right.Value
	if right.Primitive != primitives.String && reflect.TypeOf(rightVal).Kind() == reflect.String {
		rightVal, err = lang.UnmarshalDataBuffered(tree.p, []byte(rightVal.(string)), right.DataType)
		if err != nil {
			return err
		}
	}

	v, dt, err := tree.getVar(leftNode.value, varAsValue)
	if err != nil {
		if !tree.StrictTypes() && strings.Contains(err.Error(), lang.ErrDoesNotExist) {
			// var doesn't exist and we have strict types disabled so lets create var
			err = tree.setVar(leftNode.value, rightVal, right.DataType)
			if err != nil {
				return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
			}
			return tree.foldAst(&astNodeT{
				key: symbols.Calculated,
				pos: tree.ast[tree.astPos].pos,
				dt:  primitives.NewPrimitive(primitives.Null, nil),
			})
		} else {
			return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
		}
	}

	merged, err := alter.Merge(tree.p.Context, v, nil, rightVal)
	if err != nil {
		return raiseError(tree.expression, leftNode, 0, fmt.Sprintf(
			"cannot perform merge '%s' into '%s': %s",
			right.Value, leftNode.Value(),
			err.Error()))
	}

	err = tree.setVar(leftNode.value, merged, dt)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Null, nil),
	})
}
