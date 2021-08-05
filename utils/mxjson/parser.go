package mxjson

import (
	"fmt"
	"strconv"
)

// Parse converts mxjson file into a Go struct
func Parse(json []byte) (interface{}, error) {
	if len(json) == 0 {
		return nil, nil
	}
	var (
		state   parserState // a lazy way of bypassing the need to build ASTs
		i, y, x = 0, 1, 0   // cursor position
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
		return nil, fmt.Errorf("Cannot close `%s` at %d(%d,%d): %s", string([]byte{b}), i+1, y, x, err.Error())
	}

	unexpectedCharacter := func() (interface{}, error) {
		return nil, fmt.Errorf("Unexpected character `%s` at %d(%d,%d)", string([]byte{b}), i+1, y, x)
	}

	unexpectedColon := func() (interface{}, error) {
		return nil, fmt.Errorf("Unexpected `%s` at %d(%d,%d). Colons should just be used to separate keys and values", string([]byte{b}), i+1, y, x)
	}

	unexpectedComma := func() (interface{}, error) {
		return nil, fmt.Errorf("Unexpected `%s` at %d(%d,%d). Commas should just be used to separate items mid arrays and maps and not for the end value nor to separate keys and values in a map", string([]byte{b}), i+1, y, x)
	}

	invalidNewLine := func() (interface{}, error) {
		return nil, fmt.Errorf("Cannot have a new line (eg \\n) within single nor double quotes at %d(%d,%d)", i+1, y, x)
	}

	cannotOpen := func() (interface{}, error) {
		return nil, fmt.Errorf("Cannot use the brace quotes on key names at %d(%d,%d)", i+1, y, x)
	}

	cannotReOpen := func() (interface{}, error) {
		return nil, fmt.Errorf("Quote multiple strings in a key or value block at %d(%d,%d). Strings should be comma separated and inside arrays block (`[` and `]`) where multiple values are expected", i+1, y, x)
	}

	keysOutsideMap := func() (interface{}, error) {
		return nil, fmt.Errorf("Keys outside of map blocks, `{...}`, at %d(%d,%d)", i+1, y, x)
	}

	/*cannotMixArrayTypes := func() (interface{}, error) {
		return nil, fmt.Errorf("Cannot mix array types at %d(%d,%d)", i+1,x,y)
	}*/

	store := func() error {
		state++

		if state != stateEndVal {
			return nil
		}

		pos := i - current.len + 1

		switch valType {
		case objBoolean:
			s := current.String()
			switch s {
			case "true":
				objects.SetValue(true)
			case "false":
				objects.SetValue(false)
			default:
				return fmt.Errorf("Boolean values should be either 'true' or 'false', instead received '%s' at %d(%d,%d)", s, pos, y, x)
			}

		case objNumber:
			i, err := strconv.ParseFloat(current.String(), 64)
			if err != nil {
				return fmt.Errorf("%s at %d(%d,%d)", err.Error(), pos, y, x)
			}
			objects.SetValue(i)

		case objString:
			objects.SetValue(current.String())

		default:
			panic("code shouldn't fail here")
		}

		return nil
	}

	for ; i < len(json); i++ {
		b = json[i]
		x++

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

		case '\n':
			y++
			x = 0
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

		case ' ', '\t':
			switch {
			case qSingle.IsOpen(), qDouble.IsOpen():
				current.Append(b)
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
				if objects.len < 0 {
					return keysOutsideMap()
				}
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
				err = store()
				if err != nil {
					return nil, err
				}
			case state == stateBeginKey:
				if objects.len < 0 {
					return keysOutsideMap()
				}
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
				err = qBrace.Close()
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
			case qSingle.IsOpen(), qDouble.IsOpen(), qBrace.IsOpen():
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
				switch objects.GetObjType() {
				case objMap:
					state = stateBeginKey
				case objUndefined:
					return unexpectedComma()
				default:
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

	switch {
	case qSingle.IsOpen():
		return nil, fmt.Errorf("Single quote, `'`, openned at %d but not closed", qSingle.pos+1)

	case qDouble.IsOpen():
		return nil, fmt.Errorf("Double quote, `\"`, openned at %d but not closed", qDouble.pos+1)

	case qBrace.IsOpen():
		return nil, fmt.Errorf("Quote brace, `(`, openned at %d but not closed", qBrace.pos[qBrace.len]+1)

	case square.IsOpen():
		return nil, fmt.Errorf("Square brace, `(`, openned at %d but not closed", square.pos[square.len]+1)

	case curly.IsOpen():
		return nil, fmt.Errorf("Curly brace, `(`, openned at %d but not closed", curly.pos[curly.len]+1)

	default:
		return objects.nest[0].value, nil
	}
}
