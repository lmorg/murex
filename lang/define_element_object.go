package lang

import (
	"fmt"
	"strconv"
	"strings"
)

func ElementLookup(v interface{}, path string) (interface{}, error) {
	var err error

	if len(path) < 2 {
		return nil, fmt.Errorf("invalid path for element lookup: `%s` is too short", path)
	}

	pathSplit := strings.Split(path, path[0:1])
	obj := v

	for i := 1; i < len(pathSplit); i++ {
		if len(pathSplit[i]) == 0 {
			if i == len(pathSplit)-1 {
				break
			} else {
				return nil, fmt.Errorf("path element %d is a zero length string: '%s'", i-1, strings.Join(pathSplit, "/"))
			}
		}

		obj, err = elementRecursiveLookup(pathSplit, i, obj)
		if err != nil {
			return nil, err
		}
	}

	return obj, nil
}

func elementRecursiveLookup(path []string, i int, obj interface{}) (interface{}, error) {
	switch v := obj.(type) {
	case []string:
		i, err := isValidElementIndex(path[i], len(v))
		if err != nil {
			return nil, err
		}
		return v[i], nil

	case []interface{}:
		i, err := isValidElementIndex(path[i], len(v))
		if err != nil {
			return nil, err
		}
		return v[i], nil

	case []map[string]interface{}:
		i, err := isValidElementIndex(path[i], len(v))
		if err != nil {
			return nil, err
		}
		return v[i], nil

	case map[string]string:
		switch {
		case v[path[i]] != "":
			return v[path[i]], nil
		case v[strings.Title(path[i])] != "":
			return v[strings.Title(path[i])], nil
		case v[strings.ToLower(path[i])] != "":
			return v[strings.ToLower(path[i])], nil
		case v[strings.ToUpper(path[i])] != "":
			return v[strings.ToUpper(path[i])], nil
			//case v[strings.ToTitle(params[i])] != nil:
			//	return v[strings.ToTitle(path[i])], nil
		default:
			return "", nil
		}

	case map[string]interface{}:
		switch {
		case v[path[i]] != nil:
			return v[path[i]], nil
		case v[strings.Title(path[i])] != nil:
			return v[strings.Title(path[i])], nil
		case v[strings.ToLower(path[i])] != nil:
			return v[strings.ToLower(path[i])], nil
		case v[strings.ToUpper(path[i])] != nil:
			return v[strings.ToUpper(path[i])], nil
			//case v[strings.ToTitle(params[i])] != nil:
			//	return v[strings.ToTitle(path[i])], nil
		default:
			return nil, fmt.Errorf("key '%s' not found", path[i])
		}

	case map[interface{}]interface{}:
		switch {
		case v[path[i]] != nil:
			return v[path[i]], nil
		case v[strings.Title(path[i])] != nil:
			return v[strings.Title(path[i])], nil
		case v[strings.ToLower(path[i])] != nil:
			return v[strings.ToLower(path[i])], nil
		case v[strings.ToUpper(path[i])] != nil:
			return v[strings.ToUpper(path[i])], nil
			//case v[strings.ToTitle(params[i])] != nil:
			//	return v[strings.ToTitle(path[i])], nil
		default:
			return nil, fmt.Errorf("key '%s' not found", path[i])
		}

	case string, int, float64, bool, nil, []byte, []rune:
		return nil, fmt.Errorf("primitives like %T cannot be split to return property '%s'", v, path[i])

	case MxInterface:
		return elementRecursiveLookup(path, i, v.GetValue())

	default:
		return nil, fmt.Errorf("murex doesn't know how to lookup `%T` (please file a bug with on the murex Github page: https://lmorg/murex)", v)
	}
}

func isValidElementIndex(key string, length int) (int, error) {
	i, err := strconv.Atoi(key)
	if err != nil {
		return 0, fmt.Errorf("element is an array however supplied key, '%s', is not an integer", key)
	}

	if i < 0 {
		i += length
		if i < 0 {
			return 0, fmt.Errorf("element is an array however key (%s -> %d) is greater than the length (%d)", key, i, length)
		}
		return i, nil
	}

	if i >= length {
		return 0, fmt.Errorf("element is an array however key (%s) is greater than the length (%d)", key, length)
	}

	return i, nil
}
