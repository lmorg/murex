package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
)

func expGreaterThan(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.dt.Primitive != right.dt.Primitive {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot compare %s with %s", left.dt.Primitive, right.dt.Primitive,
		))
	}

	var value bool

	switch left.dt.Primitive {
	case primitives.Number:
		value = left.dt.Value.(float64) > right.dt.Value.(float64)

	case primitives.String:
		value = left.dt.Value.(string) > right.dt.Value.(string)

	default:
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s with %s types", tree.currentSymbol().key, left.dt.Primitive,
		))
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Exp(left.dt.Primitive),
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     value,
		},
	})
}

func expGreaterThanOrEqual(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.dt.Primitive != right.dt.Primitive {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot compare %s with %s", left.dt.Primitive, right.dt.Primitive,
		))
	}

	var value bool

	switch left.dt.Primitive {
	case primitives.Number:
		value = left.dt.Value.(float64) >= right.dt.Value.(float64)

	case primitives.String:
		value = left.dt.Value.(string) >= right.dt.Value.(string)

	default:
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s with %s types", tree.currentSymbol().key, left.dt.Primitive,
		))
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Exp(left.dt.Primitive),
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     value,
		},
	})
}

func expLessThan(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.dt.Primitive != right.dt.Primitive {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot compare %s with %s", left.dt.Primitive, right.dt.Primitive,
		))
	}

	var value bool

	switch left.dt.Primitive {
	case primitives.Number:
		value = left.dt.Value.(float64) < right.dt.Value.(float64)

	case primitives.String:
		value = left.dt.Value.(string) < right.dt.Value.(string)

	default:
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s with %s types", tree.currentSymbol().key, left.dt.Primitive,
		))
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Exp(left.dt.Primitive),
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     value,
		},
	})
}

func expLessThanOrEqual(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.dt.Primitive != right.dt.Primitive {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot compare %s with %s", left.dt.Primitive, right.dt.Primitive,
		))
	}

	var value bool

	switch left.dt.Primitive {
	case primitives.Number:
		value = left.dt.Value.(float64) <= right.dt.Value.(float64)

	case primitives.String:
		value = left.dt.Value.(string) <= right.dt.Value.(string)

	default:
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"cannot %s with %s types", tree.currentSymbol().key, left.dt.Primitive,
		))
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Exp(left.dt.Primitive),
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     value,
		},
	})
}
