package expressions

import (
	"encoding/json"
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func expAssign(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	if right.key <= symbols.Bareword {
		return raiseError(tree.expression, right, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, right.key))
	}

	var v interface{}
	switch right.dt.Primitive {
	case primitives.Array, primitives.Object: //, primitives.Other:
		b, err := json.Marshal(right.dt.Value)
		if err != nil {
			return err
		}
		v = string(b)

	default:
		v = right.dt.Value
	}

	err = tree.setVar(left.value, v, right.dt.DataType())
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Null,
			Value:     nil,
		},
	})
}

func expAssignAdd(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	/*if right.key != symbols.Number {
		return raiseError(tree.expression,tree.currentSymbol(), fmt.Sprintf(
			"right side should not be a %s", right.key))
	}*/

	v, dt, err := tree.getVar(left.value, varAsValue)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	var result interface{}

	switch dt {
	case types.Number, types.Float:
		if right.dt.Primitive != primitives.Number {
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}
		result = v.(float64) + right.dt.Value.(float64)

	case types.Integer:
		if right.dt.Primitive != primitives.Number {
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}
		result = float64(v.(int)) + right.dt.Value.(float64)

	case types.Boolean:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))

	case types.Null:
		switch right.dt.Primitive {
		case primitives.String:
			result = right.dt.Value.(string)
		case primitives.Number:
			result = right.dt.Value.(float64)
		default:
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}

	default:
		if right.dt.Primitive != primitives.String {
			return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}
		result = v.(string) + right.dt.Value.(string)
	}

	err = tree.setVar(left.value, result, right.dt.DataType())
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Null,
			Value:     nil,
		},
	})
}

func expAssignSubtract(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	if right.key != symbols.Number {
		return raiseError(tree.expression, right, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, right.key))
	}

	v, dt, err := tree.getVar(left.value, varAsValue)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	var f float64

	switch dt {
	case types.Number, types.Float:
		f = v.(float64) - right.dt.Value.(float64)
	case types.Integer:
		f = float64(v.(int)) - right.dt.Value.(float64)
	case types.Null:
		f = 0 - right.dt.Value.(float64)
	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))
	}

	err = tree.setVar(left.value, f, right.dt.DataType())
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Null,
			Value:     nil,
		},
	})
}

func expAssignMultiply(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	if right.key != symbols.Number {
		return raiseError(tree.expression, right, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, right.key))
	}

	v, dt, err := tree.getVar(left.value, varAsValue)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	var f float64

	switch dt {
	case types.Number, types.Float:
		f = v.(float64) * right.dt.Value.(float64)
	case types.Integer:
		f = float64(v.(int)) * right.dt.Value.(float64)
	case types.Null:
		f = 0
	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))
	}

	err = tree.setVar(left.value, f, right.dt.DataType())
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Null,
			Value:     nil,
		},
	})
}

func expAssignDivide(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.key != symbols.Bareword {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side of %s should be a bareword, instead got %s",
			tree.currentSymbol().key, left.key))
	}

	if right.key != symbols.Number {
		return raiseError(tree.expression, right, 0, fmt.Sprintf(
			"right side of %s should not be a %s",
			tree.currentSymbol().key, right.key))
	}

	v, dt, err := tree.getVar(left.value, varAsValue)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	var f float64

	switch dt {
	case types.Number, types.Float:
		f = v.(float64) / right.dt.Value.(float64)
	case types.Integer:
		f = float64(v.(int)) / right.dt.Value.(float64)
	case types.Null:
		f = 0 / right.dt.Value.(float64)
	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))
	}

	err = tree.setVar(left.value, f, right.dt.DataType())
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Calculated,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Null,
			Value:     nil,
		},
	})
}
