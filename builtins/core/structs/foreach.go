package structs

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
	"github.com/mattn/go-runewidth"
)

func init() {
	lang.DefineMethod("foreach", cmdForEach, types.ReadArrayWithType, types.Any)
}

const (
	foreachJmap = "--jmap"
	foreachStep = "--step"
)

var argsForEach = &parameters.Arguments{
	AllowAdditional: true,
	Flags: map[string]string{
		foreachJmap: types.Boolean,
		foreachStep: types.Integer,
	},
}

func cmdForEach(p *lang.Process) (err error) {
	flags, additional, err := p.Parameters.ParseFlags(argsForEach)
	if err != nil {
		return err
	}

	switch {
	case flags[foreachJmap] == types.TrueString:
		return cmdForEachJmap(p)
	default:
		return cmdForEachDefault(p, flags, additional)
	}
}

func convertToByte(v interface{}) ([]byte, error) {
	s, err := types.ConvertGoType(v, types.String)
	if err != nil {
		return nil, err
	}

	return []byte(s.(string)), nil
}

func getSteps(flags map[string]string) (int, []any, error) {
	steps, err := types.ConvertGoType(flags[foreachStep], types.Integer)
	if err != nil {
		return 0, nil, fmt.Errorf(`expecting integer for %s, instead got "%s": %s`, foreachStep, flags[foreachStep], err.Error())
	}

	return steps.(int), make([]any, steps.(int)), nil
}

func cmdForEachDefault(p *lang.Process, flags map[string]string, additional []string) (err error) {
	dt := p.Stdin.GetDataType()
	if dt == types.Json {
		p.Stdout.SetDataType(types.JsonLines)
	} else {
		p.Stdout.SetDataType(dt)
	}

	var (
		block   []rune
		varName string
	)

	switch len(additional) {
	case 1:
		block = []rune(additional[0])

	case 2:
		varName = additional[0]
		block = []rune(additional[1])

	default:
		return errors.New("invalid number of parameters")
	}
	if !types.IsBlockRune(block) {
		return fmt.Errorf("invalid code block: `%s`", runewidth.Truncate(string(block), 70, "…"))
	}

	steps, slice, err := getSteps(flags)
	if err != nil {
		return err
	}

	var step int

	err = p.Stdin.ReadArrayWithType(p.Context, func(v interface{}, dataType string) {
		if steps > 0 {
			step++
			slice[step-1] = v
			if step == steps {
				v = json.LazyLogging(slice)
				dataType = types.Json
				step = 0
			} else {
				return
			}
		}

		forEachInnerLoop(p, block, varName, v, dataType)
	})

	if steps > 0 && step > 0 {
		forEachInnerLoop(p, block, varName, slice[:step], types.Json)
	}

	return
}

func forEachInnerLoop(p *lang.Process, block []rune, varName string, varValue interface{}, dataType string) {
	var b []byte
	b, err := convertToByte(varValue)
	if err != nil {
		p.Done()
		return
	}

	if len(b) == 0 || p.HasCancelled() {
		return
	}

	if varName != "!" {
		p.Variables.Set(p, varName, varValue, dataType)
	}

	fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_CREATE_STDIN)
	fork.Stdin.SetDataType(dataType)
	fork.Stdin.Writeln(b)
	fork.Execute(block)
}

func cmdForEachJmap(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	varName, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	blockKey, err := p.Parameters.Block(2)
	if err != nil {
		return err
	}

	blockVal, err := p.Parameters.Block(3)
	if err != nil {
		return err
	}

	m := make(map[string]string)

	err = p.Stdin.ReadArrayWithType(p.Context, func(v interface{}, dt string) {
		var b []byte
		b, err = convertToByte(v)
		if err != nil {
			p.Done()
			return
		}

		if len(b) == 0 || p.HasCancelled() {
			return
		}

		if varName != "!" {
			p.Variables.Set(p, varName, v, dt)
		}

		forkKey := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
		forkKey.Execute(blockKey)
		bKey, err := forkKey.Stdout.ReadAll()
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			p.Kill()
		}

		forkVal := p.Fork(lang.F_PARENT_VARTABLE | lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
		forkVal.Execute(blockVal)
		bVal, err := forkVal.Stdout.ReadAll()
		if err != nil {
			p.Stderr.Writeln([]byte(err.Error()))
			p.Kill()
		}

		m[string(utils.CrLfTrim(bKey))] = string(utils.CrLfTrim(bVal))
	})

	if err != nil {
		return err
	}

	b, err := json.Marshal(m, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
