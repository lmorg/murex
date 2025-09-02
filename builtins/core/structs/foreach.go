package structs

import (
    "errors"
    "fmt"

    "github.com/lmorg/murex/lang"
    "github.com/lmorg/murex/lang/expressions/functions"
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
    foreachOrdered  = "--ordered"
    foreachUnordered = "--unordered"
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

        foreachOrdered:  types.Boolean,
        "-o":            foreachOrdered,
        foreachUnordered: types.Boolean,
        "-u":            foreachUnordered,
    },
}

func cmdForEach(p *lang.Process) error {
	flags, additional, err := p.Parameters.ParseFlags(argsForEach)

	if err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	switch {
	case flags[foreachJmap] == types.TrueString:
		return cmdForEachJmap(p)

	case flags[foreachParallel] != "":
		return cmdForEachParallel(p, flags, additional)

	default:
		return cmdForEachDefault(p, flags, additional)
	}
}

func convertToByte(v any) ([]byte, error) {
	s, err := types.ConvertGoType(v, types.String)
	if err != nil {
		return nil, err
	}

	return []byte(s.(string)), nil
}

func getFlagValueInt(flags map[string]string, flagName string) (int, error) {
	v, err := types.ConvertGoType(flags[flagName], types.Integer)
	if err != nil {
		return 0, fmt.Errorf(`expecting integer for %s, instead got "%s": %s`, flagName, flags[flagName], err.Error())
	}

	return v.(int), nil
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

func cmdForEachDefault(p *lang.Process, flags map[string]string, additional []string) error {
    block, varName, err := forEachInitializer(p, additional)
    if err != nil {
        return err
    }

    // Pre-parse the foreach block once
    var tree *[]functions.FunctionT
    if len(block) > 2 && block[0] == '{' && block[len(block)-1] == '}' {
        trimmed := block[1 : len(block)-1]
        t, err := lang.ParseBlock(trimmed)
        if err != nil {
            return err
        }
        tree = t
    } else {
        t, err := lang.ParseBlock(block)
        if err != nil {
            return err
        }
        tree = t
    }

	steps, err := getFlagValueInt(flags, foreachStep)
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
        forEachInnerLoopPreparsed(p, tree, varName, varValue, dataType, iteration)
    })

	if err != nil {
		return err
	}

    if steps > 0 && step > 0 {
        forEachInnerLoopPreparsed(p, tree, varName, slice[:step], types.Json, iteration+1)
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

// forEachInnerLoopPreparsed runs using a pre-parsed block tree to avoid per-iteration parsing cost.
func forEachInnerLoopPreparsed(p *lang.Process, tree *[]functions.FunctionT, varName string, varValue any, dataType string, iteration int) {
    var b []byte
    b, err := convertToByte(varValue)
    if err != nil {
        p.Done()
        return
    }

    if len(b) == 0 || p.HasCancelled() {
        return
    }

    fork := p.Fork(lang.F_PARENT_VARTABLE | lang.F_CREATE_STDIN)

    if varName != "!" {
        err = fork.Variables.Set(fork.Process, varName, varValue, dataType)
        if err != nil {
            p.Stderr.Writeln([]byte("error: " + err.Error()))
            p.Done()
            return
        }
    }

    if !setMetaValues(fork.Process, iteration) {
        return
    }

    fork.Stdin.SetDataType(dataType)
    _, err = fork.Stdin.Writeln(b)
    if err != nil {
        p.Stderr.Writeln([]byte("error: " + err.Error()))
        p.Done()
        return
    }
    _, err = fork.ExecuteTree(tree)
    if err != nil {
        p.Stderr.Writeln([]byte("error: " + err.Error()))
        p.Done()
        return
    }
}
