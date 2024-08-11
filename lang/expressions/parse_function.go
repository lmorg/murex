package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/utils"
)

func errNotAllowedInFunctions(cmd []rune, desc string, token ...rune) error {
	return fmt.Errorf("%s, `%s`, are not allowed inside inlined functions: %s(...)",
		string(desc), string(token), string(cmd))
}

func (tree *ParserT) parseFunction(exec bool, cmd []rune, strOrVal varFormatting) ([]rune, primitives.FunctionT, error) {
	params, err := tree.parseFunctionParameters(cmd)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot parse `function(parameters...)`: %s", err.Error())
	}
	r := append(cmd, params...)

	if !exec {
		return r, nil, nil
	}

	fn := func() (*primitives.Value, error) {
		val := new(primitives.Value)
		var err error

		fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
		params = append([]rune{' '}, params[1:len(params)-1]...)
		block := append(cmd, params...)
		val.ExitNum, err = fork.Execute(block)

		//val.Error = fork.Stderr

		if err != nil {
			return val, fmt.Errorf("function `%s` compilation error: %s", string(cmd), err.Error())
		}

		if val.ExitNum != 0 {
			return val, fmt.Errorf("function `%s` returned non-zero exit number (%d)", string(cmd), val.ExitNum)
		}

		b, err := fork.Stdout.ReadAll()
		if err != nil {
			return val, fmt.Errorf("function `%s` STDOUT read error: %s", string(cmd), err.Error())
		}
		b = utils.CrLfTrim(b)
		val.DataType = fork.Stdout.GetDataType()
		val.Value, err = formatBytes(tree, b, val.DataType, strOrVal)
		if err != nil {
			return nil, fmt.Errorf("function `%s` STDOUT conversion error: %s", string(cmd), err.Error())
		}

		return val, err
	}

	return r, fn, nil
}

func (tree *ParserT) parseFunctionParameters(cmd []rune) ([]rune, error) {
	var escape bool
	start := tree.charPos
	tree.charPos++

	for ; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		if escape {
			if r == '\n' {
				return nil, errNotAllowedInFunctions(cmd, "escaped line endings", []rune(`\\n`)...)
			}

			escape = false
			continue
		}

		switch r {
		case '#':
			return nil, errNotAllowedInFunctions(cmd, "line comments", r)

		case '/':
			if tree.nextChar() == '#' {
				if err := tree.parseCommentMultiLine(); err != nil {
					return nil, err
				}
			}

		case '\\':
			escape = true

		case ' ', '\t', '\r':
			// whitespace. do nothing

		case '\n':
			// '\' escaped used at end of line
			return nil, errNotAllowedInFunctions(cmd, "line feeds", []rune(`\n`)...)

		case '?':
			prev := tree.prevChar()
			next := tree.nextChar()
			if prev != ' ' && prev != '\t' &&
				next != ' ' && next != '\t' {
				continue
			}
			return nil, errNotAllowedInFunctions(cmd, "STDERR pipes", r)

		case ';':
			return nil, errNotAllowedInFunctions(cmd, "command terminators", r)

		case '|':
			if tree.nextChar() == '|' {
				return nil, errNotAllowedInFunctions(cmd, "logical operators", []rune(`||`)...)
			}
			return nil, errNotAllowedInFunctions(cmd, "pipes", r)

		case '&':
			if tree.nextChar() == '&' {
				return nil, errNotAllowedInFunctions(cmd, "logical operators", []rune(`&&`)...)
			}

		case '=':
			switch tree.nextChar() {
			case '>':
				// generic pipe
				return nil, errNotAllowedInFunctions(cmd, "generic pipes", []rune(`=>`)...)
			default:
				continue
			}

		case '>':
			switch tree.nextChar() {
			case '>':
				// redirect (append)
				return nil, errNotAllowedInFunctions(cmd, "redirection pipes", []rune(`>>`)...)
			default:
				continue
			}

		case '(':
			_, err := tree.parseParenthesis(false)
			if err != nil {
				return nil, err
			}

		case ')':
			tree.charPos++
			return tree.expression[start:tree.charPos], nil

		case '%':
			switch tree.nextChar() {
			case '[':
				// JSON array
				tree.charPos++
				_, _, err := tree.parseArray(false)
				if err != nil {
					return nil, err
				}
			case '{':
				// JSON object
				tree.charPos++
				_, _, err := tree.parseObject(false)
				if err != nil {
					return nil, err
				}
				tree.charPos++
			case '(':
				// string
				tree.charPos++
				_, err := tree.parseParenthesis(false)
				if err != nil {
					return nil, err
				}
			default:
				continue
			}

		case '{':
			// block literal
			_, err := tree.parseBlockQuote()
			if err != nil {
				return nil, err
			}

		case '}':
			return nil, raiseError(tree.expression, nil, tree.charPos,
				"unexpected closing bracket '}'")

		case '\'', '"':
			_, err := tree.parseString(r, r, false)
			if err != nil {
				return nil, err
			}
			tree.charPos++

		case '$':
			switch {
			case tree.nextChar() == '{':
				// subshell
				_, _, err := tree.parseSubShell(false, r, varAsString)
				if err != nil {
					return nil, err
				}
			default:
				// start scalar
				_, _, _, err := tree.parseVarScalar(false, varAsString)
				if err != nil {
					return nil, raiseError(tree.expression, nil, tree.charPos, err.Error())
				}
			}

		case '@':
			prev := tree.prevChar()
			next := tree.nextChar()
			switch {
			case prev != ' ' && prev != '\t' && prev != 0:
				continue
			case next == '{':
				// subshell
				_, _, err := tree.parseSubShell(false, r, varAsString)
				if err != nil {
					return nil, err
				}
			case isBareChar(tree.nextChar()):
				// start scalar
				_, _, err := tree.parseVarArray(false)
				if err != nil {
					return nil, err
				}
			default:
				continue
			}

		case '-':
			next := tree.nextChar()
			switch {
			case next == '>':
				return nil, errNotAllowedInFunctions(cmd, "pipes", []rune(`->`)...)
			default:
				// assign value
				continue
			}

		default:
			// assign value
			continue
		}
	}

	return nil, raiseError(tree.expression, nil, start,
		fmt.Sprintf("unexpected end of code. Missing closing parenthesis from inlined function `%s(...)`",
			string(cmd)))
}
