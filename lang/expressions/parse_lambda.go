package expressions

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/symbols"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

var errCancelled = errors.New("cancelled")

func treePlusPlus(tree *ParserT) { tree.charPos++ }

func (tree *ParserT) parseLambdaExecFalse(sigil rune, varName []rune) ([]rune, error) {
	defer treePlusPlus(tree)

	r, _, err := tree.parseSubShell(false, sigil, varAsValue)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (tree *ParserT) parseLambdaExecTrue(varName []rune, sigil rune, dt string) ([]rune, interface{}, string, error) {
	// no `exec` boolean here because this method should only be invoked when `exec == true`
	defer treePlusPlus(tree)
	if tree.p == nil {
		panic("`tree.p` is undefined")
	}

	r, block, err := tree.parseLambdaGetSubShellBlock()
	if err != nil {
		return nil, nil, "", err
	}

	return _parseLambdaExecTrue(tree.p, varName, sigil, dt, r, block, tree.StrictArrays())
}

type parseLambdaExecTrueT func() ([]rune, any, string, error)

func _parseLambdaExecTrue(p *lang.Process, varName []rune, sigil rune, dt string, r []rune, block []rune, strictArrays bool) ([]rune, any, string, error) {
	path := string(varName)
	v, err := p.Variables.GetValue(path)
	if err != nil {
		return nil, nil, "", err
	}

	if dt == "" {
		// TODO: test me please
		dt = p.Variables.GetDataType(path)
	}

	switch t := v.(type) {
	case nil:
		if strictArrays {
			return nil, nil, "", fmt.Errorf("cannot run lambda: value is a null object")
		}
		return parseLambdaArray(p, sigil, []string{}, dt, r, block)

		/*case string:
			return parseLambdaArray(tree, []string{t}, dt)
		case []byte:
			return parseLambdaArray(tree, []string{string(t)}, dt)
		case []rune:
			return parseLambdaArray(tree, []string{string(t)}, dt)*/

	/*case string:
		return parseLambdaString(tree, t, dt)
	case []byte:
		return parseLambdaString(tree, string(t), dt)
	case []rune:
		return parseLambdaString(tree, string(t), dt)*/

	case []string:
		return parseLambdaArray(p, sigil, t, dt, r, block)
	case []float64:
		return parseLambdaArray(p, sigil, t, dt, r, block)
	case []int:
		return parseLambdaArray(p, sigil, t, dt, r, block)
	case []bool:
		return parseLambdaArray(p, sigil, t, dt, r, block)

	case []interface{}:
		return parseLambdaArray(p, sigil, t, dt, r, block)
	case map[string]string:
		return parseLambdaMap(p, sigil, t, dt, r, block)
	case map[string]interface{}:
		return parseLambdaMap(p, sigil, t, dt, r, block)

	default:
		return nil, nil, "", fmt.Errorf("cannot run lambda: expecting an array or map, instead got '%T' in '%s' (%s)", t, path, dt)
	}
}

func (tree *ParserT) parseLambdaStatement(exec bool, sigil rune) ([]rune, interface{}, string, error) {
	if exec {
		if tree.p.IsMethod {
			return tree.parseLambdaStdinExecTrue(sigil)

		} else {
			r := tree.expression[tree.charPos]
			return nil, nil, "", raiseError(tree.expression, nil, tree.charPos, fmt.Sprintf("%s '%s' (%d)",
				errMessage[symbols.Unexpected], string(r), r))
		}

	} else {
		r, err := tree.parseLambdaExecFalse(sigil, nil)
		return r, nil, "", err
	}
}

func (tree *ParserT) parseLambdaStdinExecTrue(sigil rune) ([]rune, interface{}, string, error) {
	dataType := tree.p.Stdin.GetDataType()
	b, err := tree.p.Stdin.ReadAll()
	if err != nil {
		return nil, nil, "", err
	}

	name := fmt.Sprintf("_stdin_%d", tree.p.Id)
	err = tree.p.Variables.Set(tree.p, name, b, dataType)
	if err != nil {
		return nil, nil, "", fmt.Errorf("unable to set temporary variable '%s' for piped lambda: %s", name, err.Error())
	}

	r, v, dt, err := tree.parseLambdaExecTrue([]rune(name), sigil, dataType)
	if err != nil {
		return r, v, dt, err
	}

	err = tree.p.Variables.Unset(name)
	if err != nil {
		return r, v, dt, fmt.Errorf("unable to unset temporary variable '%s' for piped lambda: %s", name, err.Error())
	}

	return r, v, dt, nil
}

var (
	errUnableToSetLambdaVar = "unable to set `$.`: %s"
	errUnableToGetLambdaVar = "unable to retrieve value of `$.`: %s"
	errUnableToUpdateValue  = "cannot update value from lambda: %s"

	//rxLineSeparator = regexp.MustCompile(`(\r*\n)+`)
)

const (
	LAMBDA_ITERATION = "i"
	LAMBDA_KEY       = "k"
	LAMBDA_VALUE     = "v"
)

func writeKeyValVariable[K comparable](p *lang.Process, iteration int, key K, value any) error {
	element, err := json.Marshal(map[string]interface{}{
		LAMBDA_ITERATION: iteration,
		LAMBDA_KEY:       key,
		LAMBDA_VALUE:     value,
	})
	if err != nil {
		return fmt.Errorf("unable to encode element: %s", err.Error())
	}

	return p.Variables.Set(p, "", string(element), types.Json)
}

func readKeyValVariable(p *lang.Process) (any, any, error) {
	kv, err := p.Variables.GetValue("")
	if err != nil {
		return nil, nil, fmt.Errorf(errUnableToGetLambdaVar, err.Error())
	}

	switch t := kv.(type) {
	case map[string]any:
		return t[LAMBDA_KEY], t[LAMBDA_VALUE], nil
	default:
		return nil, nil, fmt.Errorf("expecting $. to be '{%s: str, %s ...}', instead got '%T'", LAMBDA_KEY, LAMBDA_VALUE, kv)
	}
}

/*func parseLambdaString(tree *ParserT, t string, path string) ([]rune, interface{}, error) {
	var (
		item interface{}
		err  error
		r    []rune
		j    int
		fn   primitives.FunctionT
	)

	split := rxLineSeparator.Split(t, -1)
	array := make([]any, len(split))

	for i := range split {
		//tree.charPos = pos
		err = tree.p.Variables.Set(tree.p, "", split[i], types.String)
		if err != nil {
			return nil, nil, fmt.Errorf(errUnableToSetLambdaVar, err.Error())
		}

		r, fn, err = tree.parseSubShell(true, '$', varAsValue)
		if err != nil {
			return nil, nil, err
		}
		val, err := fn()
		item = val.Value
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
}*/

func parseLambdaArray[V any](p *lang.Process, sigil rune, t []V, dt string, r []rune, block []rune) ([]rune, interface{}, string, error) {
	var (
		array  []any
		stdout []string
	)

	for i := range t {
		if p.HasCancelled() {
			return nil, nil, "", errCancelled
		}

		err := writeKeyValVariable(p, i+1, i, t[i])
		if err != nil {
			return nil, nil, "", fmt.Errorf(errUnableToSetLambdaVar, err.Error())
		}

		exitNum, b, err := parseLambdaRunSubShell(p, block)
		if err != nil {
			return nil, nil, "", err
		}

		if len(b) > 0 {
			stdout = append(stdout, string(utils.CrLfTrim(b)))
			continue
		}

		if exitNum > 0 {
			continue
		}

		newKey, newVal, err := readKeyValVariable(p)
		if err != nil {
			return nil, nil, "", fmt.Errorf(errUnableToUpdateValue, err.Error())
		}
		if fmt.Sprint(newKey) != fmt.Sprint(i) {
			return nil, nil, "", fmt.Errorf("arrays cannot have their $.%s changed: old key '%v', new key '%v'", LAMBDA_KEY, i, newKey)
		}

		array = append(array, newVal)
	}

	switch sigil {
	case '$':
		if len(stdout) > 0 {
			s := strings.Join(stdout, "\n")
			return r, s, types.String, nil
		}
		b, err := lang.MarshalData(p, dt, array)
		return r, string(b), dt, err

	case '@':
		if len(stdout) > 0 {
			return r, stdout, types.Json, nil
		}
		return r, array, types.Json, nil

	default:
		panic("uncaught sigil in switch statement")
	}
}

func parseLambdaMap[K comparable, V any](p *lang.Process, sigil rune, t map[K]V, dt string, r []rune, block []rune) ([]rune, interface{}, string, error) {
	var (
		object = make(map[any]interface{})
		stdout []string
	)

	var i int
	for key, value := range t {
		if p.HasCancelled() {
			return nil, nil, "", errCancelled
		}

		i++
		err := writeKeyValVariable(p, i, key, value)
		if err != nil {
			return nil, nil, "", fmt.Errorf(errUnableToSetLambdaVar, err.Error())
		}

		exitNum, b, err := parseLambdaRunSubShell(p, block)
		if err != nil {
			return nil, nil, "", err
		}

		if len(b) > 0 {
			stdout = append(stdout, string(utils.CrLfTrim(b)))
			continue
		}

		if exitNum > 0 {
			continue
		}

		newKey, newVal, err := readKeyValVariable(p)
		if err != nil {
			return nil, nil, "", fmt.Errorf(errUnableToUpdateValue, err.Error())
		}

		object[newKey] = newVal
	}

	switch sigil {
	case '$':
		if len(stdout) > 0 {
			s := strings.Join(stdout, "\n")
			return r, s, types.String, nil
		}
		b, err := lang.MarshalData(p, dt, object)
		return r, string(b), dt, err

	case '@':
		if len(stdout) > 0 {
			return r, stdout, types.Json, nil
		}
		return r, object, types.Json, nil

	default:
		panic("uncaught sigil in switch statement")
	}
}

func (tree *ParserT) parseLambdaGetSubShellBlock() ([]rune, []rune, error) {
	start := tree.charPos

	tree.charPos += 2

	_, err := tree.parseBlockQuote()
	if err != nil {
		return nil, nil, err
	}

	value := tree.expression[start : tree.charPos+1]
	block := tree.expression[start+2 : tree.charPos]

	return value, block, nil
}

func parseLambdaRunSubShell(p *lang.Process, block []rune) (int, []byte, error) {
	fork := p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_NO_STDERR | lang.F_PARENT_VARTABLE)
	exitNum, err := fork.Execute(block)
	if err != nil {
		return 1, nil, fmt.Errorf("subshell failed: %s", err.Error())
	}

	b, err := fork.Stdout.ReadAll()
	if err != nil {
		return 1, nil, fmt.Errorf("unable to read from stdout: %s", err.Error())
	}

	if fork.Stdout.GetDataType() == types.Boolean {
		if types.IsTrue(b, exitNum) {
			if string(b) == types.TrueString {
				return 0, nil, nil
			}
			return 0, b, nil
		}
		return 1, nil, nil
	}

	return exitNum, b, err
}
