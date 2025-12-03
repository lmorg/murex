package structs

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("function", cmdFunc, types.Null)
	lang.DefineFunction("!function", cmdUnfunc, types.Null)
	lang.DefineFunction("private", cmdPrivate, types.Null)
	lang.DefineFunction("!private", cmdUnprivate, types.Null)
	lang.DefineFunction("method", cmdMethod, types.Null)

	defaults.AppendProfile(`
	autocomplete set method { [
		{
			"FlagsDesc": {
				"define": "Define method"
			}
		}
	] }
`)
}

func cmdFunc(p *lang.Process) error {
	var dtParamsT []lang.MurexFuncParam

	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	blockId := 1
	if p.Parameters.Len() == 3 {
		blockId++

		dtParamsS, err := p.Parameters.String(1)
		if err != nil {
			return err
		}

		dtParamsT, err = lang.ParseMxFunctionParameters(dtParamsS)
		if err != nil {
			return fmt.Errorf("cannot parse function parameter block: %s:%scode: (%s)",
				//err.Error(), lang.EscapedColon, lang.EscapeColonInErr(dtParamsS))
				err.Error(), utils.NewLineString, dtParamsS)
		}
	}

	block, err := p.Parameters.Block(blockId)
	if err != nil {
		return err
	}

	switch {
	case len(name) == 0:
		return errors.New("function name is an empty (zero length) string")

	case strings.Contains(name, "$"):
		return errors.New("function name cannot contain a dollar, '$', character")

	default:
		lang.MxFunctions.Define(name, dtParamsT, block, p.FileRef)
		return nil
	}
}

func cmdUnfunc(p *lang.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	return lang.MxFunctions.Undefine(name)
}

func cmdPrivate(p *lang.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	switch {
	case len(name) == 0:
		return errors.New("private name is an empty (zero length) string")

	case strings.Contains(name, "$"):
		return errors.New("private name cannot contain a dollar, '$', character")

	default:
		lang.PrivateFunctions.Define(name, nil, block, p.FileRef)
		return nil
	}
}

func cmdUnprivate(p *lang.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	return lang.PrivateFunctions.Undefine(name, p.FileRef)
}

func cmdMethod(p *lang.Process) error {
	fn, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	switch fn {
	case "define":
		return cmdMethodDefine(p)
	default:
		return fmt.Errorf("invalid parameter `%s`", fn)
	}
}

type methodDefineT struct {
	Stdin  string
	Stdout string
}

func cmdMethodDefine(p *lang.Process) error {
	name, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	j, err := p.Parameters.String(2)
	if err != nil {
		return err
	}

	var mdt methodDefineT
	err = json.UnmarshalMurex([]byte(j), &mdt)
	if err != nil {
		return err
	}

	if mdt.Stdin != "" {
		lang.MethodStdin.Define(name, mdt.Stdin)
		err = lang.MethodStdin.Degroup()
		if err != nil {
			return err
		}
	}

	if mdt.Stdout != "" {
		lang.MethodStdout.Define(name, mdt.Stdout)
		err = lang.MethodStdout.Degroup()
		if err != nil {
			return err
		}
	}

	return nil
}
