package management

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/cd"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/man"
	"github.com/lmorg/murex/utils/posix"
)

func init() {
	//lang.GoFunctions["debug"] = cmdDebug
	lang.DefineMethod("debug", cmdDebug, types.Any, types.Json)
	lang.DefineFunction("exitnum", cmdExitNum, types.Integer)
	lang.DefineFunction("bexists", cmdBuiltinExists, types.Json)
	lang.DefineFunction("cd", cmdCd, types.Null)
	lang.DefineFunction("os", cmdOs, types.String)
	lang.DefineFunction("cpuarch", cmdCpuArch, types.String)
	lang.DefineFunction("cpucount", cmdCpuCount, types.Integer)
	lang.DefineFunction("murex-update-exe-list", cmdUpdateExeList, types.Null)
	lang.DefineFunction("man-summary", cmdManSummary, types.Null)
}

func cmdDebug(p *lang.Process) error {
	if p.IsMethod {
		return cmdDebugMethod(p)
	}

	p.Stdout.SetDataType(types.Boolean)

	v, err := p.Parameters.Bool(0)

	if err != nil {
		_, err = p.Stdout.Write([]byte(fmt.Sprint(debug.Enabled)))
		return err
	}
	debug.Enabled = v
	if !v {
		p.Stdout.Writeln(types.FalseByte)
		p.ExitNum = 1
		return nil
	}

	_, err = p.Stdout.Writeln(types.TrueByte)
	return err
}

func cmdDebugMethod(p *lang.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(types.Json)

	var (
		j = make(map[string]interface{})
		b []byte
	)

	obj, err := lang.UnmarshalData(p, dt)

	j["Process"] = p.Previous
	j["Data-Type"] = map[string]string{
		"Murex":             dt,
		"Go":                fmt.Sprintf("%T", obj),
		"UnmarshalData Err": fmt.Sprint(err),
	}

	b, err = json.Marshal(j, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}

func cmdExitNum(p *lang.Process) error {
	p.Stdout.SetDataType(types.Integer)
	p.Stdout.Writeln([]byte(strconv.Itoa(p.Previous.ExitNum)))
	return nil
}

func cmdBuiltinExists(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)
	if p.Parameters.Len() == 0 {
		return errors.New("missing parameters. Please name builtins you want to check")
	}

	var j struct {
		Installed []string
		Missing   []string
	}

	for _, name := range p.Parameters.StringArray() {
		if lang.GoFunctions[name] != nil {
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

func cmdCd(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)
	path, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	err = cd.Chdir(p, path)

	return err
}

func cmdOs(p *lang.Process) error {
	if p.Parameters.Len() == 0 {
		p.Stdout.SetDataType(types.String)
		_, err := p.Stdout.Write([]byte(runtime.GOOS))
		return err
	}

	for _, os := range p.Parameters.StringArray() {
		if os == runtime.GOOS || (os == "posix" && posix.IsPosix()) {
			_, err := p.Stdout.Write(types.TrueByte)
			return err
		}
	}

	p.ExitNum = 1
	_, err := p.Stdout.Write(types.FalseByte)
	return err
}

func cmdCpuArch(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.String)
	_, err = p.Stdout.Write([]byte(runtime.GOARCH))
	return
}

func cmdCpuCount(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Integer)
	_, err = p.Stdout.Write([]byte(strconv.Itoa(runtime.NumCPU())))
	return
}

func cmdUpdateExeList(p *lang.Process) error {
	p.Stdout.SetDataType(types.Null)
	autocomplete.UpdateGlobalExeList()
	return nil
}

func cmdManSummary(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.String)

	if p.Parameters.Len() == 0 {
		return errors.New("parameter expected - name of executable")
	}

	exes := p.Parameters.StringArray()

	for _, exe := range exes {
		paths := man.GetManPages(exe)
		if len(paths) == 0 {
			p.Stderr.Writeln([]byte(exe + " - no man page exists"))
			continue
		}

		s := man.ParseSummary(paths)
		if s == "" {
			p.Stderr.Writeln([]byte(exe + " - unable to parse summary"))
			continue
		}

		_, err := p.Stdout.Writeln([]byte(s))
		if err != nil {
			return err
		}
	}

	return nil
}
