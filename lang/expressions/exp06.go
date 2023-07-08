package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func expGreaterThan(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	var value bool

	lv, rv, err := compareTypes(tree, left, right)
	if err != nil {
		return err
	}

	switch lv.(type) {
	case float64, int:
		value = convertNumber(lv) > convertNumber(rv)

	case string:
		value = lv.(string) > rv.(string)

	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
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

func expGreaterThanOrEqual(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	var value bool

	lv, rv, err := compareTypes(tree, left, right)
	if err != nil {
		return err
	}

	switch lv.(type) {
	case float64, int:
		value = convertNumber(lv) >= convertNumber(rv)

	case string:
		value = lv.(string) >= rv.(string)

	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
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

func expLessThan(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	var value bool

	lv, rv, err := compareTypes(tree, left, right)
	if err != nil {
		return err
	}

	switch lv.(type) {
	case float64, int:
		value = convertNumber(lv) < convertNumber(rv)

	case string:
		value = lv.(string) < rv.(string)

	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
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

func expLessThanOrEqual(tree *ParserT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	var value bool

	lv, rv, err := compareTypes(tree, left, right)
	if err != nil {
		return err
	}

	switch lv.(type) {
	case float64, int:
		value = convertNumber(lv) <= convertNumber(rv)

	case string:
		value = lv.(string) <= rv.(string)

	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
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

func convertNumber(v any) float64 {
	f, err := types.ConvertGoType(v, types.Number)
	if err != nil {
		return 0
	}
	return f.(float64)
}
