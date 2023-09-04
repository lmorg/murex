package expressions

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/expressions/symbols"
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
	case []string:
		return parseLambdaArray(tree, t, path)
	case []float64:
		return parseLambdaArray(tree, t, path)
	case []int:
		return parseLambdaArray(tree, t, path)
	case []interface{}:
		return parseLambdaArray(tree, t, path)
	case map[string]string:
		return parseLambdaMap(tree, t, path)
	case map[string]interface{}:
		return parseLambdaMap(tree, t, path)

	default:
		return nil, nil, fmt.Errorf("cannot run lambda. Expecting an array or map, instead got '%T' in '%s'", t, path)
	}
}

func (tree *ParserT) parseLambdaStatement(exec bool) ([]rune, interface{}, error) {
	if exec {
		if tree.p.IsMethod {
			return tree.parseLambdaStdin()

		} else {
			r := tree.expression[tree.charPos]
			return nil, nil, raiseError(tree.expression, nil, tree.charPos, fmt.Sprintf("%s '%s' (%d)",
				errMessage[symbols.Unexpected], string(r), r))
		}

	} else {
		r, v, _, err := tree.parseLambdaScala(false, '@', nil, varAsValue) // just parsing source
		return r, v, err
	}
}

func (tree *ParserT) parseLambdaStdin() ([]rune, interface{}, error) {
	dataType := tree.p.Stdin.GetDataType()
	b, err := tree.p.Stdin.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	name := fmt.Sprintf("_stdin_%d", tree.p.Id)
	err = tree.p.Variables.Set(tree.p, name, b, dataType)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to set temporary variable '%s' for piped lambda: %s", name, err.Error())
	}

	r, v, err := tree.parseLambda([]rune(name))
	if err != nil {
		return r, v, err
	}

	err = tree.p.Variables.Unset(name)
	if err != nil {
		return r, v, fmt.Errorf("unable to unset temporary variable '%s' for piped lambda: %s", name, err.Error())
	}

	return r, v, nil
}

var (
	errUnableToSetLambdaVar = "unable to set `$.`: %s"
	errUnableToGetLambdaVar = "unable to retrieve value of `$.`: %s"
	rxLineSeparator         = regexp.MustCompile(`(\r*\n)+`)
)

func parseLambdaString(tree *ParserT, t string, path string) ([]rune, interface{}, error) {
	var (
		pos  = tree.charPos
		item interface{}
		err  error
		r    []rune
		j    int
		fn   primitives.FunctionT
	)

	split := rxLineSeparator.Split(t, -1)
	array := make([]any, len(split))

	for i := range split {
		tree.charPos = pos
		err = tree.p.Variables.Set(tree.p, "", split[i], types.String)
		if err != nil {
			return nil, nil, fmt.Errorf(errUnableToSetLambdaVar, err.Error())
		}

		r, fn, err = tree.parseSubShell(true, '$', varAsValue)
		if err != nil {
			return nil, nil, err
		}
		item, _, err = fn()
		if err != nil {
			return nil, nil, err
		}
		switch t := item.(type) {
		case string:
			if len(t) > 0 {
				array[j] = item
				j++
			}
		case bool:
			if t {
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
		array   = make([]any, len(t))
		pos     = tree.charPos
		exitNum int
		r       []rune
		j       int
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
			return nil, nil, fmt.Errorf(errUnableToSetLambdaVar, err.Error())
		}

		r, exitNum, err = tree.parseLambdaRunSubShell()
		if err != nil {
			return nil, nil, err
		}

		if exitNum > 0 {
			continue
		}

		item, err := tree.p.Variables.GetValue("")
		if err != nil {
			return nil, nil, fmt.Errorf(errUnableToGetLambdaVar, err.Error())
		}

		array[j] = item
		j++
	}

	return r, array[:j], nil
}

func parseLambdaMap[K comparable, V any](tree *ParserT, t map[K]V, path string) ([]rune, interface{}, error) {
	var (
		pos     = tree.charPos
		object  = make(map[string]interface{})
		exitNum int
		r       []rune
	)

	for key, value := range t {
		tree.charPos = pos

		element, err := json.Marshal(map[string]interface{}{
			"key": key,
			"val": value,
		})
		if err != nil {
			return nil, nil, fmt.Errorf("unable to encode element: %s", err.Error())
		}

		err = tree.p.Variables.Set(tree.p, "", string(element), types.Json)
		if err != nil {
			return nil, nil, fmt.Errorf(errUnableToSetLambdaVar, err.Error())
		}

		r, exitNum, err = tree.parseLambdaRunSubShell()
		if err != nil {
			return nil, nil, err
		}

		if exitNum > 0 {
			continue
		}

		kv, err := tree.p.Variables.GetValue("")
		if err != nil {
			return nil, nil, fmt.Errorf(errUnableToGetLambdaVar, err.Error())
		}

		var (
			newKey string
			newVal interface{}
		)
		switch t := kv.(type) {
		case map[string]interface{}:
			newKey = fmt.Sprint(t["key"])
			newVal = t["val"]
		default:
			if err != nil {
				return nil, nil, fmt.Errorf("$. is not %T not an object", kv)
			}
		}

		object[newKey] = newVal
	}

	return r, object, nil
}

func (tree *ParserT) parseLambdaRunSubShell() ([]rune, int, error) {
	start := tree.charPos

	tree.charPos += 2

	_, err := tree.parseBlockQuote()
	if err != nil {
		return nil, 0, err
	}

	value := tree.expression[start : tree.charPos+1]
	block := tree.expression[start+2 : tree.charPos]

	if tree.p == nil {
		panic("`tree.p` is undefined")
	}

	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR | lang.F_PARENT_VARTABLE)
	exitNum, err := fork.Execute(block)
	if err != nil {
		return value, 1, fmt.Errorf("subshell failed: %s", err.Error())
	}

	return value, exitNum, err
}
