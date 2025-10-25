package structs

import (
	"errors"
	"fmt"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/mattn/go-runewidth"
)

func init() {
	lang.DefineMethod("foreach", cmdForEach, types.ReadArrayWithType, types.Any)
}

const (
	foreachJmap     = "--jmap"
	foreachStep     = "--step"
	foreachParallel = "--parallel"
)

var argsForEach = &parameters.Arguments{
	AllowAdditional: true,
	Flags: map[string]string{
		foreachJmap: types.Boolean,
		"-j":        foreachJmap,

		foreachStep: types.Integer,
		"-s":        foreachStep,

		foreachParallel: types.Integer,
		"-p":            foreachParallel,
	},
}

func cmdForEach(p *lang.Process) error {
	flags, additional, err := p.Parameters.ParseFlags(argsForEach)

	if err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	switch {
	case flags.GetValue(foreachJmap).Boolean():
		return cmdForEachJmap(p)

	case flags.GetValue(foreachParallel).Integer() != 0:
		return cmdForEachParallel(p, flags.GetValue(foreachParallel).Integer(), additional)

	default:
		return cmdForEachDefault(p, flags.GetValue(foreachStep).Integer(), additional)
	}
}

func convertToByte(v any) ([]byte, error) {
	s, err := types.ConvertGoType(v, types.String)
	if err != nil {
		return nil, err
	}

	return []byte(s.(string)), nil
}

func forEachInitializer(p *lang.Process, additional []string) (block []rune, varName string, err error) {
	dataType := p.Stdin.GetDataType()
	if dataType == types.Json {
		p.Stdout.SetDataType(types.JsonLines)
	} else {
		p.Stdout.SetDataType(dataType)
	}

	switch len(additional) {
	case 1:
		varName = "!"
		block = []rune(additional[0])

	case 2:
		varName = additional[0]
		block = []rune(additional[1])

	default:
		return nil, "", errors.New("invalid number of parameters")
	}
	if !types.IsBlockRune(block) {
		return nil, "", fmt.Errorf("invalid code block: `%s`", runewidth.Truncate(string(block), 70, "…"))
	}

	return
}

func cmdForEachDefault(p *lang.Process, steps int, additional []string) error {
	block, varName, err := forEachInitializer(p, additional)
	if err != nil {
		return err
	}

	var (
		step      int
		iteration int
		slice     = make([]any, steps)
	)

	err = p.Stdin.ReadArrayWithType(p.Context, func(varValue any, dataType string) {
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

func forEachInnerLoop(p *lang.Process, block []rune, varName string, varValue any, dataType string, iteration int) {
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
