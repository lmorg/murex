package proc

import (
	"errors"
	"github.com/lmorg/murex/lang/proc/streams"
	"sync"
	"time"
)

type Named struct {
	pipes map[string]streams.Io
	types map[string]int
	mutex sync.Mutex
}

const (
	npStream = 1 + iota
	npFile
	npNull
)

func NewNamed() (n Named) {
	n.pipes = make(map[string]streams.Io)
	n.types = make(map[string]int)

	n.pipes["null"] = new(streams.Null)
	n.types["null"] = npNull
	return
}

func (n *Named) CreatePipe(name string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[name] != nil {
		return errors.New("Named pipe `" + name + "`already exists.")
	}

	n.pipes[name] = streams.NewStdin()
	n.pipes[name].MakePipe()
	n.types[name] = npStream
	return nil
}

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
	n.types[pipename] = npFile
	return nil
}

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
	case npStream:
		go func() {
			time.Sleep(10 * time.Second)
			delete(n.pipes, name)
			delete(n.types, name)
		}()
	case npNull:
		delete(n.pipes, name)
		delete(n.types, name)
	case npFile:
		delete(n.pipes, name)
		delete(n.types, name)
	default:
		panic("Invalid pipe ID!")
	}

	return nil
}

func (n *Named) Get(name string) (streams.Io, error) {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[name] == nil {
		return nil, errors.New("No pipe with the name `" + name + "` exists.")
	}

	return n.pipes[name], nil
}
