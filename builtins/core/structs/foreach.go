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

func cmdForEach(p *lang.Process) error {
	flags, additional, err := p.Parameters.ParseFlags(argsForEach)
	//flags := map[string]string{}
	//additional := p.Parameters.StringArray()
	//var err error
	if err != nil {
		p.Stdout.SetDataType(types.Null)
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

func getSteps(flags map[string]string) (int, []interface{}, error) {
	steps, err := types.ConvertGoType(flags[foreachStep], types.Integer)
	if err != nil {
		return 0, nil, fmt.Errorf(`expecting integer for %s, instead got "%s": %s`, foreachStep, flags[foreachStep], err.Error())
	}

	return steps.(int), make([]any, steps.(int)), nil
}

func cmdForEachDefault(p *lang.Process, flags map[string]string, additional []string) error {
	dataType := p.Stdin.GetDataType()
	if dataType == types.Json {
		p.Stdout.SetDataType(types.JsonLines)
	} else {
		p.Stdout.SetDataType(dataType)
	}

	var (
		block   []rune
		varName string
	)

	switch len(additional) {
	case 1:
		varName = "!"
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

	var (
		step      int
		iteration int
	)

	err = p.Stdin.ReadArrayWithType(p.Context, func(varValue interface{}, dataType string) {
		if steps > 0 {
			varValue, _ = marshal(p, varValue, dataType)
			slice[step] = varValue
			step++
			if step == steps {
				varValue = slice
				dataType = types.Json
				step = 0
			} else {
				return
			}
		}

		iteration++
		forEachInnerLoop(p, block, varName, varValue, dataType, iteration)
	})

	if err != nil {
		return err
	}

	if steps > 0 && step > 0 {
		forEachInnerLoop(p, block, varName, slice[:step], types.Json, iteration+1)
	}

	return nil
}

func marshal(p *lang.Process, v any, dataType string) (any, error) {
	switch v.(type) {
	case []byte:
		if dataType != types.String && dataType != types.Generic {
			return lang.UnmarshalDataBuffered(p, v.([]byte), dataType)
		}
	case string:
		if dataType != types.String && dataType != types.Generic {
			return lang.UnmarshalDataBuffered(p, []byte(v.(string)), dataType)
		}
	}
	return v, nil
}

func setMetaValues(p *lang.Process, iteration int) bool {
	meta := map[string]any{
		"i": iteration,
	}
	err := p.Variables.Set(p, "", meta, types.Json)
	if err != nil {
		p.Stderr.Writeln([]byte("unable to set meta variable: " + err.Error()))
		p.Done()
		return false
	}
	return true
}

func forEachInnerLoop(p *lang.Process, block []rune, varName string, varValue interface{}, dataType string, iteration int) {
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
		err = p.Variables.Set(p, varName, varValue, dataType)
		if err != nil {
			p.Stderr.Writeln([]byte("error: " + err.Error()))
			p.Done()
			return
		}
	}

	if !setMetaValues(p, iteration) {
		return
	}

	fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_CREATE_STDIN)
	fork.Stdin.SetDataType(dataType)
	_, err = fork.Stdin.Writeln(b)
	if err != nil {
		p.Stderr.Writeln([]byte("error: " + err.Error()))
		p.Done()
		return
	}
	_, err = fork.Execute(block)
	if err != nil {
		p.Stderr.Writeln([]byte("error: " + err.Error()))
		p.Done()
		return
	}
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

	var (
		m         = make(map[string]string)
		iteration int
	)

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

		iteration++
		if !setMetaValues(p, iteration) {
			return
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
