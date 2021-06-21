package json

import (
	"strconv"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/utils/json"
)

func readMap(read stdio.Io, _ *config.Config, callback func(key, value string, last bool)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var jObj interface{}
	err = json.Unmarshal(b, &jObj)
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

		case map[string]interface{}:
			i := 1
			for key := range jObj.(map[string]interface{}) {
				switch jObj.(map[string]interface{})[key].(type) {
				case string:
					callback(key, jObj.(map[string]interface{})[key].(string), i != len(jObj.(map[string]interface{})))

				default:
					j, err := json.Marshal(jObj.(map[string]interface{})[key], false)
					if err != nil {
						return err
					}
					callback(key, string(j), i != len(jObj.(map[string]interface{})))
				}
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
