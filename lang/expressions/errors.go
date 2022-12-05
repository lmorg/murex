package expressions

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/utils/consts"
)

func raiseError(expression []rune, node *astNodeT, message string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%v (%v:%v)\n", err, string(expression), node.pos)
		}
	}()

	pos := node.pos

	if node.pos < 0 {
		pos = 0
	}

	if node == nil {
		return fmt.Errorf("nil ast (%s)", consts.IssueTrackerURL)
	}

	if expression != nil {
		return fmt.Errorf("%s at char %d\nExpression: '%s'\n          :  %s\nSymbol    : %s\nValue     : '%s'",
			message, pos+1,
			string(expression), strings.Repeat(" ", pos)+"^",
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
