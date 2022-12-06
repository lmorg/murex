package expressions

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang/expressions/symbols"
)

func (tree *expTreeT) parse(exec bool) error {
	for ; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]
		switch r {
		case ' ', '\t', '\r':
			// whitespace. do nothing

		case '\n', ';', '|':
			// end expression
			tree.charPos--
			return nil

		case '=':
			switch tree.nextChar() {
			case '=':
				// equals
				tree.appendAst(symbols.EqualTo)
				tree.charPos++
			case '~':
				// regexp
				tree.appendAst(symbols.Regexp)
				tree.charPos++
			case '>':
				// generic pipe
				tree.charPos--
				return nil
			default:
				// assign value
				tree.appendAst(symbols.Assign)
			}

		case '!':
			switch tree.nextChar() {
			case '=':
				// not equal
				tree.appendAst(symbols.NotEqualTo)
				tree.charPos++
			case '~':
				// not regexp
				tree.appendAst(symbols.NotRegexp)
				tree.charPos++
			case '!':
				// not like
				tree.appendAst(symbols.NotLike)
				tree.charPos++
			default:
				// unexpected symbol
				tree.appendAst(symbols.Unexpected)
			}

		case '~':
			switch tree.nextChar() {
			case '~':
				// like
				tree.appendAst(symbols.Like)
				tree.charPos++
			default:
				// unexpected symbol
				tree.appendAst(symbols.Unexpected)
			}

		case '>':
			switch tree.nextChar() {
			case '=':
				// greater than or equal to
				tree.appendAst(symbols.GreaterThanOrEqual)
				tree.charPos++
			case '>':
				// redirect (append)
				tree.charPos++
				return nil
			default:
				// greater than
				tree.appendAst(symbols.GreaterThan)
			}

		case '<':
			switch tree.nextChar() {
			case '=':
				// less than or equal to
				tree.appendAst(symbols.LessThanOrEqual)
				tree.charPos++
			default:
				// less than
				tree.appendAst(symbols.LessThan)
			}

		case '(':
			// create sub expression
			if exec {
				tree.charPos++
				branch := newExpTree(tree.p, tree.expression[tree.charPos:])
				branch.charOffset = tree.charPos + tree.charOffset
				branch.isSubExp = true
				err := branch.parse(exec)
				if err != nil {
					return err
				}

				dt, err := branch.execute()
				if err != nil {
					return err
				}
				tree.appendAstWithPrimitive(symbols.Exp(dt.Primitive), dt)
				tree.charPos += branch.charPos - 1
			} else {
				i, err := ChainParser(tree.expression[tree.charPos+1:], tree.charPos+tree.charOffset+1)
				if err != nil {
					return err
				}
				tree.appendAst(symbols.Calculated)
				tree.charPos += i
			}

		case ')':
			tree.charPos++
			switch {
			case tree.isSubExp:
				// end sub expression
				return nil
			case exec:
				tree.appendAst(symbols.SubExpressionEnd, r)
			default:
				return nil
			}

		case '%':
			switch tree.nextChar() {
			case '[':
				tree.charPos++
				err := tree.createArrayAst(exec)
				if err != nil {
					return err
				}
			case '{':
				tree.charPos++
				err := tree.createObjectAst(exec)
				if err != nil {
					return err
				}
			default:
				tree.appendAst(symbols.Unexpected, r)
			}

		//case '[':
		//	tree.createArrayAst(exec)

		case ']':
			// end JSON array
			tree.appendAst(symbols.ArrayEnd, r)

		//case '{':
		// create JSON object
		// TODO

		case '}':
			// end JSON object
			tree.appendAst(symbols.ObjectEnd, r)

		case '\'', '`':
			// start string / end string
			value, nEscapes, err := tree.parseString(r)
			if err != nil {
				return err
			}
			tree.charPos -= nEscapes
			tree.appendAst(symbols.QuoteSingle, value...)
			tree.charPos += nEscapes + 1

		case '"':
			// start string / end string
			value, nEscapes, err := tree.parseString(r)
			if err != nil {
				return err
			}
			tree.charPos -= nEscapes
			tree.appendAst(symbols.QuoteDouble, value...)
			tree.charPos += nEscapes + 1

		case '$':
			// start scalar
			_, v, mxDt, err := tree.parseVarScalar(exec)
			if err != nil {
				return err
			}
			dt := scalar2Primitive(mxDt)
			dt.Value = v
			tree.appendAstWithPrimitive(symbols.Calculated, dt)
			tree.charPos--

		/*case '@':
		// start array*/

		case '+':
			switch tree.nextChar() {
			case '=':
				// equal add
				tree.appendAst(symbols.AssignAndAdd)
				tree.charPos++
			default:
				// add (+append)
				tree.appendAst(symbols.Add)
			}

		case '-':
			c := tree.nextChar()
			switch {
			case c == '=':
				// equal subtract
				tree.appendAst(symbols.AssignAndSubtract)
				tree.charPos++
			case c >= '0' && '9' >= c:
				if len(tree.ast) == 0 || tree.ast[len(tree.ast)-1].key > symbols.Operations {
					// number
					value := tree.parseNumber(r)
					tree.appendAst(symbols.Number, value...)
					tree.charPos--
				} else {
					// subtract
					tree.appendAst(symbols.Subtract)
				}
			case c == '>':
				// arrow pipe
				tree.charPos--
				return nil
			default:
				tree.appendAst(symbols.Subtract)
				// invalid hyphen
				//tree.appendAst(symbols.InvalidHyphen)
			}

		case '*':
			switch tree.nextChar() {
			case '=':
				// equal multiply
				tree.appendAst(symbols.AssignAndMultiply)
				tree.charPos++
			default:
				// multiply
				tree.appendAst(symbols.Multiply)
			}

		case '/':
			switch tree.nextChar() {
			case '=':
				// equal divide
				tree.appendAst(symbols.AssignAndDivide)
				tree.charPos++
			default:
				// divide
				tree.appendAst(symbols.Divide)
			}

		default:
			switch {
			case r >= '0' && '9' >= r:
				// number
				value := tree.parseNumber(r)
				tree.appendAst(symbols.Number, value...)
				tree.charPos--

			case isBareChar(r):
				// bareword
				value := tree.parseBareword()
				switch string(value) {
				case "true", "false":
					tree.appendAst(symbols.Boolean, value...)
				default:
					tree.appendAst(symbols.Bareword, value...)
				}
				tree.charPos--

			default:
				if !exec {
					return raiseError(tree.expression, nil, fmt.Sprintf("%s at char %d: '%s'",
						errMessage[symbols.Unexpected], tree.charPos, string(r)))
				}
				tree.charPos++
				tree.appendAst(symbols.Unexpected, r)
			}
		}
	}

	tree.charPos--
	return nil
}

