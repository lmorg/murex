package hintsummary

import (
	"fmt"
	"sync"
)

// HintSummary is a thread-safe map for storing the hint summary overrides
type HintSummary struct {
	mutex sync.Mutex
	m     map[string]string
}

// New creates a new instance of the hint summary override structure
func New() *HintSummary {
	hs := new(HintSummary)
	hs.m = make(map[string]string)
	return hs
}

// Get returns the hint summary override
func (hs *HintSummary) Get(exe string) string {
	hs.mutex.Lock()
	s := hs.m[exe]
	hs.mutex.Unlock()
	return s
}

// Set stores the hint summary override for a specific executable
func (hs *HintSummary) Set(exe, summary string) {
	hs.mutex.Lock()
	hs.m[exe] = summary
	hs.mutex.Unlock()
}

// Delete removes a hint summary for a specific executable
func (hs *HintSummary) Delete(exe string) error {
	hs.mutex.Lock()
	if hs.m[exe] == "" {
		hs.mutex.Unlock()
		return fmt.Errorf("No summary set for '%s'", exe)
	}

	delete(hs.m, exe)
	hs.mutex.Unlock()
	return nil
}

// Dump returns a JSON-able structure of the stored hint summary overrides
func (hs *HintSummary) Dump() map[string]string {
	hs.mutex.Lock()
	m := hs.m
	hs.mutex.Unlock()
	return m
}
