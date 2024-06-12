package onkeypress

import (
	"errors"
	"fmt"
	"sync"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils/readline"
)

const eventType = "onKeyPress"

func init() {
	events.AddEventType(eventType, newKeyPress(), nil)
}

type keyPressEvent struct {
	name    string
	keySeq  string
	block   []rune
	fileRef *ref.File
}

type keyPressEvents struct {
	events []keyPressEvent
	mutex  sync.Mutex
}

func newKeyPress() *keyPressEvents {
	return new(keyPressEvents)
}

// Add a key to the event list
func (evt *keyPressEvents) Add(name, keySeq string, block []rune, fileRef *ref.File) error {
	if shell.Prompt == nil {
		return errors.New("unable to register event with readline API")
	}

	shell.Prompt.AddEvent(keySeq, evt.callback)
	evt.events = append(evt.events, keyPressEvent{
		name:    name,
		keySeq:  keySeq,
		block:   block,
		fileRef: fileRef,
	})
	return nil
}

func (evt *keyPressEvents) Remove(name string) error {
	remove := func(s []keyPressEvent, i int) []keyPressEvent {
		s[len(s)-1], s[i] = s[i], s[len(s)-1]
		return s[:len(s)-1]
	}

	if shell.Prompt == nil {
		return errors.New("unable to de-register event with readline API")
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

	return fmt.Errorf("unable to delete event as no event found with the name `%s` for event type `%s`", name, eventType)
}

func callbackError(err error, line []rune, pos int) *readline.EventReturn {
	return &readline.EventReturn{
		HintText: []rune(fmt.Sprintf("callback error: %s", err.Error())),
		SetLine:  line,
		SetPos:   pos,
	}
}

func (evt *keyPressEvents) callback(keyPress string, line []rune, pos int) *readline.EventReturn {
	var i int

	evt.mutex.Lock()
	defer evt.mutex.Unlock()

	for i = range evt.events {
		if evt.events[i].keySeq == keyPress {
			goto eventFound
		}
	}
	return &readline.EventReturn{SetLine: line, SetPos: pos}

eventFound:
	block := evt.events[i].block

	interrupt := Interrupt{
		Line:        variables.ExpandString(string(line)),
		Raw:         string(line),
		Pos:         pos,
		KeySequence: keyPress,
	}

	meta, err := events.Callback(
		evt.events[i].name, interrupt, // event
		block, evt.events[i].fileRef, // script
		lang.ShellProcess.Stdout, lang.ShellProcess.Stderr, // pipes
		createMeta(line, pos), // meta
		true,                  // background
	)
	if err != nil {
		return callbackError(err, line, pos)
	}

	evtReturn, err := validateMeta(meta)
	if err != nil {
		return callbackError(err, line, pos)
	}

	return evtReturn
}

func (evt *keyPressEvents) Dump() map[string]events.DumpT {
	dump := make(map[string]events.DumpT)

	evt.mutex.Lock()

	for i := range evt.events {
		dump[evt.events[i].name] = events.DumpT{
			Interrupt: evt.events[i].keySeq,
			Block:     string(evt.events[i].block),
			FileRef:   evt.events[i].fileRef,
		}
	}

	evt.mutex.Unlock()
	return dump
}
