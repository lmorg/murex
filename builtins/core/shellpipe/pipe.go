package shellpipe

import (
	"errors"
	"io"

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
		        "Flags": [ "--file", "--udp-dial", "--tcp-dial", "--udp-listen", "--tcp-listen" ],
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
		return errors.New("Missing parameters.")
	}

	flags, additional, err := p.Parameters.ParseFlags(&parameters.Arguments{
		AllowAdditional: true,
		Flags: map[string]string{
			"--file":       types.String,
			"--tcp-dial":   types.String,
			"--udp-dial":   types.String,
			"--tcp-listen": types.String,
			"--udp-listen": types.String,
		},
	})

	if err != nil {
		return err
	}

	switch {
	case flags["--file"] != "":
		err = proc.GlobalPipes.CreateFile(flags["--create"], flags["--file"])

	case flags["--udp-dial"] != "":
		err = proc.GlobalPipes.CreateDialer(flags["--create"], "udp", flags["--udp-dial"])

	case flags["--tcp-dial"] != "":
		err = proc.GlobalPipes.CreateDialer(flags["--create"], "tcp", flags["--tcp-dial"])

	case flags["--udp-listen"] != "":
		err = proc.GlobalPipes.CreateListener(flags["--create"], "udp", flags["--udp-listen"])

	case flags["--tcp-listen"] != "":
		err = proc.GlobalPipes.CreateListener(flags["--create"], "tcp", flags["--tcp-listen"])

	case len(additional) > 0:
		for _, name := range additional {
			err := proc.GlobalPipes.CreatePipe(name)
			if err != nil {
				return err
			}
		}

	}

	return err
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
