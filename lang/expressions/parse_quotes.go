package expressions

import (
	"fmt"
)

func (tree *expTreeT) parseString(quote rune, exec bool) ([]rune, int, error) {
	if exec && quote == '"' {
		return tree.parseStringInfix(quote, exec)
	}

	var (
		value []rune
	)

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case r == quote:
			// end quote
			goto endString

		default:
			// string
			value = append(value, r)
		}
	}

	return value, 0, raiseError(
		tree.expression, nil, tree.charPos, fmt.Sprintf(
			"missing closing quote (%s)",
			string([]rune{quote})))

endString:
	tree.charPos--
	return value, 0, nil
}

func (tree *expTreeT) parseStringInfix(quote rune, exec bool) ([]rune, int, error) {
	var (
		value    []rune
		nEscapes int
		escaped  bool
	)

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case escaped:
			switch r {
			case 's':
				value = append(value, ' ')
			case 't':
				value = append(value, '\t')
			case 'r':
				value = append(value, '\r')
			case 'n':
				value = append(value, '\n')
			default:
				value = append(value, r)
			}
			// end escape
			escaped = false
			nEscapes++

		case r == '\\':
			// start escape
			escaped = true

		case r == '$':
			switch {
			case isBareChar(tree.nextChar()):
				// inline scalar
				scalar, v, _, err := tree.parseVarScalar(exec, varAsString)
				if err != nil {
					return nil, 0, err
				}
				if exec {
					value = append(value, []rune(v.(string))...)
				} else {
					value = append(value, scalar...)
				}
			case tree.nextChar() == '{':
				// subshell
				subshell, v, _, err := tree.parseSubShell(exec, varAsString)
				if err != nil {
					return nil, 0, err
				}
				if exec {
					value = append(value, []rune(v.(string))...)
				} else {
					value = append(value, subshell...)
				}
			default:
				value = append(value, r)
			}

		case r == '~':
			// tilde
			tilde := tree.parseVarTilde(true)
			value = append(value, []rune(tilde)...)

		case r == quote:
			// end quote
			goto endString

		default:
			// string
			value = append(value, r)
		}
	}

	return value, 0, raiseError(
		tree.expression, nil, tree.charPos, fmt.Sprintf(
			"missing closing quote (%s)",
			string([]rune{quote})))

endString:
	tree.charPos--
	return value, nEscapes, nil
}
