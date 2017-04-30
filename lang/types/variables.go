package types

import (
	"strconv"
	"strings"
	"sync"
)

type Vars struct {
	mutex  sync.Mutex
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
func (v *Vars) GetString(name string) (value string) {
	v.mutex.Lock()
	switch v.types[name] {
	case "":
		return ""

	case Integer:
		value = strconv.Itoa(v.values[name].(int))

	case Float, Number:
		value = strconv.FormatFloat(v.values[name].(float64), 'f', -1, 64)

	default:
		value = v.values[name].(string)
	}
	v.mutex.Unlock()
	return
}

// Set a variable.
func (v *Vars) Set(name string, value interface{}, dataType string) error {
	v.mutex.Lock()
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

	v.mutex.Unlock()
	return nil
}

func (v *Vars) Unset(name string) {
	v.mutex.Lock()
	delete(v.values, name)
	delete(v.types, name)
	v.mutex.Unlock()
}

// Replaces variable key names with values inside a string.
// Code templated here: https://play.golang.org/p/ho8RTxxe-0
func (v *Vars) KeyValueReplace(s *string) {
	replace := func(start, end int) (diff int) {
		value := v.GetString((*s)[start+1 : end])
		diff = len(value) - len((*s)[start:end])
		*s = (*s)[:start] + value + (*s)[end:]
		return
	}

	if len(*s) == 0 {
		return
	}

	*s = " " + *s + " "
	start := 0
	for i := 1; i < len(*s); i++ {

		switch {
		//case (*s)[i] == '$' && (*s)[i-1] == '\\':
		//	*s = (*s)[:i-1] + (*s)[i:]

		case (*s)[i] == '$' && (*s)[i-1] != '\\':
			if start == 0 {
				start = i
			} else {
				i += replace(start, i)
				start = 0
			}

		case (*s)[i] == '_',
			(*s)[i] <= 'z' && 'a' <= (*s)[i],
			(*s)[i] <= 'Z' && 'A' <= (*s)[i],
			(*s)[i] <= '9' && '0' <= (*s)[i]:

			continue

		default:
			if start != 0 {
				i += replace(start, i)
				start = 0
			}

		}
	}

	*s = (*s)[1 : len(*s)-1]
}
