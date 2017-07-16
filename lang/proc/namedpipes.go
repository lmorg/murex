package proc

import (
	"errors"
	"github.com/lmorg/murex/lang/proc/streams"
	"sync"
)

type Named struct {
	pipes map[string]streams.Io
	mutex sync.Mutex
}

func NewNamed() (n Named) {
	n.pipes = make(map[string]streams.Io)
	n.pipes["null"] = new(streams.Null)
	return
}

func (n *Named) Create(name string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[name] != nil {
		return errors.New("Named pipe `" + name + "`already exists.")
	}

	n.pipes[name] = streams.NewStdin()
	n.pipes[name].MakePipe()
	return nil
}

func (n *Named) Close(name string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()

	if n.pipes[name] == nil {
		return errors.New("No pipe with the name `" + name + "` exists.")
	}

	n.pipes[name].UnmakeParent()
	n.pipes[name].Close()
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
