package config

import (
	"fmt"
	"sync"

	"github.com/lmorg/murex/lang/ref"
	"github.com/lmorg/murex/lang/types"
)

// InitConf is a table of global config options
var InitConf = newGlobal()

// Properties is the Config defaults and descriptions
type Properties struct {
	Description string
	Default     interface{}
	DataType    string
	Options     []string
	Global      bool
	Dynamic     DynamicProperties
	GoFunc      GoFuncProperties
	FileRefDef  *ref.File
}

// DynamicProperties is used for dynamic values written in murex
type DynamicProperties struct {
	Read       string
	Write      string
	GetDynamic func() (interface{}, error) `json:"-"`
	SetDynamic func(interface{}) error     `json:"-"`
}

// GoFuncProperties are used for dynamic values written in Go
type GoFuncProperties struct {
	Read  func() (interface{}, error) `json:"-"`
	Write func(interface{}) error     `json:"-"`
}

// Config is used to store all the configuration settings, `config`, in a thread-safe API
type Config struct {
	mutex      sync.RWMutex
	properties map[string]map[string]Properties  // This will be the main configuration metadata for each configuration option
	fileRefSet map[string]map[string]*ref.File   // This is separate from Properties because it gets updated more frequently (eg custom setters)
	values     map[string]map[string]interface{} // This stores the values when no custom getter and setter have been defined
	global     *Config
}

func newGlobal() *Config {
	conf := new(Config)
	conf.properties = make(map[string]map[string]Properties)
	conf.values = make(map[string]map[string]interface{})
	conf.fileRefSet = make(map[string]map[string]*ref.File)
	return conf
}

func newConfiguration(global *Config) *Config {
	conf := new(Config)
	conf.properties = make(map[string]map[string]Properties)
	conf.values = make(map[string]map[string]interface{})
	conf.global = global
	return conf
}

// Copy creates a new *Config instance referenced to the parent
func (conf *Config) Copy() *Config {
	if conf.global == nil {
		return newConfiguration(conf)
	}

	return newConfiguration(conf.global)
}

// ExistsAndGlobal checks if a config option exists and/or is global
func (conf *Config) ExistsAndGlobal(app, key string) (exists, global bool) {
	conf.mutex.RLock()
	exists = conf.properties[app] != nil && conf.properties[app][key].DataType != "" && conf.properties[app][key].Description != ""
	global = exists && conf.properties[app][key].Global
	conf.mutex.RUnlock()
	return
}

// Get retrieves a setting from the Config. Returns an interface{} for the value and err for any failures.
//
//     app == tooling name
//     key == name of setting
//     dataType == what `types.dataType` to cast the return value into
func (conf *Config) Get(app, key, dataType string) (value interface{}, err error) {
	conf.mutex.RLock()

	if conf.global != nil && conf.values[app] != nil && conf.values[app][key] != nil {
		v := conf.values[app][key]
		conf.mutex.RUnlock()
		value, err = types.ConvertGoType(v, dataType)
		return
	}

	if conf.properties[app] == nil || conf.properties[app][key].DataType == "" || conf.properties[app][key].Description == "" {
		conf.mutex.RUnlock()

		if conf.global != nil {
			return conf.global.Get(app, key, dataType)
		}
		return nil, fmt.Errorf("cannot get config. No config has been defined for app `%s`, key `%s`", app, key)
	}

	var v interface{}

	switch {
	case conf.properties[app][key].Dynamic.GetDynamic != nil:
		v, err = conf.properties[app][key].Dynamic.GetDynamic()
		if err != nil {
			conf.mutex.RUnlock()
			return
		}
		conf.mutex.RUnlock()

	case conf.properties[app][key].GoFunc.Read != nil:
		v, err = conf.properties[app][key].GoFunc.Read()
		if err != nil {
			conf.mutex.RUnlock()
			return
		}
		conf.mutex.RUnlock()

	default:
		v = conf.values[app][key]
		if v == nil {
			v = conf.properties[app][key].Default
		}
		conf.mutex.RUnlock()
	}

	value, err = types.ConvertGoType(v, dataType)
	return
}

// Set changes a setting in the Config object
//
//     app == tooling name
//     key == name of setting
//     value == the setting itself
func (conf *Config) Set(app string, key string, value interface{}, fileRef *ref.File) error {
	// first check if we're in a global, and whether we should be
	if conf.global != nil {
		exists, global := conf.global.ExistsAndGlobal(app, key)
		if !exists || global {
			return conf.global.Set(app, key, value, fileRef)
		}
	}

	conf.mutex.Lock()

	if conf.global == nil {
		if conf.properties[app] == nil || conf.properties[app][key].DataType == "" || conf.properties[app][key].Description == "" {
			conf.mutex.Unlock()
			return fmt.Errorf("cannot set config. No config has been defined for app `%s`, key `%s`", app, key)
		}
	}

	switch {
	case conf.properties[app][key].Dynamic.SetDynamic != nil:
		conf.mutex.Unlock()
		err := conf.properties[app][key].Dynamic.SetDynamic(value)
		if err == nil {
			conf.mutex.Lock()
			conf.fileRefSet[app][key] = fileRef
			conf.mutex.Unlock()
		}
		return err

	case conf.properties[app][key].GoFunc.Write != nil:
		conf.mutex.Unlock()
		err := conf.properties[app][key].GoFunc.Write(value)
		if err == nil {
			conf.mutex.Lock()
			conf.fileRefSet[app][key] = fileRef
			conf.mutex.Unlock()
		}
		return err

	default:
		if len(conf.values) == 0 {
			conf.values = make(map[string]map[string]interface{})
		}
		if len(conf.values[app]) == 0 {
			conf.values[app] = make(map[string]interface{})
		}

		conf.values[app][key] = value
		conf.fileRefSet[app][key] = fileRef

		conf.mutex.Unlock()
		return nil
	}
}

