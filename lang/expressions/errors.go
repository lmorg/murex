package expressions

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/readline/v4"
)

func raiseError(expression []rune, node *astNodeT, pos int, message string) error {
	if node == nil {
		if expression != nil {
			if pos < 1 {
				pos = 1
			}

			exprRune, exprPos := cropCodeInErrMsg(expression, pos)
			expr := string(exprRune)

			return fmt.Errorf("%s\nExpression: %s\n          : %s\nCharacter : %d",
				message, expr,
				strings.Repeat(" ", exprPos)+"^", pos)
		}

		exprRune, _ := cropCodeInErrMsg(expression, pos)
		expr := string(exprRune)

		return fmt.Errorf("%s\nExpression: %s", message, expr)
	}

	pos = node.pos
	if node.pos < 0 {
		pos = 0
	}

	if expression != nil {
		exprRune, exprPos := cropCodeInErrMsg(expression, pos)
		expr := string(exprRune)

		return fmt.Errorf("%s\nExpression: %s\n          : %s\nCharacter : %d\nSymbol    : %s\nValue     : '%s'",
			message, expr,
			strings.Repeat(" ", exprPos)+"^", pos+1,
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

func cropCodeInErrMsg(code []rune, pos int) ([]rune, int) {
	width := readline.GetTermWidth()
	if width < 80 {
		width = 80
	}

	r := make([]rune, len(code))
	copy(r, code)
	replaceRune(r, '\n', ' ')
	replaceRune(r, '\r', ' ')
	replaceRune(r, '\t', ' ')

	if pos >= len(r) {
		pos = len(r) - 1
	}

	if pos < 0 {
		pos = 0
	}

	return _cropCodeInErrMsg(r, pos, width-20)
}

func _cropCodeInErrMsg(r []rune, pos, width int) ([]rune, int) {
	switch {
	case len(r) <= width:
		return r, pos
	case pos < width-2:
		return r[:width], pos
	case len(r)-pos < width:
		return r[len(r)-width:], width - (len(r) - pos)
	default:
		w := width / 2
		return r[pos-w : pos+w], w
	}
}

func replaceRune(r []rune, find, replace rune) {
	for i := range r {
		if r[i] == find {
			r[i] = replace
		}
	}
}
