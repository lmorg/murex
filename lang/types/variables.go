package types

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Vars struct {
	mutex   sync.Mutex
	strings map[string]string
	ints    map[string]int
	floats  map[string]float64
	types   map[string]string
}

func NewVariableGroup() (v Vars) {
	v.strings = make(map[string]string)
	v.ints = make(map[string]int)
	v.floats = make(map[string]float64)
	v.types = make(map[string]string)
	return
}

// Dump the entire variable structure into a JSON-able interface.
func (v *Vars) Dump() (obj map[string]interface{}) {
	v.mutex.Lock()
	obj = make(map[string]interface{}, 0)
	obj["Type"] = v.types
	obj["String"] = v.strings
	obj["Integer"] = v.ints
	obj["Float"] = v.floats
	v.mutex.Unlock()
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
	switch v.types[name] {
	case Integer:
		value = v.ints[name]

	case Float:
		value = v.floats[name]

	case Boolean:
		if IsTrue([]byte(v.strings[name]), 0) {
			value = true
		} else {
			value = false
		}

	default:
		value = v.strings[name]
	}

	v.mutex.Unlock()
	return
}

// Get variable - cast as string.
func (v *Vars) GetString(name string) (value string) {
	v.mutex.Lock()
	switch v.types[name] {
	case Integer:
		value = strconv.Itoa(v.ints[name])

	case Float:
		value = fmt.Sprint(v.floats[name]) //TODO: this lazy fix needs to be replaced by strconv.floatthingy

	default:
		value = v.strings[name]
	}
	v.mutex.Unlock()
	return
}

// Set a variable.
func (v *Vars) Set(name string, value interface{}, dataType string) {
	v.mutex.Lock()
	v.types[name] = dataType
	switch dataType {
	case Integer:
		v.ints[name] = value.(int)

	case Float:
		v.floats[name] = value.(float64)

	/*case types.Boolean:
	if types.IsTrue([]byte(v.strings[name]), 0) {
		value = true
	} else {
		value = false
	}*/

	default:
		v.strings[name] = strings.TrimSpace(value.(string))
	}

	v.mutex.Unlock()
	return
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

		case (*s)[i] == '_', (*s)[i] == '-',
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
