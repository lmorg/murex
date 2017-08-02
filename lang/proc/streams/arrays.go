package streams

import (
	"bufio"
	"bytes"
	"encoding/json"
)

func readArrayJson(read Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	j := make([]interface{}, 0)
	err = json.Unmarshal(b, &j)
	if err == nil {
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

	return readArrayDefault(read, callback)
}

func readArrayDefault(read Io, callback func([]byte)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		callback(bytes.TrimSpace(scanner.Bytes()))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
