package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
)

func (tree *ParserT) parseExpression(exec, incLogicalOps bool) error {
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

		case ';':
			// end expression
			tree.charPos--
			return nil

		case '?':
			switch tree.nextChar() {
			case '?':
				// null coalescing
				tree.appendAst(symbols.NullCoalescing)
				tree.charPos++
			case ':':
				// elvis
				tree.appendAst(symbols.Elvis)
				tree.charPos++
			default:
				// end expression
				// (deprecated: :)
				tree.charPos--
				return nil
			}

		case '|':
			if incLogicalOps && tree.nextChar() == '|' {
				// logic or
				tree.appendAst(symbols.LogicalOr)
				tree.charPos++
			} else {
				// end expression
				tree.charPos--
				return nil
			}

		case '&':
			if tree.nextChar() == '&' {
				if incLogicalOps {
					// logic and
					tree.appendAst(symbols.LogicalAnd)
					tree.charPos++
					continue
				} else {
					// end expression
					tree.charPos--
					return nil
				}
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

		case ':':
			switch tree.nextChar() {
			case '=':
				// update variable
				tree.appendAst(symbols.AssignUpdate)
				tree.charPos++
			default:
				// invalid symbol
				if !exec {
					return raiseError(tree.expression, nil, tree.charPos, fmt.Sprintf("%s '%s' (%d)",
						errMessage[symbols.Unexpected], string(r), r))
				}
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
				//  might be a function
				if !isBareChar(tree.nextChar()) {
					// unexpected symbol
					tree.appendAst(symbols.Unexpected)
				}
				value := tree.parseBareword()
				if len(tree.expression) <= tree.charPos || tree.expression[tree.charPos] != '(' {
					tree.appendAst(symbols.Unexpected)
					continue
				}
				runes, fn, err := tree.parseFunction(exec, value, varAsValue)
				if err != nil {
					return err
				}
				dt := primitives.NewFunction(fn)
				tree.appendAstWithPrimitive(symbols.Calculated, dt, runes...)
			}

		case '~':
			switch tree.nextChar() {
			case '~':
				// like
				tree.appendAst(symbols.Like)
				tree.charPos++
			case '>':
				// immutable merge
				tree.appendAst(symbols.Merge)
				tree.charPos++
			default:
				// tilde
				home, err := tree.parseVarTilde(exec)
				if err != nil {
					return err
				}
				tree.appendAstWithPrimitive(symbols.Calculated, primitives.NewPrimitive(
					primitives.String,
					home,
				))
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
			case '~':
				// assign and merge
				tree.appendAst(symbols.AssignOrMerge)
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
			err := branch.parseExpression(exec, true)
			if err != nil {
				return err
			}

			if exec {
				dt, err := branch.executeExpr()
				if err != nil {
					return err
				}
				val, err := dt.GetValue()
				if err != nil {
					return err
				}
				tree.appendAstWithPrimitive(symbols.Exp(val.Primitive), dt, '(')
			} else {
				tree.appendAst(symbols.SubExpressionBegin)
			}
			tree.charPos += branch.charPos - 1

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

		case '[':
			switch tree.nextChar() {
			case '{':
				runes, v, mxDt, err := tree.parseLambdaStatement(exec, '$')
				if err != nil {
					return err
				}
				tree.appendAstWithPrimitive(symbols.Calculated, primitives.NewScalar(mxDt, v), runes...)
			default:
				if !exec {
					return raiseError(tree.expression, nil, tree.charPos, fmt.Sprintf("%s '%s' (%d)",
						errMessage[symbols.Unexpected], string(r), r))
				}
				tree.charPos++
				tree.appendAst(symbols.Unexpected, r)
			}

		case ']':
			// end JSON array
			tree.appendAst(symbols.ArrayEnd, r)

		case '}':
			// end JSON object
			tree.appendAst(symbols.ObjectEnd, r)

		case '\'', '`':
			// start string / end string
			// (deprecated: `)
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
			next := tree.nextChar()
			switch {
			case next == '{':
				runes, fn, err := tree.parseSubShell(exec, r, varAsValue)
				if err != nil {
					return err
				}
				dt := primitives.NewFunction(fn)
				tree.appendAstWithPrimitive(symbols.Calculated, dt, runes...)
			case next == 0:
				tree.appendAst(symbols.Unexpected, r)
			default:
				// scalar
				runes, v, mxDt, fn, err := tree.parseVarScalarExpr(exec, false)
				if err != nil {
					return raiseError(tree.expression, nil, tree.charPos, fmt.Sprintf("%s: '%s'",
						err.Error(), string(r)))
				}
				if exec {
					dt := primitives.NewFunction(fn)
					tree.appendAstWithPrimitive(symbols.Scalar, dt, runes...)
				} else {
					dt := primitives.NewScalar(mxDt, v)
					tree.appendAstWithPrimitive(symbols.Scalar, dt, runes...)
				}
			}

		case '@': // TODO: test me please!
			next := tree.nextChar()
			switch {
			case next == '{':
				// sub-shell
				runes, fn, err := tree.parseSubShell(exec, r, varAsValue)
				if err != nil {
					return err
				}
				dt := primitives.NewFunction(fn)
				tree.appendAstWithPrimitive(symbols.Calculated, dt, runes...)
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
				tree.appendAstWithPrimitive(symbols.Calculated, primitives.NewPrimitive(primitives.Array, v), runes...)
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
			case '+':
				tree.appendAst(symbols.PlusPlus)
				tree.charPos++
			default:
				// add
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
			case c == '-' && (tree.charPos+2 >= len(tree.expression) ||
				!(tree.expression[tree.charPos+2] >= '0' && '9' >= tree.expression[tree.charPos+2])):
				tree.appendAst(symbols.MinusMinus)
				tree.charPos++
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
					if len(tree.expression) > tree.charPos && tree.expression[tree.charPos] == '(' {
						runes, fn, err := tree.parseFunction(exec, value, varAsValue)
						if err != nil {
							return err
						}
						dt := primitives.NewFunction(fn)
						tree.appendAstWithPrimitive(symbols.Calculated, dt, runes...)
					} else {
						tree.appendAst(symbols.Bareword, value...)
					}
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

func (tree *ParserT) parseSubExpression(exec bool) (any, error) {
	start := tree.charPos
	tree.charPos++
	branch := NewParser(tree.p, tree.expression[tree.charPos:], 0)
	branch.charOffset = tree.charPos + tree.charOffset
	branch.subExp = true
	err := branch.parseExpression(exec, true)
	if err != nil {
		return nil, err
	}
	tree.charPos += branch.charPos - 1
	if exec {
		dt, err := branch.executeExpr()
		if err != nil {
			return nil, err
		}
		val, err := dt.GetValue()
		if err != nil {
			return nil, err
		}
		return val.Value, nil
	} else {
		return string(tree.expression[start:tree.charPos]), nil
	}

}
