package types

import (
	"fmt"
	"os"
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

type jsonableVarItem struct {
	Value interface{}
	Type  string
}

type jsonableVars map[string]jsonableVarItem

// Create a new scope of variables
func NewVariableGroup() (v Vars) {
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

// This exists so we can dump variables natively into `eval` and `let`.
// It includes OS environmental variables as well as murex local variables.
// In the case where they share the same name, the local variables will override the OS env vars.
func (v *Vars) DumpMap() (m map[string]interface{}) {
	m = make(map[string]interface{})

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		m[pair[0]] = pair[1]
	}

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
	if value == nil {
		value = os.Getenv(name)
	}
	return
}

// Get variable - cast as string.
func (v *Vars) GetString(name string) (s string) {
	v.mutex.Lock()

	defer func() {
		r := recover()
		if r != nil {
			s = fmt.Sprint("Unexpected value of:", v.values[name])
		}
		v.mutex.Unlock()
	}()

	switch v.types[name] {
	case "":
		return os.Getenv(name)

	case Integer:
		return strconv.Itoa(v.values[name].(int))

	case Float, Number:
		return strconv.FormatFloat(v.values[name].(float64), 'f', -1, 64)

	default:
		return v.values[name].(string)
	}
	//return ""
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
		s, err := ConvertGoType(value, String)
		if err != nil {
			return err
		}
		v.values[name] = strings.TrimSpace(s.(string))
	}

	return nil
}

// Unset a variable
func (v *Vars) Unset(name string) {
	v.mutex.Lock()
	delete(v.values, name)
	delete(v.types, name)
	v.mutex.Unlock()
}
