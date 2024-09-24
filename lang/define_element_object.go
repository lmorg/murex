package lang

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ELEMENT_META_ROOT    = "__MUREX_ELEMENT_META_ROOT_KEY__"
	ELEMENT_META_ELEMENT = "__MUREX_ELEMENT_META_ELEMENT_KEY__"
)

func ElementLookup(v any, path string, dataType string) (any, error) {
	if len(path) < 2 {
		return nil, fmt.Errorf("invalid path for element lookup: `%s` is too short", path)
	}

	if dataType == "xml" {
		m, ok := v.(map[string]any)
		if ok && len(m) == 1 {
			for root := range m {
				path = path[0:1] + root + path
			}
		}
	}

	pathSplit := strings.Split(path, path[0:1])
	obj := v
	var err error

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

	if dataType == "xml" {
		l := len(pathSplit)
		switch t := obj.(type) {
		case map[string]any:
			t[ELEMENT_META_ROOT] = pathSplit[l-1]
			return t, nil

		case []string:
			var rootKey string
			if len(pathSplit) > 1 {
				rootKey = pathSplit[l-2]
			} else {
				rootKey = "xml"
			}
			t = append([]string{
				ELEMENT_META_ROOT + rootKey,
				ELEMENT_META_ELEMENT + pathSplit[l-1]},
				t...)
			return t, nil

		case []any:
			var rootKey string
			if len(pathSplit) > 1 {
				rootKey = pathSplit[l-2]
			} else {
				rootKey = "xml"
			}
			t = append([]any{
				ELEMENT_META_ROOT + rootKey,
				ELEMENT_META_ELEMENT + pathSplit[l-1]},
				t...)
			return t, nil

		}
	}

	return obj, nil
}

func elementRecursiveLookup(path []string, i int, obj any) (any, error) {
	switch v := obj.(type) {
	case []string:
		i, err := isValidElementIndex(path[i], len(v))
		if err != nil {
			return nil, err
		}
		return v[i], nil

	case []any:
		i, err := isValidElementIndex(path[i], len(v))
		if err != nil {
			return nil, err
		}
		return v[i], nil

	case []map[string]any:
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

	case map[string]any:
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

	case map[any]any:
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
		return 0, fmt.Errorf("negative keys are not allowed for arrays: %s", key)
	}

	if i >= length {
		return 0, fmt.Errorf("element is an array however key is greater than the length: %s", key)
	}

	return i, nil
}
