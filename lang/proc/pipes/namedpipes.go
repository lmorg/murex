package pipes

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/lang/proc/stdio"
)

// Named is a table of created named pipes
type Named struct {
	pipes map[string]pipe
	mutex sync.Mutex
}

type pipe struct {
	Pipe stdio.Io
	Type string
}

// NewNamed creates a new table of named pipes
func NewNamed() (n Named) {
	n.pipes = make(map[string]pipe)

	n.pipes["null"] = pipe{
		Pipe: new(null.Null),
		Type: "null",
	}

	return
}

// CreatePipe creates a named pipe using the stdin interface
func (n *Named) CreatePipe(name, pipeType, arguments string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[name].Pipe != nil {
		return fmt.Errorf("Named pipe `%s`already exists", name)
	}

	io, err := stdio.CreatePipe(pipeType, arguments)
	if err != nil {
		return err
	}

	n.pipes[name] = pipe{
		Pipe: io,
		Type: pipeType,
	}

	io.MakePipe()
	return nil
}

// Close a named pipe
func (n *Named) Close(name string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[name].Pipe == nil {
		return fmt.Errorf("No pipe with the name `%s` exists", name)
	}

	if name == "null" {
		return errors.New("null pipe must not be closed")
	}

	n.pipes[name].Pipe.Close()

	go func() {
		// 3 second grace period before garbage collection - just to give any buffers chance to flush within murex code
		// (really this is only needed for the standard streamer but it doesn't do any damage to have all pipes behave the same)
		time.Sleep(3 * time.Second)
		delete(n.pipes, name)
	}()

	return nil
}

// Get a named pipe interface from the named pipe table
func (n *Named) Get(name string) (stdio.Io, error) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[name].Pipe == nil {
		return nil, fmt.Errorf("No pipe with the name `%s` exists", name)
	}

	return n.pipes[name].Pipe, nil
}

// Dump returns the named pipe table in a format that can be serialised into JSON
func (n *Named) Dump() (dump map[string]string) {
	dump = make(map[string]string)
	n.mutex.Lock()
	for name := range n.pipes {
		dump[name] = n.pipes[name].Type
	}
	n.mutex.Unlock()
	return
}
