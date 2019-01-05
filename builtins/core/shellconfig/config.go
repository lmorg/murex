package shellconfig

import (
	"errors"

	"github.com/lmorg/murex/builtins/pipes/streams"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils/alter"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	proc.GoFunctions["config"] = cmdConfig
}

func cmdConfig(p *proc.Process) error {
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

	default:
		p.Stdout.SetDataType(types.Null)
		return errors.New("Unknown option. Please get, set, alter or define")
	}

}

func getConfig(p *proc.Process) error {
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

func setConfig(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)
	app, _ := p.Parameters.String(1)
	key, _ := p.Parameters.String(2)
	var val string

	if p.IsMethod == true {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		val = string(b)

	} else {
		val, _ = p.Parameters.String(3)
	}

	err := p.Config.Set(app, key, val)
	return err
}

func alterConfig(p *proc.Process) error {
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
	branch := p.BranchFID()
	defer branch.Close()
	branch.Stdin = streams.NewStdin()
	_, err = branch.Stdin.Write([]byte(v.(string)))
	if err != nil {
		return errors.New("Couldn't write to unmarshaller's buffer: " + err.Error())
	}

	v, err = define.UnmarshalData(branch.Process, dt)
	if err != nil {
		return errors.New("Couldn't unmarshal existing config: " + err.Error())
	}

	v, err = alter.Alter(p.Context, v, splitPath, new)
	if err != nil {
		return err
	}

	val, err := define.MarshalData(branch.Process, dt, v)
	if err != nil {
		return errors.New("Couldn't remarshal altered data structure: " + err.Error())
	}

	return p.Config.Set(app, key, val)
}

func defineConfig(p *proc.Process) error {
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

	if properties.DataType == "" {
		return errors.New("`DataType` not defined")
	}
	if properties.Description == "" {
		return errors.New("`Description` not defined")
	}

	proc.ShellProcess.Config.Define(app, key, properties)
	return nil
}
