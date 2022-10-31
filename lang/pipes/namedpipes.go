package pipes

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/lmorg/murex/builtins/pipes/null"
	"github.com/lmorg/murex/lang/stdio"
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

	if n.pipes[name].Pipe != nil {
		n.mutex.Unlock()
		return fmt.Errorf("named pipe `%s`already exists", name)
	}

	io, err := stdio.CreatePipe(pipeType, arguments)
	if err != nil {
		n.mutex.Unlock()
		return err
	}

	n.pipes[name] = pipe{
		Pipe: io,
		Type: pipeType,
	}

	io.Open()
	n.mutex.Unlock()
	return nil
}

// ExposePipe takes an existing stdio.Io interface and exposes it as a named pipe
func (n *Named) ExposePipe(name, pipeType string, io stdio.Io) error {
	n.mutex.Lock()

	if n.pipes[name].Pipe != nil {
		n.mutex.Unlock()
		return fmt.Errorf("named pipe `%s`already exists", name)
	}

	n.pipes[name] = pipe{
		Pipe: io,
		Type: pipeType,
	}

	n.mutex.Unlock()
	return nil
}

// Close a named pipe
func (n *Named) Close(name string) error {
	n.mutex.Lock()

	if n.pipes[name].Pipe == nil {
		n.mutex.Unlock()
		return fmt.Errorf("no pipe with the name `%s` exists", name)
	}

	if name == "null" {
		n.mutex.Unlock()
		return errors.New("null pipe must not be closed")
	}

	n.mutex.Unlock()

	go closePipe(n, name)
	return nil
}

func closePipe(n *Named, name string) {
	time.Sleep(2 * time.Second)

	n.mutex.Lock()

	n.pipes[name].Pipe.Close()
	delete(n.pipes, name)

	n.mutex.Unlock()
}

// Deletes a named pipe without closing it (careful using this!!!)
func (n *Named) Delete(name string) error {
	n.mutex.Lock()

	if n.pipes[name].Pipe == nil {
		n.mutex.Unlock()
		return fmt.Errorf("no pipe with the name `%s` exists", name)
	}

	if name == "null" {
		n.mutex.Unlock()
		return errors.New("null pipe must not be closed")
	}

	n.mutex.Unlock()

	delete(n.pipes, name)
	return nil
}

// Get a named pipe interface from the named pipe table
func (n *Named) Get(name string) (stdio.Io, error) {
	retries := 0

try:
	n.mutex.Lock()

	if n.pipes[name].Pipe == nil {
		n.mutex.Unlock()

		if retries == 5 {
			return nil, fmt.Errorf("no pipe with the name `%s` exists, timed out waiting for pipe to be created", name)
		}
		time.Sleep(100 * time.Millisecond)
		retries++
		goto try
	}

	p := n.pipes[name].Pipe
	n.mutex.Unlock()
	return p, nil
}

// Dump returns the named pipe table in a format that can be serialised into JSON
func (n *Named) Dump() map[string]string {
	dump := make(map[string]string)
	n.mutex.Lock()
	for name := range n.pipes {
		dump[name] = n.pipes[name].Type
	}
	n.mutex.Unlock()
	return dump
}
