package shellpipe

import (
	"errors"
	"io"

	"github.com/lmorg/murex/lang/proc/stdio"

	"github.com/lmorg/murex/config/defaults"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
)

func init() {
	proc.GoFunctions["pipe"] = cmdPipe
	proc.GoFunctions["!pipe"] = cmdClosePipe
	proc.GoFunctions[consts.NamedPipeProcName] = cmdReadPipe
	defaults.AppendProfile(`
		autocomplete set pipe { [
		    {
		        "Dynamic": ({
					runtime --pipes -> !match std
				}),
		        "FlagValues": {
		         	"--file": [{
		        		"IncFiles": true
		        	}]
				}
		    }
		] }

		autocomplete set !pipe { [
		    {
		        "Dynamic": "{ runtime: --pipes -> formap k v { if { = k!=` + "`null`" + ` } { $k } } }",
				"AllowMultiple": true
		    }
		] }
	`)
}

func cmdPipe(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing parameters")
	}

	// import the registered pipes
	supportedFlags := make(map[string]string)
	pipes := stdio.DumpPipes()
	for i := range pipes {
		supportedFlags["--"+pipes[i]] = types.String
	}

	// define cli flags
	flags, additional, err := p.Parameters.ParseFlags(&parameters.Arguments{
		AllowAdditional: true,
		Flags:           supportedFlags,
	})

	if err != nil {
		return err
	}

	if len(additional) == 0 {
		return errors.New("No name specified for named pipe. Usage: `pipe name [ --pipe-type creation-data ]")
	}

	if len(flags) > 1 {
		return errors.New("Too many types of pipe specified. Please use only one flag per")
	}

	for flag := range flags {
		return proc.GlobalPipes.CreatePipe(additional[0], flag[2:], flags[flag])
	}

	for _, name := range additional {
		err := proc.GlobalPipes.CreatePipe(name, "std", "")
		if err != nil {
			return err
		}
	}

	return nil
}

func cmdClosePipe(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	var names []string

	if p.IsMethod {
		p.Stdin.ReadArray(func(b []byte) {
			names = append(names, string(b))
		})

		if len(names) == 0 {
			return errors.New("Stdin contained a zero lengthed array.")
		}

	} else {
		if p.Parameters.Len() == 0 {
			return errors.New("No pipes listed for closing.")
		}

		names = p.Parameters.StringArray()
	}

	for _, name := range names {
		if err := proc.GlobalPipes.Close(name); err != nil {
			return err
		}
	}

	return nil
}

func cmdReadPipe(p *proc.Process) error {
	name, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	pipe, err := proc.GlobalPipes.Get(name)
	if err != nil {
		return err
	}

	if p.IsMethod {
		pipe.SetDataType(p.Stdin.GetDataType())
		_, err = io.Copy(pipe, p.Stdin)
		return err
	}

	p.Stdout.SetDataType(pipe.GetDataType())
	_, err = io.Copy(p.Stdout, pipe)
	return err
}
