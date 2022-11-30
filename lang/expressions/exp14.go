package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
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

	err = tree.setVar(left.Value(), right.dt.Value, right.dt.DataType())
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
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

func expAssignAdd(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.key != symbols.Bareword {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"left side should be a bareword, instead got %s", left.key))
	}

	/*if right.key != symbols.Number {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"right side should not be a %s", right.key))
	}*/

	varName := left.Value()
	v, dt, err := tree.getVar(varName)
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
	}

	var result interface{}

	switch dt {
	case types.Number, types.Float:
		if right.dt.Primitive != primitives.Number {
			return raiseError(tree.currentSymbol(), fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}
		result = v.(float64) + right.dt.Value.(float64)

	case types.Integer:
		if right.dt.Primitive != primitives.Number {
			return raiseError(tree.currentSymbol(), fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}
		result = float64(v.(int)) + right.dt.Value.(float64)

	case types.Boolean, types.Null:
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))

	case "":
		switch right.dt.Primitive {
		case primitives.String:
			result = right.dt.Value.(string)
		case primitives.Number:
			result = right.dt.Value.(float64)
		default:
			return raiseError(tree.currentSymbol(), fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}

	default:
		if right.dt.Primitive != primitives.String {
			return raiseError(tree.currentSymbol(), fmt.Sprintf(
				"cannot %s %s to %s", tree.currentSymbol().key, right.dt.Primitive, dt))
		}
		result = v.(string) + right.dt.Value.(string)
	}

	err = tree.setVar(varName, result, right.dt.DataType())
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
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

func expAssignSubtract(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.key != symbols.Bareword {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"left side should be a bareword, instead got %s", left.key))
	}

	if right.key != symbols.Number {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"right side should not be a %s", right.key))
	}

	varName := left.Value()
	v, dt, err := tree.getVar(varName)
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
	}

	var f float64

	switch dt {
	case types.Number, types.Float:
		f = v.(float64) - right.dt.Value.(float64)
	case types.Integer:
		f = float64(v.(int)) - right.dt.Value.(float64)
	case "":
		f = 0 - right.dt.Value.(float64)
	default:
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))
	}

	err = tree.setVar(varName, f, right.dt.DataType())
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
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

func expAssignMultiply(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.key != symbols.Bareword {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"left side should be a bareword, instead got %s", left.key))
	}

	if right.key != symbols.Number {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"right side should not be a %s", right.key))
	}

	varName := left.Value()
	v, dt, err := tree.getVar(varName)
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
	}

	var f float64

	switch dt {
	case types.Number, types.Float:
		f = v.(float64) * right.dt.Value.(float64)
	case types.Integer:
		f = float64(v.(int)) * right.dt.Value.(float64)
	case "":
		f = 0
	default:
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))
	}

	err = tree.setVar(varName, f, right.dt.DataType())
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
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

func expAssignDivide(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.key != symbols.Bareword {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"left side should be a bareword, instead got %s", left.key))
	}

	if right.key != symbols.Number {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"right side should not be a %s", right.key))
	}

	varName := left.Value()
	v, dt, err := tree.getVar(varName)
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
	}

	var f float64

	switch dt {
	case types.Number, types.Float:
		f = v.(float64) / right.dt.Value.(float64)
	case types.Integer:
		f = float64(v.(int)) / right.dt.Value.(float64)
	case "":
		f = 0 / right.dt.Value.(float64)
	default:
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s %s", tree.currentSymbol().key, dt))
	}

	err = tree.setVar(varName, f, right.dt.DataType())
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
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
