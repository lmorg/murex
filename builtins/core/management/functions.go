package management

import (
	corejson "encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/json"
)

func init() {
	proc.GoFunctions["debug"] = cmdDebug
	proc.GoFunctions["exitnum"] = cmdExitNum
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
			j = make(map[string]interface{})
			b []byte
		)

		dt := p.Stdin.GetDataType()
		obj, _ := define.UnmarshalData(p, dt) // For once we don't care about the error

		j["Process"] = *p.Previous // only making a readonly so the sync.Mutex is irrelevent here
		j["DataType"] = map[string]string{
			"Murex": dt,
			"Go":    fmt.Sprintf("%T", obj),
		}

		b, err = json.Marshal(j, p.Stdout.IsTTY())
		if err != nil {
			return err
		}

		_, err = p.Stdout.Writeln(b)
		return err

	}

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

	b, err := json.Marshal(s, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}

func cmdBuiltinExists(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)
	if p.Parameters.Len() == 0 {
		return errors.New("Missing parameters. Please name builtins you want to check")
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

	b, err := json.Marshal(j, p.Stdout.IsTTY())
	p.Stdout.Writeln(b)

	return err
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

	pwd, err := os.Getwd()
	if err != nil {
		p.Stderr.Writeln([]byte(err.Error()))
		pwd = s
	}

	// Update $PWD environmental variable for compatibility reasons
	err = os.Setenv("PWD", pwd)
	if err != nil {
		p.Stderr.Writeln([]byte(err.Error()))
	}

	// Update $PWDHIST murex variable - a more idiomatic approach to PWD
	hist := p.Variables.GetString("PWDHIST")
	if hist == "" {
		hist = "[]"
	}

	var v []string
	err = json.Unmarshal([]byte(hist), &v)
	if err != nil {
		return errors.New("Unable to unpack $PWDHIST: " + err.Error())
	}

	v = append(v, pwd)
	b, err := corejson.MarshalIndent(v, "", "    ")
	if err != nil {
		return errors.New("Unable to repack $PWDHIST: " + err.Error())
	}

	err = p.Variables.Set("PWDHIST", string(b), types.Json)
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
