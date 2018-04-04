package proc

import (
	"errors"
	"strings"
	"sync"
)

// MurexFuncs is a table of murex functions
type MurexFuncs struct {
	mutex   sync.Mutex
	funcs   map[string][]rune
	digests map[string]string
}

// NewMurexFuncs creates a new table of murex functions
func NewMurexFuncs() (fn MurexFuncs) {
	fn.funcs = make(map[string][]rune)
	fn.digests = make(map[string]string)
	return
}

// Define creates a function
func (fn *MurexFuncs) Define(name string, block []rune) {
	fn.mutex.Lock()
	fn.funcs[name] = block

	var (
		line1   bool
		comment bool
		digest  []rune
	)

	for _, r := range block {
		switch {
		case r == '\n' && !line1:
			line1 = true

		case r == '\n' && comment:
			goto exitLoop

		case !line1:
			continue

		case r == '#':
			comment = true

		case !comment:
			continue

		case r == '\t':
			digest = append(digest, ' ', ' ', ' ', ' ')

		case r == '\r':
			continue

		case comment:
			digest = append(digest, r)
		}
	}

exitLoop:
	fn.digests[name] = strings.TrimSpace(string(digest))

	fn.mutex.Unlock()
}

// Exists checks if function already created
func (fn *MurexFuncs) Exists(name string) bool {
	fn.mutex.Lock()
	exists := len(fn.funcs[name]) > 0
	fn.mutex.Unlock()
	return exists
}

// Block returns function code
func (fn *MurexFuncs) Block(name string) (block []rune, err error) {
	fn.mutex.Lock()
	defer fn.mutex.Unlock()
	if len(fn.funcs[name]) == 0 {
		return nil, errors.New("Cannot locate function named `" + name + "`")
	}
	block = fn.funcs[name]
	return block, err
}

// Digest returns functions digest
func (fn *MurexFuncs) Digest(name string) (digest string, err error) {
	fn.mutex.Lock()
	defer fn.mutex.Unlock()
	if len(fn.funcs[name]) == 0 {
		return "", errors.New("Cannot locate function named `" + name + "`")
	}

	digest = fn.digests[name]
	return digest, err
}

// Undefine deletes function from table
func (fn *MurexFuncs) Undefine(name string) error {
	fn.mutex.Lock()
	defer fn.mutex.Unlock()
	if len(fn.funcs[name]) == 0 {
		return errors.New("Cannot locate function named `" + name + "`")
	}
	delete(fn.funcs, name)
	return nil
}

// Dump list all murex functions in table
func (fn *MurexFuncs) Dump() interface{} {
	type t struct {
		Digest string
		Block  string
	}

	dump := make(map[string]t)

	fn.mutex.Lock()

	for name := range fn.funcs {
		dump[name] = t{
			Block:  string(fn.funcs[name]),
			Digest: fn.digests[name],
		}
	}

	fn.mutex.Unlock()
	return dump
}

// UpdateMap is used for auto-completions. It takes an existing map and updates it's values rather than copying data
func (fn *MurexFuncs) UpdateMap(m map[string]bool) {
	for name := range fn.funcs {
		m[name] = true
	}
}
