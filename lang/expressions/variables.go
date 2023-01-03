package expressions

import (
	"fmt"

	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

type varFormatting int

const (
	varAsString varFormatting = 0
	varAsValue  varFormatting = 1
)

var errEmptyRange = "range '@%s[%s]%s' is empty"

func (tree *ParserT) getVar(name []rune, strOrVal varFormatting) (interface{}, string, error) {
	var (
		value    interface{}
		dataType string
		err      error
		nameS    = string(name)
	)

	if strOrVal == varAsString {
		value, err = tree.p.Variables.GetString(nameS)
		if err != nil {
			return nil, "", err
		}
		value = utils.CrLfTrimString(value.(string))

	} else {
		dataType = tree.p.Variables.GetDataType(nameS)

		switch dataType {
		case types.String, types.Generic:
			value, err = tree.p.Variables.GetString(nameS)
			value = utils.CrLfTrimString(value.(string))

		case types.Number, types.Integer, types.Float, types.Boolean, types.Null:
			value, err = tree.p.Variables.GetValue(nameS)

		default:
			value, err = tree.p.Variables.GetString(nameS)
			if err != nil {
				return nil, "", err
			}
			fork := tree.p.Fork(lang.F_CREATE_STDIN | lang.F_NO_STDOUT | lang.F_NO_STDERR)
			_, err := fork.Stdin.Write([]byte(value.(string)))
			if err != nil {
				return nil, "", err
			}
			v, err := lang.UnmarshalData(fork.Process, dataType)
			if err != nil {
				return value, dataType, nil
			}
			return v, dataType, nil
		}
	}

	return value, dataType, err
}

const errEmptyArray = "array '@%s' is empty"

func (tree *ParserT) getArray(name []rune) (interface{}, error) {
	var nameS = string(name)

	data, err := tree.p.Variables.GetString(nameS)
	if err != nil {
		return nil, err
	}

	if data == "" && tree.StrictArrays() {
		return nil, fmt.Errorf(errEmptyArray, nameS)
	}

	var array []interface{}

	variable := streams.NewStdin()
	variable.SetDataType(tree.p.Variables.GetDataType(nameS))
	variable.Write([]byte(data))

	variable.ReadArrayWithType(tree.p.Context, func(v interface{}, _ string) {
		array = append(array, v)
	})

	if len(array) == 0 && tree.StrictArrays() {
		return nil, fmt.Errorf(errEmptyArray, nameS)
	}

	return array, nil
}

func (tree *ParserT) setVar(name []rune, value interface{}, dataType string) error {
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
	case types.String:
		return &primitives.DataType{Primitive: primitives.String}
	default:
		return &primitives.DataType{
			Primitive: primitives.Other,
			MxDT:      dt,
		}
	}
}

const (
	getVarIsIndex   = 1
	getVarIsElement = 2
)

func (tree *ParserT) getVarIndexOrElement(name, key []rune, isIorE int, strOrVal varFormatting) (interface{}, string, error) {
	var block []rune
	if isIorE == getVarIsIndex {
		block = createIndexBlock(name, key)
	} else {
		block = createElementBlock(name, key)
	}

	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
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

	block := make([]rune, 6+len(name)+len(index))
	block[0] = '$'
	copy(block[1:], name)
	copy(block[l:], []rune{'-', '>', ' ', '['})
	copy(block[l+4:], index)
	block[len(block)-1] = ']'
	return block
}

func createElementBlock(name, element []rune) []rune {
	l := len(name) + 1

	block := make([]rune, 8+len(name)+len(element))
	block[0] = '$'
	copy(block[1:], name)
	copy(block[l:], []rune{'-', '>', ' ', '[', '['})
	copy(block[l+5:], element)
	copy(block[len(block)-2:], []rune{']', ']'})
	return block
}

func createRangeBlock(name, key, flags []rune) []rune {
	l := len(name) + 1

	block := make([]rune, 7+len(name)+len(key)+len(flags))
	block[0] = '$'
	copy(block[1:], name)
	copy(block[l:], []rune{'-', '>', ' ', '@', '['})
	copy(block[l+5:], key)
	block[l+len(key)+5] = ']'
	copy(block[len(block)-len(flags):], flags)
	return block
}

func (tree *ParserT) getVarRange(name, key, flags []rune) (interface{}, error) {
	var array []interface{}

	block := createRangeBlock(name, key, flags)
	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
	fork.Execute(block)
	fork.Stdout.ReadArrayWithType(tree.p.Context, func(v interface{}, _ string) {
		array = append(array, v)
	})

	if len(array) == 0 && tree.StrictArrays() {
		return nil, fmt.Errorf(errEmptyRange, string(name), string(key), string(flags))
	}

	return array, nil

}

func formatBytes(b []byte, dataType string, strOrVal varFormatting) (interface{}, error) {
	if strOrVal == varAsString {
		return string(b), nil
	}

	switch dataType {
	case types.Number, types.String, types.Integer, types.Boolean, types.Null, types.Float:
		v, err := types.ConvertGoType(b, dataType)
		if err != nil {
			return nil, err
		}
		return v, nil

	default:
		return string(b), nil
	}
}
