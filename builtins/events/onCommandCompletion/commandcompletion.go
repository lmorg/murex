package oncommandcompletion

import (
	"fmt"
	"sync"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/lang/ref"
)

const eventType = "onCommandCompletion"

func init() {
	event := newCommandCompletion()
	events.AddEventType(eventType, event, nil)
	Callback = event.callback
}

var Callback func(string, []string)

// Interrupt is a JSONable structure passed to the murex function
type Interrupt struct {
	Name       string
	Command    string
	Parameters []string
}

type commandCompletionEvent struct {
	Command string
	Block   []rune
	FileRef *ref.File
}

type commandCompletionEvents struct {
	events map[string]commandCompletionEvent
	mutex  sync.Mutex
}

func newCommandCompletion() *commandCompletionEvents {
	cce := new(commandCompletionEvents)
	cce.events = make(map[string]commandCompletionEvent)
	return cce
}

// Add a command to the onCommandCompleteionEvents
func (evt *commandCompletionEvents) Add(name, command string, block []rune, fileRef *ref.File) error {
	evt.mutex.Lock()

	evt.events[name] = commandCompletionEvent{
		Command: command,
		Block:   block,
		FileRef: fileRef,
	}

	return nil
}

func (evt *commandCompletionEvents) Remove(name string) error {
	evt.mutex.Lock()
	if evt.events[name].FileRef == nil {
		evt.mutex.Unlock()
		return fmt.Errorf("%s event %s does not exist", eventType, name)
	}

	delete(evt.events, name)

	evt.mutex.Unlock()
	return nil
}

func (evt *commandCompletionEvents) callback(command string, parameters []string) {
	evt.mutex.Lock()

	for name, cce := range evt.events {
		if cce.Command == command {
			cce.execEvent(name, parameters)
		}
	}

	evt.mutex.Unlock()
}

func (cce *commandCompletionEvent) execEvent(name string, parameters []string) {
	interrupt := Interrupt{
		Name:       name,
		Command:    cce.Command,
		Parameters: parameters,
	}

	events.Callback(name, interrupt, cce.Block, cce.FileRef, term.NewErr(false))
}

func (evt *commandCompletionEvents) Dump() interface{} {
	dump := newCommandCompletion()

	evt.mutex.Lock()

	for name, event := range evt.events {
		dump.events[name] = event
	}

	evt.mutex.Unlock()
	return dump.events
}
