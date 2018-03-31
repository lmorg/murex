package onSecondsElapsed

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang/proc"
)

func init() {
	events.AddEventType(eventType, newTimer())
}

const eventType = "onSecondsElapsed"

type timer struct {
	error  error
	mutex  sync.Mutex
	events []timeEvent
}

type timeEvent struct {
	Name     string
	Interval int
	Block    []rune
	state    int
}

func newTimer() (t *timer) {
	t = new(timer)
	t.events = make([]timeEvent, 0)
	return
}

// Add a path to the watch event list
func (t *timer) Add(name, interrupt string, block []rune) (err error) {
	interval, err := strconv.Atoi(interrupt)
	if err != nil {
		return errors.New("Interrupt should be an integer for `" + eventType + "` events.")
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	for i := range t.events {
		if t.events[i].Name == name {
			t.events[i].Interval = interval
			t.events[i].Block = block
			return nil
		}
	}

	t.events = append(t.events, timeEvent{
		Name:     name,
		Interval: interval,
		Block:    block,
	})

	return
}

// Remove a path to the watch event list
func (t *timer) Remove(name string) (err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if len(t.events) == 0 {
		return errors.New("No events have been created for this listener.")
	}

	for i := range t.events {
		if t.events[i].Name == name {
			switch {
			case len(t.events) == 1:
				t.events = make([]timeEvent, 0)
			case i == 0:
				t.events = t.events[1:]
			case i == len(t.events)-1:
				t.events = t.events[:len(t.events)-1]
			default:
				t.events = append(t.events[:i], t.events[i+1:]...)
			}
			return nil
		}
	}

	return errors.New("No event found for this listener with the name `" + name + "`.")
}

// Init starts a new watch event loop
func (t *timer) Init() {
	for {
		time.Sleep(1 * time.Second)

		t.mutex.Lock()
		for i := range t.events {
			t.events[i].state++
			if t.events[i].state == t.events[i].Interval {
				t.events[i].state = 0
				go events.Callback(
					t.events[i].Name,
					t.events[i].Interval,
					fmt.Sprintf("%s=%d", t.events[i].Name, t.events[i].Interval),
					t.events[i].Block,
					proc.ShellProcess.Stdout,
				)
			}
		}
		t.mutex.Unlock()
	}
}

// Dump returns all the events in w
func (t *timer) Dump() interface{} {
	type te struct {
		Name     string
		Interval int
		Block    string
	}

	t.mutex.Lock()
	dump := make([]te, len(t.events))

	for i := range t.events {
		dump[i] = te{
			Name:     t.events[i].Name,
			Interval: t.events[i].Interval,
			Block:    string(t.events[i].Block),
		}
	}

	t.mutex.Unlock()

	return dump
}
