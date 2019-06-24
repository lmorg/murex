package lang

import (
	"errors"
	"sync"
)

// MurexPrivs is a table of private murex functions
type MurexPrivs struct {
	mutex sync.Mutex
	fn    []*murexPrivDetails
}

// murexPrivDetails is the properties for any given private murex function
type murexPrivDetails struct {
	Name    string
	Block   []rune
	Module  string
	Summary string
}

// NewMurexPrivs creates a new table of private murex functions
func NewMurexPrivs() (mf MurexPrivs) {
	return
}

// Define creates a private
func (mf *MurexPrivs) Define(name, module string, block []rune) error {
	//if mf.Exists(name, module) {
	//	return fmt.Errorf("private with the name `%s` already exists in module `%s`", name, module)
	//}

	summary := funcPrivSummary(block)

	mf.mutex.Lock()
	mf.fn = append(mf.fn, &murexPrivDetails{
		Name:    name,
		Block:   block,
		Module:  module,
		Summary: summary,
	})
	mf.mutex.Unlock()

	return nil
}

func (mf *MurexPrivs) get(name, module string) *murexPrivDetails {
	mf.mutex.Lock()
	for i := range mf.fn {
		if mf.fn[i].Name == name && mf.fn[i].Module == module {
			priv := mf.fn[i]
			mf.mutex.Unlock()
			return priv
		}
	}
	mf.mutex.Unlock()
	return nil
}

// Exists checks if private already created
func (mf *MurexPrivs) Exists(name, module string) (exists bool) {
	return mf.get(name, module) != nil
}

// Block returns private function code
func (mf *MurexPrivs) Block(name, module string) ([]rune, error) {
	priv := mf.get(name, module)

	if priv == nil {
		return nil, errors.New("Cannot locate private named `" + name + "`")
	}

	return priv.Block, nil
}

// Summary returns private's summary
func (mf *MurexPrivs) Summary(name, module string) (string, error) {
	priv := mf.get(name, module)

	if priv == nil {
		return "", errors.New("Cannot locate private named `" + name + "`")
	}

	return priv.Summary, nil
}

/*// Undefine deletes private from table
func (mf *MurexPrivs) Undefine(name string) error {
	mf.mutex.Lock()
	defer mf.mutex.Unlock()

	if mf.fn[name] == nil {
		return errors.New("Cannot locate function named `" + name + "`")
	}

	delete(mf.fn, name)
	return nil
}*/

// Dump list all private murex functions in table
func (mf *MurexPrivs) Dump() interface{} {
	type funcs struct {
		Name    string
		Module  string
		Summary string
		Block   string
	}

	var dump []funcs

	mf.mutex.Lock()
	for _, priv := range mf.fn {
		dump = append(dump, funcs{
			Name:    priv.Name,
			Module:  priv.Module,
			Summary: priv.Summary,
			Block:   string(priv.Block),
		})
	}
	mf.mutex.Unlock()

	return dump
}
