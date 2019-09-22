package toml

import (
	"bytes"
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/lmorg/murex/builtins/pipes/streams"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
)

const typeName = "toml"

var errNakedArrays = errors.New("The TOML specification doesn't support naked arrays")

func init() {
	stdio.RegesterReadArray(typeName, readArray)
	//stdio.RegesterReadMap(typeName, readMap)
	stdio.RegesterWriteArray(typeName, func(_ stdio.Io) (stdio.ArrayWriter, error) {
		return nil, errNakedArrays
	})

	lang.ReadIndexes[typeName] = readIndex
	lang.ReadNotIndexes[typeName] = readIndex
	lang.Marshallers[typeName] = marshal
	lang.Unmarshallers[typeName] = unmarshal

	lang.SetMime(typeName,
		"application/toml", // this is preferred but we will include others since not everyone follows standards.
		"application/x-toml",
		"text/toml",
		"text/x-toml",
	)

	lang.SetFileExtensions(typeName, "toml")
}

func tomlMarshal(v interface{}) (b []byte, err error) {
	w := streams.NewStdin()
	enc := toml.NewEncoder(w)
	err = enc.Encode(v)
	if err != nil {
		return nil, err
	}

	b, err = w.ReadAll()
	return b, err
}

func readArray(read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	j := make([]interface{}, 0)
	err = toml.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	for i := range j {
		switch j[i].(type) {
		case string:
			callback(bytes.TrimSpace([]byte(j[i].(string))))

		default:
			jBytes, err := tomlMarshal(j[i])
			if err != nil {
				return err
			}
			callback(jBytes)
		}
	}

	return nil
}

/*func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var jObj interface{}
	err = toml.Unmarshal(b, &jObj)
	if err != nil {
		return err
	}

	switch v := jObj.(type) {
	case []interface{}:
		for i := range jObj.([]interface{}) {
			j, err := tomlMarshal(jObj.([]interface{})[i])
			if err != nil {
				return err
			}
			callback(strconv.Itoa(i), string(j), i != len(jObj.([]interface{}))-1)
		}

	case map[string]interface{}, map[interface{}]interface{}:
		i := 1
		for key := range jObj.(map[string]interface{}) {
			j, err := tomlMarshal(jObj.(map[string]interface{})[key])
			if err != nil {
				return err
			}
			callback(key, string(j), i != len(jObj.(map[string]interface{})))
			i++
		}
		return nil

	default:
		if debug.Enabled {
			panic(v)
		}
	}
	return nil
}*/

func readIndex(p *lang.Process, params []string) error {
	var jInterface interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = toml.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	return lang.IndexTemplateObject(p, params, &jInterface, tomlMarshal)
}

func marshal(_ *lang.Process, v interface{}) ([]byte, error) {
	return tomlMarshal(v)
}

func unmarshal(p *lang.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = toml.Unmarshal(b, &v)
	return
}
