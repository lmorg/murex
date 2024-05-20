package onkeypress

import (
	"errors"
	"fmt"
	"sync"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/shell/variables"
	"github.com/lmorg/murex/utils/readline"
)

const eventType = "onKeyPress"

func init() {
	events.AddEventType(eventType, newKeyPress(), nil)
}

// Interrupt is a JSONable structure passed to the murex function
type Interrupt struct {
	Line        string
	Raw         string
	Pos         int
	KeySequence string
	Invoker     string
	IsMasked    bool
	Mode        string
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

func (evt *keyPressEvents) callback(keyPress string, line []rune, pos int) *readline.EventReturn {
	var i int

	evt.mutex.Lock()
	defer evt.mutex.Unlock()

	for i = range evt.events {
		if evt.events[i].keySeq == keyPress {
			goto eventFound
		}
	}
	return &readline.EventReturn{
		NewLine: line,
		NewPos:  pos,
	}

eventFound:
	block := evt.events[i].block

	interrupt := Interrupt{
		Line:        variables.ExpandString(string(line)),
		Raw:         string(line),
		Pos:         pos,
		KeySequence: keyPress,
	}

	stdout := streams.NewStdin()
	_, err := events.Callback(
		evt.events[i].name, interrupt, // event
		block, evt.events[i].fileRef, // script
		stdout, new(null.Null), // pipes
		createMeta(line, pos), // meta
		true,                  // background
	)
	if err != nil {
		return &readline.EventReturn{
			HintText: []rune("callback error: " + err.Error()),
			NewLine:  line,
			NewPos:   pos,
		}
	}

	ret := make(map[string]string)
	err = stdout.ReadMap(lang.ShellProcess.Config, func(readmap *stdio.Map) {
		v, _ := types.ConvertGoType(readmap.Value, types.String)
		ret[readmap.Key] = v.(string)
	})
	if err != nil {
		return &readline.EventReturn{
			HintText: []rune("callback error: " + err.Error()),
			NewLine:  line,
			NewPos:   pos,
		}
	}

	forwardKey, err := types.ConvertGoType(ret["ForwardKey"], types.Boolean)
	if err != nil {
		return &readline.EventReturn{
			HintText: []rune("callback error: " + err.Error()),
			NewLine:  line,
			NewPos:   pos,
		}
	}

	clearHelpers, err := types.ConvertGoType(ret["ClearHelpers"], types.Boolean)
	if err != nil {
		return &readline.EventReturn{
			HintText: []rune("callback error: " + err.Error()),
			NewLine:  line,
			NewPos:   pos,
		}
	}

	closeReadline, err := types.ConvertGoType(ret["CloseReadline"], types.Boolean)
	if err != nil {
		return &readline.EventReturn{
			HintText: []rune("callback error: " + err.Error()),
			NewLine:  line,
			NewPos:   pos,
		}
	}

	var newLine []rune
	if ret["NewLine"] != "" {
		newLine = []rune(ret["NewLine"])
	} else {
		newLine = line
	}

	var newPos int
	if ret["NewPos"] != "" {
		i, err := types.ConvertGoType(ret["NewPos"], types.Integer)
		if err != nil {
			return &readline.EventReturn{
				HintText: []rune("callback error: " + err.Error()),
				NewLine:  line,
				NewPos:   pos,
			}
		}
		newPos = i.(int)
	} else {
		newPos = pos
	}

	return &readline.EventReturn{
		ForwardKey:    forwardKey.(bool),
		ClearHelpers:  clearHelpers.(bool),
		CloseReadline: closeReadline.(bool),
		HintText:      []rune(ret["HintText"]),
		NewLine:       newLine,
		NewPos:        newPos,
	}
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
