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

func NewConfiguration() (config Config) {
	config.properties = make(map[string]map[string]Properties)
	config.values = make(map[string]map[string]interface{})
	defaults(&config)
	return
}

// Change a setting in the global configuration.
// app == tooling name
// key == name of setting
// value == the setting itself
func (config *Config) Set(app string, key string, value interface{}) error {
	config.mutex.Lock()
	defer config.mutex.Unlock()

	if config.properties[app] == nil || config.properties[app][key].DataType == "" || config.properties[app][key].Description == "" {
		return errors.New("Cannot Set() that value when no config properties have been defined for that app and key.")
	}

	config.values[app][key] = value
	return nil
}

// Retrieve a setting from the global configuration. Returns an interface{} for the value and err for conversion failures.
// app == tooling name
// key == name of setting
// dataType == what `types.dataType` to cast the return value into
func (config *Config) Get(app, key, dataType string) (value interface{}, err error) {
	config.mutex.Lock()
	defer config.mutex.Unlock()

	if config.properties[app] == nil || config.properties[app][key].DataType == "" || config.properties[app][key].Description == "" {
		return nil, errors.New("Cannot Get() that value when no config properties have been defined for that app and key.")
	}

	var v interface{}
	v = config.values[app][key]
	if v == nil {
		v = config.properties[app][key].Default
	}

	value, err = types.ConvertGoType(v, dataType)
	return
}

func (config *Config) DataType(app, key string) string {
	return config.properties[app][key].DataType
}

func (config *Config) Define(app string, key string, properties Properties) {
	config.mutex.Lock()
	if config.properties[app] == nil {
		config.properties[app] = make(map[string]Properties)
		config.values[app] = make(map[string]interface{})
	}
	config.properties[app][key] = properties
	config.mutex.Unlock()
}

func (config *Config) Dump() (obj map[string]map[string]map[string]interface{}) {
	config.mutex.Lock()
	obj = make(map[string]map[string]map[string]interface{})
	for app := range config.properties {
		obj[app] = make(map[string]map[string]interface{})
		for key := range config.properties[app] {
			obj[app][key] = make(map[string]interface{})
			obj[app][key]["Description"] = config.properties[app][key].Description
			obj[app][key]["Data-Type"] = config.properties[app][key].DataType
			obj[app][key]["Default"] = config.properties[app][key].Default
			obj[app][key]["Value"] = config.values[app][key]
		}
	}
	config.mutex.Unlock()
	return
}
