package json

import (
	"bytes"
	"encoding/json"
	"github.com/lmorg/murex/lang/proc/streams"
)

func readArray(read streams.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	j := make([]interface{}, 0)
	err = json.Unmarshal(b, &j)

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

	//return readArrayDefault(read, callback)
}
