package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lmorg/murex/utils"
)

const (
	// ErrDataTypeDefaulted is returned if the murex data type is unknown
	ErrDataTypeDefaulted = "Unexpected or unknown murex data type."

	// ErrUnexpectedGoType is returned if the Go data type is unhandled
	ErrUnexpectedGoType = "Unexpected Go type."
)

// ConvertGoType converts a Go lang variable into a murex variable
func ConvertGoType(v interface{}, dataType string) (interface{}, error) {
	// First switch:  input data type
	// Second switch: output data type

	//debug.Log("ConvertGoType:", fmt.Sprintf("%t %s %v", v, dataType, v))

	switch v.(type) {
	case nil:
		switch dataType {
		case Integer, Float, Number:
			return 0, nil
		case Boolean:
			return false, nil
		case CodeBlock:
			return "{}", nil
		default:
			return "", nil
		}

	case int:
		switch dataType {
		case Generic:
			return v, nil
		case Integer:
			return v.(int), nil
		case Float, Number:
			return float64(v.(int)), nil
		case Boolean:
			if v.(int) == 0 {
				return true, nil
			}
			return false, nil
		case CodeBlock:
			return fmt.Sprintf("out: %d", v), nil
		case String:
			return strconv.Itoa(v.(int)), nil
		case Json:
			return fmt.Sprintf(`{"Value": %d;}`, v), nil
		case Null:
			return "", nil
		default:
			return nil, errors.New(ErrDataTypeDefaulted)
		}

	case float64:
		switch dataType {
		case Generic:
			return v, nil
		case Integer:
			return int(v.(float64)), nil
		case Float, Number:
			return v.(float64), nil
		case Boolean:
			if v.(float64) == 0 {
				return true, nil
			}
			return false, nil
		case CodeBlock:
			return "out: " + FloatToString(v.(float64)), nil
		case String:
			return FloatToString(v.(float64)), nil
		case Json:
			return fmt.Sprintf(`{"Value": %s;}`, FloatToString(v.(float64))), nil
		case Null:
			return "", nil
		default:
			return nil, errors.New(ErrDataTypeDefaulted)
		}

	case bool:
		switch dataType {
		case Generic:
			return v, nil
		case Integer, Float, Number:
			if v.(bool) == true {
				return 0, nil
			}
			return 1, nil
		case Boolean:
			return v, nil
		case CodeBlock:
			if v.(bool) == true {
				return "true", nil
			}
			return "false", nil
		case String:
			if v.(bool) == true {
				return string(TrueByte), nil
			}
			return string(FalseByte), nil
		case Json:
			if v.(bool) == true {
				return `{"Value": true;}`, nil
			}
			return `{"Value": false;}`, nil
		case Null:
			return "", nil
		default:
			return nil, errors.New(ErrDataTypeDefaulted)
		}

	case string:
		return goStringRecast(v.(string), dataType)

	case []byte:
		str := string(v.([]byte))
		switch dataType {
		case Generic:
			return str, nil
		case Integer:
			if str == "" {
				str = "0"
			}
			return strconv.Atoi(strings.TrimSpace(str))
		case Float, Number:
			if str == "" {
				str = "0"
			}
			return strconv.ParseFloat(str, 64)
		case Boolean:
			return IsTrue(v.([]byte), 0), nil
		case CodeBlock:
			if str[0] == '{' && str[len(str)-1] == '}' {
				return str[1 : len(str)-1], nil
			}
			return "out: '" + str + "'", nil //errors.New("Not a valid code block: `" + str + "`")
		case String, Json:
			return v, nil
		//case Json:
		//	return fmt.Sprintf(`{"Value": "%s";}`, v), nil
		case Null:
			return "", nil
		default:
			return nil, errors.New(ErrDataTypeDefaulted)
		}

	case []rune:
		str := string(v.([]byte))
		switch dataType {
		case Generic:
			return str, nil
		case Integer:
			if str == "" {
				str = "0"
			}
			return strconv.Atoi(strings.TrimSpace(str))
		case Float, Number:
			if str == "" {
				str = "0"
			}
			return strconv.ParseFloat(str, 64)
		case Boolean:
			return IsTrue([]byte(str), 0), nil
		case CodeBlock:
			if str[0] == '{' && str[len(str)-1] == '}' {
				return str[1 : len(str)-1], nil
			}
			return "out: '" + str + "'", nil //errors.New("Not a valid code block: `" + str + "`")
		case String, Json:
			return v, nil
		//case Json:
		//	return fmt.Sprintf(`{"Value": "%s";}`, v), nil
		case Null:
			return "", nil
		default:
			return nil, errors.New(ErrDataTypeDefaulted)
		}

	default:
		switch dataType {
		//case Generic, String, Integer, Float, Number, Boolean, CodeBlock, Null:
		//	return nil, errors.New(ErrUnexpectedGoType)
		case String, Json:
			b, err := utils.JsonMarshal(v, false)
			return string(b), err
		default:
			return nil, errors.New(ErrUnexpectedGoType)
		}
	}

	return nil, errors.New(ErrUnexpectedGoType)
}

func goStringRecast(v string, dataType string) (interface{}, error) {
	switch dataType {
	case Generic:
		return v, nil

	case Integer:
		if v == "" {
			v = "0"
		}
		return strconv.Atoi(strings.TrimSpace(v))

	case Float, Number:
		if v == "" {
			v = "0"
		}
		return strconv.ParseFloat(v, 64)

	case Boolean:
		return IsTrue([]byte(v), 0), nil

	case CodeBlock:
		if v[0] == '{' && v[len(v)-1] == '}' {
			return v[1 : len(v)-1], nil
		}
		return "out: '" + v + "'", nil //errors.New("Not a valid code block: `" + v.(string) + "`")

	case String, Json:
		return v, nil
	//case Json:
	//	return fmt.Sprintf(`{"Value": "%s";}`, v), nil

	case Null:
		return "", nil

	default:
		return nil, errors.New(ErrDataTypeDefaulted)
	}
}

// FloatToString convert a Float64 (what murex numbers are stored as) into a string. Typically for outputting to Stdout/Stderr.
func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
