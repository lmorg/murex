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
	FileRef     *ref.File
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
	global     bool
	properties map[string]map[string]Properties  // This will be the main configuration metadata for each configuration option
	values     map[string]map[string]interface{} // This stores the values when no custom getter and setter have been defined
}

func newGlobal() *Config {
	conf := newConfiguration()
	conf.properties = make(map[string]map[string]Properties)
	conf.values = make(map[string]map[string]interface{})
	conf.global = true
	return conf
}

func newConfiguration() *Config {
	conf := new(Config)
	conf.properties = make(map[string]map[string]Properties)
	conf.values = make(map[string]map[string]interface{})
	return conf
}

// Set changes a setting in the Config object
//
//     app == tooling name
//     key == name of setting
//     value == the setting itself
func (conf *Config) Set(app string, key string, value interface{}) error {
	conf.mutex.Lock()

	if conf.properties[app] == nil || conf.properties[app][key].DataType == "" || conf.properties[app][key].Description == "" {
		conf.mutex.Unlock()
		return fmt.Errorf("Cannot set config. No config has been defined for app `%s`, key `%s`", app, key)
	}

	if !conf.global && conf.properties[app][key].Global {
		conf.mutex.Unlock()
		return InitConf.Set(app, key, value)
	}

	switch {
	case conf.properties[app][key].Dynamic.SetDynamic != nil:
		conf.mutex.Unlock()
		return conf.properties[app][key].Dynamic.SetDynamic(value)

	case conf.properties[app][key].GoFunc.Write != nil:
		conf.mutex.Unlock()
		return conf.properties[app][key].GoFunc.Write(value)

	default:
		conf.values[app][key] = value
		conf.mutex.Unlock()
		return nil
	}
}

// Default resets a config option back to its default
func (conf *Config) Default(app string, key string) error {
	conf.mutex.Lock()

	if conf.properties[app] == nil || conf.properties[app][key].DataType == "" || conf.properties[app][key].Description == "" {
		conf.mutex.Unlock()
		return fmt.Errorf("Cannot default config. No config has been defined for app `%s`, key `%s`", app, key)
	}

	v := conf.properties[app][key].Default
	conf.mutex.Unlock()
	return conf.Set(app, key, v)
}

// Get retrieves a setting from the Config. Returns an interface{} for the value and err for conversion failures.
//
//     app == tooling name
//     key == name of setting
//     dataType == what `types.dataType` to cast the return value into
func (conf *Config) Get(app, key, dataType string) (value interface{}, err error) {
	conf.mutex.RLock()

	//ptr := &conf.properties[app][key]

	if conf.properties[app] == nil || conf.properties[app][key].DataType == "" || conf.properties[app][key].Description == "" {
		conf.mutex.RUnlock()
		return nil, fmt.Errorf("Cannot get config. No config has been defined for app `%s`, key `%s`", app, key)
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

// DataType retrieves the murex data type for a given Config property
func (conf *Config) DataType(app, key string) string {
	conf.mutex.RLock()
	dt := conf.properties[app][key].DataType
	conf.mutex.RUnlock()
	return dt

}

// Define allows new properties to be created in the Config object
func (conf *Config) Define(app string, key string, properties Properties) {
	conf.mutex.Lock()
	if conf.properties[app] == nil {
		conf.properties[app] = make(map[string]Properties)
		conf.values[app] = make(map[string]interface{})
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

// Copy clones the structure
func (conf *Config) Copy() *Config {
	clone := newConfiguration()

	conf.mutex.RLock()

	for app := range conf.properties {

		if clone.properties[app] == nil {
			clone.properties[app] = make(map[string]Properties)
			clone.values[app] = make(map[string]interface{})
		}

		for key := range conf.properties[app] {
			//if conf.properties[app][key].Global {
			//	continue
			//}
			clone.properties[app][key] = conf.properties[app][key]
			clone.values[app][key] = conf.values[app][key]
		}
	}

	conf.mutex.RUnlock()

	return clone
}

// DumpRuntime returns an object based on Config which is optimised for JSON
// serialisation for the `runtime --config` CLI command
func (conf *Config) DumpRuntime() (obj map[string]map[string]map[string]interface{}) {
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
			obj[app][key]["FileRef"] = conf.properties[app][key].FileRef

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

			//if conf.properties[app][key].Global {
			obj[app][key]["Global"] = conf.properties[app][key].Global
			//}

			if len(conf.properties[app][key].Options) != 0 {
				obj[app][key]["Options"] = conf.properties[app][key].Options
			}

			//if len(conf.properties[app][key].Dynamic.Read) != 0 {
			//	obj[app][key]["Dynamic"] = true
			//}
			obj[app][key]["Dynamic"] = len(conf.properties[app][key].Dynamic.Read) != 0

		}
	}
	conf.mutex.RUnlock()
	return
}
