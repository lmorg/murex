package types

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/lmorg/murex/utils/json"
)

const (
	// ErrDataTypeDefaulted is returned if the murex data type is unknown
	ErrDataTypeDefaulted = "Unexpected or unknown murex data type"

	// ErrUnexpectedGoType is returned if the Go data type is unhandled
	ErrUnexpectedGoType = "Unexpected Go type"

	// ErrCannotConvertGoType is returned if the Go data type cannot be converted
	// to the murex data type (eg there is no numeric data in a string of characters)
	ErrCannotConvertGoType = "Cannot convert Go type into murex data type (eg no numeric data in a string)"
)

var (
	rxFirstInt   *regexp.Regexp = regexp.MustCompile(`([0-9]+)`)
	rxFirstFloat *regexp.Regexp = regexp.MustCompile(`([0-9]+)(\.[0-9]+|)`)
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
		return goDefaultRecast(v, dataType)
	}

	//return nil, errors.New(ErrUnexpectedGoType)
}

func goNilRecast(dataType string) (interface{}, error) {
	switch dataType {
	case Integer:
		return 0, nil

	case Float, Number:
		return float64(0), nil

	case Boolean:
		return false, nil

	case CodeBlock, Json, JsonLines:
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
			return false, nil
		}
		return true, nil

	//case CodeBlock:
	//	return fmt.Sprintf("out: %d", v), nil

	case String:
		return strconv.Itoa(v), nil

	//case Json, JsonLines:
	//	return fmt.Sprintf(`{ "Value": %d }`, v), nil

	case Null:
		return "", nil

	default:
		//	return nil, errors.New(ErrDataTypeDefaulted)
		return strconv.Itoa(v), nil
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
			return false, nil
		}
		return true, nil

	//case CodeBlock:
	//	return "out: " + FloatToString(v), nil

	case String:
		return FloatToString(v), nil

	//case Json, JsonLines:
	//	return fmt.Sprintf(`{ "Value": %s }`, FloatToString(v)), nil

	case Null:
		return "", nil

	default:
		//return nil, errors.New(ErrDataTypeDefaulted)
		return FloatToString(v), nil
	}
}

func goBooleanRecast(v bool, dataType string) (interface{}, error) {
	switch dataType {
	case Generic:
		return v, nil

	case Integer:
		if v {
			return 1, nil
		}
		return 0, nil

	case Float, Number:
		if v {
			return float64(1), nil
		}
		return float64(0), nil

	case Boolean:
		return v, nil

	case CodeBlock:
		if v {
			return "true", nil
		}
		return "false", nil

	case String:
		if v {
			return string(TrueByte), nil
		}
		return string(FalseByte), nil

	/*case Json, JsonLines:
	if v {
		return `{ "Value": true }`, nil
	}
	return `{ "Value": false }`, nil*/

	case Null:
		return "", nil

	default:
		//return nil, errors.New(ErrDataTypeDefaulted)
		if v {
			return string(TrueByte), nil
		}
		return string(FalseByte), nil
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
		//return strconv.Atoi(strings.TrimSpace(v))
		f, err := strconv.ParseFloat(v, 64)
		return int(f), err

	case Float, Number:
		if v == "" {
			v = "0"
		}
		return strconv.ParseFloat(v, 64)

	case Boolean:
		return IsTrue([]byte(v), 0), nil

	case CodeBlock:
		if len(v) > 1 && v[0] == '{' && v[len(v)-1] == '}' {
			return v[1 : len(v)-1], nil
		}
		//errors.New("Not a valid code block: `" + v.(string) + "`")
		return "out: '" + v + "'", nil

	case String:
		return v, nil

	//case Json, JsonLines:
	//	return fmt.Sprintf(`{ "Value": %s }`, strconv.Quote(v)), nil

	case Null:
		return "", nil

	default:
		//return nil, errors.New(ErrDataTypeDefaulted)
		return v, nil
	}
}

func goDefaultRecast(v interface{}, dataType string) (interface{}, error) {
	switch dataType {
	case Integer:
		s := fmt.Sprint(v)
		i := rxFirstInt.FindStringSubmatch(s)
		if len(i) > 0 {
			return i[0], nil
		}
		return 0, errors.New(ErrCannotConvertGoType)

	case Float, Number:
		s := fmt.Sprint(v)
		f := rxFirstFloat.FindStringSubmatch(s)
		if len(f) > 0 {
			return f[0], nil
		}
		return 0, errors.New(ErrCannotConvertGoType)

	case Boolean:
		s := fmt.Sprint(v)
		if s == "{}" || s == "[]" || s == "[[]]" || s == "" {
			return false, nil
		}
		if !IsTrue([]byte(s), 0) {
			return false, nil
		}
		return true, nil

	case Null:
		return nil, nil

	case String, Json, JsonLines:
		b, err := json.Marshal(v, false)
		return string(b), err

	default:
		return nil, errors.New(ErrUnexpectedGoType)
	}
}

// FloatToString convert a Float64 (what murex numbers are stored as) into a string. Typically for outputting to Stdout/Stderr.
func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
