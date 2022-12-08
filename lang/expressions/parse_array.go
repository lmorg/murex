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

func (tree *expTreeT) createArrayAst(exec bool) error {
	// create JSON array
	dt, nEscapes, err := tree.parseArray(exec)
	if err != nil {
		return err
	}
	tree.charPos -= nEscapes
	tree.appendAstWithPrimitive(symbols.ArrayBegin, dt)
	tree.charPos += nEscapes + 1
	return nil
}

func (tree *expTreeT) parseArray(exec bool) (*primitives.DataType, int, error) {
	var (
		nEscapes int
		value    = make([]rune, 0, len(tree.expression)-tree.charPos)
		slice    []interface{}
	)

	// check if valid mkarray
	dt, pos, err := tree.parseArrayMaker(exec)
	if err != nil {
		return nil, 0, err
	}
	if dt != nil {
		tree.charPos--
		return dt, 0, nil
	}
	tree.charPos = pos

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '\'', '"':
			// quoted string
			str, i, err := tree.parseString(r, exec)
			if err != nil {
				return nil, 0, err
			}
			value = append(value, str...)
			nEscapes += i
			tree.charPos++

		case '%':
			switch tree.nextChar() {
			case '[', '{':
				// do nothing because action covered in the next iteration
			default:
				// string
				value = append(value, r)
			}

		case '[':
			// start nested array
			dt, i, err := tree.parseArray(exec)
			if err != nil {
				return nil, 0, err
			}
			nEscapes += i
			slice = append(slice, dt.Value)
			tree.charPos++

		case ']':
			// end array
			if len(value) != 0 {
				slice = append(slice, string(value))
			}
			goto endArray

		case '{':
			// start nested object
			dt, i, err := tree.parseObject(exec)
			if err != nil {
				return nil, 0, err
			}
			nEscapes += i
			slice = append(slice, dt.Value)
			tree.charPos++

		case '$':
			// inline scalar
			_, v, _, err := tree.parseVarScalar(exec, varAsValue)
			if err != nil {
				return nil, 0, err
			}
			slice = append(slice, v)

		case '~':
			// tilde
			slice = append(slice, tree.parseVarTilde(true))

		case '@':
			// inline array
			name, v, err := tree.parseVarArray(exec)
			if err != nil {
				return nil, 0, err
			}
			switch t := v.(type) {
			case nil:
				slice = append(slice, t)
			case []interface{}:
				slice = append(slice, t...)
			case []string, []float64, []int:
				slice = append(slice, v.([]interface{})...)
			default:
				return nil, 0, fmt.Errorf(
					"cannot expand %T into an array type\nVariable name: @%s",
					t, string(name))
			}
			tree.charPos--

		case ',', ' ', '\t', '\r', '\n':
			if len(value) == 0 {
				continue
			}
			slice = append(slice, string(value))
			value = make([]rune, 0, len(tree.expression)-tree.charPos)

		default:
			switch {
			case r == '-':
				next := tree.nextChar()
				if next < '0' || '9' < next {
					value = append(value, r)
					continue
				}
				fallthrough
			case r >= '0' && '9' >= r:
				// number
				value := tree.parseNumber(r)
				tree.charPos--
				v, err := types.ConvertGoType(value, types.Number)
				if err != nil {
					return nil, 0, err
				}
				slice = append(slice, v)
			default:
				// string
				value = append(value, r)
			}
		}
	}

	return nil, 0, raiseError(tree.expression, nil, tree.charPos,
		"missing closing square bracket ']'")

endArray:
	tree.charPos--
	dt = &primitives.DataType{
		Primitive: primitives.Array,
		Value:     slice,
	}
	return dt, nEscapes, nil
}

func (tree *expTreeT) parseArrayMaker(exec bool) (*primitives.DataType, int, error) {
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
		return &primitives.DataType{
			Primitive: primitives.Array,
			Value:     make([]interface{}, 0),
		}, tree.charPos, nil
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

	dt := &primitives.DataType{
		Primitive: primitives.Array,
		Value:     slice,
	}
	return dt, tree.charPos, nil
}
