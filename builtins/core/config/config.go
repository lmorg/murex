package cmdconfig

import (
	"errors"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/alter"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.GoFunctions["config"] = cmdConfig
	lang.GoFunctions["!config"] = bangConfig
}

func cmdConfig(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.Json)

		b, err := json.Marshal(p.Config.Dump(), p.Stdout.IsTTY())
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)
		return err
	}

	option, _ := p.Parameters.String(0)
	switch option {
	case "get":
		return getConfig(p)

	case "set":
		return setConfig(p)

	case "alter":
		return alterConfig(p)

	case "define":
		return defineConfig(p)

	case "default":
		return defaultConfig(p)

	default:
		p.Stdout.SetDataType(types.Null)
		return errors.New("Unknown option. Please get, set, alter or define")
	}

}

func getConfig(p *lang.Process) error {
	app, _ := p.Parameters.String(1)
	key, _ := p.Parameters.String(2)
	val, err := p.Config.Get(app, key, types.String)
	if err != nil {
		return err
	}
	p.Stdout.SetDataType(p.Config.DataType(app, key))
	p.Stdout.Writeln([]byte(val.(string)))
	return nil
}

func setConfig(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)
	app, _ := p.Parameters.String(1)
	key, _ := p.Parameters.String(2)
	var val string

	if p.IsMethod {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		val = string(b)

	} else {
		val, _ = p.Parameters.String(3)
		//val = ansi.ExpandConsts(val)
	}

	err := p.Config.Set(app, key, val)
	return err
}

func alterConfig(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)

	app, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	key, err := p.Parameters.String(2)
	if err != nil {
		return err
	}

	path, err := p.Parameters.String(3)
	if err != nil {
		return err
	}

	new, err := p.Parameters.String(4)
	if err != nil {
		return err
	}

	splitPath, err := alter.SplitPath(path)
	if err != nil {
		return err
	}

	v, err := p.Config.Get(app, key, types.String)
	if err != nil {
		return err
	}

	dt := p.Config.DataType(app, key)
	//branch := p.BranchFID()
	//defer branch.Close()
	//branch.Stdin = streams.NewStdin()
	fork := p.Fork(lang.F_CREATE_STDIN)
	_, err = fork.Stdin.Write([]byte(v.(string)))
	if err != nil {
		return errors.New("Couldn't write to unmarshaller's buffer: " + err.Error())
	}

	v, err = define.UnmarshalData(fork.Process, dt)
	if err != nil {
		return errors.New("Couldn't unmarshal existing config: " + err.Error())
	}

	v, err = alter.Alter(p.Context, v, splitPath, new)
	if err != nil {
		return err
	}

	val, err := define.MarshalData(fork.Process, dt, v)
	if err != nil {
		return errors.New("Couldn't remarshal altered data structure: " + err.Error())
	}

	return p.Config.Set(app, key, val)
}

func defineConfig(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)
	app, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	key, err := p.Parameters.String(2)
	if err != nil {
		return err
	}

	var b []byte

	if p.IsMethod == true {
		b, err = p.Stdin.ReadAll()
		if err != nil {
			return err
		}

	} else {
		b, err = p.Parameters.Byte(3)
		if err != nil {
			return err
		}
	}

	var properties config.Properties
	err = json.UnmarshalMurex(b, &properties)
	if err != nil {
		return err
	}

	switch {
	case properties.DataType == "":
		return errors.New("`DataType` not defined")

	case properties.Description == "":
		return errors.New("`Description` not defined")

	case (properties.Dynamic.Read == "" && properties.Dynamic.Write != "") ||
		(properties.Dynamic.Read != "" && properties.Dynamic.Write == ""):
		return errors.New("When using dynamic values, both the `read` and `write` need to contain code blocks")

	case properties.Dynamic.Read != "" && !types.IsBlock([]byte(properties.Dynamic.Read)):
		return errors.New("Dynamic `read` is not a valid code block")

	case properties.Dynamic.Write != "" && !types.IsBlock([]byte(properties.Dynamic.Write)):
		return errors.New("Dynamic `write` is not a valid code block")
	}

	if properties.Dynamic.Read != "" {
		properties.Dynamic.GetDynamic = getDynamic([]rune(properties.Dynamic.Read))
		properties.Dynamic.SetDynamic = setDynamic([]rune(properties.Dynamic.Write))
	}

	lang.ShellProcess.Config.Define(app, key, properties)
	return nil
}

func bangConfig(p *lang.Process) error {
	app, _ := p.Parameters.String(0)
	key, _ := p.Parameters.String(1)
	err := p.Config.Default(app, key)
	return err
}

func defaultConfig(p *lang.Process) error {
	app, _ := p.Parameters.String(1)
	key, _ := p.Parameters.String(2)
	err := p.Config.Default(app, key)
	return err
}

func getDynamic(block []rune) func() (interface{}, error) {
	return func() (interface{}, error) {
		block = block[1 : len(block)-1]

		//branch := lang.ShellProcess.BranchFID()
		//branch.Scope = branch.Process
		//branch.Parent = branch.Process
		//branch.IsBackground = true

		//stdout := streams.NewStdin()
		//exitNum, err := lang.RunBlockNewConfigSpace(block, nil, stdout, lang.ShellProcess.Stderr, branch.Process)
		//branch.Close()

		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_NO_STDIN | lang.F_CREATE_STDOUT)
		exitNum, err := fork.Execute(block)

		if err != nil {
			return nil, errors.New("Dynamic config code could not compile: " + err.Error())
		}
		if exitNum != 0 && debug.Enabled {
			lang.ShellProcess.Stderr.Writeln([]byte("Dynamic config returned a none zero exit number." + utils.NewLineString))
		}

		b, err := fork.Stdout.ReadAll()
		if err != nil {
			return nil, err
		}

		return string(b), nil
	}
}

func setDynamic(block []rune) func(interface{}) error {
	return func(value interface{}) error {
		//if !types.IsBlock([]byte(stringblock)) {
		//	return nil, errors.New("Dynamic config reader is not a code block")
		//}
		block = block[1 : len(block)-1]

		//branch := lang.ShellProcess.BranchFID()
		//branch.Scope = branch.Process
		//branch.Parent = branch.Process
		//branch.IsBackground = true
		fork := lang.ShellProcess.Fork(lang.F_FUNCTION | lang.F_CREATE_STDIN)

		s, err := types.ConvertGoType(value, types.String)
		if err != nil {
			return err
		}

		_, err = fork.Stdin.Write([]byte(s.(string)))
		if err != nil {
			return err
		}

		//exitNum, err := lang.RunBlockNewConfigSpace(block, stdin, lang.ShellProcess.Stdout, lang.ShellProcess.Stderr, branch.Process)
		exitNum, err := fork.Execute(block)

		if err != nil {
			return errors.New("Dynamic config code could not compile: " + err.Error())
		}
		if exitNum != 0 && debug.Enabled {
			lang.ShellProcess.Stderr.Writeln([]byte("Dynamic config returned a none zero exit number." + utils.NewLineString))
		}

		return nil
	}
}
