package management

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/builtins/core/events"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils"
	"io/ioutil"
	"os"
	"sort"
)

func init() {
	proc.GoFunctions["args"] = cmdArgs
	proc.GoFunctions["params"] = cmdParams
	proc.GoFunctions["source"] = cmdSource
	proc.GoFunctions["."] = cmdSource
	proc.GoFunctions["autocomplete"] = cmdAutocomplete
	proc.GoFunctions["version"] = cmdVersion
	proc.GoFunctions["runtime"] = cmdRuntime
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

	err = proc.GlobalVars.Set("ARGS", string(b), types.Json)
	return err
}

func cmdParams(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	params := append([]string{p.Scope.Name}, p.Scope.Parameters.Params...)

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

	p.ExitNum, err = lang.ProcessNewBlock([]rune(string(b)), nil, p.Stdout, p.Stderr, p)
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

	//jf, err := p.Parameters.Byte(2)
	//if err != nil {
	//	return err
	//}

	var jf []byte
	jfr, err := p.Parameters.Block(2)
	if err == nil {
		jf = []byte(string(jfr))
	} else {
		jf, err = p.Parameters.Byte(2)
		if err!= nil {
			return err
		}
	}

	var flags []shell.Flags
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

func cmdRuntime(p *proc.Process) error {
	p.Stdout.SetDataType(types.Json)

	rt := make(map[string]interface{})
	rt["Vars"] = proc.GlobalVars.Dump()
	rt["Aliases"] = proc.GlobalAliases.Dump()
	rt["Config"] = proc.GlobalConf.Dump()
	rt["Pipes"] = proc.GlobalPipes.Dump()
	rt["Funcs"] = proc.MxFunctions.Dump()
	rt["Fids"] = proc.GlobalFIDs.Dump()
	rt["Arrays"] = streams.DumpArray()
	rt["Maps"] = streams.DumpMap()
	rt["Indexes"] = define.DumpIndex()
	rt["Marshallers"] = define.DumpMarshaller()
	rt["Unmarshallers"] = define.DumpUnmarshaller()
	rt["Mimes"] = define.DumpMime()
	rt["FileExts"] = define.DumpFileExts()
	rt["Events"] = events.DumpEvents()

	//var mem runtime.MemStats
	//runtime.ReadMemStats(&mem)
	//rt["MemStats"] = mem

	b, err := utils.JsonMarshal(rt, p.Stdout.IsTTY())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
