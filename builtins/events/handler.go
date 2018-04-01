package events

import (
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
	for e := range events {
		go events[e].Init()
	}
}

type eventType interface {
	Init()
	Add(name, interrupt string, block []rune) (err error)
	Remove(interrupt string) (err error)
	Dump() (dump interface{})
}

var events map[string]eventType = make(map[string]eventType)

// AddEventType registers your event type handlers
func AddEventType(eventTypeName string, handlerInterface eventType) {
	events[eventTypeName] = handlerInterface
}

type j struct {
	Interrupt   string
	Event       interface{}
	Description string
}

// Callback is a generic function your event handlers types should hook into so
// murex functions can remain consistant.
func Callback(evtName string, evtOp interface{}, evtDesc string, block []rune, stdout stdio.Io) {
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
}

// DumpEvents is used for `runtime` to output all the saved events
func DumpEvents() (dump map[string]interface{}) {
	dump = make(map[string]interface{})

	for et := range events {
		dump[et] = events[et].Dump()
	}

	return
}
