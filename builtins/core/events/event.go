package events

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["event"] = cmdEvent
	proc.GoFunctions["!event"] = cmdUnevent

	for e := range events {
		go events[e].Init()
	}
}

var args *parameters.Arguments = &parameters.Arguments{
	AllowAdditional: true,
	Flags: map[string]string{
		"-f": "--filesystem",
		"-c": "--command",
		"-s": "--interval",
		"-i": "--interrupt",

		"--filesystem": types.Boolean,
		"--command":    types.Boolean,
		"--interval":   types.Boolean,
		"--interrupt":  types.Boolean,
	},
}

type eventType interface {
	Init()
	Add(article string, block []rune) (err error)
	Remove(article string) (err error)
	Callback(path string) (block []rune)
	Dump() (dump map[string]string)
}

var events map[string]eventType = map[string]eventType{
	"--filesystem": newWatch(),
}

func cmdEvent(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	flags, params, err := p.Parameters.ParseFlags(args)
	if err != nil {
		return err
	}

	if len(flags) == 0 {
		return errors.New("Missing flag defining event type")
	}

	if len(flags) > 1 {
		return errors.New("Only 1 (one) event type flag can be used per command")
	}

	var et string
	for s := range flags {
		et = s
	}

	if len(params) < 2 {
		return errors.New("Too few parameters. You need to include articles to listen for and a callback code block")
	}

	block := []rune(params[len(params)-1])
	params = params[:len(params)-1]
	if !types.IsBlock([]byte(string(block))) {
		return errors.New("Callback parameter is not a code block")
	}

	var errs string
	for _, a := range params {
		err := events[et].Add(a, block)
		if err != nil {
			errs += " {article: " + a + ", err: " + err.Error() + "}"
		}
	}

	if errs != "" {
		err = errors.New(errs)
	}

	return err
}

func cmdUnevent(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	flags, params, err := p.Parameters.ParseFlags(args)
	if err != nil {
		return err
	}

	if len(flags) == 0 {
		return errors.New("Missing flag defining event type")
	}

	if len(flags) > 1 {
		return errors.New("Only 1 (one) event type flag can be used per command")
	}

	var et string
	for s := range flags {
		et = s
	}

	if len(params) == 0 {
		return errors.New("Too few parameters. You need to include articles you want terminated")
	}

	var errs string
	for _, a := range params {
		err := events[et].Remove(a)
		if err != nil {
			errs += " {article: " + a + ", err: " + err.Error() + "}"
		}
	}

	if errs != "" {
		err = errors.New(errs)
	}

	return err
}

// DumpEvents is used for `runtime` to output all the saved events
func DumpEvents() (dump map[string]map[string]string) {
	dump = make(map[string]map[string]string)

	for et := range events {
		dump[et] = events[et].Dump()
	}

	return
}
