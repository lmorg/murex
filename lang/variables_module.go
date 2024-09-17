package lang

import (
	"fmt"
	"sync"
)

type ModuleVars struct {
	mutex sync.Mutex
	vars  map[string]*Variables
}

func NewModuleVars() *ModuleVars {
	mod := new(ModuleVars)
	mod.vars = make(map[string]*Variables)
	return mod
}

func (mod *ModuleVars) GetValues(p *Process) interface{} {
	m := make(map[string]interface{})
	v := mod.v(p)

	v.mutex.Lock()
	for name, variable := range v.vars {
		m[name] = variable.Value
	}
	v.mutex.Unlock()

	return m
}

func (mod *ModuleVars) GetDataType(p *Process, name string) (dt string) {
	v := mod.v(p)

	v.mutex.Lock()
	variable := v.vars[name]
	if variable != nil {
		dt = variable.DataType
	}
	v.mutex.Unlock()
	return
}

func (mod *ModuleVars) Set(p *Process, value interface{}, changePath []string, dataType string) (err error) {
	if len(changePath) == 0 {
		return fmt.Errorf("invalid use of $%s. Expecting a module variable name, eg `$%s.example`", _VAR_MODULE, _VAR_MODULE)
	}

	switch t := value.(type) {
	case map[string]interface{}:
		return mod.set(p, changePath[0], t[changePath[0]], dataType)

	default:
		return fmt.Errorf("expecting a map of module variables. Instead got a %T", t)
	}
}

func (mod *ModuleVars) set(p *Process, path string, value interface{}, dataType string) error {
	return mod.v(p).Set(p, path, value, dataType)
}

func (mod *ModuleVars) v(p *Process) *Variables {
	mod.mutex.Lock()
	v, ok := mod.vars[p.FileRef.Source.Module]
	if !ok {
		v = NewVariables(p)
		mod.vars[p.FileRef.Source.Module] = v
	}
	mod.mutex.Unlock()

	return v
}
