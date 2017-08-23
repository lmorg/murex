package proc

import (
	"errors"
	"sync"
)

// MurexFuncs is a table of murex functions
type MurexFuncs struct {
	mutex sync.Mutex
	funcs map[string][]rune
}

// NewMurexFuncs creates a new table of murex functions
func NewMurexFuncs() (fn MurexFuncs) {
	fn.funcs = make(map[string][]rune)
	return
}

// Define creates a function
func (fn *MurexFuncs) Define(name string, block []rune) {
	fn.mutex.Lock()
	fn.funcs[name] = block
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
func (fn *MurexFuncs) Dump() (dump map[string]string) {
	dump = make(map[string]string)
	fn.mutex.Lock()
	for name := range fn.funcs {
		dump[name] = string(fn.funcs[name])
	}
	fn.mutex.Unlock()
	return
}
