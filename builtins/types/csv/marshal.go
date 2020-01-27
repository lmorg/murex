package csv

import (
	"bytes"
	enc "encoding/csv"
	"fmt"
	"io"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func marshal(p *lang.Process, iface interface{}) ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)
	w := enc.NewWriter(buf)

	separator, err := p.Config.Get("csv", "separator", types.String)
	if err != nil {
		return nil, err
	}
	if len(separator.(string)) != 0 {
		w.Comma = []rune(separator.(string))[0]
	}

	switch v := iface.(type) {
	case []string:
		for i := range v {
			err = w.Write([]string{v[i]})
			if err != nil {
				return buf.Bytes(), err
			}
		}
		w.Flush()
		return buf.Bytes(), w.Error()

	case [][]string:
		for i := range v {
			err = w.Write(v[i])
			if err != nil {
				return buf.Bytes(), err
			}
		}
		w.Flush()
		return buf.Bytes(), w.Error()

	case []interface{}:
		for i := range v {
			err = w.Write([]string{fmt.Sprint(v[i])})
			if err != nil {
				return buf.Bytes(), err
			}
		}
		w.Flush()
		return buf.Bytes(), w.Error()

	/*case [][]interface{}:
	for i := range v {
		a := make([]string, len(v[i]))
		for j := range v[i] {
			a[j] = fmt.Sprint(v[i][j])
		}
		b = append(b, w.ArrayToCsv(a)...)
		b = append(b, utils.NewLineByte...)
	}
	return*/

	/*case map[string]string:
		var headings []string
		for key := range v {
			headings = append(headings, key)
		}
		sort.Strings(headings)
		b = w.ArrayToCsv(headings)
		b = append(b, utils.NewLineByte...)

		records := make([]string, len(headings))
		for i := range headings {
			records[i] = v[headings[i]]
		}
		b = append(b, w.ArrayToCsv(records)...)

		return

	case map[string]interface{}:
		var headings []string
		for key := range v {
			headings = append(headings, key)
		}
		sort.Strings(headings)
		b = w.ArrayToCsv(headings)
		b = append(b, utils.NewLineByte...)

		records := make([]string, len(headings))
		for i := range headings {
			j, err := json.Marshal(v[headings[i]])
			s := string(j)
			if err != nil {
				s = strings.TrimSpace(fmt.Sprint(v[headings[i]]))
			}
			records[i] = s
		}
		b = append(b, w.ArrayToCsv(records)...)

		return

	case map[interface{}]string:
		var headings []string
		for key := range v {
			headings = append(headings, fmt.Sprint(key))
		}
		sort.Strings(headings)
		b = w.ArrayToCsv(headings)
		b = append(b, utils.NewLineByte...)

		records := make([]string, len(headings))
		for i := range headings {
			records[i] = v[headings[i]]
		}
		b = append(b, w.ArrayToCsv(records)...)

		return

	case map[interface{}]interface{}:
		var headings []string

		for key := range v {
			headings = append(headings, fmt.Sprint(key))
		}
		sort.Strings(headings)
		b = w.ArrayToCsv(headings)
		b = append(b, utils.NewLineByte...)

		records := make([]string, len(headings))
		for i := range headings {
			j, err := json.Marshal(v[headings[i]])
			s := string(j)
			if err != nil {
				s = strings.TrimSpace(fmt.Sprint(v[headings[i]]))
			}
			records[i] = s
		}
		b = append(b, w.ArrayToCsv(records)...)
		return

	case []map[string]string:
		var headings []string

		for slice := range v {
			if len(headings) == 0 {
				for key := range v[slice] {
					headings = append(headings, key)
				}
				sort.Strings(headings)
				b = w.ArrayToCsv(headings)
			}

			//order := make(map[string]int)
			//for i := range headings {
			//	order[headings[i]] = i
			//}

			records := make([]string, len(headings))
			for i := range headings {
				records[i] = v[slice][headings[i]]
			}
			b = append(b, utils.NewLineByte...)
			b = append(b, w.ArrayToCsv(records)...)
		}
		os.Stderr.WriteString("Warning: untested!\n")
		return b, nil

	case []map[string]interface{}:
		var headings []string

		for slice := range v {
			if len(headings) == 0 {
				for key := range v[slice] {
					headings = append(headings, key)
				}
				sort.Strings(headings)
				b = w.ArrayToCsv(headings)
			}

			records := make([]string, len(headings))
			for i := range headings {
				j, err := json.Marshal(v[slice][headings[i]])
				s := string(j)
				if err != nil {
					s = strings.TrimSpace(fmt.Sprint(v[slice][headings[i]]))
				}
				records[i] = s
			}
			b = append(b, utils.NewLineByte...)
			b = append(b, w.ArrayToCsv(records)...)
		}
		os.Stderr.WriteString("Warning: untested!\n")
		return b, nil

	case []map[interface{}]string:
		var headings []string

		for slice := range v {
			if len(headings) == 0 {
				for key := range v[slice] {
					headings = append(headings, fmt.Sprint(key))
				}
				sort.Strings(headings)
				b = w.ArrayToCsv(headings)
			}

			records := make([]string, len(headings))
			for i := range headings {
				j, err := json.Marshal(v[slice][headings[i]])
				s := string(j)
				if err != nil {
					s = strings.TrimSpace(fmt.Sprint(v[slice][headings[i]]))
				}
				records[i] = s
			}
			b = append(b, utils.NewLineByte...)
			b = append(b, w.ArrayToCsv(records)...)
		}
		os.Stderr.WriteString("Warning: untested!\n")
		return b, nil

	case []map[interface{}]interface{}:
		var headings []string

		for slice := range v {
			if len(headings) == 0 {
				for key := range v[slice] {
					headings = append(headings, fmt.Sprint(key))
				}
				sort.Strings(headings)
				b = w.ArrayToCsv(headings)
			}

			records := make([]string, len(headings))
			for i := range headings {
				j, err := json.Marshal(v[slice][headings[i]])
				s := string(j)
				if err != nil {
					s = strings.TrimSpace(fmt.Sprint(v[slice][headings[i]]))
				}
				records[i] = s
				records[i] = fmt.Sprint(v[slice][headings[i]])
			}
			b = append(b, utils.NewLineByte...)
			b = append(b, w.ArrayToCsv(records)...)
		}
		os.Stderr.WriteString("Warning: untested!\n")
		return b, nil*/

	default:
		err = fmt.Errorf("I don't know how to marshal %T data into a `%s`. Data possibly too complex?", v, typeName)
		return buf.Bytes(), err
	}
}

func unmarshal(p *lang.Process) (interface{}, error) {
	csvReader := enc.NewReader(p.Stdin)
	csvReader.TrimLeadingSpace = true

	v, err := p.Config.Get("csv", "separator", types.String)
	if err != nil {
		return nil, err
	}
	if len(v.(string)) != 0 {
		csvReader.Comma = rune(v.(string)[0])
	}

	v, err = p.Config.Get("csv", "comment", types.String)
	if err != nil {
		return nil, err
	}
	if len(v.(string)) != 0 {
		csvReader.Comment = rune(v.(string)[0])
	}

	var table [][]string

	for {
		record, err := csvReader.Read()
		if record == nil && err == io.EOF {
			break
		}
		table = append(table, record)
	}

	return table, nil
}
