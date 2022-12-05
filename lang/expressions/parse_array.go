package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
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

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '\'', '"':
			str, i, err := tree.parseString(r)
			value = append(value, str...)
			nEscapes += i
			if err != nil {
				return nil, 0, err
			}
			tree.charPos++

		case '%':
			switch tree.nextChar() {
			case '[':

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

		case '$':
			// inline scalar
			_, v, _, err := tree.parseVarScalar(exec)
			if err != nil {
				return nil, 0, err
			}
			//switch dataType {
			//case types.Number, types.Integer, types.Boolean, types.Float:
			//	slice = append(slice, v)
			//default:
			slice = append(slice, v)
			//}
			tree.charPos--

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

	return nil, 0, fmt.Errorf(
		"missing closing square bracket (]) at char %d:\n%s",
		tree.charPos-len(value), string(append([]rune{'['}, value...)))

endArray:
	tree.charPos--
	dt := &primitives.DataType{
		Primitive: primitives.Array,
		Value:     slice,
	}
	return dt, nEscapes, nil
}
