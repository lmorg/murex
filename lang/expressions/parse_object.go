package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
)

func (tree *ParserT) createObjectAst(exec bool) error {
	// create JSON dict
	_, dt, err := tree.parseObject(exec)
	if err != nil {
		return err
	}
	tree.appendAstWithPrimitive(symbols.ObjectBegin, dt)
	tree.charPos++

	return nil
}

func (tree *ParserT) parseObject(exec bool) ([]rune, *primitives.DataType, error) {
	start := tree.charPos
	o := newParseObjectT(tree)

	for tree.charPos++; tree.charPos < len(tree.expression); tree.charPos++ {
		r := tree.expression[tree.charPos]

		switch r {
		case '#':
			tree.parseComment()

		case '/':
			if tree.nextChar() == '#' {
				if err := tree.parseCommentMultiLine(); err != nil {
					return nil, nil, err
				}
			} else {
				err := o.AppendRune(r)
				if err != nil {
					return nil, nil, err
				}
			}

		case '\'', '"':
			// quoted string
			value, err := tree.parseString(r, r, exec)
			if err != nil {
				return nil, nil, err
			}
			err = o.UpdateInterface(string(value))
			if err != nil {
				return nil, nil, err
			}
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
				err = o.UpdateInterface(string(value))
				if err != nil {
					return nil, nil, err
				}
			default:
				// string
				err := o.AppendRune(r)
				if err != nil {
					return nil, nil, err
				}
			}

		case '[':
			// start nested array
			if o.stage == 0 {
				return nil, nil, fmt.Errorf("object keys cannot be an array")
			}
			_, dt, err := tree.parseArray(exec)
			if err != nil {
				return nil, nil, err
			}
			v, err := dt.GetValue()
			if err != nil {
				return nil, nil, err
			}
			err = o.UpdateInterface(v.Value)
			if err != nil {
				return nil, nil, err
			}
			tree.charPos++

		case '{':
			// start nested object
			if o.stage == 0 {
				return nil, nil, raiseError(
					tree.expression, nil, tree.charPos, "object keys cannot be another object")
			}
			_, dt, err := tree.parseObject(exec)
			if err != nil {
				return nil, nil, err
			}
			v, err := dt.GetValue()
			if err != nil {
				return nil, nil, err
			}
			err = o.UpdateInterface(v.Value)
			if err != nil {
				return nil, nil, err
			}
			tree.charPos++

		case '(':
			val, err := tree.parseSubExpression(exec)
			if err != nil {
				return nil, nil, err
			}
			err = o.UpdateInterface(val)
			if err != nil {
				return nil, nil, err
			}

		case '$':
			switch {
			case tree.nextChar() == '{':
				// inline sub-shell
				strOrVal := varFormatting(o.stage)
				subshell, fn, err := tree.parseSubShell(exec, r, strOrVal)
				if err != nil {
					return nil, nil, err
				}
				if exec {
					val, err := fn()
					if err != nil {
						return nil, nil, err
					}
					err = o.UpdateInterface(val.Value)
					if err != nil {
						return nil, nil, err
					}
				} else {
					err = o.UpdateInterface(string(subshell))
					if err != nil {
						return nil, nil, err
					}
				}
			default:
				// inline scalar
				strOrVal := varFormatting(o.stage)
				scalar, val, _, err := tree.parseVarScalar(exec, strOrVal)
				if err != nil {
					return nil, nil, err
				}
				if exec {
					err = o.UpdateInterface(val)
					if err != nil {
						return nil, nil, err
					}
				} else {
					err = o.UpdateInterface(string(scalar))
					if err != nil {
						return nil, nil, err
					}
				}
			}

		case '~':
			// tilde
			home, err := tree.parseVarTilde(exec)
			if err != nil {
				return nil, nil, err
			}
			err = o.AppendRune([]rune(home)...)
			if err != nil {
				return nil, nil, err
			}

		case '@':
			switch tree.nextChar() {
			case '{':
				subshell, fn, err := tree.parseSubShell(exec, r, varAsValue)
				if err != nil {
					return nil, nil, err
				}
				if exec {
					val, err := fn()
					if err != nil {
						return nil, nil, err
					}
					err = o.UpdateInterface(val.Value)
					if err != nil {
						return nil, nil, err
					}
				} else {
					err = o.UpdateInterface(string(subshell))
					if err != nil {
						return nil, nil, err
					}
				}
			default:
				_, v, err := tree.parseVarArray(exec)
				if err != nil {
					return nil, nil, err
				}
				err = o.UpdateInterface(v)
				if err != nil {
					return nil, nil, err
				}
			}

		case ':':
			if o.stage != OBJ_STAGE_KEY {
				return nil, nil, raiseError(
					tree.expression, nil, tree.charPos, "invalid symbol ':' expecting ',' or '}' instead")
			}
			o.stage++

		case '\n':
			err := o.WriteKeyValuePair()
			if err != nil {
				return nil, nil, err
			}
			tree.crLf()

		case ',':
			err := o.WriteKeyValuePair()
			if err != nil {
				return nil, nil, err
			}

		case '}':
			err := o.WriteKeyValuePair()
			if err != nil {
				return nil, nil, err
			}
			goto endObject

		case '\r':
			continue

		case ' ', '\t':
			if !o.IsValueUndefined() {
				o.keyValue[o.stage].ValueSet = true
			}
			continue

		default:
			value := tree.parseArrayBareword()
			v, err := types.ConvertGoType(value, types.Number)
			if err == nil {
				// is a number
				err = o.UpdateInterface(v)
			} else {
				// is a string
				s := string(value)
				switch s {
				case "true":
					err = o.UpdateInterface(true)
				case "false":
					err = o.UpdateInterface(false)
				case "null":
					err = o.UpdateInterface(nil)
				default:
					err = o.UpdateInterface(s)
				}
			}
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return nil, nil, raiseError(
		tree.expression, nil, tree.charPos, "missing closing bracket '}'")

endObject:
	value := tree.expression[start:tree.charPos]
	tree.charPos--
	dt := primitives.NewPrimitive(primitives.Object, o.obj)
	return value, dt, nil
}
