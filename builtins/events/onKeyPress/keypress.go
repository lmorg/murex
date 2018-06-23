package onKeyPress

import (
	"errors"
	"fmt"
	"sync"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
)

const eventType = "onKeyPress"

func init() {
	events.AddEventType(eventType, newKeyPress())
}

// Interrupt is a JSONable structure passed to the murex function
type Interrupt struct {
	Line        string
	Pos         int
	KeySequence string
}

type keyPressEvent struct {
	name   string
	keySeq string
	block  []rune
}

type keyPressEvents struct {
	events []keyPressEvent
	mutex  sync.Mutex
}

func newKeyPress() *keyPressEvents {
	return new(keyPressEvents)
}

// Add a key to the event list
func (evt *keyPressEvents) Add(name, keySeq string, block []rune) error {
	if shell.Prompt == nil {
		return errors.New("Unable to register event with readline API.")
	}

	evt.mutex.Lock()
	for i := range evt.events {
		if evt.events[i].name == name {
			evt.mutex.Unlock()
			return fmt.Errorf("Event already exists with the name `%s` for event type `%s`.", name, eventType)
		}
		if evt.events[i].keySeq == keySeq {
			evt.mutex.Unlock()
			return fmt.Errorf("Event already exists with that  key sequence for event type `%s`.", eventType)
		}
	}
	evt.mutex.Unlock()

	shell.Prompt.AddEvent(keySeq, evt.callback)
	evt.events = append(evt.events, keyPressEvent{
		name:   name,
		keySeq: keySeq,
		block:  block,
	})
	return nil
}

func (evt *keyPressEvents) Remove(name string) error {
	remove := func(s []keyPressEvent, i int) []keyPressEvent {
		s[len(s)-1], s[i] = s[i], s[len(s)-1]
		return s[:len(s)-1]
	}

	if shell.Prompt == nil {
		return errors.New("Unable to de-register event with readline API.")
	}

	evt.mutex.Lock()
	defer evt.mutex.Unlock()

	for i := range evt.events {
		if evt.events[i].name == name {
			shell.Prompt.DelEvent(evt.events[i].keySeq)
			evt.events = remove(evt.events, i)
			return nil
		}
	}

	return fmt.Errorf("Unable to delete event as no event found with the name `%s` for event type `%s`.", name, eventType)
}

func (evt *keyPressEvents) callback(keyPress string, line []rune, pos int) (bool, bool, bool, []rune) {
	var i int

	evt.mutex.Lock()
	defer evt.mutex.Unlock()

	for i = range evt.events {
		if evt.events[i].keySeq == keyPress {
			goto eventFound
		}
	}
	return false, false, false, nil

eventFound:
	block := evt.events[i].block

	interrupt := Interrupt{
		Line:        string(line),
		Pos:         pos,
		KeySequence: keyPress,
	}

	stdout := streams.NewStdin()
	events.Callback(evt.events[i].name, interrupt, block, stdout)

	ret := make(map[string]string)
	err := stdout.ReadMap(proc.ShellProcess.Config, func(key string, value string, last bool) {
		ret[key] = value
	})
	if err != nil {
		return false, false, false, []rune("Callback error: " + err.Error())
	}

	ignoreKey, err := types.ConvertGoType(ret["IgnoreKey"], types.Boolean)
	if err != nil {
		return false, false, false, []rune("Callback error: " + err.Error())
	}

	clearHelpers, err := types.ConvertGoType(ret["ClearHelpers"], types.Boolean)
	if err != nil {
		return false, false, false, []rune("Callback error: " + err.Error())
	}

	closeReadline, err := types.ConvertGoType(ret["CloseReadline"], types.Boolean)
	if err != nil {
		return false, false, false, []rune("Callback error: " + err.Error())
	}

	return ignoreKey.(bool), clearHelpers.(bool), closeReadline.(bool), []rune(ret["HintText"])
}

func (evt *keyPressEvents) Dump() interface{} {
	type kp struct {
		KeySequence string
		Block       string
	}

	dump := make(map[string]kp)

	evt.mutex.Lock()
	defer evt.mutex.Unlock()

	for i := range evt.events {
		dump[evt.events[i].name] = kp{
			KeySequence: evt.events[i].keySeq,
			Block:       string(evt.events[i].block),
		}
	}
	return dump
}
