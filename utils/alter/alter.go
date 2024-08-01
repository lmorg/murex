package alter

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/types"
)

const (
	actionAlter int = iota + 1
	actionMerge
	actionSum
)

// Alter a data structure. Requires a path (pre-split) and new structure as a
// JSON string. A more seasoned developer will see plenty of room for
// optimization however this function was largely thrown together in a "let's
// create something that works first and worry about performance later" kind of
// sense (much like a lot of murex's code base). That being said, I will accept
// any pull requests from other developers wishing to improve this - or other -
// functions. I'm also open to any breaking changes those optimisations might
// bring (at least until the project reaches version 1.0).
func Alter(ctx context.Context, v interface{}, path []string, new interface{}) (interface{}, error) {
	return loop(ctx, v, 0, path, &new, actionAlter)
}

// Merge a data structure; like Alter but merges arrays and maps where possible
func Merge(ctx context.Context, v interface{}, path []string, new interface{}) (interface{}, error) {
	if len(path) == 1 && path[0] == "" {
		path = []string{}
	}
	return loop(ctx, v, 0, path, &new, actionMerge)
}

// Sum a data structure; like Merge but sums values in arrays and maps where
// duplication exists
func Sum(ctx context.Context, v interface{}, path []string, new interface{}) (interface{}, error) {
	if len(path) == 1 && path[0] == "" {
		path = []string{}
	}
	return loop(ctx, v, 0, path, &new, actionSum)
}

var (
	errOverwritePath = errors.New("internal condition: path needs overwriting")
	errInvalidAction = errors.New("missing or invalid action. Please report this to https://github.com/lmorg/murex/issues")
)

const (
	errExpectingAnArrayIndex     = "expecting an array index in path element"
	errNegativeIndexesNotAllowed = "negative indexes not allowed in arrays: path element"
	errIndexGreaterThanArray     = "index greater than length of array in path element"
)