func (tree *expTreeT) parseNumber(first rune) []rune {
	// TODO: don't append each time, just return a range
	value := []rune{first}

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case (r >= '0' && '9' >= r) || r == '.':
			value = append(value, r)

		default:
			// not a number
			goto endNumber
		}
	}

endNumber:
	return value
}

func (tree *expTreeT) parseString(quote rune) ([]rune, int, error) {
	var (
		value    []rune
		nEscapes int
		escaped  bool
	)

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch {
		case escaped:
			// end escape
			escaped = false
			value = append(value, r)
			nEscapes++

		case r == '\\':
			// start escape
			escaped = true

		case r == quote:
			// end quote
			goto endString

		default:
			// string
			value = append(value, r)
		}
	}

	return value, 0, fmt.Errorf(
		"missing closing quote (%s) at char %d:\n%s",
		string([]rune{quote}), tree.charPos-len(value), string(append([]rune{quote}, value...)))

endString:
	tree.charPos--
	return value, nEscapes, nil
}

func isBareChar(r rune) bool {
	return r == '_' ||
		(r >= 'a' && 'z' >= r) ||
		(r >= 'A' && 'Z' >= r) ||
		(r >= '0' && '9' >= r)
}

func (tree *expTreeT) parseBareword() []rune {
	i := tree.charPos + 1
	for ; i < len(tree.expression); i++ {
		switch {
		case isBareChar(tree.expression[i]):
			// valid bareword character

		default:
			// not a valid bareword character
			goto endBareword
		}
	}

endBareword:
	value := tree.expression[tree.charPos:i]
	tree.charPos = i
	return value
}

func (tree *expTreeT) parseVarScalar(exec bool) ([]rune, interface{}, string, error) {
	if !isBareChar(tree.nextChar()) {
		return nil, nil, "", errors.New("'$' symbol found but no variable name followed")
	}

	tree.charPos++
	value := tree.parseBareword()

	if !exec {
		// don't getVar() until we come to execute the expression, skip when only
		// parsing syntax
		return nil, nil, "", nil
	}

	v, dataType, err := tree.getVar(string(value))
	return value, v, dataType, err
}

func (tree *expTreeT) parseVarArray(exec bool) ([]rune, interface{}, error) {
	if !isBareChar(tree.nextChar()) {
		return nil, nil, errors.New("'@' symbol found but no variable name followed")
	}

	tree.charPos++
	value := tree.parseBareword()

	if !exec {
		// don't getVar() until we come to execute the expression, skip when only
		// parsing syntax
		return nil, nil, nil
	}

	v, err := tree.getArray(string(value))
	return value, v, err
}
