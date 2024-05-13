package oncommandcompletion

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/lmorg/murex/builtins/events"
	"github.com/lmorg/murex/builtins/pipes/term"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/ref"
)

const eventType = "onCommandCompletion"

func init() {
	event := newCommandCompletion()
	events.AddEventType(eventType, event, nil)
	lang.ShellProcess.CCEvent = event.callback
	lang.ShellProcess.CCExists = event.exists
}

// Interrupt is a JSONable structure passed to the murex function
type Interrupt struct {
	Command    string
	Parameters []string
	Stdout     string
	Stderr     string
	ExitNum    int
}

type commandCompletionEvent struct {
	Command string
	Block   []rune
	FileRef *ref.File
}

type commandCompletionEvents struct {
	events map[string]commandCompletionEvent
	mutex  sync.Mutex
}

func newCommandCompletion() *commandCompletionEvents {
	cce := new(commandCompletionEvents)
	cce.events = make(map[string]commandCompletionEvent)
	return cce
}

// Add a command to the onCommandCompleteionEvents
func (evt *commandCompletionEvents) Add(name, command string, block []rune, fileRef *ref.File) error {
	if command == "exec" || command == "fexec" {
		return errors.New("for safety reasons, you cannot assign this event against `exec` nor `fexec`")
	}

	evt.mutex.Lock()

	evt.events[name] = commandCompletionEvent{
		Command: command,
		Block:   block,
		FileRef: fileRef,
	}

	evt.mutex.Unlock()

	return nil
}

func (evt *commandCompletionEvents) Remove(name string) error {
	evt.mutex.Lock()
	if evt.events[name].FileRef == nil {
		evt.mutex.Unlock()
		return fmt.Errorf("%s event %s does not exist", eventType, name)
	}

	delete(evt.events, name)

	evt.mutex.Unlock()
	return nil
}

func (evt *commandCompletionEvents) callback(command string, p *lang.Process) {
	if command == "exec" || command == "fexec" {
		return
	}

	evt.mutex.Lock()

	//command := p.Name.String()
	parameters := p.Parameters.StringArray()

	if p.Name.String() == "exec" && len(parameters) > 0 {
		command = parameters[0]
		parameters = parameters[1:]
	}

	for name, cce := range evt.events {
		if cce.Command == command {
			evt.mutex.Unlock()
			cce.execEvent(name, parameters, p)
			evt.mutex.Lock()
		}
	}

	evt.mutex.Unlock()
}

func (evt *commandCompletionEvents) exists(command string) bool {
	evt.mutex.Lock()

	for name := range evt.events {
		if evt.events[name].Command == command {
			evt.mutex.Unlock()
			return true
		}
	}

	evt.mutex.Unlock()
	return false
}

func writeErr(desc string, err error, name string, cmd string) {
	os.Stderr.WriteString(
		fmt.Sprintf(
			"ERROR: cannot execute event '%s' (command `%s`): %s: %s",
			name, cmd, desc, err))
}

func (cce *commandCompletionEvent) execEvent(name string, parameters []string, p *lang.Process) {
	var err error

	stdout := fmt.Sprintf("%d-out", p.Id)
	stderr := fmt.Sprintf("%d-err", p.Id)

	err = lang.GlobalPipes.ExposePipe(stdout, "std", p.CCOut)
	if err != nil {
		writeErr(fmt.Sprintf("cannot expose stdout pipe '%s'", stdout), err, name, cce.Command)
		return
	}

	err = lang.GlobalPipes.ExposePipe(stderr, "std", p.CCErr)
	if err != nil {
		writeErr(fmt.Sprintf("cannot expose stderr pipe '%s'", stderr), err, name, cce.Command)
		return
	}

	interrupt := Interrupt{
		Command:    cce.Command,
		Parameters: parameters,
		Stdout:     stdout,
		Stderr:     stderr,
		ExitNum:    p.ExitNum,
	}

	_, err = events.Callback(name, interrupt, cce.Block, cce.FileRef, term.NewErr(false), term.NewErr(false), nil, false)
	_ = lang.GlobalPipes.Delete(stdout) // we don't actually care about any errors here
	_ = lang.GlobalPipes.Delete(stderr) // nor here either
	if err != nil {
		writeErr("callback failed", err, name, cce.Command)
		return
	}
}

func (evt *commandCompletionEvents) Dump() map[string]events.DumpT {
	dump := make(map[string]events.DumpT)

	evt.mutex.Lock()

	for name, event := range evt.events {
		dump[name] = events.DumpT{
			Interrupt: event.Command,
			Block:     string(event.Block),
			FileRef:   event.FileRef,
		}
	}

	evt.mutex.Unlock()
	return dump
}
