package onkeypress

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/murex/utils/lists"
	"github.com/lmorg/murex/utils/readline"
)

const eventType = "onKeyPress"

func init() {
	events.AddEventType(eventType, newKeyPress(), nil)
}

type keyPressEvent struct {
	name    string
	keySeq  string
	escaped string
	block   []rune
	fileRef *ref.File
}

type keyPressEvents struct {
	events map[string][]*keyPressEvent
	mutex  sync.Mutex
}

func newKeyPress() *keyPressEvents {
	evt := new(keyPressEvents)
	evt.events = make(map[string][]*keyPressEvent)
	return evt
}

const shellPromptIsNil = "unable to register event with readline API: shell.Prompt is nil"

// Add a key to the event list
func (evt *keyPressEvents) Add(name, keySeq string, block []rune, fileRef *ref.File) error {
	if shell.Prompt == nil {
		return errors.New(shellPromptIsNil)
	}

	keySeqEscaped := ansi.GetConsts([]byte(keySeq))

	if evt.exists(name, keySeq) != doesNotExist {
		evt.Remove(events.CompileInterruptKey(keySeq, name))
	}

	shell.Prompt.AddEvent(keySeq, evt.readlineCallback) // doesn't matter if it's already registered

	events := append(evt.events[keySeq], &keyPressEvent{
		name:    name,
		escaped: keySeqEscaped, // purely for a human readable representation
		keySeq:  keySeq,
		block:   block,
		fileRef: fileRef,
	})

	sort.Slice(events, func(i, j int) bool {
		return events[i].name < events[j].name
	})

	evt.events[keySeq] = events
	return nil
}

func (evt *keyPressEvents) Remove(name string) (err error) {
	if shell.Prompt == nil {
		return errors.New(shellPromptIsNil)
	}

	key := events.GetInterruptFromKey(name)

	evt.mutex.Lock()
	defer evt.mutex.Unlock()

	switch key.Interrupt {
	case "":
		for key.Interrupt = range evt.events {
			err = evt._remove(key)
			if err != nil && !strings.Contains(err.Error(), "no event found") {
				return err
			}
		}
		return err

	default:
		key.Interrupt = ansi.ForceExpandConsts(key.Interrupt, false)
		return evt._remove(key)

	}
}

func (evt *keyPressEvents) _remove(key *events.Key) (err error) {
	for i := range evt.events[key.Interrupt] {
		if evt.events[key.Interrupt][i].name == key.Name {
			//shell.Prompt.DelEvent(evt.events[i].keySeq) // TODO: check if any further events exist...
			evt.events[key.Interrupt], err = lists.RemoveOrdered(evt.events[key.Interrupt], i)
			if err != nil {
				return fmt.Errorf("unable to delete event '%s': %s", key.Name, err.Error())
			}
			return nil
		}
	}

	return fmt.Errorf("unable to delete event as no event found with the name `%s` for event type `%s`", key.Name, eventType)
}

func callbackError(err error, state *readline.EventState) *readline.EventReturn {
	return &readline.EventReturn{
		HintText: []rune(fmt.Sprintf("callback error: %s", err.Error())),
		SetLine:  []rune(state.Line),
		SetPos:   state.CursorPos,
	}
}

func (evt *keyPressEvents) readlineCallback(id int, state *readline.EventState) *readline.EventReturn {
	evt.mutex.Lock()
	defer evt.mutex.Unlock()

	events := evt.events[state.KeyPress]

	if id >= len(events) {
		return &readline.EventReturn{NextEvent: true}
	}

	ret := onKeyPressEvent(events[id], state)
	if id < len(events)-1 {
		ret.MoreEvents = true
	}

	return ret
}

func onKeyPressEvent(event *keyPressEvent, state *readline.EventState) *readline.EventReturn {
	meta, err := events.Callback(
		event.name, state, // event
		event.block, event.fileRef, // script
		lang.ShellProcess.Stdout, lang.ShellProcess.Stderr, // pipes
		createMeta(state), // meta
		true,              // background
	)
	if err != nil {
		return callbackError(err, state)
	}

	evtReturn, err := validateMeta(meta)
	if err != nil {
		return callbackError(err, state)
	}

	return evtReturn
}

const doesNotExist = -1

func (evt *keyPressEvents) exists(name, keySeq string) int {
	evt.mutex.Lock()

	events := evt.events[keySeq]
	for i := range events {
		if events[i].name == name {
			return i
		}
	}

	evt.mutex.Unlock()
	return doesNotExist
}

func (evt *keyPressEvents) Dump() map[string]events.DumpT {
	dump := make(map[string]events.DumpT)

	evt.mutex.Lock()

	for _, evts := range evt.events {
		for _, event := range evts {
			dump[events.CompileInterruptKey(event.escaped, event.name)] = events.DumpT{
				Interrupt: event.escaped,
				Block:     string(event.block),
				FileRef:   event.fileRef,
			}
		}
	}

	evt.mutex.Unlock()

	return dump
}
