package pipes

import (
	"errors"
	"github.com/lmorg/murex/lang/proc/streams"
	"sync"
	"time"
)

// Named is a table of created named pipes
type Named struct {
	pipes map[string]streams.Io
	types map[string]PipeTypes
	mutex sync.Mutex
}

// NewNamed creates a new table of named pipes
func NewNamed() (n Named) {
	n.pipes = make(map[string]streams.Io)
	n.types = make(map[string]PipeTypes)

	n.pipes["null"] = new(streams.Null)
	n.types["null"] = pipeNull
	return
}

// CreatePipe creates a named pipe using the stdin interface
func (n *Named) CreatePipe(name string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[name] != nil {
		return errors.New("Named pipe `" + name + "`already exists.")
	}

	n.pipes[name] = streams.NewStdin()
	n.pipes[name].MakePipe()
	n.types[name] = pipeStream
	return nil
}

// CreateFile creates a named pipe using the file interface
func (n *Named) CreateFile(pipename string, filename string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[pipename] != nil {
		return errors.New("Named pipe `" + pipename + "`already exists.")
	}

	file, err := streams.NewFile(filename)
	if err != nil {
		return err
	}

	n.pipes[pipename] = file
	n.pipes[pipename].MakePipe()
	n.types[pipename] = pipeFileWriter
	return nil
}

// CreateDialer creates a named pipe using the net dialer interface
func (n *Named) CreateDialer(pipename, protocol, address string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[pipename] != nil {
		return errors.New("Named pipe `" + pipename + "`already exists.")
	}

	file, err := streams.NewDialer(protocol, address)
	if err != nil {
		return err
	}

	n.pipes[pipename] = file
	n.pipes[pipename].MakePipe()
	n.types[pipename] = pipeNetDialer
	return nil
}

// CreateListener creates a named pipe using the net listener interface
func (n *Named) CreateListener(pipename, protocol, address string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[pipename] != nil {
		return errors.New("Named pipe `" + pipename + "`already exists.")
	}

	file, err := streams.NewListener(protocol, address)
	if err != nil {
		return err
	}

	n.pipes[pipename] = file
	n.pipes[pipename].MakePipe()
	n.types[pipename] = pipeNetListener
	return nil
}

// Close a named pipe
func (n *Named) Close(name string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[name] == nil {
		return errors.New("No pipe with the name `" + name + "` exists.")
	}

	if name == "null" {
		return errors.New("I will not close the `null` device!")
	}

	n.pipes[name].UnmakeParent()
	n.pipes[name].Close()

	switch n.types[name] {
	case pipeStream:
		go func() {
			time.Sleep(10 * time.Second)
			delete(n.pipes, name)
			delete(n.types, name)
		}()
	case pipeNull, pipeFileWriter, pipeNetDialer, pipeNetListener:
		delete(n.pipes, name)
		delete(n.types, name)
	default:
		return errors.New("Invalid pipe ID!")
	}

	return nil
}

// Get a named pipe interface from the named pipe table
func (n *Named) Get(name string) (streams.Io, error) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[name] == nil {
		return nil, errors.New("No pipe with the name `" + name + "` exists.")
	}

	return n.pipes[name], nil
}

// Dump returns the named pipe table in a format that can be serialised into JSON
func (n *Named) Dump() (dump map[string]string) {
	dump = make(map[string]string)
	n.mutex.Lock()
	for name := range n.types {
		dump[name] = n.types[name].String()
	}
	n.mutex.Unlock()
	return
}
