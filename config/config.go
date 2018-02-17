package config

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/lmorg/murex/lang/types"
)

// Properties is the Config defaults and descriptions
type Properties struct {
	Description string
	Default     interface{}
	DataType    string
	//GetValue    func() interface{}      // Getter to override murex default
	//SetValue    func(interface{}) error // Setter to override murex default
	//IsGlobal bool // If set it means configuration settings are global and thus not thread safe
}

// Config is used to store all the configuration settings, `config`, in a thread-safe API
type Config struct {
	mutex      sync.Mutex
	properties map[string]map[string]Properties  // This will be the main configuration metadata for each configuration option
	values     map[string]map[string]interface{} // This stores the values when no custom getter and setter have been defined
}

// NewConfiguration creates an new Config object (see above)
func NewConfiguration() (conf *Config) {
	conf = new(Config)
	conf.properties = make(map[string]map[string]Properties)
	conf.values = make(map[string]map[string]interface{})
	return
}

// Set changes a setting in the Config object
//
//     app == tooling name
//     key == name of setting
//     value == the setting itself
func (conf *Config) Set(app string, key string, value interface{}) error {
	conf.mutex.Lock()
	defer conf.mutex.Unlock()

	if conf.properties[app] == nil || conf.properties[app][key].DataType == "" || conf.properties[app][key].Description == "" {
		return errors.New("Cannot Get() `" + app + "`:`" + key + "` when no config properties have been defined for that app and key.")
	}

	switch conf.values[app][key].(type) {
	case []string:
		//if config.properties[app][key].DataType == types.Json {
		var iface interface{}
		err := json.Unmarshal([]byte(value.(string)), &iface)
		if err != nil {
			return errors.New("Unable to set config with that data: " + err.Error())
		}

		for i := range iface.([]string) {
			conf.values[app][key].([]string)[i] = iface.([]string)[i]
		}
	//}

	case map[string]string:
		//if config.properties[app][key].DataType == types.Json {
		var iface interface{}
		err := json.Unmarshal([]byte(value.(string)), &iface)
		if err != nil {
			return errors.New("Unable to set config with that data: " + err.Error())
		}

		for k := range conf.values[app][key].(map[string]string) {
			delete(conf.values[app][key].(map[string]string), k)
		}

		for k := range iface.(map[string]string) {
			conf.values[app][key].(map[string]string)[k] = iface.(map[string]string)[k]
		}

	default:
		conf.values[app][key] = value
	}

	return nil
}

// Get retrieves a setting from the Config. Returns an interface{} for the value and err for conversion failures.
//
//     app == tooling name
//     key == name of setting
//     dataType == what `types.dataType` to cast the return value into
func (conf Config) Get(app, key, dataType string) (value interface{}, err error) {
	conf.mutex.Lock()
	defer conf.mutex.Unlock()

	if conf.properties[app] == nil || conf.properties[app][key].DataType == "" || conf.properties[app][key].Description == "" {
		return nil, errors.New("Cannot Get() `" + app + "`:`" + key + "` when no config properties have been defined for that app and key.")
	}

	var v interface{}
	v = conf.values[app][key]
	if v == nil {
		v = conf.properties[app][key].Default
	}

	value, err = types.ConvertGoType(v, dataType)
	return
}

// DataType retrieves the murex data type for a given Config property
func (conf *Config) DataType(app, key string) string {
	return conf.properties[app][key].DataType
}

// Define allows new properties to be created in the Config object
func (config *Config) Define(app string, key string, properties Properties) {
	config.mutex.Lock()
	if config.properties[app] == nil {
		config.properties[app] = make(map[string]Properties)
		config.values[app] = make(map[string]interface{})
	}
	config.properties[app][key] = properties
	config.values[app][key] = properties.Default
	config.mutex.Unlock()
}

// Copy clones the structure
func (conf *Config) Copy() *Config {
	clone := NewConfiguration()

	for app := range conf.properties {

		if clone.properties[app] == nil {
			clone.properties[app] = make(map[string]Properties)
			clone.values[app] = make(map[string]interface{})
		}

		for key := range conf.properties[app] {
			clone.properties[app][key] = conf.properties[app][key]
			clone.values[app][key] = conf.values[app][key]
		}
	}

	return clone
}

// Dump returns an object based on Config which is optimised for JSON serialisation
func (conf *Config) Dump() (obj map[string]map[string]map[string]interface{}) {
	conf.mutex.Lock()
	obj = make(map[string]map[string]map[string]interface{})
	for app := range conf.properties {
		obj[app] = make(map[string]map[string]interface{})
		for key := range conf.properties[app] {
			obj[app][key] = make(map[string]interface{})
			obj[app][key]["Description"] = conf.properties[app][key].Description
			obj[app][key]["Data-Type"] = conf.properties[app][key].DataType
			obj[app][key]["Default"] = conf.properties[app][key].Default
			obj[app][key]["Value"] = conf.values[app][key]
		}
	}
	conf.mutex.Unlock()
	return
}
