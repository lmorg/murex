package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func (tree *ParserT) createObjectAst(exec bool) error {
	// create JSON dict
	_, dt, nEscapes, err := tree.parseObject(exec)
	if err != nil {
		return err
	}
	tree.charPos -= nEscapes
	tree.appendAstWithPrimitive(symbols.ObjectBegin, dt)
	tree.charPos += nEscapes + 1
	return nil
}

type parseObjectT struct {
	keyValueR [2][]rune
	keyValueI [2]interface{}
	stage     int
	obj       map[string]interface{}
}

func newParseObjectT() *parseObjectT {
	o := new(parseObjectT)
	o.obj = make(map[string]interface{})
	return o
}

func (o *parseObjectT) WriteKeyValuePair(pos int) error {
	if o.keyValueI[0] == nil {
		return fmt.Errorf("object key cannot be null before %d", pos)
	}
	if len(o.keyValueR[1]) != 0 {
		o.keyValueI[1] = string(o.keyValueR[1])
	}

	s, err := types.ConvertGoType(o.keyValueI[0], types.String)
	if err != nil {
		return err
	}
	o.obj[s.(string)] = o.keyValueI[1]
	o.keyValueR = [2][]rune{nil, nil}
	o.keyValueI = [2]interface{}{nil, nil}
	o.stage++

	return nil
}

func (tree *ParserT) parseObject(exec bool) ([]rune, *primitives.DataType, int, error) {
	var (
		start    = tree.charPos
		nEscapes int
		o        = newParseObjectT()
	)

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '#':
			tree.parseComment()

		case '\'', '"':
			// quoted string
			str, i, err := tree.parseString(r, r, exec)
			if err != nil {
				return nil, nil, 0, err
			}
			o.keyValueR[o.stage&1] = append(o.keyValueR[o.stage&1], str...)
			nEscapes += i
			tree.charPos++

		case '%':
			switch tree.nextChar() {
			case '[', '{':
				// do nothing because action covered in the next iteration
			case '(':
				// start nested string
				_, i, err := tree.parseParen(exec)
				if err != nil {
					return nil, nil, 0, err
				}
				nEscapes += i
				o.keyValueR[o.stage&1] = append(o.keyValueR[o.stage&1], r)
				tree.charPos++
			default:
				// string
				o.keyValueR[o.stage&1] = append(o.keyValueR[o.stage&1], r)
			}

		case '[':
			// start nested array
			if o.stage&1 == 0 {
				return nil, nil, 0, fmt.Errorf("object keys cannot be an array")
			}
			_, dt, i, err := tree.parseArray(exec)
			if err != nil {
				return nil, nil, 0, err
			}
			nEscapes += i
			o.keyValueI[1] = dt.Value
			tree.charPos++

		case '{':
			// start nested object
			if o.stage&1 == 0 {
				return nil, nil, 0, raiseError(
					tree.expression, nil, tree.charPos, "object keys cannot be another object")
			}
			_, dt, i, err := tree.parseObject(exec)
			if err != nil {
				return nil, nil, 0, err
			}
			nEscapes += i
			o.keyValueI[1] = dt.Value
			tree.charPos++

		case '$':
			switch {
			case isBareChar(tree.nextChar()):
				// inline scalar
				strOrVal := varFormatting(o.stage & 1)
				scalar, v, _, err := tree.parseVarScalar(exec, strOrVal)
				if err != nil {
					return nil, nil, 0, err
				}
				if exec {
					o.keyValueI[o.stage&1] = v
				} else {
					o.keyValueI[o.stage&1] = string(scalar)
				}
			case tree.nextChar() == '{':
				// inline subshell
				strOrVal := varFormatting(o.stage & 1)
				subshell, v, _, err := tree.parseSubShell(exec, r, strOrVal)
				if err != nil {
					return nil, nil, 0, err
				}
				if exec {
					o.keyValueI[o.stage&1] = v
				} else {
					o.keyValueI[o.stage&1] = string(subshell)
				}
			default:
				o.keyValueR[o.stage&1] = append(o.keyValueR[o.stage&1], r)
			}

		case '~':
			// tilde
			o.keyValueI[o.stage&1] = tree.parseVarTilde(exec)

		case '@':
			// inline array
			if o.stage&1 == 0 {
				return nil, nil, 0, raiseError(
					tree.expression, nil, tree.charPos, "arrays cannot be object keys")
			}
			switch tree.nextChar() {
			case '{':
				_, v, _, err := tree.parseSubShell(exec, r, varAsValue)
				if err != nil {
					return nil, nil, 0, err
				}
				o.keyValueI[1] = v
			default:
				_, v, err := tree.parseVarArray(exec)
				if err != nil {
					return nil, nil, 0, err
				}
				o.keyValueI[1] = v
			}

		case ':':
			if o.stage&1 == 1 {
				return nil, nil, 0, raiseError(
					tree.expression, nil, tree.charPos, "invalid symbol ':' expecting ',' or '}' instead")
			}
			o.stage++
			if o.keyValueI[0] != nil {
				continue
			}
			o.keyValueI[0] = string(o.keyValueR[0])

		case '\n':
			if o.stage&1 == 0 && (len(o.keyValueR[0]) > 0 || o.keyValueI[0] != nil) {
				return nil, nil, 0, raiseError(
					tree.expression, nil, tree.charPos,
					"unexpected new line, expecting ':' instead")
			}

			err := o.WriteKeyValuePair(tree.charPos)
			if err != nil {
				return nil, nil, 0, err
			}

		case ',':
			if o.stage&1 == 0 && (len(o.keyValueR[0]) > 0 || o.keyValueI[0] != nil) {
				return nil, nil, 0, raiseError(
					tree.expression, nil, tree.charPos, fmt.Sprintf(
						"invalid symbol '%s', expecting ':' instead",
						string(r)))
			}

			err := o.WriteKeyValuePair(tree.charPos)
			if err != nil {
				return nil, nil, 0, err
			}

		case '}':
			if o.stage&1 == 0 {
				if len(o.keyValueR[0]) > 0 || o.keyValueI[0] != nil {
					return nil, nil, 0, raiseError(
						tree.expression, nil, tree.charPos, fmt.Sprintf(
							"invalid symbol '%s', expecting ':' instead",
							string(r)))
				} else {
					// empty object
					goto endObject
				}
			}

			err := o.WriteKeyValuePair(tree.charPos)
			if err != nil {
				return nil, nil, 0, err
			}

			if r == '}' {
				goto endObject
			}

		case ' ', '\t', '\r':
			continue

		default:
			switch {
			case r == '-':
				next := tree.nextChar()
				if next < '0' || '9' < next {
					o.keyValueR[o.stage&1] = append(o.keyValueR[o.stage&1], r)
					continue
				}
				fallthrough
			case r >= '0' && '9' >= r:
				// number
				value := tree.parseNumber(r)
				tree.charPos--
				v, err := types.ConvertGoType(value, types.Number)
				if err != nil {
					return nil, nil, 0, err
				}
				o.keyValueI[o.stage&1] = v

			default:
				// string
				o.keyValueR[o.stage&1] = append(o.keyValueR[o.stage&1], r)
			}
		}
	}

	return nil, nil, 0, raiseError(
		tree.expression, nil, tree.charPos, "missing closing bracket (})")

endObject:
	value := tree.expression[start:tree.charPos]
	tree.charPos--
	dt := &primitives.DataType{
		Primitive: primitives.Object,
		Value:     o.obj,
	}
	return value, dt, nEscapes, nil
}
