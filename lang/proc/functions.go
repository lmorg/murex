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
	summary map[string]string
}

// NewMurexFuncs creates a new table of murex functions
func NewMurexFuncs() (fn MurexFuncs) {
	fn.funcs = make(map[string][]rune)
	fn.summary = make(map[string]string)
	return
}

// Define creates a function
func (fn *MurexFuncs) Define(name string, block []rune) {
	fn.mutex.Lock()
	fn.funcs[name] = block

	var (
		line1   bool
		comment bool
		summary []rune
	)

	for _, r := range block {
		switch {
		case r == '\n' && !line1:
			line1 = true

		case r == '\n':
			goto exitLoop

		case !line1:
			continue

		case r == '#':
			comment = true

		case !comment:
			continue

		case r == '\t':
			summary = append(summary, ' ', ' ', ' ', ' ')

		case r == '\r':
			continue

		case comment:
			summary = append(summary, r)
		}
	}

exitLoop:
	fn.summary[name] = strings.TrimSpace(string(summary))

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

// Summary returns functions summary
func (fn *MurexFuncs) Summary(name string) (summary string, err error) {
	fn.mutex.Lock()
	defer fn.mutex.Unlock()
	if len(fn.funcs[name]) == 0 {
		return "", errors.New("Cannot locate function named `" + name + "`")
	}

	summary = fn.summary[name]
	return summary, err
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
		Summary string
		Block   string
	}

	dump := make(map[string]t)

	fn.mutex.Lock()

	for name := range fn.funcs {
		dump[name] = t{
			Block:   string(fn.funcs[name]),
			Summary: fn.summary[name],
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
