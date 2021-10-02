package cmdconfig

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	lang.DefineFunction("config", cmdConfig, types.Any)
	lang.DefineFunction("!config", bangConfig, types.Null)
}

func cmdConfig(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		return allConfig(p)
	}

	option, _ := p.Parameters.String(0)
	switch option {
	case "get":
		return getConfig(p)

	case "set":
		return setConfig(p)

	case "eval":
		return evalConfig(p)

	case "define":
		return defineConfig(p)

	case "default":
		return defaultConfig(p)

	default:
		p.Stdout.SetDataType(types.Null)
		return errors.New("Unknown option. Please get, set, eval, default or define")
	}
}

func allConfig(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	b, err := json.Marshal(p.Config.DumpConfig(), p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
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
	}

	return p.Config.Set(app, key, val)
}

func defaultConfig(p *lang.Process) error {
	app, _ := p.Parameters.String(1)
	key, _ := p.Parameters.String(2)
	return p.Config.Default(app, key)
}

func bangConfig(p *lang.Process) error {
	app, _ := p.Parameters.String(0)
	key, _ := p.Parameters.String(1)
	return p.Config.Default(app, key)
}
