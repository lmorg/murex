package types

import (
	"strings"
	"sync"
)

// Vars is used to store a table of variables
type Vars struct {
	mutex sync.Mutex
	//mutex  debug.Mutex
	values map[string]interface{}
	types  map[string]string
}

type jsonableVarItem struct {
	Value interface{}
	Type  string
}

type jsonableVars map[string]jsonableVarItem

// NewVariableGroup creates a new scope of variables
func NewVariableGroup() (v *Vars) {
	v = new(Vars)
	v.values = make(map[string]interface{})
	v.types = make(map[string]string)
	return
}

// Dump the entire variable structure into a JSON-able interface.
func (v *Vars) Dump() (obj jsonableVars) {
	v.mutex.Lock()
	obj = make(map[string]jsonableVarItem, 0)
	for name := range v.values {
		obj[name] = jsonableVarItem{
			Value: v.values[name],
			Type:  v.types[name],
		}
	}
	v.mutex.Unlock()
	return
}

// DumpMap exists so we can dump variables natively into `eval` and `let`.
func (v *Vars) DumpMap(m map[string]interface{}) {
	m = make(map[string]interface{})

	for k, v := range v.values {
		m[k] = v
	}

	return
}

// GetType gets the murex variable's murex data type.
func (v *Vars) GetType(name string) (t string) {
	v.mutex.Lock()
	t = v.types[name]
	v.mutex.Unlock()

	if t == "" {
		return Null
	}
	return
}

// GetValue gets a murex variable in a Go native type.
func (v *Vars) GetValue(name string) (value interface{}) {
	v.mutex.Lock()
	value = v.values[name]
	v.mutex.Unlock()

	return
}

// Set murex a variable.
func (v *Vars) Set(name string, value interface{}, dataType string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	v.types[name] = dataType

	switch dataType {
	case Integer:
		i, err := ConvertGoType(value, dataType)
		if err != nil {
			return err
		}
		v.values[name] = i.(int)

	case Float, Number:
		f, err := ConvertGoType(value, dataType)
		if err != nil {
			return err
		}
		v.values[name] = f.(float64)

	case Boolean:
		b, err := ConvertGoType(value, Boolean)
		if err != nil {
			return err
		}
		v.values[name] = b.(bool)
	/*if IsTrue([]byte(s.(string)), 0) {
		value = true
	} else {
		value = false
	}*/

	default:
		s, err := ConvertGoType(value, String)
		if err != nil {
			return err
		}
		v.values[name] = strings.TrimSpace(s.(string))
	}

	return nil
}

// Unset a murex variable
func (v *Vars) Unset(name string) {
	v.mutex.Lock()
	delete(v.values, name)
	delete(v.types, name)
	v.mutex.Unlock()
}

// Copy clones the structure
func (v *Vars) Copy() *Vars {
	clone := NewVariableGroup()

	v.mutex.Lock()

	for name := range v.values {
		clone.values[name] = v.values[name]
		clone.types[name] = v.types[name]
	}

	v.mutex.Unlock()

	return clone
}
