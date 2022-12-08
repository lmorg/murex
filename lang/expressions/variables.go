package expressions

import (
	"fmt"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

const (
	varAsString = true
	varAsValue  = false
)

func (tree *expTreeT) getVar(name []rune, strOrVal bool) (interface{}, string, error) {
	var (
		value    interface{}
		dataType string
		err      error
		nameS    = string(name)
	)

	if strOrVal { // == varAsString
		value, err = tree.p.Variables.GetString(nameS)

	} else {
		dataType = tree.p.Variables.GetDataType(nameS)

		switch dataType {
		case types.Number, types.Integer, types.Boolean, types.Null, types.Float:
			value, err = tree.p.Variables.GetValue(nameS)

		default:
			value, err = tree.p.Variables.GetString(nameS)
			value = utils.CrLfTrimString(value.(string))
		}
	}

	return value, dataType, err
}

const errEmptyArray = "array '@%s' is empty"

func (tree *expTreeT) getArray(name []rune) (interface{}, error) {
	var nameS = string(name)

	data, err := tree.p.Variables.GetString(nameS)
	if err != nil {
		return nil, err
	}

	strictArrays, err := tree.p.Config.Get("proc", "strict-arrays", "bool")
	if err != nil {
		strictArrays = true
	}

	if data == "" && strictArrays.(bool) {
		return nil, fmt.Errorf(errEmptyArray, nameS)
	}

	var array []interface{}

	variable := streams.NewStdin()
	variable.SetDataType(tree.p.Variables.GetDataType(nameS))
	variable.Write([]byte(data))

	variable.ReadArrayWithType(tree.p.Context, func(v interface{}, _ string) {
		array = append(array, v)
	})

	if len(array) == 0 && strictArrays.(bool) {
		return nil, fmt.Errorf(errEmptyArray, nameS)
	}

	return array, nil
}

func (tree *expTreeT) setVar(name []rune, value interface{}, dataType string) error {
	nameS := string(name)
	return tree.p.Variables.Set(tree.p, nameS, value, dataType)
}

func scalar2Primitive(dt string) *primitives.DataType {
	switch dt {
	case types.Number, types.Integer, types.Float:
		return &primitives.DataType{Primitive: primitives.Number}
	case types.Boolean:
		return &primitives.DataType{Primitive: primitives.Boolean}
	case types.Null:
		return &primitives.DataType{Primitive: primitives.Null}
	default:
		return &primitives.DataType{Primitive: primitives.String}
	}
}

const (
	getVarIsIndex   = 1
	getVarIsElement = 2
)

func (tree *expTreeT) getVarIndexOrElement(name, key []rune, isIorE int, strOrVal bool) (interface{}, string, error) {
	var block []rune
	if isIorE == getVarIsIndex {
		block = createIndexBlock(name, key)
	} else {
		block = createElementBlock(name, key)
	}

	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_PARENT_VARTABLE)
	fork.Execute(block)
	b, err := fork.Stdout.ReadAll()
	if err != nil {
		return "", "", err
	}

	b = utils.CrLfTrim(b)
	dataType := fork.Stdout.GetDataType()

	v, err := formatBytes(b, dataType, strOrVal)
	return v, dataType, err
}

func createIndexBlock(name, index []rune) []rune {
	l := len(name) + 1

	block := make([]rune, 5+len(name)+len(index))
	block[0] = '$'
	copy(block[1:], name)
	copy(block[l:], []rune{'-', '>', '['})
	copy(block[l+3:], index)
	block[len(block)-1] = ']'
	return block
}

func createElementBlock(name, element []rune) []rune {
	l := len(name) + 1

	block := make([]rune, 7+len(name)+len(element))
	block[0] = '$'
	copy(block[1:], name)
	copy(block[l:], []rune{'-', '>', '[', '['})
	copy(block[l+4:], element)
	copy(block[len(block)-2:], []rune{']', ']'})
	return block
}

func formatBytes(b []byte, dataType string, strOrVal bool) (interface{}, error) {
	if strOrVal { // == varAsString
		return string(b), nil
	}

	switch dataType {
	case types.Number, types.Integer, types.Boolean, types.Null, types.Float:
		v, err := types.ConvertGoType(b, dataType)
		if err != nil {
			return nil, err
		}
		return v, nil

	default:
		return string(b), nil
	}
}
