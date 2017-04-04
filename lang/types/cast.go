package types

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	ErrConversionFailed  = "Conversion failed. No reason given. Please review shell source code for point of failure."
	ErrDataTypeDefaulted = "Unexpected or unknown shell data type."
	ErrUnexpectedGoType  = "Unexpected Go type."
)

func ConvertGoType(v interface{}, dataType string) (interface{}, error) {
	switch v.(type) {
	case nil:
		switch dataType {
		case Integer, Float:
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
		case Float:
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
		case Xml:
			return nil, errors.New(ErrConversionFailed)
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
		case Float:
			return v.(float64), nil
		case Boolean:
			if v.(int) == 0 {
				return true, nil
			}
			return false, nil
		case CodeBlock:
			return fmt.Sprintf("out: %f", v), nil
		case String:
			return fmt.Sprintf("%f", v), nil
		case Json:
			return fmt.Sprintf(`{"Value": %d;}`, v), nil
		case Xml:
			return nil, errors.New(ErrConversionFailed)
		case Null:
			return "", nil
		default:
			return nil, errors.New(ErrDataTypeDefaulted)
		}

	case bool:
		switch dataType {
		case Generic:
			return v, nil
		case Integer, Float:
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
		case Xml:
			return nil, errors.New(ErrConversionFailed)
		case Null:
			return "", nil
		default:
			return nil, errors.New(ErrDataTypeDefaulted)
		}

	case string:
		switch dataType {
		case Generic:
			return v, nil
		case Integer:
			return strconv.Itoa(v.(int)), nil
		case Float:
			return strconv.ParseFloat(v.(string), 64)
		case Boolean:
			return IsTrue([]byte(v.(string)), 0), nil
		case CodeBlock:
			if v.(string)[0] == '{' && v.(string)[len(v.(string))-1] == '}' {
				return v.(string)[1 : len(v.(string))-1], nil
			}
			return "out: '" + v.(string) + "'", nil
		case String, Json:
			return v, nil
		//case Json:
		//	return fmt.Sprintf(`{"Value": "%s";}`, v), nil
		case Xml:
			return nil, errors.New(ErrConversionFailed)
		case Null:
			return "", nil
		default:
			return nil, errors.New(ErrDataTypeDefaulted)
		}

	}

	return nil, errors.New(ErrUnexpectedGoType)
}
