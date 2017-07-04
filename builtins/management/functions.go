package management

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"os"
	"strconv"
)

func init() {
	proc.GoFunctions["debug"] = proc.GoFunction{Func: cmdDebug, TypeIn: types.Generic, TypeOut: types.Json}
	proc.GoFunctions["exitnum"] = proc.GoFunction{Func: cmdExitNum, TypeIn: types.Generic, TypeOut: types.Integer}
	proc.GoFunctions["config"] = proc.GoFunction{Func: cmdConfig, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["builtins"] = proc.GoFunction{Func: cmdListBuiltins, TypeIn: types.Null, TypeOut: types.String}
	proc.GoFunctions["bexists"] = proc.GoFunction{Func: cmdBuiltinExists, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["cd"] = proc.GoFunction{Func: cmdCd, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["alias"] = proc.GoFunction{Func: cmdAlias, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["!alias"] = proc.GoFunction{Func: cmdUnalias, TypeIn: types.Null, TypeOut: types.Null}
}

func cmdDebug(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	if p.IsMethod {
		var (
			obj proc.Process = *p.Previous
			b   []byte
		)

		b, err = json.MarshalIndent(obj, "", "\t")
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)

	} else {

		var v bool
		v, err = p.Parameters.Bool(0)

		if err != nil {
			p.Stdout.Writeln(types.FalseByte)
			p.ExitNum = 1
			return nil
		}
		debug.Enable = v
		if v == false {
			p.Stdout.Writeln(types.FalseByte)
			p.ExitNum = 1
			return nil
		}

		_, err = p.Stdout.Writeln(types.TrueByte)
	}

	return
}

func cmdExitNum(p *proc.Process) error {
	p.Stdout.SetDataType(types.Integer)
	p.Stdout.Writeln([]byte(strconv.Itoa(p.Previous.ExitNum)))
	return nil
}

func cmdListBuiltins(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)
	//for name := range proc.GoFunctions {
	//	p.Stdout.Writeln([]byte(name))
	//}
	utils.JsonMarshal(proc.GoFunctions)
	return nil
}

func cmdBuiltinExists(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)
	if p.Parameters.Len() == 0 {
		return errors.New("Missing parameters. Please name builtins you want to check.")
	}

	var j struct {
		Installed []string
		Missing   []string
	}

	for _, name := range p.Parameters.StringArray() {
		if proc.GoFunctions[name].Func != nil {
			j.Installed = append(j.Installed, name)
		} else {
			j.Missing = append(j.Missing, name)
			p.ExitNum++
		}
	}

	b, err := json.MarshalIndent(j, "", "\t")
	p.Stdout.Writeln(b)

	return err
}

func cmdConfig(p *proc.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.Json)

		b, err := json.MarshalIndent(proc.GlobalConf.Dump(), "", "\t")
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)
		return err
	}

	option, _ := p.Parameters.String(0)
	switch option {
	case "get":
		app, _ := p.Parameters.String(1)
		key, _ := p.Parameters.String(2)
		val, err := proc.GlobalConf.Get(app, key, types.String)
		if err != nil {
			return err
		}
		p.Stdout.SetDataType(proc.GlobalConf.DataType(app, key))
		p.Stdout.Writeln([]byte(val.(string)))

	case "set":
		p.Stdout.SetDataType(types.Null)
		app, _ := p.Parameters.String(1)
		key, _ := p.Parameters.String(2)
		val, _ := p.Parameters.String(3)
		err := proc.GlobalConf.Set(app, key, val)
		return err

		/*case "stdin":
		err := proc.GlobalConf.Set(p.Parameters.String(1), p.Parameters.String(2), p.Stdin.ReadAll())
		return err*/
	default:
		p.Stdout.SetDataType(types.Null)
		return errors.New("Unknown option. Please get or set.")
	}

	return nil
}

func cmdCd(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)
	s, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	err = os.Chdir(s)
	return err
}

func cmdAlias(p *proc.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.Json)
		b, err := utils.JsonMarshal(proc.GlobalAliases.Dump())
		if err != nil {
			return err
		}
		_, err = p.Stdout.Writeln(b)
		return err

	}

	p.Stdout.SetDataType(types.Null)

	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	block, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	proc.GlobalAliases.Add(name, block)
	return nil
}

func cmdUnalias(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	for _, name := range p.Parameters.StringArray() {
		err := proc.GlobalAliases.Delete(name)
		if err != nil {
			return err
		}
	}
	return nil
}
