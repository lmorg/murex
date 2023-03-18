package expressions

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func raiseError(expression []rune, node *astNodeT, pos int, message string) error {
	if pos < 0 {
		pos = 0
	}
	expr := string(expression)
	expr = strings.ReplaceAll(expr, "\r", " ")
	expr = strings.ReplaceAll(expr, "\n", " ")

	if pos < 80 {
		if len(expr) > 80 {
			expr = expr[:80] + "... (long expression cropped)"
		}
	} else {
		expr = expr[:pos]
	}

	if node == nil {
		if expression != nil {
			if pos < 1 {
				pos = 1
			}
			return fmt.Errorf("%s\nExpression: %s\n          : %s\nCharacter : %d",
				message, expr,
				strings.Repeat(" ", pos-1)+"^", pos)
		}
		return fmt.Errorf("%s\nExpression: %s", message, expr)
	}

	pos = node.pos
	if node.pos < 0 {
		pos = 0
	}

	if expression != nil {
		return fmt.Errorf("%s\nExpression: %s\n          : %s\nCharacter : %d\nSymbol    : %s\nValue     : '%s'",
			message, expr,
			strings.Repeat(" ", pos)+"^", pos+1,
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
