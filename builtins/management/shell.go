package management

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
	"io/ioutil"
	"os"
	"sort"
)

func init() {
	proc.GoFunctions["args"] = proc.GoFunction{Func: cmdArgs, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["source"] = proc.GoFunction{Func: cmdSource, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["."] = proc.GoFunction{Func: cmdSource, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["autocomplete"] = proc.GoFunction{Func: cmdAutocomplete, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["version"] = proc.GoFunction{Func: cmdVersion, TypeIn: types.Null, TypeOut: types.String}
}

func cmdArgs(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Boolean)

	if p.Parameters.Len() != 1 {
		return errors.New("Invalid parameters! Expecting JSON input.")
	}

	var args parameters.Arguments
	err = json.Unmarshal(p.Parameters.ByteAll(), &args)
	if err != nil {
		return err
	}

	type flags struct {
		Self       string
		Flags      map[string]string
		Additional []string
		Error      string
	}
	var jObj flags

	//margs := flag.Args()
	//if len(margs) == 0 {
	//	return errors.New("Empty args. Was this run inside a murex shell script?")
	//}

	//jObj.Flags, jObj.Additional, err = parameters.ParseFlags(margs[1:], &args)
	jObj.Flags, jObj.Additional, err = parameters.ParseFlags(p.Scope.Parameters.Params, &args)
	if err != nil {
		jObj.Error = err.Error()
		p.ExitNum = 1
	}
	//jObj.Self = margs[0]
	jObj.Self = p.Scope.Name

	b, err := utils.JsonMarshal(jObj, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	err = proc.GlobalVars.Set("ARGS", string(b), types.Json)
	return err
}

func cmdSource(p *proc.Process) error {
	filename, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	p.ExitNum, err = lang.ProcessNewBlock([]rune(string(b)), nil, p.Stdout, p.Stderr, p)
	return err
}

func cmdAutocomplete(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	mode, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	switch mode {
	case "get":
		return listAutocomplete(p)
	case "set":
	default:
		return errors.New("Not a valid mode. Please use `get` or `set`.")
	}

	exe, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	jf, err := p.Parameters.Byte(2)
	if err != nil {
		return err
	}

	var flags shell.Flags
	err = json.Unmarshal(jf, &flags)
	if err != nil {
		return err
	}

	sort.Strings(flags.Flags)
	shell.ExesFlags[exe] = flags
	return nil
}

func listAutocomplete(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	b, err := utils.JsonMarshal(shell.ExesFlags, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}

func cmdVersion(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)
	_, err := p.Stdout.Writeln([]byte(config.AppName + ": " + config.Version))
	return err
}
