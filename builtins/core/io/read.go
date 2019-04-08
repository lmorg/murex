package io

import (
	"errors"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/readline"
)

func init() {
	lang.GoFunctions["read"] = cmdRead
	lang.GoFunctions["tread"] = cmdTread
}

func cmdRead(p *lang.Process) error {
	return read(p, types.String, 0)
}

func cmdTread(p *lang.Process) error {
	dt, err := p.Parameters.String(0)
	if err != nil {
		return err
	}
	return read(p, dt, 1)
}

func read(p *lang.Process, dt string, paramAdjust int) error {
	p.Stdout.SetDataType(types.Null)

	if p.IsBackground {
		return errors.New("Background processes cannot read from stdin")
	}

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

	return p.Parent.Variables.Set(varName, s, dt)
}
