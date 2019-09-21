package management

import (
	corejson "encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/escape"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/man"
	"github.com/lmorg/murex/utils/posix"
)

func init() {
	lang.GoFunctions["debug"] = cmdDebug
	lang.GoFunctions["exitnum"] = cmdExitNum
	lang.GoFunctions["builtins"] = cmdListBuiltins
	lang.GoFunctions["bexists"] = cmdBuiltinExists
	lang.GoFunctions["cd"] = cmdCd
	lang.GoFunctions["os"] = cmdOs
	lang.GoFunctions["cpuarch"] = cmdCpuArch
	lang.GoFunctions["cpucount"] = cmdCpuCount
	lang.GoFunctions["murex-update-exe-list"] = cmdUpdateExeList
	lang.GoFunctions["man-summary"] = cmdManSummary
	lang.GoFunctions["esccli"] = cmdEscapeCli
}

func cmdDebug(p *lang.Process) (err error) {
	p.Stdout.SetDataType(types.Json)
	if p.IsMethod {
		var (
			j = make(map[string]interface{})
			b []byte
		)

		dt := p.Stdin.GetDataType()
		obj, _ := lang.UnmarshalData(p, dt) // For once we don't care about the error

		j["Process"] = *p.Previous // only making a readonly so the sync.Mutex is irrelevant here
		j["Data-Type"] = map[string]string{
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
	return
}

func cmdExitNum(p *lang.Process) error {
	p.Stdout.SetDataType(types.Integer)
	p.Stdout.Writeln([]byte(strconv.Itoa(p.Previous.ExitNum)))
	return nil
}

func cmdListBuiltins(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)

	var s []string
	for name := range lang.GoFunctions {
		s = append(s, name)
	}

	b, err := json.Marshal(s, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}

func cmdBuiltinExists(p *lang.Process) error {
	p.Stdout.SetDataType(types.Json)
	if p.Parameters.Len() == 0 {
		return errors.New("Missing parameters. Please name builtins you want to check")
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
		return errors.New("Parameter expected - name of executable")
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

func cmdEscapeCli(p *lang.Process) error {
	p.Stdout.SetDataType(types.String)

	var s []string

	if p.IsMethod {
		err := p.Stdin.ReadArray(func(b []byte) {
			s = append(s, string(b))
		})
		if err != nil {
			return err
		}
	} else {
		s = p.Parameters.StringArray()
	}

	escape.CommandLine(s)

	_, err := p.Stdout.Writeln([]byte(strings.Join(s, " ")))
	return err
}