// Default resets a config option back to its default
func (conf *Config) Default(app string, key string, fileRef *ref.File) error {
	c := conf.global
	if c == nil {
		c = conf
	}

	exists, _ := c.ExistsAndGlobal(app, key)

	if !exists {
		return fmt.Errorf("cannot default config. No config has been defined for app `%s`, key `%s`", app, key)
	}

	c.mutex.RLock()
	v := c.properties[app][key].Default
	c.mutex.RUnlock()

	return conf.Set(app, key, v, fileRef)
}

// DataType retrieves the murex data type for a given Config property
func (conf *Config) DataType(app, key string) string {
	if conf.global != nil {
		return conf.global.DataType(app, key)
	}

	conf.mutex.Lock()
	dt := conf.properties[app][key].DataType
	conf.mutex.Unlock()
	return dt
}

// Define allows new properties to be created in the Config object
func (conf *Config) Define(app string, key string, properties Properties) {
	if conf.global != nil {
		conf.global.Define(app, key, properties)
		return
	}

	conf.mutex.Lock()
	if conf.properties[app] == nil {
		conf.properties[app] = make(map[string]Properties)
		conf.values[app] = make(map[string]interface{})
		conf.fileRefSet[app] = make(map[string]*ref.File)
	}

	// don't set the value to the default if it's a dynamic property
	if properties.Dynamic.Read == "" && properties.GoFunc.Read == nil {
		conf.values[app][key] = properties.Default
	} else {
		properties.Global = true
	}
	conf.properties[app][key] = properties

	conf.mutex.Unlock()
}

// DumpRuntime returns an object based on Config which is optimised for JSON
// serialisation for the `runtime --config` CLI command
func (conf *Config) DumpRuntime() (obj map[string]map[string]map[string]interface{}) {
	if conf.global != nil {
		return conf.global.DumpRuntime()
	}

	conf.mutex.RLock()
	obj = make(map[string]map[string]map[string]interface{})
	for app := range conf.properties {
		obj[app] = make(map[string]map[string]interface{})
		for key := range conf.properties[app] {
			obj[app][key] = make(map[string]interface{})
			obj[app][key]["Description"] = conf.properties[app][key].Description
			obj[app][key]["Data-Type"] = conf.properties[app][key].DataType
			obj[app][key]["Default"] = conf.properties[app][key].Default
			obj[app][key]["Value"] = conf.values[app][key]
			obj[app][key]["FileRefDefined"] = conf.properties[app][key].FileRefDef
			obj[app][key]["FileRefSet"] = conf.fileRefSet[app][key]

			//if conf.properties[app][key].Global {
			obj[app][key]["Global"] = conf.properties[app][key].Global
			//}

			//if len(conf.properties[app][key].Options) != 0 {
			obj[app][key]["Options"] = conf.properties[app][key].Options
			//}

			//if len(conf.properties[app][key].Dynamic.Read) != 0 {
			obj[app][key]["Dynamic"] = conf.properties[app][key].Dynamic
			//}

			//if conf.properties[app][key].GoFunc.Read != nil {
			obj[app][key]["GoFunc"] = map[string]bool{
				"Read":  conf.properties[app][key].GoFunc.Read != nil,
				"Write": conf.properties[app][key].GoFunc.Write != nil,
			}
			//}

		}
	}
	conf.mutex.RUnlock()
	return
}

// DumpConfig returns an object based on Config which is optimised for JSON
// serialisation for the `config` CLI command
func (conf *Config) DumpConfig() (obj map[string]map[string]map[string]interface{}) {
	if conf.global != nil {
		return conf.global.DumpConfig()
	}

	conf.mutex.RLock()
	obj = make(map[string]map[string]map[string]interface{})
	for app := range conf.properties {
		obj[app] = make(map[string]map[string]interface{})
		for key := range conf.properties[app] {
			obj[app][key] = make(map[string]interface{})
			obj[app][key]["Description"] = conf.properties[app][key].Description
			obj[app][key]["Data-Type"] = conf.properties[app][key].DataType
			obj[app][key]["Default"] = conf.properties[app][key].Default

			switch {
			case conf.properties[app][key].GoFunc.Read != nil:
				v, err := conf.properties[app][key].GoFunc.Read()
				if err == nil {
					obj[app][key]["Value"] = v
				}
			case len(conf.properties[app][key].Dynamic.Read) == 0:
				obj[app][key]["Value"] = conf.values[app][key]
			}

			obj[app][key]["Global"] = conf.properties[app][key].Global

			if len(conf.properties[app][key].Options) != 0 {
				obj[app][key]["Options"] = conf.properties[app][key].Options
			}

			obj[app][key]["Dynamic"] = len(conf.properties[app][key].Dynamic.Read) != 0

		}
	}
	conf.mutex.RUnlock()
	return
}
