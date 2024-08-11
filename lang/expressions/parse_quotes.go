package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/utils/ansi"
)

func (tree *ParserT) createStringAst(qStart, qEnd rune, exec bool) error {
	// create JSON dict
	value, err := tree.parseString(qStart, qEnd, exec)
	if err != nil {
		return err
	}
	tree.appendAst(symbols.QuoteParenthesis, value...)
	tree.charPos++
	return nil
}

func (tree *ParserT) parseParenthesis(exec bool) ([]rune, error) {
	value, err := tree.parseString('(', ')', exec)
	if err != nil {
		return nil, err
	}

	tree.charPos++

	if exec {
		value = []rune(ansi.ExpandConsts(string(value)))
	}

	return value, nil
}

func (tree *ParserT) parseString(qStart, qEnd rune, exec bool) ([]rune, error) {
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
		case r == '(' && qEnd == ')':
			v, err := tree.parseParenthesis(exec)
			if err != nil {
				return nil, err
			}
			value = append(value, v...)

		case r == '\n':
			value = append(value, r)
			tree.crLf()

		case r == qEnd:
			// end quote
			goto endString

		default:
			// string
			value = append(value, r)
		}
	}

	return value, raiseError(
		tree.expression, nil, tree.charPos, fmt.Sprintf(
			"missing closing quote (%s)",
			string([]rune{qEnd})))

endString:
	tree.charPos--
	if !exec {
		value = append(value, qEnd)
	}

	return value, nil
}

func (tree *ParserT) parseStringInfix(qEnd rune, exec bool) ([]rune, error) {
	var (
		value   []rune
		escaped bool
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

		case r == '\\' && qEnd != ')':
			// start escape
			escaped = true

		case r == '\n':
			value = append(value, r)
			tree.crLf()

		case r == '$':
			switch {
			case tree.nextChar() == '{':
				// subshell
				subshell, fn, err := tree.parseSubShell(exec, r, varAsString)
				if err != nil {
					return nil, err
				}
				if exec {
					val, err := fn()
					if err != nil {
						return nil, err
					}
					value = append(value, []rune(val.Value.(string))...)
				} else {
					value = append(value, subshell...)
				}
			default:
				// inline scalar
				scalar, v, _, err := tree.parseVarScalar(exec, varAsString)
				if err != nil {
					return nil, err
				}
				if exec {
					value = append(value, []rune(v.(string))...)
				} else {
					value = append(value, scalar...)
				}
			}

		case r == '~':
			// tilde
			home, err := tree.parseVarTilde(exec)
			if err != nil {
				return nil, err
			}
			value = append(value, []rune(home)...)

		case r == '(' && qEnd == ')':
			v, err := tree.parseParenthesis(exec)
			if err != nil {
				return nil, err
			}
			value = append(value, '(')
			value = append(value, v...)
			value = append(value, ')')

		case r == qEnd:
			// end quote
			goto endString

		default:
			// string
			value = append(value, r)
		}
	}

	return value, raiseError(
		tree.expression, nil, tree.charPos, fmt.Sprintf(
			"missing closing quote '%s'",
			string([]rune{qEnd})))

endString:
	tree.charPos--
	return value, nil
}

func (tree *ParserT) parseBackTick(quote rune, exec bool) ([]rune, error) {
	if exec {
		lang.FeatureDeprecated("Automatic backtick (`) conversions to single quote (')", tree.p.FileRef)
		quote = '\''
	}

	var value []rune

	value = []rune{quote}

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '`':
			// end quote
			goto endBackTick

		default:
			// string
			value = append(value, r)
		}
	}

	return value, raiseError(
		tree.expression, nil, tree.charPos, "missing closing backtick, '`'")

endBackTick:
	tree.charPos--
	value = append(value, quote)

	return value, nil
}

func (tree *ParserT) parseBlockQuote() ([]rune, error) {
	start := tree.charPos

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '#':
			if tree.prevChar() == '/' {
				if err := tree.parseCommentMultiLine(); err != nil {
					return nil, err
				}
			} else {
				tree.parseComment()
			}

		case '\n':
			tree.crLf()

		case '%':
			switch tree.nextChar() {
			case '[':
				tree.charPos++
				_, _, err := tree.parseArray(false)
				if err != nil {
					return nil, err
				}
			case '{':
				tree.charPos++
				_, _, err := tree.parseObject(false)
				if err != nil {
					return nil, err
				}
				tree.charPos++
			}

		case '\'':
			_, err := tree.parseString('\'', '\'', false)
			if err != nil {
				return nil, err
			}
			tree.charPos++

		case '"':
			_, err := tree.parseStringInfix('"', false)
			if err != nil {
				return nil, err
			}
			tree.charPos++

		case '(':
			_, err := tree.parseString('(', ')', false)
			if err != nil {
				return nil, err
			}
			tree.charPos++

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

func (tree *ParserT) parseNamedPipe() []rune {
	start := tree.charPos

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		if isBareChar(r) || r == '!' || r == ':' || r == '=' || r == '.' {
			continue
		}

		if r == '>' {
			return tree.expression[start+1 : tree.charPos]
		}

		break
	}

	tree.charPos = start
	return nil
}
