package lang

import (
	"fmt"
	"strings"

	"github.com/lmorg/murex/lang/ref"
)

type privateFunctions struct {
	module map[string]*MurexFuncs
}

func NewMurexPrivs() *privateFunctions {
	pf := new(privateFunctions)
	pf.module = make(map[string]*MurexFuncs)
	return pf
}

func (pf *privateFunctions) Define(name string, parameters []MxFunctionParams, block []rune, fileRef *ref.File) {
	if pf.module[fileRef.Source.Module] == nil {
		pf.module[fileRef.Source.Module] = NewMurexFuncs()
	}

	pf.module[fileRef.Source.Module].Define(name, parameters, block, fileRef)
}

func (pf *privateFunctions) get(name string, fileRef *ref.File) *murexFuncDetails {
	if pf.module[fileRef.Source.Module] == nil {
		return nil
	}

	return pf.module[fileRef.Source.Module].get(name)
}

func (pf *privateFunctions) Exists(name string, fileRef *ref.File) bool {
	if pf.module[fileRef.Source.Module] == nil {
		return false
	}

	return pf.module[fileRef.Source.Module].Exists(name)
}

func (pf *privateFunctions) ExistsString(name string, module string) bool {
	if pf.module[module] == nil {
		return false
	}

	return pf.module[module].Exists(name)
}

func (pf *privateFunctions) BlockString(name string, module string) ([]rune, error) {
	if pf.module[module] == nil {
		return nil, fmt.Errorf("no private functions exist for module `%s`", module)
	}

	return pf.module[module].Block(name)
}

func (pf *privateFunctions) Summary(name string, fileRef *ref.File) (string, error) {
	if pf.module[fileRef.Source.Module] == nil {
		return "", fmt.Errorf("no private functions exist for module `%s`", fileRef.Source.Module)
	}

	return pf.module[fileRef.Source.Module].Summary(name)
}

func (pf *privateFunctions) Undefine(name string, fileRef *ref.File) error {
	if pf.module[fileRef.Source.Module] == nil {
		return fmt.Errorf("no private functions exist for module `%s`", fileRef.Source.Module)
	}

	return pf.module[fileRef.Source.Module].Undefine(name)
}

func (pf *privateFunctions) Dump() interface{} {
	dump := make(map[string]map[string]interface{})
	for name, module := range pf.module {
		path := strings.SplitN(name, "/", 2)
		if len(path) != 2 {
			return fmt.Sprintf("error: module path doesn't follow standard 'package/module' format: '%s'", name)
		}

		if len(dump[path[0]]) == 0 {
			dump[path[0]] = make(map[string]interface{})
		}

		dump[path[0]][path[1]] = module.Dump()
	}

	return dump
}
