package open

import (
	"errors"
	"sync"
)

// OpenAgents is the exported table of `open`'s helper functions
var OpenAgents = newOpenAgents()

func newOpenAgents() *openAgents {
	oa := new(openAgents)
	oa.agents = make(map[string][]rune)
	return oa
}

type openAgents struct {
	mutex  sync.Mutex
	agents map[string][]rune
}

// Get the murex code block for a particular murex data type
func (oa *openAgents) Get(dataType string) ([]rune, error) {
	oa.mutex.Lock()
	r := oa.agents[dataType]
	oa.mutex.Unlock()

	if r == nil {
		return nil, errors.New("No agent set for that data type")
	}

	return r, nil
}

// Set the murex code block for a particular murex data type
func (oa *openAgents) Set(dataType string, block []rune) {
	oa.mutex.Lock()
	defer oa.mutex.Unlock()

	oa.agents[dataType] = block
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
func (oa *openAgents) Dump() map[string]string {
	oa.mutex.Lock()
	defer oa.mutex.Unlock()

	dump := make(map[string]string)
	for dt := range oa.agents {
		dump[dt] = string(oa.agents[dt])
	}

	return dump
}
