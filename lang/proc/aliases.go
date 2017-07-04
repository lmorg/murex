package proc

import (
	"errors"
	"sync"
)

type Aliases struct {
	aliases map[string][]rune
	mutex   sync.Mutex
}

func NewAliases() (a Aliases) {
	//a = new(Aliases)
	a.aliases = make(map[string][]rune)
	return
}

func (a *Aliases) Add(name string, block []rune) {
	a.mutex.Lock()
	a.aliases[name] = block
	a.mutex.Unlock()
}

func (a *Aliases) Exists(name string) bool {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return len(a.aliases[name]) > 0
}

func (a *Aliases) Get(name string) (block []rune) {
	a.mutex.Lock()
	block = a.aliases[name]
	a.mutex.Unlock()
	return
}

func (a *Aliases) Delete(name string) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if len(a.aliases[name]) == 0 {
		return errors.New("Alias does not exist.")
	}
	delete(a.aliases, name)
	return nil
}

func (a *Aliases) Dump() map[string]string {
	dump := make(map[string]string)
	a.mutex.Lock()
	for name := range a.aliases {
		dump[name] = string(a.aliases[name])
	}
	a.mutex.Unlock()
	return dump
}
