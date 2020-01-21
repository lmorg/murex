package alter

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

// SplitPath takes a string with a prefixed delimiter and separates it into a slice of path elements
func SplitPath(path string) ([]string, error) {
	split := strings.Split(path, string(path[0]))
	if len(split) == 0 || (len(split) == 1 && split[0] == "") {
		return nil, errors.New("Empty path")
	}

	if split[0] == "" {
		split = split[1:]
	}

	return split, nil
}

// Alter a data structure. Requires a path (pre-split) and new structure as a
// JSON string. A more seasoned developer will see plenty of room for
// optimisation however this function was largely thrown together in a "let's
// create something that works first and worry about performance later" kind of
// sense (much like a lot of murex's code base). That being said, I will accept
// any pull requests from other developers wishing to improve this - or other -
// functions. I'm also open to any breaking changes those optimisations might
// bring (at least until the project reaches version 1.0).
func Alter(ctx context.Context, v interface{}, path []string, new string) (interface{}, error) {
	return loop(ctx, v, 0, path, &new, false)
}

// Merge a data structure; like Alter but merges arrays and maps where possible
func Merge(ctx context.Context, v interface{}, path []string, new string) (interface{}, error) {
	return loop(ctx, v, 0, path, &new, true)
}

var errOverwritePath = errors.New("internal condition: path needs overwriting")

func loop(ctx context.Context, v interface{}, i int, path []string, new *string, merge bool) (ret interface{}, err error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("Cancelled")
	default:
	}

	switch {
	case i < len(path):
		switch v.(type) {
		case []interface{}:
			pathI, err := strconv.Atoi(path[i])
			if err != nil {
				return nil, fmt.Errorf("Expecting an array index in path element '%s': %s", path[i], err)
			}

			if pathI < 0 {
				return nil, fmt.Errorf("Negative indexes not allowed in arrays: path element '%d'", pathI)
			}

			if pathI >= len(v.([]interface{})) {
				return nil, fmt.Errorf("Index greater than length of array in path element '%d' (array length '%d')", pathI, len(v.([]interface{})))
			}

			ret, err = loop(ctx, v.([]interface{})[pathI], i+1, path, new, merge)
			if err == errOverwritePath {
				v.([]interface{})[pathI] = parseString(new)
			}
			if err == nil {
				v.([]interface{})[pathI] = ret
				ret = v
			}

		case map[interface{}]interface{}:
			ret, err = loop(ctx, v.(map[interface{}]interface{})[path[i]], i+1, path, new, merge)
			if err == errOverwritePath {
				v.(map[interface{}]interface{})[path[i]] = parseString(new)
			}
			if err == nil {
				v.(map[interface{}]interface{})[path[i]] = ret
				ret = v
			}

		case map[string]interface{}:
			ret, err = loop(ctx, v.(map[string]interface{})[path[i]], i+1, path, new, merge)
			if err == errOverwritePath {
				v.(map[string]interface{})[path[i]] = parseString(new)
			}
			if err == nil {
				v.(map[string]interface{})[path[i]] = ret
				ret = v
			}

		case map[interface{}]string:
			ret, err = loop(ctx, v.(map[interface{}]string)[path[i]], i+1, path, new, merge)
			if err == errOverwritePath {
				v.(map[interface{}]string)[path[i]] = fmt.Sprint(parseString(new))
			}
			if err == nil {
				v.(map[interface{}]string)[path[i]] = ret.(string)
				ret = v
			}

		case nil:
			// Let's overwrite part of the path
			return nil, errOverwritePath

		case string, int, float64, bool:
			return nil, fmt.Errorf("Unable to alter data structure using that path because one of the path elements is an end of tree (%T) rather than a map. Instead please have the full path you want to add as part of the amend JSON string in `alter`", v)

		default:
			return nil, fmt.Errorf("murex code error: No condition is made for `%T`. Please report this bug to https://github.com/lmorg/murex/issues", v)
		}

	case i == len(path):
		switch t := v.(type) {
		case string:
			ret = *new

		case int:
			num, err := strconv.Atoi(*new)
			if err != nil {
				return nil, err
			}
			ret = num

		case float64:
			num, err := strconv.ParseFloat(*new, 64)
			if err != nil {
				return nil, err
			}
			ret = num

		case bool:
			ret = types.IsTrue([]byte(*new), 0)

		case nil, map[string]interface{}:
			ret = parseString(new)

		default:
			return nil, fmt.Errorf("Cannot locate `%s` in object path or no condition is made for `%T`. Please report this bug to https://github.com/lmorg/murex/issues", path[i-1], t)
		}

	default:
		return nil, errors.New("murex code error: default condition calculating the length of the path. I don't know how I got here. Please report this bug to https://github.com/lmorg/murex/issues")
	}

	if err == errOverwritePath {
		err = nil
	}
	return
}

func parseString(new *string) (v interface{}) {
	// Regardless of the output format of `alter` - we only really want to accept JSON input
	// because that's how murex's parser is designed to read multiline structures.
	err := json.Unmarshal([]byte(*new), &v)
	if err == nil {
		return v
	}

	return new
}
