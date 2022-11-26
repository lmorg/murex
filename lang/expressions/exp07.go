package expressions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func expEqualTo(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.DataValues,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     left.dt.Value == right.dt.Value,
		},
	})
}

func expNotEqualTo(tree *expTreeT) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.DataValues,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     left.dt.Value != right.dt.Value,
		},
	})
}

func expLike(tree *expTreeT, eq bool) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	leftL, err := types.ConvertGoType(left.dt.Value, types.String)
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
	}

	rightL, err := types.ConvertGoType(right.dt.Value, types.String)
	if err != nil {
		return raiseError(tree.currentSymbol(), err.Error())
	}

	leftL = strings.TrimSpace(strings.ToLower(leftL.(string)))
	rightL = strings.TrimSpace(strings.ToLower(rightL.(string)))

	return tree.foldAst(&astNodeT{
		key: symbols.DataValues,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     (leftL == rightL) == eq,
		},
	})
}

func expRegexp(tree *expTreeT, eq bool) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.dt.Primitive != primitives.String {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"left side should be %s, instead received %s",
			primitives.String, left.dt.Primitive))
	}

	if right.dt.Primitive != primitives.String {
		return raiseError(tree.currentSymbol(), fmt.Sprintf(
			"right side should be a regexp expression, instead received %s",
			right.dt.Primitive))
	}

	rx, err := regexp.Compile(right.dt.Value.(string))
	if err != nil {
		raiseError(tree.currentSymbol(), err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.DataValues,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     rx.MatchString(left.dt.Value.(string)) == eq,
		},
	})
}
