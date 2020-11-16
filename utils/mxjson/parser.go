package mxjson

import (
	"fmt"
	"strconv"
)

// Unmarshal converts mxjson file into a Go struct
func Unmarshal(json []byte) (interface{}, error) {
	if len(json) == 0 {
		return nil, nil
	}
	var (
		state   parserState // a lazy way of bypassing the need to build ASTs
		i       int         // cursor position
		b       byte        // current character
		err     error       // any errors
		current *str        // pointer for strings
		value   = newStr()  // current value stored as a string
		valType objectType  // data type for value
		objects = newObjs() // cursor inside nested objects
		comment bool        // cursor inside a comment?
		escape  bool        // next character escaped?
		unquote quote       // cursor inside an unquoted block?
		qSingle quote       // cursor inside a ' quote?
		qDouble quote       // cursor inside a " quote?
		qBrace  = newPair() // cursor inside a ( quote?
		square  = newPair() // cursor inside a [ block?
		curly   = newPair() // cursor inside a { block?
	)

	cannotClose := func() (interface{}, error) {
		return nil, fmt.Errorf("Cannot close `%s` at %d: %s", string([]byte{b}), i, err.Error())
	}

	unexpectedCharacter := func() (interface{}, error) {
		return nil, fmt.Errorf("Unexpected character `%s` at %d", string([]byte{b}), i)
	}

	unexpectedColon := func() (interface{}, error) {
		return nil, fmt.Errorf("Unexpected `%s` at %d. Colons should just be used to separate keys and values", string([]byte{b}), i)
	}

	unexpectedComma := func() (interface{}, error) {
		return nil, fmt.Errorf("Unexpected `%s` at %d. Commas should just be used to separate items mid arrays and maps and not for the end value", string([]byte{b}), i)
	}

	invalidNewLine := func() (interface{}, error) {
		return nil, fmt.Errorf("Cannot have a new line (eg \\n) within single nor double quotes at %d", i)
	}

	cannotOpen := func() (interface{}, error) {
		return nil, fmt.Errorf("Cannot use the brace quotes on key names at %d", i)
	}

	cannotReOpen := func() (interface{}, error) {
		return nil, fmt.Errorf("Quote multiple strings in a key or value block at %d. Should use arrays (`[` and `]`) if multiple values expected", i)
	}

	/*cannotMixArrayTypes := func() (interface{}, error) {
		return nil, fmt.Errorf("Cannot mix array types at %d", i)
	}*/

	store := func() error {
		state++
		if state == stateEndVal {
			switch valType {
			case objBoolean:
				s := current.String()
				switch s {
				case "true":
					objects.SetValue(true)
				case "false":
					objects.SetValue(false)
				default:
					return fmt.Errorf("Boolean values should be either 'true' or 'false', instead received '%s'", s)
				}

			case objNumber:
				i, err := strconv.Atoi(current.String())
				if err != nil {
					return err
				}
				objects.SetValue(i)

			case objString:
				objects.SetValue(current.String())

			default:
				panic("code shouldn't fail here")
			}
		}

		return nil
	}

	for ; i < len(json); i++ {
		b = json[i]

		if comment {
			if b == '\n' {
				comment = false
			}
			continue
		}

		switch b {
		case '#':
			comment = true

		case '\r':
			// do nothing

		case '\n', ' ', '\t':
			switch {
			case qSingle.IsOpen(), qDouble.IsOpen():
				return invalidNewLine()
			case qBrace.IsOpen():
				current.Append(b)
			case unquote.IsOpen():
				unquote.Close()
				err = store()
				if err != nil {
					return nil, err
				}
			default:
				// do nothing
			}

		case '\\':
			switch {
			case qSingle.IsOpen(), qDouble.IsOpen(), qBrace.IsOpen():
				escape = !escape
				if !escape {
					current.Append(b)
				}
			default:
				return unexpectedCharacter()
			}

		case '\'':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case unquote.IsOpen():
				return unexpectedCharacter()
			case qDouble.IsOpen(), qBrace.IsOpen():
				current.Append(b)
			case qSingle.IsOpen():
				qSingle.Close()
				state++
				if state == stateEndVal {
					objects.SetValue(current.String())
				}
			case state == stateBeginKey:
				qSingle.Open(i)
				current = objects.GetKeyPtr()
			case state == stateBeginVal:
				qSingle.Open(i)
				current = value
				valType = objString
			default:
				return cannotReOpen()
			}

		case '"':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case unquote.IsOpen():
				return unexpectedCharacter()
			case qSingle.IsOpen(), qBrace.IsOpen():
				current.Append(b)
			case qDouble.IsOpen():
				qDouble.Close()
				//state++
				//if state == stateEndVal {
				//	objects.SetValue(current.String())
				//}
				err = store()
				if err != nil {
					return nil, err
				}
			case state == stateBeginKey:
				/*switch objects.GetObjType() {
				case objArrayUndefined:
					objects.SetObjType(objArrayString)
				case objArrayNumber, objArrayArray, objArrayMap:
					return cannotMixArrayTypes()
				}*/
				qDouble.Open(i)
				current = objects.GetKeyPtr()
			case state == stateBeginVal:
				qDouble.Open(i)
				current = value
				valType = objString
			default:
				return cannotReOpen()
			}

		case '(':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case unquote.IsOpen():
				return unexpectedCharacter()
			case qSingle.IsOpen(), qDouble.IsOpen():
				current.Append(b)
			case qBrace.IsOpen():
				current.Append(b)
				qBrace.Open(i)
			default:
				if state != stateBeginKey && state != stateBeginVal {
					return cannotOpen()
				}
				qBrace.Open(i)
				current = value
				valType = objString
			}

		case ')':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case unquote.IsOpen():
				return unexpectedCharacter()
			case qSingle.IsOpen(), qDouble.IsOpen():
				current.Append(b)
			case qBrace.len > 1:
				current.Append(b)
				qBrace.Close()
			default:
				err = curly.Close()
				if err != nil {
					return cannotClose()
				}
				//state++
				//objects.SetValue(current.String())
				err = store()
				if err != nil {
					return nil, err
				}
			}

		case '{':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case unquote.IsOpen():
				return unexpectedCharacter()
			case qSingle.IsOpen(), qDouble.IsOpen(), qBrace.IsOpen():
				current.Append(b)
			default:
				state = stateBeginKey
				curly.Open(i)
				objects.New(objMap)
			}

		case '}':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case qSingle.IsOpen(), qDouble.IsOpen(), qBrace.IsOpen():
				current.Append(b)
			case unquote.IsOpen():
				unquote.Close()
				err = store()
				if err != nil {
					return nil, err
				}
				fallthrough
			default:
				err = curly.Close()
				if err != nil {
					return cannotClose()
				}
				state++
				objects.MergeDown()
			}

		case '[':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case unquote.IsOpen():
				return unexpectedCharacter()
			case qSingle.IsOpen(), qDouble.IsOpen(), qBrace.IsOpen():
				current.Append(b)
			default:
				state = stateBeginVal
				square.Open(i)
				objects.New(objArrayUndefined)
			}

		case ']':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case qSingle.IsOpen(), qDouble.IsOpen(), qBrace.IsOpen():
				current.Append(b)
			case unquote.IsOpen():
				unquote.Close()
				err = store()
				if err != nil {
					return nil, err
				}
				fallthrough
			default:
				err = square.Close()
				if err != nil {
					return cannotClose()
				}
				state++
				objects.MergeDown()
			}

		case ':':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case unquote.IsOpen():
				return unexpectedCharacter()
			case state != stateEndKey:
				return unexpectedColon()
			default:
				state++
			}

		case ',':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case qSingle.IsOpen(), qDouble.IsOpen(), qBrace.IsOpen():
				current.Append(b)
			case unquote.IsOpen():
				unquote.Close()
				err = store()
				if err != nil {
					return nil, err
				}
				fallthrough
			case state > stateBeginVal:
				if objects.GetObjType() == objMap {
					state = stateBeginKey
				} else {
					state = stateBeginVal
				}
			default:
				return unexpectedComma()
			}

		case 't', 'r', 'u', 'e',
			'f', 'a', 'l', 's':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case qSingle.IsOpen(), qDouble.IsOpen(), qBrace.IsOpen():
				current.Append(b)
			case unquote.IsOpen():
				current.Append(b)
			case state == stateBeginVal:
				unquote.Open(i)
				current = value
				current.Append(b)
				valType = objBoolean
			default:
				return unexpectedCharacter()
			}

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', '-':
			switch {
			case escape:
				escape = false
				current.Append(b)
			case qSingle.IsOpen(), qDouble.IsOpen(), qBrace.IsOpen():
				current.Append(b)
			case unquote.IsOpen():
				current.Append(b)
			case state == stateBeginVal:
				unquote.Open(i)
				current = value
				current.Append(b)
				valType = objNumber
			default:
				return unexpectedCharacter()
			}

		default:
			switch {
			case escape:
				escape = false
				current.Append(b)
			case unquote.IsOpen():
				return unexpectedCharacter()
			case qSingle.IsOpen(), qDouble.IsOpen(), qBrace.IsOpen():
				current.Append(b)
			default:
				return unexpectedCharacter()
			}
		}

	}

	return objects.nest[0].value, nil
}
