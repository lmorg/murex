package alter

import (
	"fmt"

	"github.com/lmorg/murex/debug"
)

func mergeArray(v interface{}, new *interface{}) (ret interface{}, err error) {
	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("cannot merge new array into old. Likely this is because of a data-type mismatch: %s", r)
			}
		}()
	}

	switch v := v.(type) {
	case []string:
		ret = append(v, (*new).([]string)...)

	case []float64:
		ret = append(v, (*new).([]float64)...)

	case []int:
		ret = append(v, (*new).([]int)...)

	case []bool:
		ret = append(v, (*new).([]bool)...)

	case []interface{}:
		ret = append(v, (*new).([]interface{})...)

	default:
		return v, fmt.Errorf("path either points to an object that's not an array or no condition has been made for %T", v)
	}

	return
}

func mergeMap(v interface{}, new *interface{}) (ret interface{}, err error) {
	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("cannot merge new map into old. Likely this is because of a data-type mismatch: %s", r)
			}
		}()
	}

	switch v.(type) {

	//interface

	case map[string]interface{}:
		ret = v
		for key, val := range (*new).(map[string]interface{}) {
			ret.(map[string]interface{})[key] = val
		}

	case map[interface{}]interface{}:
		ret = v
		for key, val := range (*new).(map[interface{}]interface{}) {
			ret.(map[interface{}]interface{})[key] = val
		}

		// string

	case map[string]string:
		ret = v
		for key, val := range (*new).(map[string]interface{}) {
			ret.(map[string]interface{})[key] = val
		}

	case map[interface{}]string:
		ret = v
		for key, val := range (*new).(map[interface{}]interface{}) {
			ret.(map[interface{}]interface{})[key] = val
		}

		// integer

	case map[string]int:
		ret = v
		for key, val := range (*new).(map[string]interface{}) {
			ret.(map[string]interface{})[key] = val
		}

	case map[interface{}]int:
		ret = v
		for key, val := range (*new).(map[interface{}]interface{}) {
			ret.(map[interface{}]interface{})[key] = val
		}

		// float

	case map[string]float64:
		ret = v
		for key, val := range (*new).(map[string]interface{}) {
			ret.(map[string]interface{})[key] = val
		}

	case map[interface{}]float64:
		ret = v
		for key, val := range (*new).(map[interface{}]interface{}) {
			ret.(map[interface{}]interface{})[key] = val
		}

		// bool

	case map[string]bool:
		ret = v
		for key, val := range (*new).(map[string]interface{}) {
			ret.(map[string]interface{})[key] = val
		}

	case map[interface{}]bool:
		ret = v
		for key, val := range (*new).(map[interface{}]interface{}) {
			ret.(map[interface{}]interface{})[key] = val
		}

	default:
		return v, fmt.Errorf("path either points to an object that's not an map or no condition has been made for %T", v)
	}

	return
}
