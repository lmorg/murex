package config

import (
	"errors"
	"github.com/lmorg/murex/lang/types"
	"sync"
)

type Properties struct {
	Description string
	Default     interface{}
	DataType    string
}

type Config struct {
	mutex      sync.Mutex
	properties map[string]map[string]Properties
	values     map[string]map[string]interface{}
}

func NewConfiguration() (gc Config) {
	gc.properties = make(map[string]map[string]Properties)
	gc.values = make(map[string]map[string]interface{})
	return
}

// Change a setting in the global configuration.
// app == tooling name
// key == name of setting
// value == the setting itself
func (gc *Config) Set(app string, key string, value interface{}) error {
	gc.mutex.Lock()
	if gc.properties[app] == nil || gc.properties[app][key].DataType == "" || gc.properties[app][key].Description == "" {
		return errors.New("Cannot Set() that value when no config properties have been defined for that app and key.")
	}

	gc.values[app][key] = value
	gc.mutex.Unlock()
	return nil
}

// Retrieve a setting from the global configuration. Returns an interface{} for the value and err for conversion failures.
// app == tooling name
// key == name of setting
// dataType == what `types.dataType` to cast the return value into
func (gc *Config) Get(app, key, dataType string) (value interface{}, err error) {
	gc.mutex.Lock()
	if gc.properties[app] == nil || gc.properties[app][key].DataType == "" || gc.properties[app][key].Description == "" {
		return nil, errors.New("Cannot Get() that value when no config properties have been defined for that app and key.")
	}

	var v interface{}
	v = gc.values[app][key]
	if v == nil {
		v = gc.properties[app][key].Default
	}

	value, err = types.ConvertGoType(v, dataType)
	gc.mutex.Unlock()
	return
}

func (gc *Config) Define(app string, key string, properties Properties) {
	gc.mutex.Lock()
	if gc.properties[app] == nil {
		gc.properties[app] = make(map[string]Properties)
		gc.values[app] = make(map[string]interface{})
	}
	gc.properties[app][key] = properties
	gc.mutex.Unlock()
}

func (gc *Config) Dump() (obj map[string]map[string]map[string]interface{}) {
	gc.mutex.Lock()
	obj = make(map[string]map[string]map[string]interface{})
	for app := range gc.properties {
		obj[app] = make(map[string]map[string]interface{})
		for key := range gc.properties[app] {
			obj[app][key] = make(map[string]interface{})
			obj[app][key]["Description"] = gc.properties[app][key].Description
			obj[app][key]["Data-Type"] = gc.properties[app][key].DataType
			obj[app][key]["Default"] = gc.properties[app][key].Default
			obj[app][key]["Value"] = gc.values[app][key]
		}
	}
	gc.mutex.Unlock()
	return
}
