package streams

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/csv"
	"strconv"
	"strings"
)

// Stream arrays regardless of data type.
func readArray(read Io, callback func([]byte)) error {
	dt := read.GetDataType()
	switch dt {
	case types.Json:
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
					jBytes, err := utils.JsonMarshal(j[i])
					if err != nil {
						return err
					}
					callback(jBytes)
				}
			}
		}
		fallthrough

	default:
		scanner := bufio.NewScanner(read)
		for scanner.Scan() {
			callback(bytes.TrimSpace(scanner.Bytes()))
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}

func readMap(read Io, config *config.Config, callback func(key, value string, last bool)) error {
	dt := read.GetDataType()
	switch dt {
	case types.Json:
		b, err := read.ReadAll()
		if err != nil {
			return err
		}

		var jObj interface{}
		err = json.Unmarshal(b, &jObj)
		if err == nil {

			switch jObj.(type) {
			case map[string]interface{}:
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

			case []interface{}:
				for i := range jObj.([]interface{}) {
					j, err := json.Marshal(jObj.([]interface{})[i])
					if err != nil {
						return err
					}
					callback(strconv.Itoa(i), string(j), i != len(jObj.([]interface{}))-1)
				}
			}
		}
		fallthrough

	case types.Csv:
		csvParser, err := csv.NewParser(read, config)
		if err != nil {
			return err
		}

		err = csvParser.ReadLine(func(records []string, headings []string) {
			for i := range records {
				callback(headings[i], records[i], i == len(records)-1)
			}
		})

		return err

	default:
		scanner := bufio.NewScanner(read)
		i := -1
		for scanner.Scan() {
			i++
			callback(strconv.Itoa(i), strings.TrimSpace(string(scanner.Bytes())), false)
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}
