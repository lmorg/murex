package expressions

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/lang/types"
)

func treePlusPlus(tree *ParserT) { tree.charPos++ }

func (tree *ParserT) parseLambda(varName []rune) ([]rune, interface{}, error) {
	// no `exec` boolean here because this method should only be invoked when `exec == true`
	defer treePlusPlus(tree)
	if tree.p == nil {
		panic("`tree.p` is undefined")
	}

	path := string(varName)
	v, err := tree.p.Variables.GetValue(path)
	if err != nil {
		return nil, nil, err
	}

	switch t := v.(type) {
	case string:
		return parseLambdaString(tree, t, path)
	case []byte:
		return parseLambdaString(tree, string(t), path)
	case []rune:
		return parseLambdaString(tree, string(t), path)
	case []interface{}:
		return parseLambdaArray(tree, t, path)
	case []string:
		return parseLambdaArray(tree, t, path)
	case []float64:
		return parseLambdaArray(tree, t, path)
	case []int:
		return parseLambdaArray(tree, t, path)
	case map[string]interface{}:
		return parseLambdaMap(tree, t, path)
	case map[string]string:
		return parseLambdaMap(tree, t, path)
	default:
		return nil, nil, fmt.Errorf("cannot run lambda. Expecting an array, instead got '%T' in '%s'", t, path)
	}
}

var rxLineSeparator = regexp.MustCompile(`(\r*\n)+`)

func parseLambdaString(tree *ParserT, t string, path string) ([]rune, interface{}, error) {
	var (
		pos  = tree.charPos
		item interface{}
		err  error
		r    []rune
		j    int
	)

	split := rxLineSeparator.Split(t, -1)
	array := make([]any, len(split))

	for i := range split {
		tree.charPos = pos
		err = tree.p.Variables.Set(tree.p, "", split[i], types.String)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to set `$.`: %s", err.Error())
		}

		r, item, _, err = tree.parseSubShell(true, '$', varAsValue)
		if err != nil {
			return nil, nil, err
		}
		switch item.(type) {
		case string:
			if len(item.(string)) > 0 {
				array[j] = item
				j++
			}
		case bool:
			if item.(bool) {
				array[j] = split[i]
				j++
			}
		default:
			array[j] = item
			j++
		}
	}

	return r, array[:j], nil
}

func parseLambdaArray[V any](tree *ParserT, t []V, path string) ([]rune, interface{}, error) {
	var (
		array = make([]any, len(t))
		pos   = tree.charPos
		item  interface{}
		r     []rune
		j     int
	)

	for i := range t {
		tree.charPos = pos
		element := fmt.Sprintf("%s.%d", path, i)

		value, err := tree.p.Variables.GetValue(element)
		if err != nil {
			return nil, nil, err
		}

		dataType := tree.p.Variables.GetDataType(element)

		err = tree.p.Variables.Set(tree.p, "", value, dataType)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to set `$.`: %s", err.Error())
		}

		r, item, _, err = tree.parseSubShell(true, '$', varAsValue)
		if err != nil {
			return nil, nil, err
		}
		switch item.(type) {
		case string:
			if len(item.(string)) > 0 {
				array[j] = item
				j++
			}
		case bool:
			if item.(bool) {
				array[j] = value
				j++
			}
		default:
			array[j] = item
			j++
		}
	}

	return r, array[:j], nil
}

func parseLambdaMap[K comparable, V any](tree *ParserT, t map[K]V, path string) ([]rune, interface{}, error) {
	var (
		pos    = tree.charPos
		object = make(map[K]V)
		item   interface{}
		r      []rune
	)

	for key, value := range t {
		tree.charPos = pos
		//element := fmt.Sprintf("%s.%d", path, key)

		/*value, err := tree.p.Variables.GetValue(element)
		if err != nil {
			return nil, nil, err
		}*/

		//dataType := tree.p.Variables.GetDataType(element)

		element, err := json.Marshal(map[string]interface{}{
			"key": key,
			"val": value,
		})
		if err != nil {
			return nil, nil, fmt.Errorf("unable to encode element: %s", err.Error())
		}

		err = tree.p.Variables.Set(tree.p, "", string(element), types.Json)
		if err != nil {
			return nil, nil, fmt.Errorf("unable to set `$.`: %s", err.Error())
		}

		r, item, _, err = tree.parseSubShell(true, '$', varAsValue)
		if err != nil {
			return nil, nil, err
		}
		switch item.(type) {
		//case string:
		//if len(item.(string)) > 0 {
		//	array[j] = item
		//}
		case bool:
			if item.(bool) {
				object[key] = value
			}
			//default:
			//	object[key] = value
		}
	}

	return r, object, nil
}
