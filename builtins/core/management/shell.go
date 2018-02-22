package management

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"runtime"
	"sort"

	"github.com/lmorg/murex/builtins/core/events"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils"
)

func init() {
	proc.GoFunctions["args"] = cmdArgs
	proc.GoFunctions["params"] = cmdParams
	proc.GoFunctions["source"] = cmdSource
	proc.GoFunctions["."] = cmdSource
	proc.GoFunctions["autocomplete"] = cmdAutocomplete
	proc.GoFunctions["version"] = cmdVersion
	proc.GoFunctions["runtime"] = cmdRuntime
	proc.GoFunctions["murex-runtime"] = cmdRuntime
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

	jObj.Flags, jObj.Additional, err = parameters.ParseFlags(p.Scope.Parameters.Params, &args)
	if err != nil {
		jObj.Error = err.Error()
		p.ExitNum = 1
	}

	jObj.Self = p.Scope.Name

	b, err := utils.JsonMarshal(jObj, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	err = p.ScopedVars.Set("ARGS", string(b), types.Json)
	return err
}

func cmdParams(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	params := append([]string{p.Scope.Name}, p.Scope.Parameters.Params...)

	debug.Json("builtin.params:", params)

	b, err := utils.JsonMarshal(&params, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
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

	p.RunMode = runmode.Shell

	p.ExitNum, err = lang.RunBlockShellNamespace([]rune(string(b)), nil, p.Stdout, p.Stderr)
	return err
}

func cmdAutocomplete(p *proc.Process) error {
	mode, err := p.Parameters.String(0)
	if err != nil {
		p.Stdout.SetDataType(types.Null)
		return err
	}

	switch mode {
	case "get":
		p.Stdout.SetDataType(types.Json)
		return listAutocomplete(p)
	case "set":
	default:
		p.Stdout.SetDataType(types.Null)
		return errors.New("Not a valid mode. Please use `get` or `set`.")
	}

	p.Stdout.SetDataType(types.Null)

	exe, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	var jf []byte

	if p.IsMethod {

		jf, err = p.Stdin.ReadAll()
		if err != nil {
			return err
		}

	} else {

		jfr, err := p.Parameters.Block(2)
		if err == nil {
			jf = []byte(string(jfr))
		} else {
			jf, err = p.Parameters.Byte(2)
			if err != nil {
				return err
			}
		}
	}

	var flags []autocomplete.Flags
	err = json.Unmarshal(jf, &flags)
	if err != nil {
		return err
	}

	for i := range flags {
		// So we don't have nil values in JSON
		if len(flags[i].Flags) == 0 {
			flags[i].Flags = make([]string, 0)
		}

		sort.Strings(flags[i].Flags)
	}

	autocomplete.ExesFlags[exe] = flags
	return nil
}

func listAutocomplete(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	b, err := utils.JsonMarshal(autocomplete.ExesFlags, p.Stdout.IsTTY())
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

func cmdRuntime(p *proc.Process) error {
	const (
		fVars          = "--vars"
		fAliases       = "--aliases"
		fConfig        = "--config"
		fPipes         = "--pipes"
		fFuncs         = "--funcs"
		fFids          = "--fids"
		fArrays        = "--arrays"
		fMaps          = "--maps"
		fIndexes       = "--indexes"
		fMarshallers   = "--marshallers"
		fUnmarshallers = "--unmarshallers"
		fEvents        = "--events"
		fFlags         = "--flags"
		fMemstats      = "--memstats"
	)
	p.Stdout.SetDataType(types.Json)

	f, _, err := p.Parameters.ParseFlags(
		&parameters.Arguments{
			Flags: map[string]string{
				//"all":           types.Boolean,
				fVars:          types.Boolean,
				fAliases:       types.Boolean,
				fConfig:        types.Boolean,
				fPipes:         types.Boolean,
				fFuncs:         types.Boolean,
				fFids:          types.Boolean,
				fArrays:        types.Boolean,
				fMaps:          types.Boolean,
				fIndexes:       types.Boolean,
				fMarshallers:   types.Boolean,
				fUnmarshallers: types.Boolean,
				fEvents:        types.Boolean,
				fFlags:         types.Boolean,
				fMemstats:      types.Boolean,
			},
			AllowAdditional: false,
		},
	)

	if err != nil {
		return err
	}

	if len(f) == 0 {
		return errors.New("Please include one or more parameters.")
	}

	ret := make(map[string]interface{})
	for flag := range f {
		switch flag {
		case fVars:
			ret[fVars[2:]] = proc.ShellProcess.ScopedVars.Dump()
		case fAliases:
			ret[fAliases[2:]] = proc.GlobalAliases.Dump()
		case fConfig:
			ret[fConfig[2:]] = proc.ShellProcess.Config.Dump()
		case fPipes:
			ret[fPipes[2:]] = proc.GlobalPipes.Dump()
		case fFuncs:
			ret[fFuncs[2:]] = proc.MxFunctions.Dump()
		case fFids:
			ret[fFids[2:]] = proc.GlobalFIDs.Dump()
		case fArrays:
			ret[fArrays[2:]] = streams.DumpArray()
		case fMaps:
			ret[fMaps[2:]] = streams.DumpMap()
		case fIndexes:
			ret[fIndexes[2:]] = define.DumpIndex()
		case fMarshallers:
			ret[fMarshallers[2:]] = define.DumpMarshaller()
		case fUnmarshallers:
			ret[fUnmarshallers[2:]] = define.DumpUnmarshaller()
		case fEvents:
			ret[fEvents[2:]] = events.DumpEvents()
		case fFlags:
			ret[fFlags[2:]] = autocomplete.ExesFlags
		case fMemstats:
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			ret[fMemstats[2:]] = mem
		default:
			return errors.New("Unrecognised parameter: " + flag)
		}
	}

	var b []byte
	if len(ret) == 1 {
		var obj interface{}
		for _, obj = range ret {
		}
		b, err = utils.JsonMarshal(obj, p.Stdout.IsTTY())
		if err != nil {
			return err
		}

	} else {
		b, err = utils.JsonMarshal(ret, p.Stdout.IsTTY())
		if err != nil {
			return err
		}
	}

	_, err = p.Stdout.Write(b)
	return err
}
