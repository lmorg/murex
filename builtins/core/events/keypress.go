package events

import (
	"errors"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/types"
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
	if /*!shell.Interactive ||*/ shell.Prompt == nil {
		return errors.New("Unable to register event with readline API.")
	}

	shell.Prompt.AddEvent(keyPress, evt.callback)
	evt.events[keyPress] = block
	return nil
}

func (evt *keyPressEvents) Remove(keyPress string) error {
	if /*!shell.Interactive ||*/ shell.Prompt == nil {
		return errors.New("Unable to de-register event with readline API.")
	}

	shell.Prompt.DelEvent(keyPress)
	delete(evt.events, keyPress)
	return nil
}

func (evt keyPressEvents) callback(keyPress string, line []rune, pos int) (bool, bool, []rune) {
	block := evt.events[keyPress]
	stdout := streams.NewStdin()
	callback(keyPress, pos, string(line), block, stdout)
	defer stdout.Close()

	//fmt.Print(stdout.GetDataType(), "<--\r\n")

	ret := make(map[string]string)
	err := stdout.ReadMap(proc.ShellProcess.Config, func(key string, value string, last bool) {
		ret[key] = value
	})
	if err != nil {
		return false, false, []rune("Callback error: " + err.Error())
	}

	//fmt.Print(ret, stdout.GetDataType(), "<-\r\n")

	ignoreKey, err := types.ConvertGoType(ret["IgnoreKey"], types.Boolean)
	if err != nil {
		return false, false, []rune("Callback error: " + err.Error())
	}

	closeReadline, err := types.ConvertGoType(ret["CloseReadline"], types.Boolean)
	if err != nil {
		return false, false, []rune("Callback error: " + err.Error())
	}

	return ignoreKey.(bool), closeReadline.(bool), []rune(ret["HintText"])
}

func (evt *keyPressEvents) Dump() interface{} {
	dump := make(map[string]string)
	for e := range evt.events {
		dump[e] = string(evt.events[e])
	}
	return dump
}
