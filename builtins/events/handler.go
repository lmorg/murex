package events

import (
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/json"
)

type eventType interface {
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
	Name      string
	Interrupt interface{}
}

// Callback is a generic function your event handlers types should hook into so
// murex functions can remain consistent.
func Callback(name string, interrupt interface{}, block []rune, stdout stdio.Io) {
	json, err := json.Marshal(&j{
		Name:      name,
		Interrupt: interrupt,
	}, false)
	if err != nil {
		//ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "error building event input: "+err.Error())
		proc.ShellProcess.Stderr.Writeln([]byte("error building event input: " + err.Error()))

		return
	}

	stdin := streams.NewStdin()
	stdin.SetDataType(types.Json)
	_, err = stdin.Write(json)
	if err != nil {
		//ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "error writing event input: "+err.Error())
		proc.ShellProcess.Stderr.Writeln([]byte("error writing event input: " + err.Error()))
		return
	}

	debug.Log("Event callback:", string(json), string(block))
	branch := proc.ShellProcess.BranchFID()
	branch.Process.IsBackground = true
	defer branch.Close()
	_, err = lang.RunBlockExistingConfigSpace(block, stdin, stdout, proc.ShellProcess.Stderr, branch.Process)
	if err != nil {
		//ansi.Stderrln(proc.ShellProcess, ansi.FgRed, "error compiling event callback: "+err.Error())
		proc.ShellProcess.Stderr.Writeln([]byte("error compiling event callback: " + err.Error()))
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
