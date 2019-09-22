package bson

import (
	"bytes"
	"strconv"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"labix.org/v2/mgo/bson"
)

const typeName = "bson"

func init() {
	stdio.RegesterReadArray(typeName, readArray)
	stdio.RegesterReadMap(typeName, readMap)

	lang.ReadIndexes[typeName] = readIndex
	lang.ReadNotIndexes[typeName] = readIndex
	lang.Marshallers[typeName] = marshal
	lang.Unmarshallers[typeName] = unmarshal

	// These are just guessed at as I couldn't find any formally named MIMEs
	lang.SetMime(typeName,
		"application/bson",
		"application/x-bson",
		"text/bson",
		"text/x-bson",
	)

	lang.SetFileExtensions(typeName, "bson")
}

func readArray(read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	j := make([]interface{}, 0)
	err = bson.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	for i := range j {
		switch j[i].(type) {
		case string:
			callback(bytes.TrimSpace([]byte(j[i].(string))))

		default:
			jBytes, err := bson.Marshal(j[i])
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
	err = bson.Unmarshal(b, &jObj)
	if err == nil {

		switch v := jObj.(type) {
		case []interface{}:
			for i := range jObj.([]interface{}) {
				j, err := bson.Marshal(jObj.([]interface{})[i])
				if err != nil {
					return err
				}
				callback(strconv.Itoa(i), string(j), i != len(jObj.([]interface{}))-1)
			}

		case map[string]interface{}, map[interface{}]interface{}:
			i := 1
			for key := range jObj.(map[string]interface{}) {
				j, err := bson.Marshal(jObj.(map[string]interface{})[key])
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
	}
	return err
}

func marshal(_ *lang.Process, v interface{}) ([]byte, error) {
	return bson.Marshal(v)
}

func unmarshal(p *lang.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = bson.Unmarshal(b, &v)
	return
}
