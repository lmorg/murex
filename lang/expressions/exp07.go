package expressions

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func expEqualFunc(tree *ParserT) (*primitives.DataType, error) {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return nil, err
	}

	var lv *interface{}
	var rv *interface{}

	switch left.dt.Value.(type) {
	case string, float64, int, bool:
		lv = &left.dt.Value
	default:
		r, err := left.dt.Marshal()
		if err != nil {
			return nil, err
		}
		var s interface{}
		s = string(r)
		lv = &s
	}

	switch right.dt.Value.(type) {
	case string, float64, int, bool:
		rv = &right.dt.Value
	default:
		r, err := right.dt.Marshal()
		if err != nil {
			return nil, err
		}
		var s interface{}
		s = string(r)
		rv = &s
	}

	return &primitives.DataType{
		Primitive: primitives.Boolean,
		Value:     *lv == *rv,
	}, nil
}

func expEqualTo(tree *ParserT) error {
	dt, err := expEqualFunc(tree)
	if err != nil {
		return err
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Boolean,
		pos: tree.ast[tree.astPos].pos,
		dt:  dt,
	})
}

func expNotEqualTo(tree *ParserT) error {
	dt, err := expEqualFunc(tree)
	if err != nil {
		return err
	}

	dt.Value = !dt.Value.(bool)

	return tree.foldAst(&astNodeT{
		key: symbols.Boolean,
		pos: tree.ast[tree.astPos].pos,
		dt:  dt,
	})
}

func expLike(tree *ParserT, eq bool) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	leftL, err := types.ConvertGoType(left.dt.Value, types.String)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	rightL, err := types.ConvertGoType(right.dt.Value, types.String)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	leftL = strings.TrimSpace(strings.ToLower(leftL.(string)))
	rightL = strings.TrimSpace(strings.ToLower(rightL.(string)))

	return tree.foldAst(&astNodeT{
		key: symbols.Boolean,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     (leftL == rightL) == eq,
		},
	})
}

func expRegexp(tree *ParserT, eq bool) error {
	left, right, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	if left.dt.Primitive != primitives.String {
		return raiseError(tree.expression, left, 0, fmt.Sprintf(
			"left side should be %s, instead received %s",
			primitives.String, left.dt.Primitive))
	}

	if right.dt.Primitive != primitives.String {
		return raiseError(tree.expression, right, 0, fmt.Sprintf(
			"right side should be a regexp expression, instead received %s",
			right.dt.Primitive))
	}

	rx, err := regexp.Compile(right.dt.Value.(string))
	if err != nil {
		raiseError(tree.expression, right, 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Boolean,
		pos: tree.ast[tree.astPos].pos,
		dt: &primitives.DataType{
			Primitive: primitives.Boolean,
			Value:     rx.MatchString(left.dt.Value.(string)) == eq,
		},
	})
}
