package open

import (
	"errors"
	"sync"

	"github.com/lmorg/murex/lang/ref"
)

// OpenAgents is the exported table of `open`'s helper functions
var OpenAgents = newOpenAgents()

type openBlocks struct {
	Block   []rune
	FileRef *ref.File
}

func newOpenAgents() *openAgents {
	oa := new(openAgents)
	oa.agents = make(map[string]*openBlocks)
	return oa
}

type openAgents struct {
	mutex  sync.Mutex
	agents map[string]*openBlocks
}

// Get the murex code block for a particular murex data type
func (oa *openAgents) Get(dataType string) (*openBlocks, error) {
	oa.mutex.Lock()
	ob := oa.agents[dataType]
	oa.mutex.Unlock()

	if ob == nil {
		return nil, errors.New("No agent set for that data type")
	}

	return ob, nil
}

// Set the murex code block for a particular murex data type
func (oa *openAgents) Set(dataType string, block []rune, fileRef *ref.File) {
	oa.mutex.Lock()
	defer oa.mutex.Unlock()

	oa.agents[dataType] = &openBlocks{
		Block:   block,
		FileRef: fileRef,
	}
}

// Unset removes an associated code block for a particular data type
func (oa *openAgents) Unset(dataType string) error {
	oa.mutex.Lock()
	defer oa.mutex.Unlock()

	if oa.agents[dataType] == nil {
		return errors.New("No agent set for that data type")
	}

	oa.agents[dataType] = nil
	return nil
}

// Dump returns the entire OpenAgent table
func (oa *openAgents) Dump() interface{} {
	oa.mutex.Lock()
	defer oa.mutex.Unlock()

	type dumpedBlocks struct {
		Block   string
		FileRef *ref.File
	}

	dump := make(map[string]dumpedBlocks)
	for dt, ob := range oa.agents {
		dump[dt] = dumpedBlocks{
			Block:   string(ob.Block),
			FileRef: ob.FileRef,
		}
	}

	return dump
}
