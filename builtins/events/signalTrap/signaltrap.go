package signaltrap

import (
	"fmt"
	"os"
	"sort"
	"syscall"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/utils/lists"
)

var eventType = "signalTrap"

const errInvalidInterrupt = `invalid interrupt, '%s'. Expecting a signal name, like "SIGINT"`

func init() {
	event := newSignalTrap()
	events.AddEventType(eventType, event, nil)
	go event.signalHandler()
}

// Interrupt is a JSONable structure passed to the murex function
type Interrupt struct {
	Name   string
	Signal string
}

type sigEvent struct {
	Key     string
	Block   []rune
	FileRef *ref.File
}

type sigEvents struct {
	events []sigEvent
	//mutex  sync.Mutex
}

func newSignalTrap() *sigEvents {
	return new(sigEvents)
}

// Add a command to the onPrompt
func (evt *sigEvents) Add(name, interrupt string, block []rune, fileRef *ref.File) error {
	if !isValidInterrupt(interrupt) {
		return fmt.Errorf(errInvalidInterrupt, interrupt)
	}

	//evt.mutex.Lock()

	key := compileInterruptKey(interrupt, name)
	event := sigEvent{
		Key:     key,
		Block:   block,
		FileRef: fileRef,
	}

	if evt.noRegisteredSignal(interrupt) {
		if err := register(interrupt); err != nil {
			return err
		}
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

func (evt *sigEvents) Remove(key string) error {
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

	var (
		success bool
		remove  []syscall.Signal
	)

	for name := range interrupts {
		newKey := compileInterruptKey(name, key)
		i = evt.exists(newKey)
		if i != doesNotExist {
			events, err := lists.RemoveOrdered(evt.events, i)
			if err != nil {
				return fmt.Errorf("unable to remove %s: %s", newKey, err.Error())
			}
			evt.events = events
			success = true
			remove = append(remove, interrupts[name])
		}
	}

	if success {
		for i := range remove {
			if evt.noRegisteredSignal(remove[i].String()) {
				deregister(remove[i])
			}
		}
		return nil
	}
	return fmt.Errorf("no %s event found called `%s`", eventType, key)
}

func (evt *sigEvents) callback(sig os.Signal) {
	var interrupt string
	for name, signal := range interrupts {
		if signal.String() == sig.String() {
			interrupt = name
			goto event
		}
	}

	panic(fmt.Sprintf("unknown signal: %V", sig))

event:

	//evt.mutex.Lock()

	for i := range evt.events {
		split := getInterruptFromKey(evt.events[i].Key)
		if split[0] == interrupt {
			interruptValue := Interrupt{
				Name:   split[1],
				Signal: interrupt,
			}
			events.Callback(evt.events[i].Key, interruptValue, evt.events[i].Block, evt.events[i].FileRef, lang.ShellProcess.Stdout, false)
		}
	}

	//evt.mutex.Unlock()
}

const doesNotExist = -1

func (evt *sigEvents) exists(key string) int {
	//evt.mutex.Lock()

	for i := range evt.events {
		if evt.events[i].Key == key {
			return i
		}
	}

	//evt.mutex.Unlock()

	return doesNotExist
}

func (evt *sigEvents) noRegisteredSignal(sig string) bool {
	//evt.mutex.Lock()

	for i := range evt.events {
		split := getInterruptFromKey(evt.events[i].Key)
		if split[0] == sig {
			return false
		}
	}

	//evt.mutex.Unlock()

	return true
}

func (evt *sigEvents) signalHandler() {
	for {
		sig := <-signalChan
		evt.callback(sig)
	}
}

func (evt *sigEvents) Dump() map[string]events.DumpT {
	dump := make(map[string]events.DumpT)

	//evt.mutex.Lock()

	for i := range evt.events {
		dump[evt.events[i].Key] = events.DumpT{
			Interrupt: getInterruptFromKey(evt.events[i].Key)[0],
			Block:     string(evt.events[i].Block),
			FileRef:   evt.events[i].FileRef,
		}
	}

	//evt.mutex.Unlock()

	return dump
}
