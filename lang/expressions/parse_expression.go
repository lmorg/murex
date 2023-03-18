package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
)

func (tree *ParserT) parseExpression(exec bool) error {
	for ; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]
		switch r {
		case '#':
			tree.charPos--
			return nil

		case ' ', '\t', '\r':
			// whitespace. do nothing

		case '\n':
			if len(tree.ast) == 0 {
				// do nothing if just empty lines
				tree.crLf()
				continue
			}
			tree.charPos--
			return nil

		case ';', '|', '?':
			// end expression
			tree.charPos--
			return nil

		case '&':
			if tree.nextChar() == '&' {
				tree.charPos--
				return nil
			}
			tree.appendAst(symbols.Unexpected, r)
			raiseError(tree.expression, nil, tree.charPos, errMessage[symbols.Unexpected])

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
				// tilde
				tree.appendAstWithPrimitive(symbols.Calculated, &primitives.DataType{
					Primitive: primitives.String,
					Value:     tree.parseVarTilde(exec),
				})
			}

		case '>':
			switch tree.nextChar() {
			case '=':
				// greater than or equal to
				tree.appendAst(symbols.GreaterThanOrEqual)
				tree.charPos++
			case '>':
				// redirect (append)
				tree.charPos--
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
			tree.charPos++
			branch := NewParser(tree.p, tree.expression[tree.charPos:], 0)
			branch.charOffset = tree.charPos + tree.charOffset
			branch.subExp = true
			err := branch.parseExpression(exec)
			if err != nil {
				return err
			}
			tree.charPos += branch.charPos - 1
			if exec {
				dt, err := branch.executeExpr()
				if err != nil {
					return err
				}
				tree.appendAstWithPrimitive(symbols.Exp(dt.Primitive), dt)
			} else {
				tree.appendAst(symbols.SubExpressionBegin)
				//tree.appendAst(symbols.Calculated)
			}

		case ')':
			tree.charPos++
			switch {
			case tree.subExp:
				// end sub expression
				return nil
			default:
				raiseError(tree.expression, nil, tree.charPos, errMessage[symbols.SubExpressionEnd])
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
			case '(':
				tree.charPos++
				err := tree.createStringAst('(', ')', exec)
				if err != nil {
					return err
				}
			default:
				tree.appendAst(symbols.Unexpected, r)
				raiseError(tree.expression, nil, tree.charPos, errMessage[symbols.Unexpected])
			}

		//case '[':
		//	tree.createArrayAst(exec)

		case ']':
			// end JSON array
			tree.appendAst(symbols.ArrayEnd, r)

		case '}':
			// end JSON object
			tree.appendAst(symbols.ObjectEnd, r)

		case '\'', '`':
			// start string / end string
			value, err := tree.parseString(r, r, exec)
			if err != nil {
				return err
			}
			tree.appendAst(symbols.QuoteSingle, value...)
			tree.charPos++

		case '"':
			// start string / end string
			value, err := tree.parseString(r, r, exec)
			if err != nil {
				return err
			}
			tree.appendAst(symbols.QuoteDouble, value...)
			tree.charPos++

		case '$':
			switch {
			case tree.nextChar() == '{':
				// subshell
				runes, v, mxDt, err := tree.parseSubShell(exec, r, varAsValue)
				if err != nil {
					return err
				}
				dt := scalar2Primitive(mxDt)
				dt.Value = v
				tree.appendAstWithPrimitive(symbols.Calculated, dt, runes...)
			default:
				// scalar
				runes, v, mxDt, err := tree.parseVarScalar(exec, false, varAsValue)
				if err != nil {
					return raiseError(tree.expression, nil, tree.charPos, fmt.Sprintf("%s: '%s'",
						err.Error(), string(r)))
				}
				if exec && v == nil && mxDt == "" {
					tree.appendAst(symbols.Scalar, runes...)
				} else {
					dt := scalar2Primitive(mxDt)
					dt.Value = v
					tree.appendAstWithPrimitive(symbols.Calculated, dt, runes...)
				}
			}

		case '@': // TODO: test me please!
			next := tree.nextChar()
			switch {
			case next == '{':
				// subshell
				runes, v, _, err := tree.parseSubShell(exec, r, varAsValue)
				if err != nil {
					return err
				}
				tree.appendAstWithPrimitive(symbols.Calculated, &primitives.DataType{
					Primitive: primitives.Array,
					Value:     v,
				}, runes...)
			case next == '[':
				// range (this needs to be a statement)
				return raiseError(tree.expression, nil, tree.charPos, fmt.Sprintf("%s: '%s'",
					errMessage[symbols.Unexpected], string(r)))
			case isBareChar(next):
				// start array
				runes, v, err := tree.parseVarArray(exec)
				if err != nil {
					return err
				}
				tree.appendAstWithPrimitive(symbols.Calculated, &primitives.DataType{
					Primitive: primitives.Array,
					Value:     v,
				}, runes...)
			default:
				if !exec {
					return raiseError(tree.expression, nil, tree.charPos, fmt.Sprintf("%s: '%s'",
						errMessage[symbols.Unexpected], string(r)))
				}
				tree.appendAst(symbols.Unexpected, r)
			}

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
					value := tree.parseNumber()
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
			case '#':
				// multi-line comment
				if err := tree.parseCommentMultiLine(); err != nil {
					return err
				}
			default:
				// divide
				tree.appendAst(symbols.Divide)
			}

		default:
			switch {
			case r >= '0' && '9' >= r:
				// number
				value := tree.parseNumber()
				tree.appendAst(symbols.Number, value...)
				tree.charPos--

			case isBareChar(r):
				// bareword
				value := tree.parseBareword()
				switch string(value) {
				case "true", "false":
					tree.appendAst(symbols.Boolean, value...)
				case "null":
					tree.appendAst(symbols.Null, value...)
				default:
					tree.appendAst(symbols.Bareword, value...)
				}
				tree.charPos--

			default:
				if !exec {
					return raiseError(tree.expression, nil, tree.charPos, fmt.Sprintf("%s '%s' (%d)",
						errMessage[symbols.Unexpected], string(r), r))
				}
				tree.charPos++
				tree.appendAst(symbols.Unexpected, r)
			}
		}
	}

	if tree.charPos >= len(tree.expression)-1 || tree.expression[tree.charPos] == 0 {
		tree.charPos--
	}

	return nil
}
