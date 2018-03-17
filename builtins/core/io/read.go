package io

import (
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/readline"
)

func init() {
	proc.GoFunctions["read"] = cmdRead
	proc.GoFunctions["tread"] = cmdTread
}

func cmdRead(p *proc.Process) error {
	return read(p, types.String, 0)
}

func cmdTread(p *proc.Process) error {
	dt, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	return read(p, dt, 1)
}

func read(p *proc.Process, dt string, paramAdjust int) error {
	p.Stdout.SetDataType(types.Null)

	var prompt string
	if p.IsMethod {
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}
		prompt = string(b)
	} else {
		prompt = p.Parameters.StringAllRange(1+paramAdjust, -1)
	}

	varName, err := p.Parameters.String(0 + paramAdjust)
	if err != nil {
		return err
	}

	rl := readline.NewInstance()

	rl.SetPrompt(prompt)
	rl.History = new(readline.NullHistory)

	s, err := rl.Readline()
	if err != nil {
		return err
	}

	err = p.Variables.Set(varName, s, dt)
	return err
}