func loop(ctx context.Context, v interface{}, i int, path []string, new *interface{}, action int) (ret interface{}, err error) {
	if !debug.Enabled {
		defer func() {
			r := recover()
			if r != nil {
				err = fmt.Errorf("unhandled error in type conversion: %v", r)
			}
		}()
	}

	select {
	case <-ctx.Done():
		return nil, errors.New("cancelled")
	default:
	}

	switch {
	case i < len(path):
		switch v := v.(type) {
		case []interface{}:
			pathI, err := strconv.Atoi(path[i])
			if err != nil {
				return nil, fmt.Errorf("%s '%s': %s", errExpectingAnArrayIndex, path[i], err)
			}

			if pathI < 0 {
				return nil, fmt.Errorf("%s '%d'", errNegativeIndexesNotAllowed, pathI)
			}

			if pathI >= len(v) {
				return nil, fmt.Errorf("%s '%d' (array length '%d')", errIndexGreaterThanArray, pathI, len(v))
			}

			ret, err = loop(ctx, v[pathI], i+1, path, new, action)
			if err == errOverwritePath {
				v[pathI] = *new

			}
			if err == nil {
				v[pathI] = ret
				ret = v
			}

		case []string:
			pathI, err := strconv.Atoi(path[i])
			if err != nil {
				return nil, fmt.Errorf("%s '%s': %s", errExpectingAnArrayIndex, path[i], err)
			}

			if pathI < 0 {
				return nil, fmt.Errorf("%s '%d'", errNegativeIndexesNotAllowed, pathI)
			}

			if pathI >= len(v) {
				return nil, fmt.Errorf("%s '%d' (array length '%d')", errIndexGreaterThanArray, pathI, len(v))
			}

			ret, err = loop(ctx, v[pathI], i+1, path, new, action)
			if err == errOverwritePath {
				s, err := types.ConvertGoType(*new, types.String)
				if err != nil {
					return nil, err
				}
				v[pathI] = s.(string)

			}
			if err == nil {
				v[pathI] = ret.(string)
				ret = v
			}

		case []int:
			pathI, err := strconv.Atoi(path[i])
			if err != nil {
				return nil, fmt.Errorf("%s '%s': %s", errExpectingAnArrayIndex, path[i], err)
			}

			if pathI < 0 {
				return nil, fmt.Errorf("%s '%d'", errNegativeIndexesNotAllowed, pathI)
			}

			if pathI >= len(v) {
				return nil, fmt.Errorf("%s '%d' (array length '%d')", errIndexGreaterThanArray, pathI, len(v))
			}

			ret, err = loop(ctx, v[pathI], i+1, path, new, action)
			if err == errOverwritePath {
				i, err := types.ConvertGoType(*new, types.Integer)
				if err != nil {
					return nil, err
				}
				v[pathI] = i.(int)

			}
			if err == nil {
				v[pathI] = ret.(int)
				ret = v
			}

		case []float64:
			pathI, err := strconv.Atoi(path[i])
			if err != nil {
				return nil, fmt.Errorf("%s '%s': %s", errExpectingAnArrayIndex, path[i], err)
			}

			if pathI < 0 {
				return nil, fmt.Errorf("%s '%d'", errNegativeIndexesNotAllowed, pathI)
			}

			if pathI >= len(v) {
				return nil, fmt.Errorf("%s '%d' (array length '%d')", errIndexGreaterThanArray, pathI, len(v))
			}

			ret, err = loop(ctx, v[pathI], i+1, path, new, action)
			if err == errOverwritePath {
				f, err := types.ConvertGoType(*new, types.Float)
				if err != nil {
					return nil, err
				}
				v[pathI] = f.(float64)

			}
			if err == nil {
				v[pathI] = ret.(float64)
				ret = v
			}

		case []bool:
			pathI, err := strconv.Atoi(path[i])
			if err != nil {
				return nil, fmt.Errorf("%s '%s': %s", errExpectingAnArrayIndex, path[i], err)
			}

			if pathI < 0 {
				return nil, fmt.Errorf("%s '%d'", errNegativeIndexesNotAllowed, pathI)
			}

			if pathI >= len(v) {
				return nil, fmt.Errorf("%s '%d' (array length '%d')", errIndexGreaterThanArray, pathI, len(v))
			}

			ret, err = loop(ctx, v[pathI], i+1, path, new, action)
			if err == errOverwritePath {
				b, err := types.ConvertGoType(*new, types.Boolean)
				if err != nil {
					return nil, err
				}
				v[pathI] = b.(bool)

			}
			if err == nil {
				v[pathI] = ret.(bool)
				ret = v
			}

		case map[interface{}]interface{}:
			ret, err = loop(ctx, v[path[i]], i+1, path, new, action)
			if err == errOverwritePath {
				v[path[i]] = *new

			}
			if err == nil {
				v[path[i]] = ret
				ret = v
			}

		case map[string]interface{}:
			ret, err = loop(ctx, v[path[i]], i+1, path, new, action)
			if err == errOverwritePath {
				v[path[i]] = *new

			}
			if err == nil {
				v[path[i]] = ret
				ret = v
			}

		case map[interface{}]string:
			ret, err = loop(ctx, v[path[i]], i+1, path, new, action)
			if err == errOverwritePath {
				s, err := types.ConvertGoType(*new, types.String)
				if err != nil {
					return nil, err
				}
				v[path[i]] = s.(string)

			}
			if err == nil {
				v[path[i]] = ret.(string)
				ret = v
			}

		case nil:
			// Let's overwrite part of the path
			return nil, errOverwritePath

		case string, int, float64, bool:
			return nil, fmt.Errorf("unable to alter data structure using that path because one of the path elements is an end of tree (%T) rather than a map. Instead please have the full path you want to add as part of the amend JSON string in `alter`", v)

		default:
			return nil, fmt.Errorf("murex code error: No condition is made for `%T`. Please report this bug to https://github.com/lmorg/murex/issues", v)
		}

	case i == len(path):
		switch v.(type) {
		case string:
			s, err := types.ConvertGoType(*new, types.String)
			if err != nil {
				return nil, err
			}
			ret = s.(string)

		case int:
			i, err := types.ConvertGoType(*new, types.Integer)
			if err != nil {
				return nil, err
			}
			ret = i.(int)

		case float64:
			f, err := types.ConvertGoType(*new, types.Float)
			if err != nil {
				return nil, err
			}
			ret = f.(float64)

		case bool:
			b, err := types.ConvertGoType(*new, types.Boolean)
			if err != nil {
				return nil, err
			}
			ret = b.(bool)

		case nil:
			ret = *new

		case []string, []bool, []float64, []int, []interface{}:
			switch action {
			case actionMerge, actionSum:
				return mergeArray(v, new)
			case actionAlter:
				ret = *new
			default:
				return nil, errInvalidAction
			}

		case map[string]interface{}, map[interface{}]interface{},
			map[string]int, map[interface{}]int,
			map[string]float64, map[interface{}]float64:
			switch action {
			case actionMerge:
				return mergeMap(v, new)
			case actionSum:
				return sumMap(v, new)
			case actionAlter:
				ret = *new
			default:
				return nil, errInvalidAction
			}

		case map[string]string, map[interface{}]string,
			map[string]bool, map[interface{}]bool:
			switch action {
			case actionMerge, actionSum:
				return mergeMap(v, new)
			case actionAlter:
				ret = *new
			default:
				return nil, errInvalidAction
			}

		default:
			if len(path) == 0 {
				return nil, fmt.Errorf("path is 0 (zero) length and unable to construct an object path for %T. Possibly due to bad parameters supplied", v)
			}
			return nil, fmt.Errorf("cannot locate `%s` in object path or no condition is made for `%T`. Please report this bug to https://github.com/lmorg/murex/issues", path[i-1], v)
		}

	default:
		return nil, errors.New("murex code error: default condition calculating the length of the path. I don't know how I got here. Please report this bug to https://github.com/lmorg/murex/issues")
	}

	if err == errOverwritePath {
		err = nil
	}
	return
}
