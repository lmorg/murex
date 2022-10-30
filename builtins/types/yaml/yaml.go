package yaml

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"gopkg.in/yaml.v3"
)

const typeName = "yaml"

func init() {
	stdio.RegisterReadArray(typeName, readArray)
	stdio.RegisterReadArrayWithType(typeName, readArrayWithType)
	stdio.RegisterReadMap(typeName, readMap)
	stdio.RegisterWriteArray(typeName, newArrayWriter)
	lang.ReadIndexes[typeName] = readIndex
	lang.ReadNotIndexes[typeName] = readIndex
	lang.Marshallers[typeName] = marshal
	lang.Unmarshallers[typeName] = unmarshal

	lang.SetMime(typeName,
		"application/yaml", // this is preferred but we will include others since not everyone follows standards.
		"application/x-yaml",
		"text/yaml",
		"text/x-yaml",
	)

	lang.SetFileExtensions(typeName, "yaml", "yml")
}

func readArray(ctx context.Context, read stdio.Io, callback func([]byte)) error {
	return lang.ArrayTemplate(ctx, yaml.Marshal, yaml.Unmarshal, read, callback)
}

func readArrayWithType(ctx context.Context, read stdio.Io, callback func([]byte, string)) error {
	return lang.ArrayWithTypeTemplate(ctx, typeName, yaml.Marshal, yaml.Unmarshal, read, callback)
}

func noCrLf(b []byte) []byte {
	if len(b) > 0 && b[len(b)-1] == '\n' {
		b = b[:len(b)-1]
	}

	if len(b) > 0 && b[len(b)-1] == '\r' {
		b = b[:len(b)-1]
	}

	return b
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
				callback(strconv.Itoa(i), string(noCrLf(j)), i != len(jObj.([]interface{}))-1)
			}

		case map[string]interface{}:
			i := 1
			for key := range jObj.(map[string]interface{}) {
				j, err := yaml.Marshal(jObj.(map[string]interface{})[key])
				if err != nil {
					return err
				}
				callback(key, string(noCrLf(j)), i != len(jObj.(map[string]interface{})))
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
				callback(fmt.Sprint(key), string(noCrLf(j)), i != len(jObj.(map[interface{}]interface{})))
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

	return lang.IndexTemplateObject(p, params, &jInterface, yaml.Marshal)
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
