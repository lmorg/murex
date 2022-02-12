package alter

import (
	"fmt"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/utils/json"
)

const errCannotUnmarshalNewArray = "cannot unmarshal new array: %s"

func mergeArray(v interface{}, new *string) (ret interface{}, err error) {
	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("cannot merge new array into old. Likely this is because of a data-type mismatch: %s", r)
			}
		}()
	}

	switch v := v.(type) {
	case []string:
		var newV []string
		err = json.Unmarshal([]byte(*new), &newV)
		if err != nil {
			return v, fmt.Errorf(errCannotUnmarshalNewArray, err)
		}
		ret = append(v, newV...)

	case []float64:
		var newV []float64
		err = json.Unmarshal([]byte(*new), &newV)
		if err != nil {
			return v, fmt.Errorf(errCannotUnmarshalNewArray, err)
		}
		ret = append(v, newV...)

	case []int:
		var newV []int
		err = json.Unmarshal([]byte(*new), &newV)
		if err != nil {
			return v, fmt.Errorf(errCannotUnmarshalNewArray, err)
		}
		ret = append(v, newV...)

	case []bool:
		var newV []bool
		err = json.Unmarshal([]byte(*new), &newV)
		if err != nil {
			return v, fmt.Errorf(errCannotUnmarshalNewArray, err)
		}
		ret = append(v, newV...)

	case []interface{}:
		var newV []interface{}
		err = json.Unmarshal([]byte(*new), &newV)
		if err != nil {
			return v, fmt.Errorf(errCannotUnmarshalNewArray, err)
		}
		ret = append(v, newV...)

	default:
		return v, fmt.Errorf("path either points to an object that's not an array or no condition has been made for %T", v)
	}

	return
}

func mergeMap(v interface{}, new *string) (ret interface{}, err error) {
	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("cannot merge new map into old. Likely this is because of a data-type mismatch: %s", r)
			}
		}()
	}

	var newV interface{}
	err = json.Unmarshal([]byte(*new), &newV)
	if err != nil {
		return v, fmt.Errorf("cannot unmarshal new map: %s", err)
	}

	switch v.(type) {

	//interface

	case map[string]interface{}:
		ret = v
		for key, val := range newV.(map[string]interface{}) {
			ret.(map[string]interface{})[key] = val
		}

	case map[interface{}]interface{}:
		ret = v
		for key, val := range newV.(map[interface{}]interface{}) {
			ret.(map[interface{}]interface{})[key] = val
		}

		// string

	case map[string]string:
		ret = v
		for key, val := range newV.(map[string]interface{}) {
			ret.(map[string]interface{})[key] = val
		}

	case map[interface{}]string:
		ret = v
		for key, val := range newV.(map[interface{}]interface{}) {
			ret.(map[interface{}]interface{})[key] = val
		}

		// integer

	case map[string]int:
		ret = v
		for key, val := range newV.(map[string]interface{}) {
			ret.(map[string]interface{})[key] = val
		}

	case map[interface{}]int:
		ret = v
		for key, val := range newV.(map[interface{}]interface{}) {
			ret.(map[interface{}]interface{})[key] = val
		}

		// float

	case map[string]float64:
		ret = v
		for key, val := range newV.(map[string]interface{}) {
			ret.(map[string]interface{})[key] = val
		}

	case map[interface{}]float64:
		ret = v
		for key, val := range newV.(map[interface{}]interface{}) {
			ret.(map[interface{}]interface{})[key] = val
		}

		// bool

	case map[string]bool:
		ret = v
		for key, val := range newV.(map[string]interface{}) {
			ret.(map[string]interface{})[key] = val
		}

	case map[interface{}]bool:
		ret = v
		for key, val := range newV.(map[interface{}]interface{}) {
			ret.(map[interface{}]interface{})[key] = val
		}

	default:
		return v, fmt.Errorf("path either points to an object that's not an map or no condition has been made for %T", v)
	}

	return
}
