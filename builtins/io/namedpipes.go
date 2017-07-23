package io

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
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

	flag, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	args := p.Parameters.StringArray()
	if len(args) < 1 {
		return errors.New("Not enough parameters!")
	}
	args = args[1:]

	switch flag {
	case "--create", "-c":
		for i := range args {
			err := proc.GlobalPipes.CreatePipe(args[i])
			if err != nil {
				return err
			}
		}

	case "--close", "-x":
		for i := range args {
			err := proc.GlobalPipes.Close(args[i])
			if err != nil {
				return err
			}
		}

	case "--file", "-f":
		if len(args) < 2 {
			return errors.New("Not enough parameters!")
		}

		err := proc.GlobalPipes.CreateFile(args[0], args[1])
		if err != nil {
			return err
		}

	case "--udp-dial":
		if len(args) < 2 {
			return errors.New("Not enough parameters!")
		}

		err := proc.GlobalPipes.CreateDialer(args[0], "udp", args[1])
		if err != nil {
			return err
		}

	case "--tcp-dial":
		if len(args) < 2 {
			return errors.New("Not enough parameters!")
		}

		err := proc.GlobalPipes.CreateDialer(args[0], "tcp", args[1])
		if err != nil {
			return err
		}

	case "--udp-listen":
		if len(args) < 2 {
			return errors.New("Not enough parameters!")
		}

		err := proc.GlobalPipes.CreateListener(args[0], "udp", args[1])
		if err != nil {
			return err
		}

	case "--tcp-listen":
		if len(args) < 2 {
			return errors.New("Not enough parameters!")
		}

		err := proc.GlobalPipes.CreateListener(args[0], "tcp", args[1])
		if err != nil {
			return err
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
