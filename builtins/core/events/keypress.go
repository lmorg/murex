package events

import (
	"errors"
	"github.com/lmorg/murex/shell"
)

type keyPressEvents struct {
	events map[string][]rune
}

func newKeyPress() *keyPressEvents {
	return new(keyPressEvents)
}

func (evt *keyPressEvents) Init() {
	evt.events = make(map[string][]rune)
}

// Add a key to the event list
func (evt *keyPressEvents) Add(keyPress string, block []rune) error {
	if !shell.Interactive || shell.Prompt == nil {
		return errors.New("Key press events can only be created in an interactive shell.")
	}

	shell.Prompt.AddEvent(keyPress, evt.callback)
	evt.events[keyPress] = block
	return nil
}

func (evt *keyPressEvents) Remove(keyPress string) error {
	if !shell.Interactive || shell.Prompt == nil {
		return errors.New("Key press events can only be created in an interactive shell.")
	}

	shell.Prompt.DelEvent(keyPress)
	delete(evt.events, keyPress)
	return nil
}

func (evt keyPressEvents) callback(keyPress string, line []rune, pos int) (bool, bool) {
	block := evt.events[keyPress]
	callback(keyPress, pos, string(line), block)
	return false, false
}

func (evt *keyPressEvents) Dump() interface{} {
	return evt.events
}
