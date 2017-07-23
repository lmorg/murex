package io

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
	"io"
)

func init() {
	proc.GoFunctions["pipe"] = proc.GoFunction{Func: cmdPipe, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions[consts.NamedPipeProcName] = proc.GoFunction{Func: cmdReadPipe, TypeIn: types.Null, TypeOut: types.Generic}
}

func cmdPipe(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	flags, _, err := p.Parameters.ParseFlags(&parameters.Arguments{
		AllowAdditional: false,
		Flags: map[string]string{
			"--create":     types.String,
			"-c":           "--create",
			"--close":      types.String,
			"-x":           "--close",
			"--file":       types.String,
			"--f":          "--file",
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
	case flags["--close"] != "":
		if err := proc.GlobalPipes.Close(flags["--close"]); err != nil {
			return err
		}

	case flags["--create"] != "":
		switch {
		case flags["--file"] != "":
			if err := proc.GlobalPipes.CreateFile(flags["--create"], flags["--file"]); err != nil {
				return err
			}

		case flags["--udp-dial"] != "":
			if err := proc.GlobalPipes.CreateDialer(flags["--create"], "udp", flags["--udp-dial"]); err != nil {
				return err
			}

		case flags["--tcp-dial"] != "":
			if err := proc.GlobalPipes.CreateDialer(flags["--create"], "tcp", flags["--tcp-dial"]); err != nil {
				return err
			}

		case flags["--udp-listen"] != "":
			if err := proc.GlobalPipes.CreateListener(flags["--create"], "udp", flags["--udp-listen"]); err != nil {
				return err
			}

		case flags["--tcp-listen"] != "":
			if err := proc.GlobalPipes.CreateListener(flags["--create"], "tcp", flags["--tcp-listen"]); err != nil {
				return err
			}

		default:
			if err := proc.GlobalPipes.CreatePipe(flags["--create"]); err != nil {
				return err
			}

		}
	default:
		return errors.New("Invalid parameters.")
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
