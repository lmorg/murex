package expressions

import (
	"fmt"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/types"
)

func (tree *expTreeT) getVar(name string) (interface{}, string, error) {
	var (
		value interface{}
		err   error
	)

	dataType := tree.p.Variables.GetDataType(name)

	switch dataType {
	case types.Number, types.Integer, types.Boolean, types.Null, types.Float:
		value, err = tree.p.Variables.GetValue(name)

	default:
		value, err = tree.p.Variables.GetString(name)
	}

	return value, dataType, err
}

const errEmptyArray = "array '@%s' is empty"

func (tree *expTreeT) getArray(name string) (interface{}, error) {
	data, err := tree.p.Variables.GetString(name)
	if err != nil {
		return nil, err
	}

	strictArrays, err := tree.p.Config.Get("proc", "strict-arrays", "bool")
	if err != nil {
		strictArrays = true
	}

	if data == "" && strictArrays.(bool) {
		return nil, fmt.Errorf(errEmptyArray, name)
	}

	var array []interface{}

	variable := streams.NewStdin()
	variable.SetDataType(tree.p.Variables.GetDataType(name))
	variable.Write([]byte(data))

	variable.ReadArrayWithType(tree.p.Context, func(v interface{}, _ string) {
		array = append(array, v)
	})

	if len(array) == 0 && strictArrays.(bool) {
		return nil, fmt.Errorf(errEmptyArray, name)
	}

	return array, nil
}

func (tree *expTreeT) setVar(name string, value interface{}, dataType string) error {
	return tree.p.Variables.Set(tree.p, name, value, dataType)
}
