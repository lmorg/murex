package structs

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("alias", cmdAlias, types.Null)
	lang.DefineFunction("!alias", cmdUnalias, types.Null)
	lang.DefineFunction("function", cmdFunc, types.Null)
	lang.DefineFunction("!function", cmdUnfunc, types.Null)
	lang.DefineFunction("private", cmdPrivate, types.Null)
	//lang.DefineFunction("!private", cmdUnprivate, types.Null)
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

var rxAlias = regexp.MustCompile(`^([-_.a-zA-Z0-9]+)=(.*?)$`)

func cmdAlias(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.Json)
		b, err := json.Marshal(lang.GlobalAliases.Dump(), p.Stdout.IsTTY())
		if err != nil {
			return err
		}
		_, err = p.Stdout.Writeln(b)
		return err

	}

	p.Stdout.SetDataType(types.Null)

	s, _ := p.Parameters.String(0)
	eq, _ := p.Parameters.String(1)

	if !rxAlias.MatchString(s) && len(eq) > 0 && eq[0] != '=' {
		return errors.New("invalid syntax. Expecting `alias new_name=original_name parameter1 parameter2 ...`")
	}

	var (
		split  = rxAlias.FindStringSubmatch(s)
		name   string
		params []string
	)

	if len(split) == 0 {
		name = s
		params = p.Parameters.StringArray()[1:]
		switch {
		case len(params) == 0:
			return fmt.Errorf("no command supplied")
		case len(params[0]) == 1 && params[0] == "=":
			params = params[1:]
		case len(params[0]) > 0 && params[0][0] == '=':
			params[0] = params[0][1:]
		default:
			return fmt.Errorf("unknown error. Please check syntax follows `alias new_name=original_name parameter1 parameter2 ...`")
		}

	} else {
		name = split[1]
		params = append([]string{split[2]}, p.Parameters.StringArray()[1:]...)
	}

	if len(params) == 0 {
		return fmt.Errorf("no command supplied")
	}

	if params[0] == "" && len(params) > 0 {
		params = params[1:]
	}

	if len(params) == 0 || params[0] == "" {
		return fmt.Errorf("no command supplied")
	}

	lang.GlobalAliases.Add(name, params)
	return nil
}

func cmdUnalias(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	for _, name := range p.Parameters.StringArray() {
		err := lang.GlobalAliases.Delete(name)
		if err != nil {
			return err
		}
	}
	return nil
}

func cmdFunc(p *lang.Process) error {
	var dtParamsT []lang.MxFunctionParams

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

/*func cmdUnprivate(p *lang.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	return lang.PrivateFunctions.Undefine(name)
}*/

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
		lang.MethodStdin.Degroup()
	}

	if mdt.Stdout != "" {
		lang.MethodStdout.Define(name, mdt.Stdout)
		lang.MethodStdout.Degroup()
	}

	return nil
}
