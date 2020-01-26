package csv

import (
	enc "encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

func marshal(p *lang.Process, iface interface{}) (b []byte, err error) {
	w, err := NewParser(nil, p.Config)
	if err != nil {
		return
	}

	switch v := iface.(type) {
	case []string:
		for i := range v {
			s := strings.TrimSpace(v[i])
			b = append(b, w.ArrayToCsv([]string{s})...)
			b = append(b, utils.NewLineByte...)
		}
		return

	case [][]string:
		for i := range v {
			//s := strings.TrimSpace(v[i])
			b = append(b, w.ArrayToCsv(v[i])...)
			b = append(b, utils.NewLineByte...)
		}
		return

	case []interface{}:
		for i := range v {
			j, err := json.Marshal(v[i])
			s := string(j)
			if err != nil {
				s = strings.TrimSpace(fmt.Sprintln(v[i]))
			}
			b = append(b, w.ArrayToCsv([]string{s})...)
			b = append(b, utils.NewLineByte...)
		}
		return

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

	case map[string]string:
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
		return b, nil

	default:
		err = fmt.Errorf("I don't know how to marshal %T data into a `%s`. Data possibly too complex?", v, typeName)
		return
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
