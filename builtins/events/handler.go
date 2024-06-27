package events

import (
	"fmt"
	"os"
	"time"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/consts"
	"github.com/lmorg/murex/utils/json"
)

type eventType interface {
	Add(name, interrupt string, block []rune, fileRef *ref.File) (err error)
	Remove(interrupt string) (err error)
	Dump() (dump map[string]DumpT)
}

var events = make(map[string]eventType)

// AddEventType registers your event type handlers
func AddEventType(eventTypeName string, handlerInterface eventType, err error) error {
	if err != nil {
		os.Stderr.WriteString(
			fmt.Sprintf("cannot add event module %s: %s", eventTypeName, err),
		)
	}
	events[eventTypeName] = handlerInterface
	return nil
}

type j struct {
	Name      string
	Interrupt interface{}
}

// Callback is a generic function your event handlers types should hook into so
// murex functions can remain consistent.
func Callback(name string, interrupt interface{}, block []rune, fileRef *ref.File, stdout, stderr stdio.Io, evtReturn any, background bool) (any, error) {
	if fileRef == nil {
		if debug.Enabled {
			panic("fileRef should not be nil value")
		}
		os.Stderr.WriteString("Murex error with `event`: '" + name + "'. fileRef should not be nil value. Creating empty object to continue. Please report this https://github.com/lmorg/murex/issues\n")
		fileRef = &ref.File{
			Source: &ref.Source{
				Filename: "UNKNOWN: forked from `event` " + name,
				Module:   "UNKNOWN: forked from `event` " + name,
				DateTime: time.Now(),
			},
		}
	}

	json, err := json.Marshal(&j{
		Name:      name,
		Interrupt: interrupt,
	}, false)
	if err != nil {
		return nil, fmt.Errorf("error building event input: %s", err.Error())
	}

	var bgProc int
	if background {
		bgProc = lang.F_BACKGROUND
	}

	fork := lang.ShellProcess.Fork(lang.F_FUNCTION | bgProc | lang.F_CREATE_STDIN)
	fork.Stdin.SetDataType(types.Json)
	fork.Name.Set("(event)")
	fork.FileRef = fileRef
	fork.CCEvent = nil
	fork.CCExists = nil
	_, err = fork.Stdin.Write(json)
	if err != nil {
		return nil, fmt.Errorf("error writing event input: %s", err.Error())
	}

	debug.Log("Event callback:", string(json), string(block))

	fork.Stdout = stdout
	fork.Stderr = stderr
	if evtReturn != nil {
		err = fork.Variables.Set(fork.Process, consts.EventReturn, evtReturn, types.Json)
		if err != nil {
			return nil, fmt.Errorf("cannot compile %s for event '%s' (interrupt: %v): %s", consts.EventReturn, name, interrupt, err.Error())
		}
	}

	_, err = fork.Execute(block)
	if err != nil {
		return nil, fmt.Errorf("error compiling event callback: %s", err.Error())
	}

	if evtReturn != nil {
		return fork.Variables.GetValue(consts.EventReturn)
	}
	return nil, nil
}

// DumpEvents is used for `runtime` to output all the saved events
func DumpEvents() (dump map[string]interface{}) {
	dump = make(map[string]interface{})

	for et := range events {
		dump[et] = events[et].Dump()
	}

	return
}

func DumpEventTypes() []string {
	var (
		s = make([]string, len(events))
		i int
	)

	for name := range events {
		s[i] = name
		i++
	}

	return s
}

type DumpT struct {
	Interrupt string
	Block     string
	FileRef   *ref.File
}
