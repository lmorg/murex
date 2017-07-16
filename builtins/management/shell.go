package management

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
	"io"
	"os"
)

func init() {
	proc.GoFunctions["history"] = proc.GoFunction{Func: cmdHistory, TypeIn: types.Null, TypeOut: types.Json}
	proc.GoFunctions["args"] = proc.GoFunction{Func: cmdArgs, TypeIn: types.Null, TypeOut: types.Json}
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
