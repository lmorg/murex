package streams

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"io"
	"strconv"
)

func (in *Stdin) GetDataType() (dt string) {
	for dt == "" {
		in.mutex.Lock()
		dt = in.dataType
		in.mutex.Unlock()
	}
	return
}

func (in *Stdin) SetDataType(dt string) {
	in.mutex.Lock()
	in.dataType = dt
	in.mutex.Unlock()
	return
}

func (in *Stdin) DefaultDataType(err bool) {
	in.mutex.Lock()
	if in.dataType == "" {
		if err {
			in.dataType = types.Null
		} else {
			in.dataType = types.Generic
		}
	}
	in.mutex.Unlock()
}

// Stream arrays regardless of data type.
// Though currently only 'strings' support streaming, but since this is now a single API it gives an easy place to
// upgrade multiple builtins.
func (read *Stdin) ReadArray(callback func([]byte)) {
	switch read.GetDataType() {
	case types.Json:
		b := read.ReadAll()
		j := make([]string, 0)
		err := json.Unmarshal(b, &j)
		if err == nil {
			for i := range j {
				callback([]byte(j[i]))
			}
			return
		}
		fallthrough

	default:
		read.ReadLine(callback)
	}

	return
}

func (read *Stdin) ReadMap(config *config.Config, callback func(key, value string, last bool)) error {
	switch read.GetDataType() {
	/*case types.Json:
	b := read.ReadAll()
	j := make(map[string]string, 0)
	err := json.Unmarshal(b, &j)
	if err == nil {
		for i := range j {
			callback([]byte(j[i]))
		}
		return
	}
	fallthrough*/

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
					recHeadings = fields
					//r.FieldsPerRecord = len(fields)
					continue
				}

				l := len(fields) - 2
				for i := range fields {
					if i < len(recHeadings) {
						callback(recHeadings[i], fields[i], i == l)
					} else {
						callback(strconv.Itoa(i), fields[i], i == l)
					}
				}

			} else {
				l := len(fields) - 2
				for i := range fields {
					callback(fmt.Sprintf("%d:%d", recNum, i), fields[i], i == l)
				}
			}
		}

	default:
		scanner := bufio.NewScanner(read)
		var i int
		for scanner.Scan() {
			i++
			callback(strconv.Itoa(i), string(scanner.Bytes()), false)
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}
