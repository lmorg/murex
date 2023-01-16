package expressions

import (
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
			appendToParam(tree, r)

		case '?':
			prev := tree.prevChar()
			next := tree.nextChar()
			if prev != ' ' && prev != '\t' &&
				next != ' ' && next != '\t' {
				tree.statement.possibleGlob = exec
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
			appendToParam(tree, r)

		case ':':
			if err := processStatementColon(tree, exec); err != nil {
				return err
			}

		case '=':
			switch tree.nextChar() {
			case '>':
				// generic pipe
				err := tree.nextParameter()
				tree.charPos--
				return err
			default:
				// assign value
				appendToParam(tree, r)
			}

		case '~':
			// tilde
			appendToParam(tree, []rune(tree.parseVarTilde(exec))...)

		case '<':
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
			if len(tree.statement.command) == 0 && len(tree.statement.paramTemp) == 0 {
				appendToParam(tree, r)
				if err := tree.nextParameter(); err != nil {
					return err
				}
				continue
			}
			prev := tree.prevChar()
			if prev == ' ' || prev == '\t' {
				// quotes
				value, err := tree.parseParen(exec)
				if err != nil {
					return err
				}
				appendToParam(tree, value...)
				continue
			}
			appendToParam(tree, r)

		case '%':
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
				value, err := tree.parseParen(exec)
				if err != nil {
					return err
				}
				appendToParam(tree, value...)
			default:
				appendToParam(tree, r)
			}

		case '{':
			// block literal
			value, err := tree.parseBlockQuote()
			if err != nil {
				return err
			}
			appendToParam(tree, value...)

		case '[':
			switch {
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
			value, err := tree.parseString(r, r, exec)
			if err != nil {
				return err
			}
			appendToParam(tree, value...)
			tree.statement.canHaveZeroLenStr = true
			tree.charPos++

		case '`':
			value, err := tree.parseBackTick(r, exec)
			if err != nil {
				return err
			}
			appendToParam(tree, value...)
			tree.charPos++

		case '$':
			switch {
			case tree.nextChar() == '{':
				// subshell
				value, v, _, err := tree.parseSubShell(exec, r, varAsString)
				if err != nil {
					return err
				}
				if exec {
					appendToParam(tree, []rune(v.(string))...)
					tree.statement.canHaveZeroLenStr = true
				} else {
					appendToParam(tree, value...)
				}
			case isBareChar(tree.nextChar()):
				// start scalar
				value, v, _, err := tree.parseVarScalar(exec, varAsString)
				if err != nil {
					return err
				}
				if exec {
					appendToParam(tree, []rune(v.(string))...)
					tree.statement.canHaveZeroLenStr = true
				} else {
					appendToParam(tree, value...)
				}
			default:
				appendToParam(tree, r)
			}

		case '@':
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
				value, v, _, err := tree.parseSubShell(exec, r, varAsString)
				if err != nil {
					return err
				}
				processStatementArrays(tree, value, v, exec)
			case next == '[' && len(tree.statement.command) == 0 && len(tree.statement.paramTemp) == 0:
				// @[ command
				appendToParam(tree, '@', '[')
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
			c := tree.nextChar()
			switch {
			case c == '>':
				err := tree.nextParameter()
				tree.charPos--
				return err
			default:
				// assign value
				appendToParam(tree, r)
			}

		default:
			// assign value
			appendToParam(tree, r)
		}
	}

	err := tree.nextParameter()
	tree.charPos--
	return err
}

func processStatementArrays(tree *ParserT, value []rune, v interface{}, exec bool) error {
	if exec {
		switch v.(type) {
		case []string:
			for i := range v.([]string) {
				value = []rune(v.([]string)[i])
				appendToParam(tree, value...)
				if err := tree.nextParameter(); err != nil {
					return err
				}
			}
		case [][]rune:
			for i := range v.([][]rune) {
				value = v.([][]rune)[i]
				appendToParam(tree, value...)
				if err := tree.nextParameter(); err != nil {
					return err
				}
			}
		case [][]byte:
			for i := range v.([][]byte) {
				value = []rune(string(v.([][]rune)[i]))
				appendToParam(tree, value...)
				if err := tree.nextParameter(); err != nil {
					return err
				}
			}
		case []interface{}:
			for i := range v.([]interface{}) {
				s, err := types.ConvertGoType(v.([]interface{})[i], types.String)
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
			s, err := types.ConvertGoType(v.([]interface{}), types.String)
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
		value, err = dt.Marshal()
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
