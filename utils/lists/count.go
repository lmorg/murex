package lists

import (
	"fmt"
	"strconv"

	"github.com/lmorg/murex/lang/types"
)

func Count(slice interface{}) (map[string]int, error) {
	switch v := slice.(type) {
	case []int:
		return countInt(v), nil
	case []float64:
		return countFloat64(v), nil
	case []string:
		return countString(v), nil
	case []bool:
		return countBool(v), nil
	case []interface{}:
		return countInterface(v)

	default:
		return make(map[string]int), fmt.Errorf("data type '%T' not supported in lists.Count(). Please report this at https://github.com/lmorg/murex/issues", v)
	}
}

func countInt(s []int) map[string]int {
	m := make(map[string]int)
	for _, i := range s {
		m[strconv.Itoa(i)]++
	}

	return m
}

func countFloat64(s []float64) map[string]int {
	m := make(map[string]int)
	for _, f := range s {
		m[types.FloatToString(f)]++
	}

	return m
}

func countString(s []string) map[string]int {
	m := make(map[string]int)
	for i := range s {
		m[s[i]]++
	}

	return m
}

func countBool(s []bool) map[string]int {
	m := make(map[string]int)
	for i := range s {
		if s[i] {
			m[types.TrueString]++
		} else {
			m[types.FalseString]++
		}
	}

	return m
}

func countInterface(s []interface{}) (map[string]int, error) {
	m := make(map[string]int)
	for i := range s {
		v, err := types.ConvertGoType(s[i], types.String)
		if err != nil {
			return make(map[string]int), err
		}
		m[v.(string)]++
	}

	return m, nil
}
