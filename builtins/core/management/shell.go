package management

import (
	"errors"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/runmode"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/parser"
)

func init() {
	proc.GoFunctions["args"] = cmdArgs
	proc.GoFunctions["params"] = cmdParams
	proc.GoFunctions["source"] = cmdSource
	proc.GoFunctions["."] = cmdSource
	proc.GoFunctions["version"] = cmdVersion
	proc.GoFunctions["runtime"] = cmdRuntime
	proc.GoFunctions["murex-parser"] = cmdParser
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

	b, err := json.Marshal(jObj, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	err = p.Variables.Set("ARGS", string(b), types.Json)
	return err
}

func cmdParams(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	params := append([]string{p.Scope.Name}, p.Scope.Parameters.Params...)

	//debug.Json("builtin.params:", params)

	b, err := json.Marshal(&params, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdSource(p *proc.Process) error {
	var block []rune

	if p.IsMethod {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		block = []rune(string(b))

	} else {
		var err error
		block, err = p.Parameters.Block(0)

		if err != nil {
			// get block from file
			name, err := p.Parameters.String(0)
			if err != nil {
				return err
			}

			file, err := os.Open(name)
			if err != nil {
				return err
			}

			b, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			block = []rune(string(b))
		}

	}

	var err error
	p.RunMode = runmode.Shell
	p.ExitNum, err = lang.RunBlockShellConfigSpace(block, nil, p.Stdout, p.Stderr)
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
		fAllVars       = "--all-vars"
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
				fAllVars:       types.Boolean,
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
			ret[fVars[2:]] = p.Variables.Dump()
		case fAllVars:
			ret[fAllVars[2:]] = p.Variables.DumpEntireTable()
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
		b, err = json.Marshal(obj, p.Stdout.IsTTY())
		if err != nil {
			return err
		}

	} else {
		b, err = json.Marshal(ret, p.Stdout.IsTTY())
		if err != nil {
			return err
		}
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdParser(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	var (
		block []rune
		pos   int
		//syntaxHighlight bool
	)

	if p.IsMethod {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		block = []rune(string(b))

		//syntaxHighlight, _ = p.Parameters.Bool(0)

	} else {
		r, err := p.Parameters.Block(0)
		if err != nil {
			return err
		}
		block = r
		pos, _ = p.Parameters.Int(1)
	}

	pt, _ /*ansiHighlighted*/ := parser.Parse(block, pos)

	/*if syntaxHighlight {
		_, err := p.Stdout.Write([]byte(ansiHighlighted))
		return err
	}*/

	b, err := json.Marshal(pt, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
