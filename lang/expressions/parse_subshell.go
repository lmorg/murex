package expressions

import (
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions/primitives"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/consts"
)

func (tree *ParserT) parseSubShell(exec bool, sigil rune, strOrVal varFormatting) ([]rune, primitives.FunctionT, error) {
	start := tree.charPos

	tree.charPos += 2

	_, err := tree.parseBlockQuote()
	if err != nil {
		return nil, nil, err
	}

	value := tree.expression[start : tree.charPos+1]
	block := tree.expression[start+2 : tree.charPos]

	if !exec {
		return value, nil, nil
	}

	switch sigil {
	case '$':
		fn := func() (*primitives.Value, error) {
			return execSubShellScalar(tree, block, strOrVal)
		}
		return value, fn, nil

	case '@':
		fn := func() (*primitives.Value, error) {
			return execSubShellArray(tree, block, strOrVal)
		}
		return value, fn, nil
	default:
		err = fmt.Errorf("invalid prefix in expression '%s'. %s", string(sigil), consts.IssueTrackerURL)
		return nil, nil, err
	}
}

func execSubShellScalar(tree *ParserT, block []rune, strOrVal varFormatting) (*primitives.Value, error) {
	if tree.p == nil {
		panic("`tree.p` is undefined")
	}

	val := new(primitives.Value)
	var err error

	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_PARENT_VARTABLE | lang.F_CREATE_STDOUT)
	val.ExitNum, err = fork.Execute(block)
	if err != nil {
		return val, fmt.Errorf("subshell failed: %s", err.Error())
	}
	if val.ExitNum > 0 && tree.p.RunMode.IsStrict() {
		return val, fmt.Errorf("subshell exit status %d", val.ExitNum)
	}
	b, err := fork.Stdout.ReadAll()
	if err != nil {
		return val, fmt.Errorf("cannot read from sub-shell's stdout: %s", err.Error())
	}

	b = utils.CrLfTrim(b)
	val.DataType = fork.Stdout.GetDataType()

	val.Value, err = formatBytes(tree, b, val.DataType, strOrVal)
	return val, err
}

func execSubShellArray(tree *ParserT, block []rune, strOrVal varFormatting) (*primitives.Value, error) {
	var slice []interface{}
	val := new(primitives.Value)
	var err error

	fork := tree.p.Fork(lang.F_NO_STDIN | lang.F_CREATE_STDOUT | lang.F_PARENT_VARTABLE)
	val.ExitNum, err = fork.Execute(block)
	if err != nil {
		return val, fmt.Errorf("subshell failed: %s", err.Error())
	}
	if val.ExitNum > 0 && tree.p.RunMode.IsStrict() {
		return val, fmt.Errorf("subshell exit status %d", val.ExitNum)
	}

	switch strOrVal {
	case varAsString:
		err = fork.Stdout.ReadArray(tree.p.Context, func(b []byte) {
			slice = append(slice, string(b))
		})
	case varAsValue:
		err = fork.Stdout.ReadArrayWithType(tree.p.Context, func(v interface{}, _ string) {
			slice = append(slice, v)
		})
	default:
		panic("invalid value set for strOrVal")
	}
	if err != nil {
		return val, fmt.Errorf("cannot read from subshell's STDOUT: %s", err.Error())
	}

	if len(slice) == 0 && tree.StrictArrays() {
		return nil, fmt.Errorf(errEmptyArray, "{"+string(block)+"}")
	}

	val.Value = slice
	val.DataType = types.Json
	return val, nil
}
