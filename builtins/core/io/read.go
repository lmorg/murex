package io

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/readline"
)

func init() {
	proc.GoFunctions["read"] = cmdRead
	proc.GoFunctions["tread"] = cmdTread
}

func cmdRead(p *proc.Process) error {
	return read(p, types.Generic, 1)
}

func cmdTread(p *proc.Process) error {
	dt, err := p.Parameters.String(1)
	if err != nil {
		return err
	}
	return read(p, dt, 2)
}

func read(p *proc.Process, dt string, paramStart int) error {
	p.Stdout.SetDataType(types.Null)

	var prompt string
	if p.IsMethod {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		prompt = string(b)
	} else {
		prompt = p.Parameters.StringAllRange(paramStart, -1)
	}

	varName, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	rl, err := readline.NewEx(&readline.Config{
		InterruptPrompt:        " ",
		DisableAutoSaveHistory: true,
		NoEofOnEmptyDelete:     false,
		Prompt:                 prompt,
	})

	if err != nil {
		return err
	}

	s, err := rl.Readline()
	if err != nil {
		return err
	}

	err = proc.GlobalVars.Set(varName, s, dt)
	return err
}
