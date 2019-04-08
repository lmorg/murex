package lang

import (
	"errors"
	"strings"
	"sync"
)

// MurexFuncs is a table of murex functions
type MurexFuncs struct {
	mutex sync.Mutex
	fn    map[string]*murexFuncDetails
}

// MurexFuncDetails is the properties for any given murex function
type murexFuncDetails struct {
	Block   []rune
	Module  string
	Summary string
}

// NewMurexFuncs creates a new table of murex functions
func NewMurexFuncs() (mf MurexFuncs) {
	mf.fn = make(map[string]*murexFuncDetails)

	return
}

func funcPrivSummary(block []rune) string {
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
	return strings.TrimSpace(string(summary))
}

// Define creates a function
func (mf *MurexFuncs) Define(name, module string, block []rune) {
	summary := funcPrivSummary(block)

	mf.mutex.Lock()
	mf.fn[name] = &murexFuncDetails{
		Block:   block,
		Module:  module,
		Summary: summary,
	}

	mf.mutex.Unlock()
}

// get returns the function's details
func (mf *MurexFuncs) get(name string) *murexFuncDetails {
	mf.mutex.Lock()
	fn := mf.fn[name]
	mf.mutex.Unlock()
	return fn
}

// Exists checks if function already created
func (mf *MurexFuncs) Exists(name string) bool {
	mf.mutex.Lock()
	exists := mf.fn[name] != nil
	mf.mutex.Unlock()
	return exists
}

// Block returns function code
func (mf *MurexFuncs) Block(name string) ([]rune, error) {
	mf.mutex.Lock()
	fn := mf.fn[name]
	mf.mutex.Unlock()

	if fn == nil {
		return nil, errors.New("Cannot locate function named `" + name + "`")
	}

	return fn.Block, nil
}

// Summary returns functions summary
func (mf *MurexFuncs) Summary(name string) (string, error) {
	mf.mutex.Lock()
	fn := mf.fn[name]
	mf.mutex.Unlock()

	if fn == nil {
		return "", errors.New("Cannot locate function named `" + name + "`")
	}

	return fn.Summary, nil
}

// Undefine deletes function from table
func (mf *MurexFuncs) Undefine(name string) error {
	mf.mutex.Lock()
	defer mf.mutex.Unlock()

	if mf.fn[name] == nil {
		return errors.New("Cannot locate function named `" + name + "`")
	}

	delete(mf.fn, name)
	return nil
}

// Dump list all murex functions in table
func (mf *MurexFuncs) Dump() interface{} {
	type funcs struct {
		Summary string
		Module  string
		Block   string
	}

	dump := make(map[string]funcs)

	mf.mutex.Lock()
	for name, fn := range mf.fn {
		dump[name] = funcs{
			Summary: fn.Summary,
			Module:  fn.Module,
			Block:   string(fn.Block),
		}
	}
	mf.mutex.Unlock()

	return dump
}

// UpdateMap is used for auto-completions. It takes an existing map and updates it's values rather than copying data
func (mf *MurexFuncs) UpdateMap(m map[string]bool) {
	for name := range mf.fn {
		m[name] = true
	}
}
