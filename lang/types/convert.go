package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/utils/json"
)

const (
	// ErrDataTypeDefaulted is returned if the murex data type is unknown
	ErrDataTypeDefaulted = "unexpected or unknown murex data type"

	// ErrUnexpectedGoType is returned if the Go data type is unhandled
	ErrUnexpectedGoType = "unexpected Go type"

	errCannotConvertInterfaceToNumeric = "cannot convert '%T' into the numeric type '%s'"
)

// ConvertGoType converts a Go lang variable into a murex variable
func ConvertGoType(v any, dataType string) (any, error) {
	switch t := v.(type) {
	case nil:
		return goNilRecast(dataType)

	case int:
		return goIntegerRecast(t, dataType)

	case float64:
		return goFloatRecast(t, dataType)

	case bool:
		return goBooleanRecast(t, dataType)

	case string:
		return goStringRecast(t, dataType)

	case []byte:
		return goStringRecast(string(t), dataType)

	case []rune:
		return goStringRecast(string(t), dataType)

	default:
		return goDefaultRecast(v, dataType)
	}
}

func goNilRecast(dataType string) (any, error) {
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

func goIntegerRecast(v int, dataType string) (any, error) {
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

	case String:
		return strconv.Itoa(v), nil

	case Null:
		return "", nil

	default:
		return strconv.Itoa(v), nil
	}
}

func goFloatRecast(v float64, dataType string) (any, error) {
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

	case String:
		return FloatToString(v), nil

	case Null:
		return "", nil

	default:
		return FloatToString(v), nil
	}
}

func goBooleanRecast(v bool, dataType string) (any, error) {
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

	case Null:
		return "", nil

	default:
		if v {
			return string(TrueByte), nil
		}
		return string(FalseByte), nil
	}
}

func goStringRecast(v string, dataType string) (any, error) {
	switch dataType {
	case Generic:
		return v, nil

	case Integer:
		v = strings.TrimSpace(v)
		if v == "" {
			v = "0"
		}

		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return int(f), fmt.Errorf("cannot convert '%s' to an integer: %s", v, err.Error())
		}
		return int(f), nil

	case Float, Number:
		v = strings.TrimSpace(v)
		if v == "" {
			v = "0"
		}

		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return f, fmt.Errorf("cannot convert '%s' to a floating point number: %s", v, err.Error())
		}
		return f, nil

	case Boolean:
		return IsTrue([]byte(v), 0), nil

	case CodeBlock:
		if len(v) > 1 && v[0] == '{' && v[len(v)-1] == '}' {
			return v[1 : len(v)-1], nil
		}
		return "out: '" + v + "'", nil

	case String:
		return v, nil

	case Null:
		return "", nil

	default:
		return v, nil
	}
}

func goDefaultRecast(v any, dataType string) (any, error) {
	switch dataType {
	case Generic:
		switch t := v.(type) {
		case []string:
			return strings.Join(t, " "), nil
		case []any:
			a := make([]string, len(t))
			for i := range t {
				a[i] = fmt.Sprint(t[i])
			}
			return strings.Join(a, " "), nil
		case []int:
			a := make([]string, len(t))
			for i := range t {
				a[i] = strconv.Itoa(t[i])
			}
			return strings.Join(a, " "), nil
		case []float64:
			a := make([]string, len(t))
			for i := range t {
				a[i] = FloatToString(t[i])
			}
			return strings.Join(a, " "), nil
		case []bool:
			a := make([]string, len(t))
			for i := range t {
				if t[i] {
					a[i] = TrueString
				} else {
					a[i] = FalseString
				}
			}
			return strings.Join(a, " "), nil
		default:
			b, err := json.Marshal(v, false)
			return string(b), err
		}

	case String:
		switch t := v.(type) {
		case []string:
			return strings.Join(t, "\n"), nil
		case []any:
			a := make([]string, len(t))
			for i := range t {
				a[i] = fmt.Sprint(t[i])
			}
			return strings.Join(a, "\n"), nil
		case []int:
			a := make([]string, len(t))
			for i := range t {
				a[i] = strconv.Itoa(t[i])
			}
			return strings.Join(a, "\n"), nil
		case []float64:
			a := make([]string, len(t))
			for i := range t {
				a[i] = FloatToString(t[i])
			}
			return strings.Join(a, "\n"), nil
		case []bool:
			a := make([]string, len(t))
			for i := range t {
				if t[i] {
					a[i] = TrueString
				} else {
					a[i] = FalseString
				}
			}
			return strings.Join(a, "\n"), nil
		default:
			b, err := json.Marshal(v, false)
			return string(b), err
		}

	case Integer:
		s, err := Flatten(v)
		if err != nil {
			return 0, fmt.Errorf(errCannotConvertInterfaceToNumeric, v, dataType)
		}

		s = strings.TrimSpace(s)
		if s == "" {
			s = "0"
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return int(f), fmt.Errorf("cannot convert '%s' to an integer: %s", s, err.Error())
		}
		return int(f), nil

	case Float, Number:
		s, err := Flatten(v)
		if err != nil {
			return 0, fmt.Errorf(errCannotConvertInterfaceToNumeric, v, dataType)
		}

		s = strings.TrimSpace(s)
		if s == "" {
			s = "0"
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return f, fmt.Errorf("cannot convert '%s' to a floating point number: %s", s, err.Error())
		}
		return f, nil

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

	case Json, JsonLines:
		b, err := json.Marshal(v, false)
		return string(b), err

	default:
		return nil, fmt.Errorf(ErrUnexpectedGoType+": Go '%T', %s '%s'", v, app.Name, dataType)
	}
}

// FloatToString convert a Float64 (what murex numbers are stored as) into a string. Typically for outputting to Stdout/Stderr.
func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
