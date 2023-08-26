package lang

import (
	"fmt"
	"sync"

	"github.com/lmorg/murex/lang/ref"
)

type Alias struct {
	Alias   []string
	FileRef *ref.File
}

// Aliases is a table of aliases
type Aliases struct {
	aliases map[string]Alias
	mutex   sync.Mutex
}

// NewAliases creates a new table of aliases
func NewAliases() (a Aliases) {
	a.aliases = make(map[string]Alias)
	return
}

// Add creates an alias
func (a *Aliases) Add(name string, alias []string, fileRef *ref.File) {
	a.mutex.Lock()
	a.aliases[name] = Alias{
		Alias:   alias,
		FileRef: fileRef,
	}
	a.mutex.Unlock()
}

// Exists checks if alias exists in table
func (a *Aliases) Exists(name string) (exists bool) {
	a.mutex.Lock()
	_, exists = a.aliases[name]
	a.mutex.Unlock()
	return exists
}

// Get the aliased code
func (a *Aliases) Get(name string) []string {
	a.mutex.Lock()
	alias, ok := a.aliases[name]
	a.mutex.Unlock()

	if !ok {
		return nil
	}

	return alias.Alias
}

// Delete an alias
func (a *Aliases) Delete(name string) error {
	a.mutex.Lock()

	if _, ok := a.aliases[name]; !ok {
		a.mutex.Unlock()
		return fmt.Errorf("no alias named '%s' exists", name)
	}
	delete(a.aliases, name)
	a.mutex.Unlock()
	return nil
}

// Dump returns the complete alias table
func (a *Aliases) Dump() map[string]Alias {
	a.mutex.Lock()
	dump := a.aliases
	a.mutex.Unlock()
	return dump
}

// UpdateMap is used for auto-completions. It takes an existing map and updates it's values rather than copying data
func (a *Aliases) UpdateMap(m map[string]bool) {
	for name := range a.aliases {
		m[name] = true
	}
}
