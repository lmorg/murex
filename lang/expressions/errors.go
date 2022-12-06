package expressions

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func raiseError(expression []rune, node *astNodeT, message string) error {
	expr := string(expression)
	if len(expr) > 30 {
		expr = expr[:30] + "... (long expression cropped)"
	}
	if node == nil {
		return fmt.Errorf("%s\nExpression: '%s'", message, expr)
	}

	pos := node.pos
	if node.pos < 0 {
		pos = 0
	}

	if expression != nil {
		return fmt.Errorf("%s at char %d\nExpression: '%s'\n          :  %s\nSymbol    : %s\nValue     : '%s'",
			message, pos+1,
			expr, strings.Repeat(" ", pos)+"^",
			node.key.String(), node.Value())
	} else {
		return fmt.Errorf("%s at char %d\nSymbol    : %s\nValue     : '%s'",
			message, pos+1, node.key.String(), node.Value())

	}
}

var errMessage = map[symbols.Exp]string{
	symbols.Undefined:        "parser error",
	symbols.Unexpected:       "unexpected symbol",
	symbols.SubExpressionEnd: "more closing parenthesis then opening parenthesis",
	symbols.ObjectEnd:        "more closing curly braces then opening braces",
	symbols.ArrayEnd:         "more closing square brackets then opening brackets",
	symbols.InvalidHyphen:    "unexpected hyphen",
}
