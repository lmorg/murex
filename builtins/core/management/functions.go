package management

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils"
	"os"
	"runtime"
	"strconv"
)

func init() {
	proc.GoFunctions["debug"] = cmdDebug
	proc.GoFunctions["exitnum"] = cmdExitNum
	proc.GoFunctions["config"] = cmdConfig
	proc.GoFunctions["builtins"] = cmdListBuiltins
	proc.GoFunctions["bexists"] = cmdBuiltinExists
	proc.GoFunctions["cd"] = cmdCd
	proc.GoFunctions["os"] = cmdOs
	proc.GoFunctions["cpuarch"] = cmdCpuArch
	proc.GoFunctions["cpucount"] = cmdCpuCount
	proc.GoFunctions["murex-update-exe-list"] = cmdUpdateExeList
}

func cmdDebug(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	if p.IsMethod {
		var (
			json map[string]interface{} = make(map[string]interface{})
			b    []byte
		)

		dt := p.Stdin.GetDataType()
		obj, err := define.UnmarshalData(p, dt)

		json["Process"] = *p.Previous
		json["DataType"] = map[string]string{
			"Murex": dt,
			"Go":    fmt.Sprintf("%T", obj),
		}

		b, err = utils.JsonMarshal(json, p.Stdout.IsTTY())
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)
		return err

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
		return
	}
}

func cmdExitNum(p *proc.Process) error {
	p.Stdout.SetDataType(types.Integer)
	p.Stdout.Writeln([]byte(strconv.Itoa(p.Previous.ExitNum)))
	return nil
}

func cmdListBuiltins(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	var s []string
	for name := range proc.GoFunctions {
		s = append(s, name)
	}

	b, err := utils.JsonMarshal(s, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
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
		if proc.GoFunctions[name] != nil {
			j.Installed = append(j.Installed, name)
		} else {
			j.Missing = append(j.Missing, name)
			p.ExitNum++
		}
	}

	b, err := utils.JsonMarshal(j, p.Stdout.IsTTY())
	p.Stdout.Writeln(b)

	return err
}

func cmdConfig(p *proc.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.Json)

		b, err := utils.JsonMarshal(proc.GlobalConf.Dump(), p.Stdout.IsTTY())
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
	if err != nil {
		return err
	}

	hist := proc.GlobalVars.GetString("PWDHIST")
	if hist == "" {
		hist = "[]"
	}

	var v []string
	err = json.Unmarshal([]byte(hist), &v)
	if err != nil {
		return errors.New("Unable to unpack $PWDHIST: " + err.Error())
	}

	pwd, err := os.Getwd()
	if err != nil {
		pwd = s
	}

	v = append(v, pwd)
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return errors.New("Unable to repack $PWDHIST: " + err.Error())
	}

	err = proc.GlobalVars.Set("PWDHIST", string(b), types.Json)
	return err
}

func cmdOs(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.String)
	_, err = p.Stdout.Write([]byte(runtime.GOOS))
	return
}

func cmdCpuArch(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.String)
	_, err = p.Stdout.Write([]byte(runtime.GOARCH))
	return
}

func cmdCpuCount(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Integer)
	_, err = p.Stdout.Write([]byte(strconv.Itoa(runtime.NumCPU())))
	return
}

func cmdUpdateExeList(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)
	autocomplete.UpdateGlobalExeList()
	return nil
}
