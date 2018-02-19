package csv

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
)

func marshal(p *proc.Process, iface interface{}) (b []byte, err error) {
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
		err = errors.New("I don't know how to marshal that data into a `csv`. Data possibly too complex?")
		return
	}
}

func unmarshal(p *proc.Process) (interface{}, error) {
	csvReader, err := NewParser(p.Stdin, p.Config)
	if err != nil {
		return nil, err
	}

	table := make([]map[string]string, 0)
	csvReader.ReadLine(func(recs []string, heads []string) {
		record := make(map[string]string)
		for i := range recs {
			record[heads[i]] = recs[i]
		}
		table = append(table, record)
	})

	return table, nil
}
