package lang

import (
	"fmt"
	"strings"
	"sync"

	"github.com/lmorg/murex/lang/ref"
)

func ErrPrivateNotFound(module string) error {
	return fmt.Errorf("no private functions exist for module `%s`", module)
}

type privateFunctions struct {
	module map[string]*MurexFuncs
	mutex  sync.RWMutex
}

func NewMurexPrivs() *privateFunctions {
	pf := new(privateFunctions)
	pf.module = make(map[string]*MurexFuncs)
	return pf
}

func (pf *privateFunctions) Define(name string, parameters []MurexFuncParam, block []rune, fileRef *ref.File) {
	pf.mutex.Lock()

	if pf.module[fileRef.Source.Module] == nil {
		pf.module[fileRef.Source.Module] = NewMurexFuncs()
	}

	pf.module[fileRef.Source.Module].Define(name, parameters, block, fileRef)

	pf.mutex.Unlock()
}

func (pf *privateFunctions) get(name string, fileRef *ref.File) *murexFuncDetails {
	pf.mutex.RLock()
	defer pf.mutex.RUnlock()

	if pf.module[fileRef.Source.Module] == nil {
		return nil
	}

	return pf.module[fileRef.Source.Module].get(name)
}

func (pf *privateFunctions) GetString(name string, module string) *murexFuncDetails {
	pf.mutex.RLock()
	defer pf.mutex.RUnlock()

	if pf.module[module] == nil {
		return nil
	}

	return pf.module[module].get(name)
}

func (pf *privateFunctions) Exists(name string, fileRef *ref.File) bool {
	pf.mutex.RLock()
	defer pf.mutex.RUnlock()

	if pf.module[fileRef.Source.Module] == nil {
		return false
	}

	return pf.module[fileRef.Source.Module].Exists(name)
}

func (pf *privateFunctions) ExistsString(name string, module string) bool {
	pf.mutex.RLock()
	defer pf.mutex.RUnlock()

	if pf.module[module] == nil {
		return false
	}

	return pf.module[module].Exists(name)
}

func (pf *privateFunctions) BlockString(name string, module string) ([]rune, error) {
	pf.mutex.RLock()
	defer pf.mutex.RUnlock()

	if pf.module[module] == nil {
		return nil, ErrPrivateNotFound(module)
	}

	return pf.module[module].Block(name)
}

func (pf *privateFunctions) Summary(name string, fileRef *ref.File) (string, error) {
	pf.mutex.RLock()
	defer pf.mutex.RUnlock()

	if pf.module[fileRef.Source.Module] == nil {
		return "", ErrPrivateNotFound(fileRef.Source.Module)
	}

	return pf.module[fileRef.Source.Module].Summary(name)
}

func (pf *privateFunctions) Undefine(name string, fileRef *ref.File) error {
	pf.mutex.Lock()
	defer pf.mutex.Unlock()

	if pf.module[fileRef.Source.Module] == nil {
		return ErrPrivateNotFound(fileRef.Source.Module)
	}

	return pf.module[fileRef.Source.Module].Undefine(name)
}

func (pf *privateFunctions) Dump() any {
	pf.mutex.RLock()
	defer pf.mutex.RUnlock()

	dump := make(map[string]map[string]any)
	for name, module := range pf.module {
		path := strings.SplitN(name, "/", 2)
		if len(path) != 2 {
			return fmt.Sprintf("error: module path doesn't follow standard 'package/module' format: '%s'", name)
		}

		if len(dump[path[0]]) == 0 {
			dump[path[0]] = make(map[string]any)
		}

		dump[path[0]][path[1]] = module.Dump()
	}

	return dump
}
