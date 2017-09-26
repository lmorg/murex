package hcl

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/hashicorp/hcl"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/proc/streams"
	"github.com/lmorg/murex/lang/proc/streams/stdio"
	"github.com/lmorg/murex/lang/types/define"
	"github.com/lmorg/murex/utils"
	"strconv"
)

const typeName = "hcl"

func init() {
	streams.ReadArray[typeName] = readArray
	streams.ReadMap[typeName] = readMap
	define.ReadIndexes[typeName] = readIndex
	define.Marshal[typeName] = marshal
	define.Unmarshal[typeName] = unmarshal

	// These are just guessed at as I couldn't find any formally named MIMEs
	define.SetMime(typeName,
		"application/hcl",
		"application/x-hcl",
		"text/hcl",
		"text/x-hcl",
	)

	define.SetFileExtensions(typeName, "hcl")
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
			jBytes, err := json.Marshal(j[i])
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
				j, err := json.Marshal(jObj.([]interface{})[i])
				if err != nil {
					return err
				}
				callback(strconv.Itoa(i), string(j), i != len(jObj.([]interface{}))-1)
			}

		case map[string]interface{}, map[interface{}]interface{}:
			i := 1
			for key := range jObj.(map[string]interface{}) {
				j, err := json.Marshal(jObj.(map[string]interface{})[key])
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

	var jArray []interface{}
	switch v := jInterface.(type) {
	case []interface{}:
		for _, key := range params {
			i, err := strconv.Atoi(key)
			if err != nil {
				return err
			}
			if i < 0 {
				return errors.New("Cannot have negative keys in array.")
			}
			if i >= len(v) {
				return errors.New("Key '" + key + "' greater than number of items in array.")
			}

			if len(params) > 1 {
				jArray = append(jArray, v[i])

			} else {
				switch v[i].(type) {
				case string:
					p.Stdout.Write([]byte(v[i].(string)))
				default:
					b, err := utils.JsonMarshal(v[i], p.Stdout.IsTTY())
					if err != nil {
						return err
					}
					p.Stdout.Writeln(b)
				}
			}
		}
		if len(jArray) > 0 {
			b, err := utils.JsonMarshal(jArray, p.Stdout.IsTTY())
			if err != nil {
				return err
			}
			p.Stdout.Writeln(b)
		}
		return nil

	case map[string]interface{}:
		for _, key := range params {
			if v[key] == nil {
				return errors.New("Key '" + key + "' not found.")
			}

			if len(params) > 1 {
				jArray = append(jArray, v[key])

			} else {
				switch v[key].(type) {
				case string:
					p.Stdout.Write([]byte(v[key].(string)))
				default:
					b, err := utils.JsonMarshal(v[key], p.Stdout.IsTTY())
					if err != nil {
						return err
					}
					p.Stdout.Writeln(b)
				}
			}
		}
		if len(jArray) > 0 {
			b, err := utils.JsonMarshal(jArray, p.Stdout.IsTTY())
			if err != nil {
				return err
			}
			p.Stdout.Writeln(b)
		}
		return nil

	case map[interface{}]interface{}:
		for _, key := range params {
			if v[key] == nil {
				return errors.New("Key '" + key + "' not found.")
			}

			if len(params) > 1 {
				jArray = append(jArray, v[key])

			} else {
				switch v[key].(type) {
				case string:
					p.Stdout.Write([]byte(v[key].(string)))
				default:
					b, err := utils.JsonMarshal(v[key], p.Stdout.IsTTY())
					if err != nil {
						return err
					}
					p.Stdout.Writeln(b)
				}
			}
		}
		if len(jArray) > 0 {
			b, err := utils.JsonMarshal(jArray, p.Stdout.IsTTY())
			if err != nil {
				return err
			}
			p.Stdout.Writeln(b)
		}
		return nil

	default:
		return errors.New("HCL object cannot be indexed.")
	}
}

func marshal(p *proc.Process, v interface{}) ([]byte, error) {
	return utils.JsonMarshal(v, p.Stdout.IsTTY())
}

func unmarshal(p *proc.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = hcl.Unmarshal(b, &v)
	return
}
