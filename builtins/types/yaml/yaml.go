package yaml

import (
	"fmt"
	"strconv"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types/define"
	"gopkg.in/yaml.v2"
)

const typeName = "yaml"

func init() {
	stdio.RegesterReadArray(typeName, readArray)
	stdio.RegesterReadMap(typeName, readMap)
	stdio.RegesterWriteArray(typeName, newArrayWriter)
	define.ReadIndexes[typeName] = readIndex
	define.ReadNotIndexes[typeName] = readIndex
	define.Marshallers[typeName] = marshal
	define.Unmarshallers[typeName] = unmarshal

	define.SetMime(typeName,
		"application/yaml", // this is preferred but we will include others since not everyone follows standards.
		"application/x-yaml",
		"text/yaml",
		"text/x-yaml",
	)

	define.SetFileExtensions(typeName, "yaml", "yml")
}

func readArray(read stdio.Io, callback func([]byte)) error {
	return define.ArrayTemplate(yaml.Marshal, yaml.Unmarshal, read, callback)
}

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var jObj interface{}
	err = yaml.Unmarshal(b, &jObj)
	if err == nil {

		switch v := jObj.(type) {
		case []interface{}:
			for i := range jObj.([]interface{}) {
				j, err := yaml.Marshal(jObj.([]interface{})[i])
				if err != nil {
					return err
				}
				callback(strconv.Itoa(i), string(j), i != len(jObj.([]interface{}))-1)
			}

		case map[string]interface{}:
			i := 1
			for key := range jObj.(map[string]interface{}) {
				j, err := yaml.Marshal(jObj.(map[string]interface{})[key])
				if err != nil {
					return err
				}
				callback(key, string(j), i != len(jObj.(map[string]interface{})))
				i++
			}
			return nil

		case map[interface{}]interface{}:
			i := 1
			for key := range jObj.(map[interface{}]interface{}) {
				j, err := yaml.Marshal(jObj.(map[interface{}]interface{})[key])
				if err != nil {
					return err
				}
				callback(fmt.Sprint(key), string(j), i != len(jObj.(map[interface{}]interface{})))
				i++
			}
			return nil

		default:
			if debug.Enabled {
				panic(v)
			}
		}
		return nil
	}
	return err
}

func readIndex(p *lang.Process, params []string) error {
	var jInterface interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	return define.IndexTemplateObject(p, params, &jInterface, yaml.Marshal)
}

func marshal(_ *lang.Process, v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func unmarshal(p *lang.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = yaml.Unmarshal(b, &v)
	return
}
