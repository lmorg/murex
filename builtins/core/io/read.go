package io

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/readline"
)

func init() {
	lang.DefineMethod("read", cmdRead, types.String, types.Null)
	lang.DefineMethod("tread", cmdTread, types.String, types.Null)
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

const (
	flagReadDefault  = "--default"
	flagReadPrompt   = "--prompt"
	flagReadVariable = "--variable"
	flagReadDataType = "--datatype"
	flagReadMask     = "--mask"
)

var readArguments = parameters.Arguments{
	Flags: map[string]string{
		flagReadDefault:  types.String,
		flagReadPrompt:   types.String,
		flagReadVariable: types.String,
		flagReadDataType: types.String,
		flagReadMask:     types.String,
	},
	AllowAdditional: true,
}

func read(p *lang.Process, dt string, paramAdjust int) error {
	p.Stdout.SetDataType(types.Null)

	if p.Background.Get() {
		return errors.New("background processes cannot read from stdin")
	}

	var prompt, varName, defaultVal, mask string

	flags, additional, err := p.Parameters.ParseFlags(&readArguments)
	if err != nil {
		return fmt.Errorf("cannot parse parameters: %s", err.Error())
	}

	if len(additional) == 0 {
		prompt = flags[flagReadPrompt]
		varName = flags[flagReadVariable]
		defaultVal = flags[flagReadDefault]
		datatype := flags[flagReadDataType]
		mask = flags[flagReadMask]

		if datatype != "" {
			dt = datatype
		}

		if varName == "" {
			varName = "read"
		}

	} else {
		if p.IsMethod {
			b, err := p.Stdin.ReadAll()
			if err != nil {
				return err
			}
			prompt = string(b)

			varName, err = p.Parameters.String(0 + paramAdjust)
			if err != nil {
				return err
			}
		} else {
			varName, err = p.Parameters.String(0 + paramAdjust)
			if err != nil {
				return err
			}
			prompt = p.Parameters.StringAllRange(1+paramAdjust, -1)
		}
	}

	prompt = ansi.ExpandConsts(prompt)

	rl := readline.NewInstance()
	rl.SetPrompt(prompt)
	rl.History = new(readline.NullHistory)

	if len(mask) > 0 {
		rl.PasswordMask, _ = utf8.DecodeRuneInString(mask)
	}

	s, err := rl.Readline()
	if err != nil {
		return err
	}

	if s == "" {
		s = defaultVal
		//os.Stdout.WriteString(s)
	}

	v, err := types.ConvertGoType(s, dt)
	if err != nil {
		return err
	}

	return p.Variables.Set(p, varName, v, dt)
}
