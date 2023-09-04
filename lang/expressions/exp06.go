package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func expGtLt(tree *ParserT, compareFloat ltGtFT, compareString ltGtST) error {
	leftNode, rightNode, err := tree.getLeftAndRightSymbols()
	if err != nil {
		return err
	}

	var value bool

	lv, rv, err := compareTypes(tree, leftNode, rightNode)
	if err != nil {
		return err
	}

	left, err := leftNode.dt.GetValue()
	if err != nil { // error should have been captured with compareTypes() but doesn't hurt to be cautious
		return err
	}

	switch lv.(type) {
	case float64, int:
		value = compareFloat(convertNumber(lv), convertNumber(rv))

	case string:
		value = compareString(lv.(string), rv.(string))

	default:
		return raiseError(tree.expression, tree.currentSymbol(), 0, fmt.Sprintf(
			"cannot %s with %s types", tree.currentSymbol().key, left.Primitive,
		))
	}

	return tree.foldAst(&astNodeT{
		key: symbols.Exp(left.Primitive),
		pos: tree.ast[tree.astPos].pos,
		dt:  primitives.NewPrimitive(primitives.Boolean, value),
	})
}

type ltGtFT func(float64, float64) bool
type ltGtST func(string, string) bool

func _ltF(lv, rv float64) bool   { return lv < rv }
func _ltEqF(lv, rv float64) bool { return lv <= rv }
func _gtEqF(lv, rv float64) bool { return lv >= rv }
func _gtF(lv, rv float64) bool   { return lv > rv }
func _ltS(lv, rv string) bool    { return lv < rv }
func _ltEqS(lv, rv string) bool  { return lv <= rv }
func _gtEqS(lv, rv string) bool  { return lv >= rv }
func _gtS(lv, rv string) bool    { return lv > rv }

func convertNumber(v any) float64 {
	f, err := types.ConvertGoType(v, types.Number)
	if err != nil {
		return 0
	}
	return f.(float64)
}
