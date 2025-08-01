package alter

import (
	"fmt"
	"reflect"

	"github.com/lmorg/murex/debug"
)

func mergeArray(v any, new *any) (ret any, err error) {
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

	case []any:
		ret = append(v, (*new).([]any)...)

	default:
		return v, fmt.Errorf("path either points to an object that's not an array or no condition has been made for %T", v)
	}

	return
}

func mergeMap(v any, new *any) (ret any, err error) {
	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("cannot merge new map into old. Likely this is because of a data-type mismatch: %s", r)
			}
		}()
	}

	switch v.(type) {

	//interface

	case map[string]any:
		ret = v
		for key, val := range (*new).(map[string]any) {
			switch t := val.(type) {
			case string, int, float64, bool, nil:
				ret.(map[string]any)[key] = t
			default:
				mapVal, ok := ret.(map[string]any)[key]
				if !ok {
					ret.(map[string]any)[key] = t
					return
				}
				oldKind := reflect.TypeOf(mapVal).Kind()
				newKind := reflect.TypeOf(val).Kind()
				switch {
				case oldKind != newKind:
					ret.(map[string]any)[key] = t
				case newKind == reflect.Slice:
					ret.(map[string]any)[key], err = mergeArray(ret.(map[string]any)[key], &val)
				case newKind == reflect.Map:
					ret.(map[string]any)[key], err = mergeMap(ret.(map[string]any)[key], &val)
				default:
					// possibly not an object so lets just overwrite...
					ret.(map[string]any)[key] = t
				}
			}
		}

	case map[any]any:
		ret = v
		for key, val := range (*new).(map[any]any) {
			ret.(map[any]any)[key] = val
		}

		// string

	case map[string]string:
		ret = v
		for key, val := range (*new).(map[string]any) {
			ret.(map[string]any)[key] = val
		}

	case map[any]string:
		ret = v
		for key, val := range (*new).(map[any]any) {
			ret.(map[any]any)[key] = val
		}

		// integer

	case map[string]int:
		ret = v
		for key, val := range (*new).(map[string]any) {
			ret.(map[string]any)[key] = val
		}

	case map[any]int:
		ret = v
		for key, val := range (*new).(map[any]any) {
			ret.(map[any]any)[key] = val
		}

		// float

	case map[string]float64:
		ret = v
		for key, val := range (*new).(map[string]any) {
			ret.(map[string]any)[key] = val
		}

	case map[any]float64:
		ret = v
		for key, val := range (*new).(map[any]any) {
			ret.(map[any]any)[key] = val
		}

		// bool

	case map[string]bool:
		ret = v
		for key, val := range (*new).(map[string]any) {
			ret.(map[string]any)[key] = val
		}

	case map[any]bool:
		ret = v
		for key, val := range (*new).(map[any]any) {
			ret.(map[any]any)[key] = val
		}

	default:
		return v, fmt.Errorf("path either points to an object that's not an map or no condition has been made for %T", v)
	}

	return
}
