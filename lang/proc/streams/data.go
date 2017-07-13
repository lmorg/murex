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

func (in *Stdin) GetDataType() (dt string) {
	for {
		in.dtLock.Lock()
		dt = in.dataType
		in.dtLock.Unlock()
		if dt != "" {
			return
		}
	}
}

func (in *Stdin) SetDataType(dt string) {
	in.dtLock.Lock()
	in.dataType = dt
	in.dtLock.Unlock()
	return
}

func (in *Stdin) DefaultDataType(err bool) {
	return
	in.dtLock.Lock()
	dt := in.dataType
	in.dtLock.Unlock()

	if dt == "" {
		if err {
			in.dtLock.Lock()
			in.dataType = types.Null
			in.dtLock.Unlock()
		} else {
			in.dtLock.Lock()
			in.dataType = types.Generic
			in.dtLock.Unlock()
		}
	}
}

// Stream arrays regardless of data type.
func (read *Stdin) ReadArray(callback func([]byte)) {
	dt := read.GetDataType()
	switch dt {
	case types.Json:
		b := read.ReadAll()
		j := make([]interface{}, 0)
		err := json.Unmarshal(b, &j)
		if err == nil {
			for i := range j {
				switch j[i].(type) {
				case string:
					callback(bytes.TrimSpace([]byte(j[i].(string))))

				default:
					jBytes, _ := utils.JsonMarshal(j[i])
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
	}

	return
}

func (read *Stdin) ReadMap(config *config.Config, callback func(key, value string, last bool)) error {
	dt := read.GetDataType()
	switch dt {
	case types.Json:
		b := read.ReadAll()

		var jObj interface{}
		err := json.Unmarshal(b, &jObj)
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

		/*v, err := config.Get("shell", "Csv-Headings", types.Boolean)
		if err != nil {
			return err
		}
		headings := v.(bool)

		var (
			recHeadings []string
			recNum      int
		)

		err = csvParser.ReadLine(func(fields []string) {
			recNum++

			if headings {
				if recNum == 1 {
					for i := range fields {
						recHeadings = append(recHeadings, strings.TrimSpace(fields[i]))
					}
					return
				}

				l := len(fields) - 1
				for i := range fields {
					if i < len(recHeadings) {
						callback(recHeadings[i], strings.TrimSpace(fields[i]), i == l)
					} else {
						callback(strconv.Itoa(i), strings.TrimSpace(fields[i]), i == l)
					}
				}

			} else {
				l := len(fields) - 1
				for i := range fields {
					callback(strconv.Itoa(i), strings.TrimSpace(fields[i]), i == l)
				}
			}
		})*/

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
