package management

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
	"io"
	"io/ioutil"
	"os"
	"sort"
)

func init() {
	proc.GoFunctions["history"] = proc.GoFunction{Func: cmdHistory, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["args"] = proc.GoFunction{Func: cmdArgs, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["fork"] = proc.GoFunction{Func: cmdFork, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["source"] = proc.GoFunction{Func: cmdSource, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["."] = proc.GoFunction{Func: cmdSource, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["set-cached-flags"] = proc.GoFunction{Func: cmdSetFlags, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["list-cached-flags"] = proc.GoFunction{Func: cmdListFlags, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["version"] = proc.GoFunction{Func: cmdVersion, TypeIn: types.Null, TypeOut: types.String}
}

func cmdHistory(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.String)
	if shell.Instance == nil {
		return errors.New("This is only designed to be run when the shell is in interactive mode.")
	}

	file, err := os.Open(shell.Instance.Config.HistoryFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(p.Stdout, file)
	return err
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

	margs := flag.Args()
	jObj.Flags, jObj.Additional, err = parameters.ParseFlags(margs[1:], &args)
	if err != nil {
		jObj.Error = err.Error()
		p.ExitNum = 1
	}
	jObj.Self = margs[0]

	b, err := utils.JsonMarshal(jObj)
	if err != nil {
		return err
	}

	err = proc.GlobalVars.Set("ARGS", string(b), types.Json)
	return err
}

func cmdFork(p *proc.Process) (err error) {
	block, err := p.Parameters.Block(0)
	if err != nil {
		return err
	}

	go lang.ProcessNewBlock(block, p.Stdin, p.Stdout, p.Stderr, "fork")

	return
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

	p.ExitNum, err = lang.ProcessNewBlock([]rune(string(b)), nil, p.Stdout, p.Stderr, "source")
	return err
}

func cmdListFlags(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	b, err := utils.JsonMarshal(shell.ExesFlags)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Writeln(b)
	return err
}

func cmdSetFlags(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	exe, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	jf, err := p.Parameters.Block(1)
	if err != nil {
		return err
	}

	flags := make(shell.Flags, 0)
	err = json.Unmarshal([]byte(string(jf)), &flags)
	if err != nil {
		return err
	}

	sort.Strings(flags)
	shell.ExesFlags[exe] = flags
	return nil
}

func cmdVersion(p *proc.Process) error {
	p.Stdout.SetDataType(types.String)
	_, err := p.Stdout.Writeln([]byte(config.AppName + ": " + config.Version))
	return err
}
