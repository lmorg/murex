package shellconfig

import (
	"errors"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func init() {
	proc.GoFunctions["config"] = cmdConfig
}

func cmdConfig(p *proc.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.Json)

		b, err := utils.JsonMarshal(p.Config.Dump(), p.Stdout.IsTTY())
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)
		return err
	}

	option, _ := p.Parameters.String(0)
	switch option {
	case "get":
		return get(p)

	case "set":
		return set(p)

	case "define":
		return define(p)

	default:
		p.Stdout.SetDataType(types.Null)
		return errors.New("Unknown option. Please get, set or define.")
	}

	return nil
}

func get(p *proc.Process) error {
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

func set(p *proc.Process) error {
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

func define(p *proc.Process) error {
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
	err = utils.JsonUnmarshal(b, &properties)
	if err != nil {
		return err
	}

	if properties.DataType == "" {
		return errors.New("`DataType` not defined.")
	}
	if properties.Description == "" {
		return errors.New("`Description` not defined.")
	}

	proc.ShellProcess.Config.Define(app, key, properties)
	return nil
}
