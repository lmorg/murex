package types

import (
	"strconv"
	"strings"
	"sync"
)

type Vars struct {
	mutex sync.Mutex
	//mutex  debug.Mutex
	values map[string]interface{}
	types  map[string]string
}

type JsonableVarItem struct {
	Value interface{}
	Type  string
}

type JsonableVars map[string]JsonableVarItem

func NewVariableGroup() (v Vars) {
	v.values = make(map[string]interface{})
	v.types = make(map[string]string)
	return
}

// Dump the entire variable structure into a JSON-able interface.
func (v *Vars) Dump() (obj JsonableVars) {
	v.mutex.Lock()
	obj = make(map[string]JsonableVarItem, 0)
	for name := range v.values {
		obj[name] = JsonableVarItem{
			Value: v.values[name],
			Type:  v.types[name],
		}
	}
	v.mutex.Unlock()
	return
}

/*
func (v *Vars) Dump() (obj map[string]interface{}) {
	v.mutex.Lock()
	obj = make(map[string]interface{}, 0)
	obj["Type"] = v.types
	obj["Values"] = v.values
	v.mutex.Unlock()
	return
}
*/

// This exists so we can dump variables natively into `eval` and `let`.
func (v *Vars) DumpMap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	for k, v := range v.values {
		m[k] = v
	}
	return
}

// Get the variable type.
func (v *Vars) GetType(name string) (t string) {
	v.mutex.Lock()
	t = v.types[name]
	v.mutex.Unlock()

	if t == "" {
		return Null
	}
	return
}

// Get variable in native type.
func (v *Vars) GetValue(name string) (value interface{}) {
	v.mutex.Lock()
	value = v.values[name]
	v.mutex.Unlock()
	return
}

// Get variable - cast as string.
func (v *Vars) GetString(name string) string {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	switch v.types[name] {
	case "":
		return ""

	case Integer:
		return strconv.Itoa(v.values[name].(int))

	case Float, Number:
		return strconv.FormatFloat(v.values[name].(float64), 'f', -1, 64)

	default:
		return v.values[name].(string)
	}
	return ""
}

// Set a variable.
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

	/*case types.Boolean:
	if types.IsTrue([]byte(v.strings[name]), 0) {
		value = true
	} else {
		value = false
	}*/

	default:
		s, err := ConvertGoType(value, dataType)
		if err != nil {
			return err
		}
		v.values[name] = strings.TrimSpace(s.(string))
	}

	return nil
}

func (v *Vars) Unset(name string) {
	v.mutex.Lock()
	delete(v.values, name)
	delete(v.types, name)
	v.mutex.Unlock()
}
