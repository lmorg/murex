package expressions

import (
	"github.com/lmorg/murex/lang/expressions/primitives"
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
				tree.statement.NextParameter()
				escape = false
				continue
			}

			if !exec {
				appendToParam(tree, '\\', r)
				escape = false
				continue
			}

			switch r {
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

		case '\\':
			escape = true

		case ' ', '\t', '\r':
			// whitespace. do nothing
			tree.statement.NextParameter()

		case '\n':
			// ignore empty lines while in the statement parser
			if len(tree.statement.command) > 0 || len(tree.statement.paramTemp) > 0 {
				tree.statement.NextParameter()
				tree.charPos--
				return nil
			}

		case '?':
			prev := tree.prevChar()
			if prev != ' ' && prev != '\t' {
				appendToParam(tree, r)
				continue
			}
			fallthrough

		case ';', '|':
			// end expression
			tree.statement.NextParameter()
			tree.charPos--
			return nil

		case '&':
			if tree.nextChar() == '&' {
				tree.statement.NextParameter()
				tree.charPos--
				return nil
			}
			appendToParam(tree, r)

		case ':':
			processStatementColon(tree, exec)

		case '=':
			switch tree.nextChar() {
			case '>':
				// generic pipe
				tree.statement.NextParameter()
				tree.charPos--
				return nil
			default:
				// assign value
				appendToParam(tree, r)
			}

		/*case '!':
		switch tree.nextChar() {
		default:
			// unexpected symbol
			tree.appendAst(symbols.Unexpected)
		}*/

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
					tree.statement.NextParameter()
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
					tree.statement.NextParameter()
				} else {
					tree.charPos++
					tree.statement.NextParameter()
					return nil
				}
			default:
				appendToParam(tree, r)
			}

		case '(':
			if len(tree.statement.command) == 0 && len(tree.statement.paramTemp) == 0 {
				appendToParam(tree, r)
				tree.statement.NextParameter()
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
				tree.statement.NextParameter()
			default:
				// index
				appendToParam(tree, r)
				tree.statement.NextParameter()
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
				tree.statement.NextParameter()
				value, v, _, err := tree.parseSubShell(exec, r, varAsString)
				if err != nil {
					return err
				}
				processStatementArrays(tree, value, v, exec)
			case next == '[' && len(tree.statement.command) == 0 && len(tree.statement.paramTemp) == 0:
				// @[ command
				appendToParam(tree, '@', '[')
				tree.charPos++
				tree.statement.NextParameter()
			case isBareChar(tree.nextChar()):
				// start scalar
				tree.statement.NextParameter()
				value, v, err := tree.parseVarArray(exec)
				if err != nil {
					return err
				}
				processStatementArrays(tree, value, v, exec)
			default:
				appendToParam(tree, r)
			}

		case '-':
			c := tree.nextChar()
			switch {
			case c == '>':
				tree.statement.NextParameter()
				tree.charPos--
				return nil
			default:
				// assign value
				appendToParam(tree, r)
			}

		default:
			// assign value
			appendToParam(tree, r)
		}
	}

	tree.statement.NextParameter()
	tree.charPos--
	return nil
}

func processStatementArrays(tree *ParserT, value []rune, v interface{}, exec bool) {
	if exec {
		switch v.(type) {
		case []string:
			for i := range v.([]string) {
				value = []rune(v.([]string)[i])
				appendToParam(tree, value...)
				tree.statement.NextParameter()
			}
		case [][]rune:
			for i := range v.([][]rune) {
				value = v.([][]rune)[i]
				appendToParam(tree, value...)
				tree.statement.NextParameter()
			}
		case [][]byte:
			for i := range v.([][]byte) {
				value = []rune(string(v.([][]rune)[i]))
				appendToParam(tree, value...)
				tree.statement.NextParameter()
			}
		case []interface{}:
			for i := range v.([]interface{}) {
				switch v.([]interface{})[i].(type) {
				case string:
					value = []rune(v.([]interface{})[i].(string))
				case []rune:
					value = v.([]interface{})[i].([]rune)
				case []byte:
					value = []rune(string(v.([]interface{})[i].([]byte)))
				default:
					panic("unexpected data type")
				}
				appendToParam(tree, value...)
				tree.statement.NextParameter()
			}
		case string:
			appendToParam(tree, []rune(value)...)
			tree.statement.NextParameter()
		default:
			panic("unexpected data type")
		}

	} else {
		appendToParam(tree, value...)
		tree.statement.NextParameter()
	}
	tree.statement.NextParameter()
}

func processStatementColon(tree *ParserT, exec bool) {
	switch {
	case len(tree.statement.command) == 0:
		if len(tree.statement.paramTemp) > 0 {
			// is a command
			//if !exec {
			//	appendToParam(tree, ':')
			//}
			tree.statement.NextParameter()
		} else {
			// TODO: is a cast
		}
	default:
		// is a value
		appendToParam(tree, ':')
	}
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
