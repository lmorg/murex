package events

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/ansi"
)

func init() {
	proc.GoFunctions["event"] = cmdEvent
	//proc.GoFunctions["!event"] = cmdUnevent

	for e := range events {
		go events[e].Init()
	}
}

var rxNameInterruptSyntax *regexp.Regexp = regexp.MustCompile(`^([-_a-zA-Z0-9]+)=(.*)$`)

type eventType interface {
	Init()
	Add(name, interrupt string, block []rune) (err error)
	Remove(interrupt string) (err error)
	Dump() (dump interface{})
}

var events map[string]eventType = make(map[string]eventType)

//"onFilesystemChange":  newWatch(),

func AddEventType(eventTypeName string, handlerInterface eventType) {
	events[eventTypeName] = handlerInterface
}

func cmdEvent(p *proc.Process) error {
	p.Stdout.SetDataType(types.Null)

	et, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	if events[et] == nil {
		return fmt.Errorf("No event-type known for `%s`.\nRun `runtime --events` to view which events are compiled in.", et)
	}

	nameInterrupt, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	split := rxNameInterruptSyntax.FindAllStringSubmatch(nameInterrupt, 1)
	if len(split) != 1 || len(split[0]) != 3 {
		return errors.New("Invalid syntax: " + nameInterrupt + ". Expected: `name=interrupt`.")
	}

	name := split[0][1]
	interrupt := split[0][2]

	block, err := p.Parameters.Block(2)
	if err != nil {
		return err
	}

	err = events[et].Add(name, interrupt, block)
	return err
}

/*func cmdUnevent(p *proc.Process) error {
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
}*/

type j struct {
	Interrupt   string
	Event       interface{}
	Description string
}

func Callback(evtName string, evtOp interface{}, evtDesc string, block []rune, stdout stdio.Io) {
	//go func() {
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
	stdin.SetDataType(types.Json)
	_, err = stdin.Write(json)
	if err != nil {
		ansi.Stderrln(ansi.FgRed, "error writing event input: "+err.Error())
		return
	}
	stdin.Close()

	debug.Log("Event callback:", string(json), string(block))
	_, err = lang.RunBlockShellNamespace(block, stdin, stdout, proc.ShellProcess.Stderr)
	if err != nil {
		ansi.Stderrln(ansi.FgRed, "error compiling event callback: "+err.Error())
	}

	return
	//}()
}

// DumpEvents is used for `runtime` to output all the saved events
func DumpEvents() (dump map[string]interface{}) {
	dump = make(map[string]interface{})

	for et := range events {
		dump[et] = events[et].Dump()
	}

	return
}
