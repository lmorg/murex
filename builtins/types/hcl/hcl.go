package hcl

import (
	"bytes"
	"strconv"

	"github.com/hashicorp/hcl"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils/json"
)

const typeName = "hcl"

func init() {
	define.ReadIndexes[typeName] = readIndex
	define.ReadNotIndexes[typeName] = readIndex

	stdio.RegesterReadArray(typeName, readArray)
	stdio.RegesterReadMap(typeName, readMap)
	stdio.RegesterWriteArray(typeName, newArrayWriter)

	define.Marshallers[typeName] = marshal
	define.Unmarshallers[typeName] = unmarshal

	// These are just guessed at as I couldn't find any formally named MIMEs
	define.SetMime(typeName,
		"application/hcl",
		"application/x-hcl",
		"text/hcl",
		"text/x-hcl",
	)

	define.SetFileExtensions(typeName, "hcl", "tf", "tfvars")
}

func readArray(read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	j := make([]interface{}, 0)
	err = hcl.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	for i := range j {
		switch j[i].(type) {
		case string:
			callback(bytes.TrimSpace([]byte(j[i].(string))))

		default:
			jBytes, err := json.Marshal(j[i], false)
			if err != nil {
				return err
			}
			callback(jBytes)
		}
	}

	return nil
}

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var jObj interface{}
	err = hcl.Unmarshal(b, &jObj)
	if err == nil {

		switch v := jObj.(type) {
		case []interface{}:
			for i := range jObj.([]interface{}) {
				j, err := json.Marshal(jObj.([]interface{})[i], false)
				if err != nil {
					return err
				}
				callback(strconv.Itoa(i), string(j), i != len(jObj.([]interface{}))-1)
			}

		case map[string]interface{}, map[interface{}]interface{}:
			i := 1
			for key := range jObj.(map[string]interface{}) {
				j, err := json.Marshal(jObj.(map[string]interface{})[key], false)
				if err != nil {
					return err
				}
				callback(key, string(j), i != len(jObj.(map[string]interface{})))
				i++
			}
			return nil

		default:
			if debug.Enable {
				panic(v)
			}
		}
		return nil
	}
	return err
}

func readIndex(p *proc.Process, params []string) error {
	var jInterface interface{}

	b, err := p.Stdin.ReadAll()
	if err != nil {
		return err
	}

	err = hcl.Unmarshal(b, &jInterface)
	if err != nil {
		return err
	}

	marshaller := func(iface interface{}) ([]byte, error) {
		return json.Marshal(iface, p.Stdout.IsTTY())
	}

	return define.IndexTemplateObject(p, params, &jInterface, marshaller)
}

func marshal(p *proc.Process, v interface{}) ([]byte, error) {
	return json.Marshal(v, p.Stdout.IsTTY())
}

func unmarshal(p *proc.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = hcl.Unmarshal(b, &v)
	return
}
