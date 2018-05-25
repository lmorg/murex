package open

import (
	"errors"
	"sync"
)

var OpenAgents *openAgents = newOpenAgents()

func newOpenAgents() *openAgents {
	oa := new(openAgents)
	oa.agents = make(map[string][]rune)
	return oa
}

type openAgents struct {
	mutex  sync.Mutex
	agents map[string][]rune
}

func (oa *openAgents) Get(dataType string) ([]rune, error) {
	oa.mutex.Lock()
	r := oa.agents[dataType]
	oa.mutex.Unlock()

	if r == nil {
		return nil, errors.New("No agent set for that data type")
	}

	return r, nil
}

func (oa *openAgents) Set(dataType string, block []rune) {
	oa.mutex.Lock()
	defer oa.mutex.Unlock()

	oa.agents[dataType] = block
}

func (oa *openAgents) Unset(dataType string) error {
	oa.mutex.Lock()
	defer oa.mutex.Unlock()

	if oa.agents[dataType] == nil {
		return errors.New("No agent set for that data type")
	}

	oa.agents[dataType] = nil
	return nil
}

func (oa *openAgents) Dump() map[string]string {
	oa.mutex.Lock()
	defer oa.mutex.Unlock()

	dump := make(map[string]string)
	for dt := range oa.agents {
		dump[dt] = string(oa.agents[dt])
	}

	return dump
}
