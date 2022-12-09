package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func (tree *expTreeT) createStringAst(qStart, qEnd rune, exec bool) error {
	// create JSON dict
	value, nEscapes, err := tree.parseString(qStart, qEnd, exec)
	if err != nil {
		return err
	}
	tree.charPos -= nEscapes
	tree.appendAst(symbols.QuoteDouble, value...)
	tree.charPos += nEscapes + 1
	return nil
}

func (tree *expTreeT) parseParen(exec bool) ([]rune, int, error) {
	value, nEscape, err := tree.parseString('(', ')', exec)
	if err != nil {
		return nil, 0, err
	}

	tree.charPos++

	if len(value) > 0 && exec {
		return value, nEscape, nil
	}

	return value, nEscape, nil
}

func (tree *expTreeT) parseString(qStart, qEnd rune, exec bool) ([]rune, int, error) {
	if exec && qStart != '\'' {
		return tree.parseStringInfix(qEnd, exec)
	}

	var value []rune

	if !exec {
		value = []rune{qStart}
	}

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case r == qEnd:
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
			string([]rune{qEnd})))

endString:
	tree.charPos--
	if !exec {
		value = append(value, qEnd)
	}

	return value, 0, nil
}

func (tree *expTreeT) parseStringInfix(qEnd rune, exec bool) ([]rune, int, error) {
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

		case r == '\\' && qEnd != ')':
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
				subshell, v, _, err := tree.parseSubShell(exec, r, varAsString)
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
			tilde := tree.parseVarTilde(exec)
			value = append(value, []rune(tilde)...)

		case r == qEnd:
			// end quote
			goto endString

		default:
			// string
			value = append(value, r)
		}
	}

	return value, 0, raiseError(
		tree.expression, nil, tree.charPos, fmt.Sprintf(
			"missing closing quote '%s'",
			string([]rune{qEnd})))

endString:
	tree.charPos--
	return value, nEscapes, nil
}

func (tree *expTreeT) parseBlockQuote() ([]rune, error) {
	start := tree.charPos

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '\'', '"':
			_, _, err := tree.parseString(r, r, false)
			if err != nil {
				return nil, err
			}
			tree.charPos++

		case '%':
			switch tree.nextChar() {
			case '[':
				tree.charPos++
				_, _, _, err := tree.parseArray(false)
				if err != nil {
					return nil, err
				}
			case '{':
				tree.charPos++
				_, _, _, err := tree.parseObject(false)
				if err != nil {
					return nil, err
				}
			}

		case '{':
			_, err := tree.parseBlockQuote()
			if err != nil {
				return nil, err
			}

		case '}':
			// end quote
			return tree.expression[start : tree.charPos+1], nil

		default:
			// nothing to do
		}
	}

	return nil, raiseError(tree.expression, nil, tree.charPos, "missing closing brace '}'")
}

func (tree *expTreeT) parseNamedPipe() []rune {
	start := tree.charPos

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		if isBareChar(r) || r == '!' || r == ':' || r == '.' {
			continue
		}

		if r == '>' {
			return tree.expression[start : tree.charPos+1]
		}

		break
	}

	tree.charPos = start
	return nil
}
