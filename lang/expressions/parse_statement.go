package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

func appendToParam(tree *ParserT, r ...rune) {
	tree.statement.paramTemp = append(tree.statement.paramTemp, r...)
}

var namedPipeFn = []rune(consts.NamedPipeProcName)

func (tree *ParserT) parseStatement(exec bool) error {
	var escape bool

	for ; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		if escape {
			if r == '\n' {
				tree.crLf()
				if err := tree.nextParameter(); err != nil {
					return err
				}
				escape = false
				continue
			}

			if !exec {
				appendToParam(tree, '\\', r)
				escape = false
				if (r == ' ' || r == '\t') && tree.nextChar() == '#' {
					tree.statement.ignoreCrLf = true
				}
				continue
			}

			switch r {
			case ' ', '\t':
				if tree.nextChar() == '#' {
					tree.statement.ignoreCrLf = true
				} else {
					appendToParam(tree, r)
				}
			case 's':
				appendToParam(tree, ' ')
			case 't':
				appendToParam(tree, '\t')
			case 'r':
				appendToParam(tree, '\r')
			case 'n':
				appendToParam(tree, '\n')
			case '\r':
				continue
			default:
				appendToParam(tree, r)
			}
			escape = false
			continue
		}

		switch r {
		case '#':
			tree.statement.validFunction = false
			tree.parseComment()

		case '/':
			if tree.nextChar() == '#' {
				if err := tree.parseCommentMultiLine(); err != nil {
					return err
				}
			} else {
				appendToParam(tree, r)
			}

		case '\\':
			tree.statement.validFunction = false
			escape = true

		case ' ', '\t', '\r':
			// whitespace. do nothing
			if err := tree.nextParameter(); err != nil {
				return err
			}

		case '\n':
			// '\' escaped used at end of line
			tree.crLf()
			if tree.statement.ignoreCrLf {
				tree.statement.ignoreCrLf = false
				if err := tree.nextParameter(); err != nil {
					return err
				}
				continue
			}
			// ignore empty lines while in the statement parser
			if len(tree.statement.command) > 0 || len(tree.statement.paramTemp) > 0 {
				err := tree.nextParameter()
				tree.charPos--
				return err
			}

		case '*':
			tree.statement.possibleGlob = exec
			tree.statement.validFunction = false
			appendToParam(tree, r)

		case '?':
			prev := tree.prevChar()
			next := tree.nextChar()
			if prev != ' ' && prev != '\t' &&
				next != ' ' && next != '\t' {
				tree.statement.possibleGlob = exec
				tree.statement.validFunction = false
				appendToParam(tree, r)
				continue
			}
			fallthrough

		case ';', '|':
			// end expression
			err := tree.nextParameter()
			tree.charPos--
			return err

		case '&':
			if tree.nextChar() == '&' {
				err := tree.nextParameter()
				tree.charPos--
				return err
			}
			tree.statement.validFunction = false
			appendToParam(tree, r)

		case ':':
			tree.statement.validFunction = false
			if err := processStatementColon(tree, exec); err != nil {
				return err
			}

		case '=':
			tree.statement.validFunction = false
			switch tree.nextChar() {
			case '>':
				// generic pipe
				err := tree.nextParameter()
				tree.charPos--
				return err
			default:
				appendToParam(tree, r)
			}

		case '~':
			tree.statement.validFunction = false
			switch tree.nextChar() {
			case '>':
				// redirect (merge into)
				if len(tree.statement.command) == 0 && len(tree.statement.paramTemp) == 0 {
					appendToParam(tree, r, '>')
					tree.charPos++
					if err := tree.nextParameter(); err != nil {
						return err
					}
				} else {
					if len(tree.statement.paramTemp) > 0 {
						tree.statement.paramTemp = tree.statement.paramTemp[:len(tree.statement.paramTemp)-2]
						if err := tree.nextParameter(); err != nil {
							return err
						}
					}
					tree.charPos--
					return nil
				}
			default:
				// tilde
				home, err := tree.parseVarTilde(exec)
				if err != nil {
					return err
				}
				appendToParam(tree, []rune(home)...)
			}

		case '<':
			tree.statement.validFunction = false
			switch {
			case len(tree.statement.paramTemp) > 0:
				appendToParam(tree, r)
			case len(tree.statement.command) == 0:
				// check if named pipe
				value := tree.parseNamedPipe()
				if len(value) == 0 {
					appendToParam(tree, r)
				} else {
					tree.statement.command = namedPipeFn
					tree.statement.paramTemp = value
					if err := tree.nextParameter(); err != nil {
						return err
					}
				}
			case len(tree.statement.parameters) == 0:
				// check if named pipe
				value := tree.parseNamedPipe()
				if len(value) == 0 {
					appendToParam(tree, r)
				} else {
					tree.statement.namedPipes = append(tree.statement.namedPipes, string(value))
				}
			default:
				appendToParam(tree, r)
			}

		case '>':
			tree.statement.validFunction = false
			switch tree.nextChar() {
			case '>':
				// redirect (append)
				if len(tree.statement.command) == 0 && len(tree.statement.paramTemp) == 0 {
					appendToParam(tree, r, r)
					tree.charPos++
					if err := tree.nextParameter(); err != nil {
						return err
					}
				} else {
					if len(tree.statement.paramTemp) > 0 {
						tree.statement.paramTemp = tree.statement.paramTemp[:len(tree.statement.paramTemp)-2]
						if err := tree.nextParameter(); err != nil {
							return err
						}
					}
					tree.charPos--
					return nil
				}
			default:
				appendToParam(tree, r)
			}

		case '(':
			prev := tree.prevChar()
			switch {
			case len(tree.statement.command) == 0 && len(tree.statement.paramTemp) == 0:
				// command (deprecated)
				appendToParam(tree, r)
				if err := tree.nextParameter(); err != nil {
					return err
				}

			case prev == ' ', prev == '\t':
				// parenthesis quotes
				if exec {
					pos := tree.charPos
					expr, err := tree.parseParenthesis(false)
					if err != nil {
						return err
					}
					dt, err := ExecuteExpr(tree.p, expr)
					if err == nil {
						// parenthesis is an expression
						v, err := dt.GetValue()
						if err != nil {
							return err
						}
						r, err := v.Marshal()
						if err != nil {
							return fmt.Errorf("cannot marshal output of inlined expression in `%s`: %s",
								string(tree.statement.command), err.Error())
						}
						appendToParam(tree, r...)
						continue
					}
					tree.charPos = pos
				}
				// parenthesis is a string (deprecated)
				value, err := tree.parseParenthesis(exec)
				if err != nil {
					return err
				}
				appendToParam(tree, value...)

			case tree.statement.validFunction:
				// function(parameters...)
				value, fn, err := tree.parseFunction(exec, tree.statement.paramTemp, varAsString)
				if err != nil {
					return err
				}
				tree.statement.paramTemp = nil
				if exec {
					val, err := fn()
					if err != nil {
						return err
					}
					appendToParam(tree, []rune(val.Value.(string))...)
				} else {
					appendToParam(tree, value...)
				}
				tree.charPos--
				if err := tree.nextParameter(); err != nil {
					return err
				}
			default:
				tree.statement.validFunction = false
				appendToParam(tree, r)
			}

		case '%':
			tree.statement.validFunction = false
			if !exec {
				appendToParam(tree, '%')
			}
			switch tree.nextChar() {
			case '[':
				// JSON array
				err := processStatementFromExpr(tree, tree.parseArray, exec)
				if err != nil {
					return err
				}
			case '{':
				// JSON object
				err := processStatementFromExpr(tree, tree.parseObject, exec)
				if err != nil {
					return err
				}
				// i don't know why I need the next 4 lines, but tests fail without it
				tree.charPos++
				if !exec {
					appendToParam(tree, '}')
				}
			case '(':
				// string
				tree.charPos++
				value, err := tree.parseParenthesis(exec)
				if err != nil {
					return err
				}
				appendToParam(tree, value...)
			default:
				appendToParam(tree, r)
			}

		case '{':
			tree.statement.validFunction = false
			// block literal
			value, err := tree.parseBlockQuote()
			if err != nil {
				return err
			}
			// was this the start of a parameter...
			var nextParam bool
			if len(tree.statement.paramTemp) == 0 && tree.tokeniseCurlyBrace() {
				nextParam = true
			}
			appendToParam(tree, value...)
			// ...if so lets create a new parameter
			if nextParam {
				if err := tree.nextParameter(); err != nil {
					return err
				}
			}

		case '[':
			tree.statement.validFunction = false
			switch {
			//case len(tree.statement.command) == 1 && tree.prevChar() != '!':
			//	kjhkjhkjhkjhkjhkh
			case len(tree.statement.command) > 0 || len(tree.statement.paramTemp) > 0:
				appendToParam(tree, r)
			case tree.nextChar() == '[':
				// element
				appendToParam(tree, '[', '[')
				tree.charPos++
				if err := tree.nextParameter(); err != nil {
					return err
				}
			default:
				// index
				appendToParam(tree, r)
				if err := tree.nextParameter(); err != nil {
					return err
				}
			}

		case '}':
			return raiseError(tree.expression, nil, tree.charPos,
				"unexpected closing bracket '}'")

		case '\'', '"':
			tree.statement.validFunction = false
			value, err := tree.parseString(r, r, exec)
			if err != nil {
				return err
			}
			appendToParam(tree, value...)
			tree.statement.canHaveZeroLenStr = true
			tree.charPos++

		case '`':
			tree.statement.validFunction = false
			value, err := tree.parseBackTick(r, exec)
			if err != nil {
				return err
			}
			appendToParam(tree, value...)
			tree.charPos++

		case '$':
			tree.statement.validFunction = false
			switch {
			case tree.nextChar() == '{':
				// subshell
				value, fn, err := tree.parseSubShell(exec, r, varAsString)
				if err != nil {
					return err
				}
				if exec {
					val, err := fn()
					if err != nil {
						return err
					}
					appendToParam(tree, []rune(val.Value.(string))...)
					tree.statement.canHaveZeroLenStr = true
				} else {
					appendToParam(tree, value...)
				}
			default:
				// start scalar
				var tokenise bool
				tokenise = tree.tokeniseScalar()
				execScalar := exec && tokenise
				value, v, _, err := tree.parseVarScalar(execScalar, varAsString)
				if err != nil {
					return raiseError(tree.expression, nil, tree.charPos, err.Error())
				}
				switch {
				case execScalar:
					appendToParam(tree, []rune(v.(string))...)
					tree.statement.canHaveZeroLenStr = true
				case !tokenise:
					appendToParam(tree, value[1:]...)
				default:
					appendToParam(tree, value...)
				}
			}

		case '@':
			tree.statement.validFunction = false
			prev := tree.prevChar()
			next := tree.nextChar()
			switch {
			case prev != ' ' && prev != '\t' && prev != 0:
				appendToParam(tree, r)
			case next == '{':
				// subshell
				if err := tree.nextParameter(); err != nil {
					return err
				}
				value, fn, err := tree.parseSubShell(exec, r, varAsString)
				if err != nil {
					return err
				}
				var v any
				if exec {
					val, err := fn()
					if err != nil {
						return err
					}
					v = val.Value
				}
				processStatementArrays(tree, value, v, exec)
			case next == '[' && len(tree.statement.command) == 0 && len(tree.statement.paramTemp) == 0:
				// @[ command
				appendToParam(tree, '@', '[')
				//lang.Deprecated(`@[`, tree.p.FileRef)
				tree.charPos++
				if err := tree.nextParameter(); err != nil {
					return err
				}
			case isBareChar(tree.nextChar()):
				// start scalar
				if err := tree.nextParameter(); err != nil {
					return err
				}
				value, v, err := tree.parseVarArray(exec)
				if err != nil {
					return err
				}
				if exec {
					processStatementArrays(tree, value, v, exec)
				} else {
					appendToParam(tree, value...)
				}
			default:
				appendToParam(tree, r)
			}

		case '-':
			next := tree.nextChar()
			switch {
			case next == '>':
				err := tree.nextParameter()
				tree.charPos--
				return err
			default:
				// assign value
				appendToParam(tree, r)
			}

		default:
			// assign value
			if !isBareChar(r) || r == '!' {
				tree.statement.validFunction = false
			}
			appendToParam(tree, r)
		}
	}

	err := tree.nextParameter()
	tree.charPos--
	return err
}

