package io

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/parameters"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/json"
	"github.com/lmorg/murex/utils/lists"
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
	lang.FeatureDeprecatedBuiltin(p)

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
	flagReadComplete = "--autocomplete"
)

var readArguments = parameters.Arguments{
	Flags: map[string]string{
		flagReadDefault:  types.String,
		flagReadPrompt:   types.String,
		flagReadVariable: types.String,
		flagReadDataType: types.String,
		flagReadMask:     types.String,
		flagReadComplete: types.String,
	},
	AllowAdditional: true,
}

func read(p *lang.Process, dt string, paramAdjust int) error {
	p.Stdout.SetDataType(types.Null)

	if p.Background.Get() {
		return errors.New("background processes cannot read from stdin")
	}

	var prompt, varName, defaultVal, mask, complete string

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
		complete = flags[flagReadComplete]

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

	err = tabCompleter(rl, []byte(complete))
	if err != nil {
		return err
	}

	if len(mask) > 0 {
		rl.PasswordMask, _ = utf8.DecodeRuneInString(mask)
	}

	s, err := rl.Readline()
	if err != nil {
		return err
	}

	if s == "" {
		s = defaultVal
		//tty.Stdout.WriteString(s)
	}

	v, err := types.ConvertGoType(s, dt)
	if err != nil {
		return err
	}

	return p.Variables.Set(p, varName, v, dt)
}

func tabCompleter(rl *readline.Instance, b []byte) error {
	if len(b) == 0 {
		return nil
	}

	maxRows, _ := lang.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	rl.MaxTabCompleterRows = maxRows.(int)

	var v interface{}
	err := json.UnmarshalMurex(b, &v)
	if err != nil {
		return fmt.Errorf("cannot unmarshal JSON input for `read`'s autocomplete: %s", err.Error())
	}

	switch t := v.(type) {
	case []string:
		rl.TabCompleter = func(r []rune, i int, dtc readline.DelayedTabContext) *readline.TabCompleterReturnT {
			tcr := new(readline.TabCompleterReturnT)
			if i > len(r) {
				return tcr
			}
			tcr.Prefix = string(r[:i])
			tcr.Suggestions = lists.CropPartial(t, tcr.Prefix)
			return tcr
		}

	case []interface{}:
		// this is horribly inefficient POC code
		s := make([]string, len(t))
		for i := range t {
			s[i] = fmt.Sprint(t[i])
		}
		rl.TabCompleter = func(r []rune, i int, dtc readline.DelayedTabContext) *readline.TabCompleterReturnT {
			tcr := new(readline.TabCompleterReturnT)
			if i > len(r) {
				return tcr
			}
			tcr.Prefix = string(r[:i])
			tcr.Suggestions = lists.CropPartial(s, tcr.Prefix)
			return tcr
		}

	case map[string]string:
		// this is horribly inefficient POC code
		s := make([]string, len(t))
		var i int
		for key := range t {
			s[i] = key
			i++
		}
		rl.TabCompleter = func(r []rune, i int, dtc readline.DelayedTabContext) *readline.TabCompleterReturnT {
			tcr := new(readline.TabCompleterReturnT)
			if i > len(r) {
				return tcr
			}
			tcr.Prefix = string(r[:i])
			tcr.Suggestions = lists.CropPartial(s, tcr.Prefix)
			tcr.Descriptions = lists.CropPartialMapKeys(t, tcr.Prefix)
			tcr.DisplayType = readline.TabDisplayList
			return tcr
		}

	case map[string]interface{}:
		// this is horribly inefficient POC code
		s := make([]string, len(t))
		var i int
		m := make(map[string]string)
		for key, val := range t {
			s[i] = key
			m[key] = fmt.Sprint(val)
			i++
		}
		rl.TabCompleter = func(r []rune, i int, dtc readline.DelayedTabContext) *readline.TabCompleterReturnT {
			tcr := new(readline.TabCompleterReturnT)
			if i > len(r) {
				return tcr
			}
			tcr.Prefix = string(r[:i])
			tcr.Suggestions = lists.CropPartial(s, tcr.Prefix)
			tcr.Descriptions = lists.CropPartialMapKeys(m, tcr.Prefix)
			tcr.DisplayType = readline.TabDisplayList
			return tcr
		}

	default:
		return fmt.Errorf("autocomplete JSON unmarshalled to unsupported object %T. Expecting either a string or a map", t)
	}

	return nil
}
