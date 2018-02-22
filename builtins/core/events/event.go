package events

import (
	"errors"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/parameters"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
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
		"-t": "--timer",
		"-f": "--filesystem",
		//"-c": "--command",
		//"-i": "--interrupt",

		"--timer":      types.Boolean,
		"--filesystem": types.Boolean,
		//"--command":    types.Boolean,
		//"--interrupt": types.Boolean,
	},
}

type eventType interface {
	Init()
	Add(interrupt string, block []rune) (err error)
	Remove(interrupt string) (err error)
	Dump() (dump interface{})
}

var events map[string]eventType = map[string]eventType{
	"--filesystem": newWatch(),
	"--timer":      newTimer(),
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
		return errors.New("Too few parameters. You need to include interrupts to listen for and a callback code block")
	}

	//block := []rune(params[len(params)-1])
	block, _ := types.ConvertGoType(params[len(params)-1], types.CodeBlock)
	//if err != nil {
	//	return err
	//}

	params = params[:len(params)-1]
	//if !types.IsBlock([]byte(string(block))) {
	//	return errors.New("Callback parameter is not a code block")
	//}

	var errs string
	for _, a := range params {
		err := events[et].Add(a, []rune(block.(string)))
		if err != nil {
			errs += " {interrupt: " + a + ", err: " + err.Error() + "}"
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
		return errors.New("Too few parameters. You need to include interrupts you want terminated")
	}

	var errs string
	for _, a := range params {
		err := events[et].Remove(a)
		if err != nil {
			errs += " {interrupt: " + a + ", err: " + err.Error() + "}"
		}
	}

	if errs != "" {
		err = errors.New(errs)
	}

	return err
}

type j struct {
	Interrupt   string
	Event       interface{}
	Description string
}

func callback(evtName string, evtOp interface{}, evtDesc string, block []rune) {
	go func() {
		json, err := utils.JsonMarshal(&j{
			Interrupt:   evtName,
			Event:       evtOp,
			Description: evtDesc,
		}, false)
		if err != nil {
			ansi.Stderrln(ansi.FgRed, "error building event input: "+err.Error())
			return
		}

		stdin := streams.NewStdin()
		_, err = stdin.Write(json)
		if err != nil {
			ansi.Stderrln(ansi.FgRed, "error writing event input: "+err.Error())
			return
		}
		stdin.Close()

		/*stdout := streams.NewStdin()
		go func() {
			b, _ := stdout.ReadAll()
			os.Stdout.Write(b)
			stdout.Close()
		}()

		stderr := streams.NewStdin()
		go func() {
			b, _ := stderr.ReadAll()
			os.Stderr.Write(b)
			stderr.Close()
		}()*/

		debug.Log("Event callback:", string(json), string(block))
		_, err = lang.RunBlockShellNamespace(block, stdin, proc.ShellProcess.Stdout, proc.ShellProcess.Stderr)
		if err != nil {
			ansi.Stderrln(ansi.FgRed, "error compiling event callback: "+err.Error())
			return
		}
	}()
}

// DumpEvents is used for `runtime` to output all the saved events
func DumpEvents() (dump map[string]interface{}) {
	dump = make(map[string]interface{})

	for et := range events {
		dump[et] = events[et].Dump()
	}

	return
}