func processStatementArrays(tree *ParserT, value []rune, v interface{}, exec bool) error {
	if exec {
		switch t := v.(type) {
		case []string:
			for i := range t {
				value = []rune(t[i])
				appendToParam(tree, value...)
				if err := tree.nextParameter(); err != nil {
					return err
				}
			}
		case [][]rune:
			for i := range t {
				value = t[i]
				appendToParam(tree, value...)
				if err := tree.nextParameter(); err != nil {
					return err
				}
			}
		case [][]byte:
			for i := range t {
				value = []rune(string(t[i]))
				appendToParam(tree, value...)
				if err := tree.nextParameter(); err != nil {
					return err
				}
			}
		case []interface{}:
			for i := range t {
				s, err := types.ConvertGoType(t[i], types.String)
				if err != nil {
					return err
				}
				value = []rune(s.(string))
				appendToParam(tree, value...)
				if err := tree.nextParameter(); err != nil {
					return err
				}
			}
		case string:
			appendToParam(tree, []rune(value)...)
			if err := tree.nextParameter(); err != nil {
				return err
			}
		default:
			s, err := types.ConvertGoType(t, types.String)
			if err != nil {
				return err
			}
			value = []rune(s.(string))
			appendToParam(tree, value...)
		}

	} else {
		appendToParam(tree, value...)
		if err := tree.nextParameter(); err != nil {
			return err
		}
	}
	return tree.nextParameter()
}

func processStatementColon(tree *ParserT, exec bool) error {
	tree.statement.asStatement = true

	switch {
	case len(tree.statement.command) == 0:
		if len(tree.statement.paramTemp) > 0 {
			// is a command
			if !exec {
				appendToParam(tree, ':')
			}
			return tree.nextParameter()
		} else {
			// is a cast
			tree.charPos++
			tree.statement.cast = tree.parseBareword()
		}
	default:
		// is a value
		appendToParam(tree, ':')
	}

	return nil
}

type parserMethodT func(bool) ([]rune, *primitives.DataType, error)

func processStatementFromExpr(tree *ParserT, method parserMethodT, exec bool) error {
	tree.charPos++
	value, dt, err := method(exec)
	if err != nil {
		return err
	}

	if exec {
		val, err := dt.GetValue()
		if err != nil {
			return err
		}
		value, err = val.Marshal()
		if err != nil {
			return err
		}
		appendToParam(tree, value...)
	} else {
		appendToParam(tree, value...)
	}

	if exec {
		tree.charPos++
	}
	return nil
}
