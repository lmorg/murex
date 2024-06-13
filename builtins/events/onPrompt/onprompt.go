package onprompt

import (
	"fmt"
	"sort"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/shell"
	"github.com/lmorg/murex/utils/lists"
)

const eventType = "onPrompt"

func init() {
	event := newOnPrompt()
	events.AddEventType(eventType, event, nil)
	shell.EventsPrompt = event.callback
}

// Interrupt is a JSONable structure passed to the murex function
type Interrupt struct {
	Name      string
	Operation string
	CmdLine   string
}

type promptEvent struct {
	Key     string
	Block   []rune
	FileRef *ref.File
}

type promptEvents struct {
	events []promptEvent
	//mutex  sync.Mutex
}

func newOnPrompt() *promptEvents {
	return new(promptEvents)
}

// Add a command to the onPrompt
func (evt *promptEvents) Add(name, interrupt string, block []rune, fileRef *ref.File) error {
	if err := isValidInterrupt(interrupt); err != nil {
		return err
	}

	//evt.mutex.Lock()

	key := events.CompileInterruptKey(interrupt, name)
	event := promptEvent{
		Key:     key,
		Block:   block,
		FileRef: fileRef,
	}

	i := evt.exists(key)
	if i == doesNotExist {
		evt.events = append(evt.events, event)
		sort.SliceStable(evt.events, func(i, j int) bool {
			return evt.events[i].Key < evt.events[j].Key
		})
	} else {
		evt.events[i] = event
	}

	//evt.mutex.Unlock()

	return nil
}

func (evt *promptEvents) Remove(key string) error {
	//evt.mutex.Lock()
	//defer evt.mutex.Unlock()

	i := evt.exists(key)
	if i != doesNotExist {
		events, err := lists.RemoveOrdered(evt.events, i)
		if err != nil {
			return fmt.Errorf("unable to remove %s: %s", key, err.Error())
		}
		evt.events = events
		return nil
	}

	var success bool
	for _, interrupt := range interrupts {
		newKey := events.CompileInterruptKey(interrupt, key)
		i = evt.exists(newKey)
		if i != doesNotExist {
			events, err := lists.RemoveOrdered(evt.events, i)
			if err != nil {
				return fmt.Errorf("unable to remove %s: %s", newKey, err.Error())
			}
			evt.events = events
			success = true
		}
	}

	if success {
		return nil
	}
	return fmt.Errorf("no %s event found called `%s`", eventType, key)
}

func (evt *promptEvents) callback(interrupt string, cmdLine []rune) {

	//evt.mutex.Lock()

	for i := range evt.events {
		key := events.GetInterruptFromKey(evt.events[i].Key)
		if key.Interrupt == interrupt {
			interruptValue := Interrupt{
				Name:      key.Name,
				Operation: interrupt,
				CmdLine:   string(cmdLine),
			}
			_, err := events.Callback(evt.events[i].Key, interruptValue, evt.events[i].Block, evt.events[i].FileRef, lang.ShellProcess.Stdout, lang.ShellProcess.Stderr, nil, false)
			if err != nil {
				lang.ShellProcess.Stderr.Writeln([]byte(fmt.Sprintf(
					"error in event callback: %s", err.Error(),
				)))
			}
		}
	}

	//evt.mutex.Unlock()
}

const doesNotExist = -1

func (evt *promptEvents) exists(key string) int {
	//evt.mutex.Lock()

	for i := range evt.events {
		if evt.events[i].Key == key {
			return i
		}
	}

	//evt.mutex.Unlock()

	return doesNotExist
}

func (evt *promptEvents) Dump() map[string]events.DumpT {
	dump := make(map[string]events.DumpT)

	//evt.mutex.Lock()

	for i := range evt.events {
		dump[evt.events[i].Key] = events.DumpT{
			Interrupt: events.GetInterruptFromKey(evt.events[i].Key).Interrupt,
			Block:     string(evt.events[i].Block),
			FileRef:   evt.events[i].FileRef,
		}
	}

	//evt.mutex.Unlock()

	return dump
}
