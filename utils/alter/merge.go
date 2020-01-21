package alter

import (
	"fmt"

	"github.com/lmorg/murex/utils/json"
)

func mergeArray(v interface{}, new *string) (ret interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Cannot merge new array into old. Likely this is because of a data-type mismatch: %s", r)
		}
	}()

	var newV interface{}
	err = json.Unmarshal([]byte(*new), &newV)
	if err != nil {
		return v, fmt.Errorf("Cannot unmarshal new array: %s", err)
	}

	switch v.(type) {
	case []string:
		ret = append(v.([]string), newV.([]string)...)

	case []int:
		ret = append(v.([]int), newV.([]int)...)

	case []bool:
		ret = append(v.([]bool), newV.([]bool)...)

	case []interface{}:
		ret = append(v.([]interface{}), newV.([]interface{})...)

	default:
		return v, fmt.Errorf("Path either points to an object that's not an array or no condition has been made for %T", v)
	}

	return
}

func mergeMap(v interface{}, new *string) (ret interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Cannot merge new map into old. Likely this is because of a data-type mismatch: %s", r)
		}
	}()

	var newV interface{}
	err = json.Unmarshal([]byte(*new), &newV)
	if err != nil {
		return v, fmt.Errorf("Cannot unmarshal new map: %s", err)
	}

	switch v.(type) {
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

	default:
		return v, fmt.Errorf("Path either points to an object that's not an map or no condition has been made for %T", v)
	}

	return
}
