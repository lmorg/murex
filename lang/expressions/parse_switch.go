package expressions

import (
	"errors"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

type SwitchT struct {
	Condition  string
	Parameters [][]rune
}

func (s *SwitchT) ParametersLen() int {
	return len(s.Parameters)
}

func (s *SwitchT) ParametersString(i int) string {
	return string(s.Parameters[i])
}

func (s *SwitchT) ParametersAll() []string {
	params := make([]string, len(s.Parameters))

	for i := range s.Parameters {
		params[i] = string(s.Parameters[i])
	}

	return params
}

func (s *SwitchT) ParametersStringAll() string {
	return strings.Join(s.ParametersAll(), " ")
}

func (s *SwitchT) Block(i int) ([]rune, error) {
	if i > len(s.Parameters)-1 {
		return nil, errors.New("too few parameters")
	}
	if types.IsBlockRune(s.Parameters[i]) {
		return s.Parameters[i], nil
	}
	return nil, errors.New("not a code block")
}

func ParseSwitch(p *lang.Process, expression []rune) ([]*SwitchT, error) {
	var swt []*SwitchT

	for i := 0; i < len(expression); i++ {

		switch expression[i] {
		case ' ', '\t', '\r', '\n', ';':
			continue

		default:
			tree := NewParser(p, expression[i:], 0)
			tree.statement = new(StatementT)
			newPos, err := tree.parseSwitch()
			if err != nil {
				return nil, err
			}
			if len(tree.statement.command) > 0 || len(tree.statement.parameters) > 0 {
				swt = append(swt, &SwitchT{
					Condition:  tree.statement.String(),
					Parameters: tree.statement.parameters,
				})
				i += newPos
			}
		}

	}

	return swt, nil
}

func (tree *ParserT) parseSwitch() (int, error) {
	var escape bool

	for ; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		if escape {
			escape = false
			switch r {
			case 's':
				appendToParam(tree, ' ')
			case 't':
				appendToParam(tree, '\t')
			case 'r':
				appendToParam(tree, '\r')
			case 'n':
				appendToParam(tree, '\n')
			default:
				appendToParam(tree, r)
			}
			continue
		}

		switch r {
		case '#':
			tree.parseComment()

		case '/':
			if tree.nextChar() == '#' {
				if err := tree.parseCommentMultiLine(); err != nil {
					return 0, err
				}
			} else {
				appendToParam(tree, r)
			}

		case '\\':
			escape = true

		case ' ', '\t', '\r':
			// whitespace. do nothing
			if err := tree.nextParameter(); err != nil {
				return 0, err
			}

		case '\n', ';':
			// ignore empty lines while in the statement parser
			if len(tree.statement.command) > 0 {
				if err := tree.nextParameter(); err != nil {
					return 0, err
				}
				tree.charPos--
				return tree.charPos, nil
			}

		case ':':
			if err := processStatementColon(tree, true); err != nil {
				return 0, err
			}

		case '~':
			// tilde
			home, err := tree.parseVarTilde(true)
			if err != nil {
				return 0, err
			}
			appendToParam(tree, []rune(home)...)
			if err := tree.nextParameter(); err != nil {
				return 0, err
			}

		case '(':
			if len(tree.statement.command) == 0 && len(tree.statement.paramTemp) == 0 {
				appendToParam(tree, r)
				if err := tree.nextParameter(); err != nil {
					return 0, err
				}
				continue
			}
			prev := tree.prevChar()
			if prev == ' ' || prev == '\t' {
				// quotes
				value, err := tree.parseParenthesis(true)
				if err != nil {
					return 0, err
				}
				appendToParam(tree, value...)
				continue
			}
			appendToParam(tree, r)

		case '%':
			switch tree.nextChar() {
			case '[':
				// JSON array
				err := processStatementFromExpr(tree, tree.parseArray, true)
				if err != nil {
					return 0, err
				}
			case '{':
				// JSON object
				err := processStatementFromExpr(tree, tree.parseObject, true)
				if err != nil {
					return 0, err
				}
			case '(':
				tree.charPos++
				value, err := tree.parseParenthesis(true)
				if err != nil {
					return 0, err
				}
				appendToParam(tree, value...)
			default:
				appendToParam(tree, r)
			}

		case '{':
			// block literal
			value, err := tree.parseBlockQuote()
			if err != nil {
				return 0, err
			}
			appendToParam(tree, value...)

		case '}':
			return 0, raiseError(tree.expression, nil, tree.charPos,
				"unexpected closing bracket '}'")

		case '\'', '"':
			value, err := tree.parseString(r, r, true)
			if err != nil {
				return 0, err
			}
			appendToParam(tree, value...)
			tree.statement.canHaveZeroLenStr = true
			tree.charPos++

		case '$':
			switch {
			case tree.nextChar() == '{':
				// subshell
				_, fn, err := tree.parseSubShell(true, r, varAsString)
				if err != nil {
					return 0, err
				}
				val, err := fn()
				if err != nil {
					return 0, err
				}
				appendToParam(tree, []rune(val.Value.(string))...)
				tree.statement.canHaveZeroLenStr = true
			case isBareChar(tree.nextChar()):
				// start scalar
				_, v, _, err := tree.parseVarScalar(true, varAsString)
				if err != nil {
					return 0, err
				}
				appendToParam(tree, []rune(v.(string))...)
				tree.statement.canHaveZeroLenStr = true
			default:
				appendToParam(tree, r)
			}

		case '@':
			prev := tree.prevChar()
			switch {
			case prev != ' ' && prev != '\t':
				appendToParam(tree, r)
			case tree.nextChar() == '{':
				// subshell
				if err := tree.nextParameter(); err != nil {
					return 0, err
				}
				value, fn, err := tree.parseSubShell(true, r, varAsString)
				if err != nil {
					return 0, err
				}
				val, err := fn()
				if err != nil {
					return 0, err
				}
				processStatementArrays(tree, value, val.Value, true)
			case isBareChar(tree.nextChar()):
				// start scalar
				if err := tree.nextParameter(); err != nil {
					return 0, err
				}
				value, v, err := tree.parseVarArray(true)
				if err != nil {
					return 0, err
				}
				processStatementArrays(tree, value, v, true)
			default:
				appendToParam(tree, r)
			}

		default:
			// assign value
			appendToParam(tree, r)
		}
	}

	if err := tree.nextParameter(); err != nil {
		return 0, err
	}
	tree.charPos--
	return tree.charPos, nil
}
