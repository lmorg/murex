package expressions

import (
	"github.com/lmorg/murex/builtins/core/expressions/symbols"
)

func (tree *expTreeT) parse() error {
	for ; tree.charPos < len(tree.expression); tree.charPos++ {
		b := tree.expression[tree.charPos]
		switch b {
		case ' ', '\t', '\r', '\n':
			// whitespace. do nothing

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
			branch := newExpTree(tree.expression[tree.charOffset:])
			branch.charOffset = tree.charPos + tree.charOffset + 1
			err := branch.parse()
			if err != nil {
				return err
			}
			dt, err := branch.execute()
			if err != nil {
				return err
			}
			tree.appendAstWithPrimitive(symbols.Exp(dt.Primitive), dt)
			tree.charPos += branch.charPos

		case ')':
			if tree.charOffset != 0 {
				// end sub expression
				return nil
			}
			tree.appendAst(symbols.SubExpressionEnd, b)

		case '{':
			// create JSON object
			// TODO

		case '}':
			// end JSON object
			tree.appendAst(symbols.ObjectEnd, b)

		case '[':
			// create JSON array
			// TODO

		case ']':
			// end JSON array
			tree.appendAst(symbols.ArrayEnd, b)

		case '\'':
			// start string / end string
			value := tree.parseString(b)
			tree.appendAst(symbols.QuoteSingle, value...)
			tree.charPos++

		case '"':
			// start string / end string
			value := tree.parseString(b)
			tree.appendAst(symbols.QuoteDouble, value...)
			tree.charPos++

		case '$':
			// start scalar

		case '@':
			// start array

		case '%':
			// start object

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
			case c > '0' && '9' > c:
				prev := tree.prevSymbol()
				if prev == nil || prev.key > symbols.Operations {
					// number
					value := tree.parseNumber(b)
					tree.appendAst(symbols.Number, value...)
					tree.charPos--
				} else {
					// subtract
					tree.appendAst(symbols.Subtract)
				}
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
			case b >= '0' && '9' >= b:
				// number
				value := tree.parseNumber(b)
				tree.appendAst(symbols.Number, value...)
				tree.charPos--
			case isBareChar(b):
				// bareword
				value := tree.parseBareword()
				switch string(value) {
				case "true", "false":
					tree.appendAst(symbols.Boolean, value...)
				default:
					tree.appendAst(symbols.Bareword, value...)
				}
			default:
				tree.appendAst(symbols.Unexpected, b)
			}
		}
	}

	return nil
}

func (tree *expTreeT) parseNumber(first byte) []byte {
	// TODO: don't append each time, just return a range
	value := []byte{first}

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		b := tree.expression[tree.charPos]

		switch {
		case (b >= '0' && '9' >= b) || b == '.':
			value = append(value, b)

		case b == ',':
			// TODO: do nothing

		default:
			// not a number
			return value
		}
	}

	return value
}

func (tree *expTreeT) parseString(quote byte) []byte {
	var (
		value   []byte
		escaped bool
	)

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		b := tree.expression[tree.charPos]

		switch {
		case escaped:
			// end escape
			escaped = false
			value = append(value, b)

		case b == '\\':
			// start escape
			escaped = true

		case b == quote:
			// end quote
			goto exit

		default:
			// string
			value = append(value, b)
		}
	}

exit:
	tree.charPos--
	return value
}

func isBareChar(b byte) bool {
	return (b >= 'a' && 'z' >= b) || (b >= 'A' && 'Z' >= b) ||
		(b >= '0' && '9' >= b) || b == '_'
}

func (tree *expTreeT) parseBareword() []byte {
	i := tree.charPos + 1
	for ; i < len(tree.expression); i++ {
		switch {
		case isBareChar(tree.expression[i]):
			// valid bareword character

		default:
			// not a valid bareword character
			goto exit
		}
	}

exit:
	value := tree.expression[tree.charPos:i]
	tree.charPos = i
	return value
}
