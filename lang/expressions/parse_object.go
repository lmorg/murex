package expressions

const (
	expectKey   = 0
	expectColon = 2
	expectValue = 1
)

/*func (tree *expTreeT) parseObject(exec bool) (*primitives.DataType, int, error) {
	var (
		nEscapes int
		keyValue = make([][]rune, 2, 2)
		stage    int
		obj      = make(map[string]interface{})
	)

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '\'', '"':
			str, i, err := tree.parseString(r)
			if err != nil {
				return nil, 0, err
			}
			keyValue[stage] = str
			nEscapes += i
			//stage=
			//tree.charPos++

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
			goto endObject

		case '{':
			// start nested array
			dt, i, err := tree.parseArray(exec)
			if err != nil {
				return nil, 0, err
			}
			nEscapes += i
			slice = append(slice, dt.Value)
			tree.charPos++

		case '}':
			// end object
			goto endObject

		case '$':
			// inline scalar
			_, v, dataType, err := tree.parseVarScalar(exec)
			if err != nil {
				return nil, 0, err
			}
			switch dataType {
			case types.Number, types.Integer, types.Boolean, types.Float:
				slice = append(slice, v)
			default:
				slice = append(slice, v)
			}
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
				tree.charPos--
			}
		}
	}

	return nil, 0, fmt.Errorf(
		"missing closing square bracket (]) at char %d:\n%s",
		tree.charPos-len(value), string(append([]rune{'['}, value...)))

endObject:
	tree.charPos--
	dt := &primitives.DataType{
		Primitive: primitives.Object,
		Value:     obj,
	}
	return dt, nEscapes, nil
}
*/
