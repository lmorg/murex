package io

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
	"io"
)

func init() {
	proc.GoFunctions["pipe"] = proc.GoFunction{Func: cmdPipe, TypeIn: types.Null, TypeOut: types.Null}
	proc.GoFunctions["<read-pipe>"] = proc.GoFunction{Func: cmdReadPipe, TypeIn: types.Null, TypeOut: types.Generic}
}

func cmdPipe(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	flags, args, err := p.Parameters.ParseFlags(&parameters.Arguments{
		AllowAdditional: true,
		Flags: map[string]string{
			"--create": types.Boolean,
			"--close":  types.Boolean,
		},
	})

	if err != nil {
		return err
	}

	if flags["--create"] == types.TrueString && flags["--close"] == types.TrueString {
		return errors.New("Cannot `--create` and `--close` in the same command.")
	}

	switch types.TrueString {
	case flags["--create"]:
		if len(args) == 0 {
			return errors.New("Not enough parameters. Please include the name(s) of pipe(s) you wish to create.")
		}

		for i := range args {
			err := proc.GlobalPipes.Create(args[i])
			if err != nil {
				return err
			}
		}

	case flags["--close"]:
		if len(args) == 0 {
			return errors.New("Not enough parameters. Please include the name(s) of pipe(s) you wish to close.")
		}

		for i := range args {
			err := proc.GlobalPipes.Close(args[i])
			if err != nil {
				return err
			}
		}

	default:
		return errors.New("Not enough parameters. Please include either `--create` or `--close`.")
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

	p.Stdout.SetDataType(pipe.GetDataType())
	_, err = io.Copy(p.Stdout, pipe)
	return err
}
