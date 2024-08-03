package expressions

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func (tree *ParserT) createArrayAst(exec bool) error {
	// create JSON array
	_, dt, err := tree.parseArray(exec)
	if err != nil {
		return err
	}
	tree.appendAstWithPrimitive(symbols.ArrayBegin, dt)
	tree.charPos++
	return nil
}

func (tree *ParserT) parseArray(exec bool) ([]rune, *primitives.DataType, error) {
	var (
		start = tree.charPos
		slice = []interface{}{}
	)

	// check if valid mkarray
	dt, pos, err := tree.parseArrayMaker(exec)
	if err != nil {
		return nil, nil, err
	}
	if dt != nil {
		tree.charPos--
		return nil, dt, nil
	}
	tree.charPos = pos

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '#':
			tree.parseComment()

		case '/':
			// multiline comment
			if tree.nextChar() == '#' {
				if err := tree.parseCommentMultiLine(); err != nil {
					return nil, nil, err
				}
			} else {
				// string
				value := tree.parseArrayBareword()
				slice = append(slice, formatArrayValue(value))
			}

		case '\'', '"':
			// quoted string
			value, err := tree.parseString(r, r, exec)
			if err != nil {
				return nil, nil, err
			}
			slice = append(slice, string(value))
			tree.charPos++

		case '%':
			switch tree.nextChar() {
			case '[', '{':
				// do nothing because action covered in the next iteration
			case '(':
				// start nested string
				tree.charPos++
				value, err := tree.parseParenthesis(exec)
				if err != nil {
					return nil, nil, err
				}
				slice = append(slice, string(value))
			default:
				return nil, nil, fmt.Errorf("'%%' token should be followed with '[', '{', or '(', instead got '%s'", string(r))
			}

		case '[':
			// start nested array
			_, dt, err := tree.parseArray(exec)
			if err != nil {
				return nil, nil, err
			}
			v, err := dt.GetValue()
			if err != nil {
				return nil, nil, err
			}
			slice = append(slice, v.Value)
			tree.charPos++

		case ']':
			// end array
			goto endArray

		case '{':
			// start nested object
			_, dt, err := tree.parseObject(exec)
			if err != nil {
				return nil, nil, err
			}
			v, err := dt.GetValue()
			if err != nil {
				return nil, nil, err
			}
			slice = append(slice, v.Value)
			tree.charPos++

		case '(':
			val, err := tree.parseSubExpression(exec)
			if err != nil {
				return nil, nil, err
			}
			slice = append(slice, val)

		case '$':
			// inline scalar
			switch tree.nextChar() {
			case '{':
				r, fn, err := tree.parseSubShell(exec, r, varAsValue)
				if err != nil {
					return nil, nil, err
				}
				if exec {
					v, err := fn()
					if err != nil {
						return nil, nil, err
					}
					slice = append(slice, v.Value)
				} else {
					slice = append(slice, string(r))
				}

			default:
				_, v, _, err := tree.parseVarScalar(exec, varAsValue)
				if err != nil {
					return nil, nil, err
				}
				slice = append(slice, v)
			}

		case '~':
			// tilde
			home, err := tree.parseVarTilde(exec)
			if err != nil {
				return nil, nil, err
			}
			slice = append(slice, home)

		case '@':
			// inline array
			var (
				name []rune
				v    interface{}
			)
			switch tree.nextChar() {
			case '{':
				r, fn, err := tree.parseSubShell(exec, r, varAsValue)
				if err != nil {
					return nil, nil, err
				}
				if exec {
					val, err := fn()
					v = val.Value
					if err != nil {
						return nil, nil, err
					}
				} else {
					v = string(r)
				}
			default:
				name, v, err = tree.parseVarArray(exec)
				if err != nil {
					return nil, nil, err
				}
				//tree.charPos++
			}
			switch t := v.(type) {
			case nil:
				slice = append(slice, t)
			case []interface{}:
				slice = append(slice, t...)
			case []string, []float64, []int:
				slice = append(slice, v.([]interface{})...)
			case string:
				slice = append(slice, t)
			default:
				return nil, nil, raiseError(tree.expression, nil, tree.charPos,
					fmt.Sprintf(
						"cannot expand %T into an array type\nVariable name: @%s",
						t, string(name)))
			}

		case '\n':
			tree.crLf()
			//fallthrough

		case ',', ' ', '\t', '\r':
			// do nothing

		default:
			value := tree.parseArrayBareword()
			v, err := types.ConvertGoType(value, types.Number)
			if err == nil {
				// is a number
				slice = append(slice, v)
			} else {
				// is a string
				slice = append(slice, formatArrayValue(value))
			}
		}
	}

	return nil, nil, raiseError(tree.expression, nil, tree.charPos,
		"missing closing square bracket ']'")

endArray:
	value := tree.expression[start:tree.charPos]
	tree.charPos--
	dt = primitives.NewPrimitive(primitives.Array, slice)
	return value, dt, nil
}

func (tree *ParserT) parseArrayBareword() []rune {
	var start = tree.charPos

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case ',', ' ', '\t', '\r', '\n', '[', ']', '{', '}', ':', '$', '~', '@', '"', '\'', '%':
			goto endArrayBareword
		case '/':
			if tree.nextChar() == '*' {
				goto endArrayBareword
			}
		}
	}

endArrayBareword:
	value := tree.expression[start:tree.charPos]
	tree.charPos--
	return value
}

func formatArrayValue(value []rune) interface{} {
	s := string(value)
	switch s {
	case "true":
		return true
	case "false":
		return false
	case "null":
		return nil
	default:
		return s
	}
}

func (tree *ParserT) parseArrayMaker(exec bool) (*primitives.DataType, int, error) {
	start := tree.charPos
	var (
		mkarray     bool
		brackets    int = 1
		twoBrackets bool
	)

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '\'', '"', '(', '{', '%':
			return nil, start, nil

		case '.':
			if tree.nextChar() == '.' {
				tree.charPos++
				mkarray = true
			}

		case '[':
			brackets++
			if brackets == 3 {
				return nil, start, nil
			}
			if brackets == 2 {
				twoBrackets = true
			}

		case ']':
			brackets--
			if brackets == 0 {
				goto endParseArrayMaker
			}

		}
	}

	return nil, start, raiseError(
		tree.expression, nil, tree.charPos, "missing closing bracket `]` ")

endParseArrayMaker:
	if !mkarray {
		return nil, start, nil
	}

	if !exec {
		return primitives.NewPrimitive(primitives.Array, make([]any, 0)), tree.charPos, nil
	}

	var block []rune
	if twoBrackets {
		block = append([]rune{'j', 'a', ':', ' '}, tree.expression[start+1:tree.charPos]...)
	} else {
		block = append([]rune{'j', 'a', ':', ' ', '['}, tree.expression[start+1:tree.charPos]...)
		block = append(block, ']')
	}

	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDERR | lang.F_CREATE_STDOUT)
	_, err := fork.Execute(block)
	if err != nil {
		return nil, start, err
	}

	b, err := fork.Stderr.ReadAll()
	if err != nil {
		return nil, start, err
	}
	if len(b) > 0 {
		b = utils.CrLfTrim(b)
		return nil, start, errors.New(string(b))
	}

	var slice []interface{}
	err = fork.Stdout.ReadArrayWithType(tree.p.Context, func(v interface{}, _ string) {
		slice = append(slice, v)
	})

	if err != nil {
		return nil, start, err
	}

	dt := primitives.NewPrimitive(primitives.Array, slice)
	return dt, tree.charPos, nil
}
