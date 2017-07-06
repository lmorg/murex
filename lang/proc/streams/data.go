package streams

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"io"
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
		var jMap map[string]interface{}
		err := json.Unmarshal(b, &jMap)
		if err == nil {
			i := 1
			for key := range jMap {
				j, err := json.Marshal(jMap[key])
				if err != nil {
					return err
				}
				callback(key, string(j), i != len(jMap))
				i++
			}
			return nil
		}
		fallthrough

	case types.Csv:
		r := csv.NewReader(read)
		r.LazyQuotes = true
		r.TrimLeadingSpace = false
		//r.FieldsPerRecord = -1

		v, err := config.Get("shell", "Csv-Separator", types.String)
		if err != nil {
			return err
		}
		if len(v.(string)) > 0 {
			r.Comma = []rune(v.(string))[0]
		}

		v, err = config.Get("shell", "Csv-Comment", types.String)
		if err != nil {
			return err
		}
		if len(v.(string)) > 0 {
			r.Comment = []rune(v.(string))[0]
		}

		v, err = config.Get("shell", "Csv-Headings", types.Boolean)
		if err != nil {
			return err
		}

		var (
			useHeadings bool = v.(bool)
			recHeadings []string
			recNum      int
		)

		for {
			recNum++
			fields, err := r.Read()
			switch {
			case err == io.EOF:
				return nil
			case err != nil:
				return err
			}

			if useHeadings {
				if recNum == 1 {
					for i := range fields {
						recHeadings = append(recHeadings, strings.TrimSpace(fields[i]))
					}
					//r.FieldsPerRecord = len(fields)
					continue
				}

				l := len(fields) - 2
				for i := range fields {
					if i < len(recHeadings) {
						callback(recHeadings[i], strings.TrimSpace(fields[i]), i == l)
					} else {
						callback(strconv.Itoa(i), strings.TrimSpace(fields[i]), i == l)
					}
				}

			} else {
				l := len(fields) - 2
				for i := range fields {
					//callback(fmt.Sprintf("%d:%d", recNum, i), strings.TrimSpace(fields[i]), i == l)
					callback(strconv.Itoa(i), strings.TrimSpace(fields[i]), i == l)
				}
			}
		}

	default:
		scanner := bufio.NewScanner(read)
		var i int
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

/*func (out *Stdin) WriteArray(item string) error {
	out.
	return nil
}*/
