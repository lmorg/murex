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

	lv, rv, err := compareTypes(tree, left, right)
	if err != nil {
		return nil, err
	}

	return primitives.NewPrimitive(primitives.Boolean, lv == rv), nil
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

	dt.NotValue()

	return tree.foldAst(&astNodeT{
		key: symbols.Boolean,
		pos: tree.ast[tree.astPos].pos,
		dt:  dt,
	})
}

func expLike(tree *ParserT, eq bool) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	left, err := leftNode.dt.GetValue()
	if err != nil {
		return err
	}

	right, err := rightNode.dt.GetValue()
	if err != nil {
		return err
	}

	lv, rv := left.Value, right.Value
	// convert to number, if possible
	if left.Primitive == primitives.String {
		if v, err := types.ConvertGoType(left.Value, types.Number); err == nil {
			lv = v
		}
	}
	if right.Primitive == primitives.String {
		if v, err := types.ConvertGoType(right.Value, types.Number); err == nil {
			rv = v
		}
	}

	// convert to string
	lv, err = types.ConvertGoType(lv, types.String)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}
	rv, err = types.ConvertGoType(rv, types.String)
	if err != nil {
		return raiseError(tree.expression, tree.currentSymbol(), 0, err.Error())
	}

	// trim and lowercase string
	lv = strings.TrimSpace(strings.ToLower(lv.(string)))
	rv = strings.TrimSpace(strings.ToLower(rv.(string)))

	return tree.foldAst(&astNodeT{
		key: symbols.Boolean,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Boolean, (lv == rv) == eq),
	})
}

func expRegexp(tree *ParserT, eq bool) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	left, err := leftNode.dt.GetValue()
	if err != nil {
		return err
	}

	right, err := rightNode.dt.GetValue()
	if err != nil {
		return err
	}

	var lv string

	if tree.StrictTypes() {
		if left.Primitive != primitives.String {
			return raiseError(tree.expression, leftNode, 0, fmt.Sprintf(
				"left side should be %s, instead received %s",
				primitives.String, left.Primitive))
		}
		lv = left.Value.(string)
	} else {
		v, err := types.ConvertGoType(left.Value, types.String)
		if err != nil {
			return fmt.Errorf("cannot convert left side %s into a %s: %s",
				left.Primitive, primitives.String, err.Error())
		}
		lv = v.(string)
	}

	if right.Primitive != primitives.String {
		return raiseError(tree.expression, rightNode, 0, fmt.Sprintf(
			"right side should be a regexp expression, instead received %s",
			right.Primitive))
	}

	rx, err := regexp.Compile(right.Value.(string))
	if err != nil {
		return raiseError(tree.expression, rightNode, 0, err.Error())
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Boolean,
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Boolean, rx.MatchString(lv) == eq),
	})
}
