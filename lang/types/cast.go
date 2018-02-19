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
	//debug.Log("ConvertGoType:", fmt.Sprintf("%t %s %v", v, dataType, v))

	switch v.(type) {
	case nil:
		return goNilRecast(dataType)

	case int:
		return goIntegerRecast(v.(int), dataType)

	//case float32:
	//	return goFloatRecast(float64(v.(float64)), dataType)

	case float64:
		return goFloatRecast(v.(float64), dataType)

	case bool:
		return goBooleanRecast(v.(bool), dataType)

	case string:
		return goStringRecast(v.(string), dataType)

	case []byte:
		return goStringRecast(string(v.([]byte)), dataType)

	case []rune:
		return goStringRecast(string(v.([]rune)), dataType)

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

func goNilRecast(dataType string) (interface{}, error) {
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
}

func goIntegerRecast(v int, dataType string) (interface{}, error) {
	switch dataType {
	case Generic:
		return v, nil

	case Integer:
		return v, nil

	case Float, Number:
		return float64(v), nil

	case Boolean:
		if v == 0 {
			return true, nil
		}
		return false, nil

	case CodeBlock:
		return fmt.Sprintf("out: %d", v), nil

	case String:
		return strconv.Itoa(v), nil

	case Json:
		return fmt.Sprintf(`{"Value": %d;}`, v), nil

	case Null:
		return "", nil

	default:
		return nil, errors.New(ErrDataTypeDefaulted)
	}
}

func goFloatRecast(v float64, dataType string) (interface{}, error) {
	switch dataType {
	case Generic:
		return v, nil

	case Integer:
		return int(v), nil

	case Float, Number:
		return v, nil

	case Boolean:
		if v == 0 {
			return true, nil
		}
		return false, nil

	case CodeBlock:
		return "out: " + FloatToString(v), nil

	case String:
		return FloatToString(v), nil

	case Json:
		return fmt.Sprintf(`{"Value": %s;}`, FloatToString(v)), nil

	case Null:
		return "", nil

	default:
		return nil, errors.New(ErrDataTypeDefaulted)
	}
}

func goBooleanRecast(v bool, dataType string) (interface{}, error) {
	switch dataType {
	case Generic:
		return v, nil

	case Integer, Float, Number:
		if v == true {
			return 0, nil
		}
		return 1, nil

	case Boolean:
		return v, nil

	case CodeBlock:
		if v == true {
			return "true", nil
		}
		return "false", nil

	case String:
		if v == true {
			return string(TrueByte), nil
		}
		return string(FalseByte), nil

	case Json:
		if v == true {
			return `{"Value": true;}`, nil
		}
		return `{"Value": false;}`, nil

	case Null:
		return "", nil

	default:
		return nil, errors.New(ErrDataTypeDefaulted)
	}
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

/*func goBytesRecast(v []byte, dataType string) (interface{}, error) {
	str := string(v)
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
		return IsTrue(v, 0), nil

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
}

func goRunesRecast(v []rune, dataType string) (interface{}, error) {
	str := string(v)
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
}*/

// FloatToString convert a Float64 (what murex numbers are stored as) into a string. Typically for outputting to Stdout/Stderr.
func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
