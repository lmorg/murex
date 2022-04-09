package onsecondselapsed

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
)

const eventType = "onSecondsElapsed"

func init() {
	t := newTimer()
	events.AddEventType(eventType, t, nil)
	go t.init()
}

type timer struct {
	mutex  sync.Mutex
	events []timeEvent
}

type timeEvent struct {
	Name     string
	Interval int
	Block    []rune
	state    int
	FileRef  *ref.File
}

func newTimer() (t *timer) {
	t = new(timer)
	t.events = make([]timeEvent, 0)
	return
}

func (t *timer) init() {
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
					t.events[i].Block,
					t.events[i].FileRef,
					lang.ShellProcess.Stdout,
					true,
				)
			}
		}
		t.mutex.Unlock()
	}
}

// Add a path to the watch event list
func (t *timer) Add(name, interrupt string, block []rune, fileRef *ref.File) (err error) {
	interval, err := strconv.Atoi(interrupt)
	if err != nil {
		return errors.New("interrupt should be an integer for `" + eventType + "` events")
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
		FileRef:  fileRef,
	})

	return
}

// Remove a path to the watch event list
func (t *timer) Remove(name string) (err error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if len(t.events) == 0 {
		return errors.New("no events have been created for this listener")
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

	return errors.New("no event found for this listener with the name `" + name + "`.")
}

// Dump returns all the events in w
func (t *timer) Dump() interface{} {
	type te struct {
		Interval int
		Block    string
		FileRef  *ref.File
	}

	dump := make(map[string]te)

	t.mutex.Lock()

	for i := range t.events {
		dump[t.events[i].Name] = te{
			Interval: t.events[i].Interval,
			Block:    string(t.events[i].Block),
			FileRef:  t.events[i].FileRef,
		}
	}

	t.mutex.Unlock()

	return dump
}
