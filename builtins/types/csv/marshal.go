package csv

import (
	"bytes"
	enc "encoding/csv"
	"fmt"
	"io"
	"reflect"
	"sort"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

type T comparable

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
		if len(v) == 0 {
			w.Flush()
			return buf.Bytes(), w.Error()
		}
		if reflect.TypeOf(v[0]).Kind() == reflect.Map {
			return marshalSliceOfMap(v, buf, w)
		}

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
		tty.Stderr.WriteString("Warning: untested!\n")
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
		tty.Stderr.WriteString("Warning: untested!\n")
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
		tty.Stderr.WriteString("Warning: untested!\n")
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
		tty.Stderr.WriteString("Warning: untested!\n")
		return b, nil*/

	default:
		err = fmt.Errorf("cannot marshal %T data into a `%s`", v, typeName)
		return buf.Bytes(), err
	}
}

func marshalSliceOfMap(v []interface{}, buf *bytes.Buffer, w *enc.Writer) ([]byte, error) {
	headings, err := getMapKeys(v[0].(map[string]any))
	if err != nil {
		return nil, err
	}

	err = w.Write(headings)
	if err != nil {
		return nil, err
	}

	lenHeadings := len(headings)
	slice := make([]string, lenHeadings)
	var j int

	for i := range v {
		if reflect.TypeOf(v[i]).Kind() != reflect.Map {
			return nil, fmt.Errorf("expecting map on row %d, instead got a %s", i, reflect.TypeOf(v[i]).Kind().String())
		}

		if len(v[i].(map[string]any)) != len(headings) {
			return nil, fmt.Errorf("row %d has a different number of records to the first row:\nrow 0 == %d records,\nrow %d == %d records",
				i, lenHeadings, i, len(v))
		}

		for j = 0; j < lenHeadings; j++ {
			val, ok := v[i].(map[string]any)[headings[j]]
			if !ok {
				return nil, fmt.Errorf("row %d is missing a record name found in the first row: '%s'", i, headings[j])
			}
			s, err := types.ConvertGoType(val, types.String)
			if err != nil {
				return nil, fmt.Errorf("cannot convert a %T (%v) to a %s in record %d: %s", val, val, types.String, i, err.Error())
			}
			slice[j] = s.(string)
		}

		err = w.Write(slice)
		if err != nil {
			return nil, err
		}
	}

	w.Flush()
	return buf.Bytes(), w.Error()
}

func getMapKeys[T comparable](v map[string]T) ([]string, error) {
	slice := make([]string, len(v))
	var i int

	for k := range v {
		s, err := types.ConvertGoType(k, types.String)
		if err != nil {
			return nil, fmt.Errorf("cannot convert a %T (%v) to a %s: %s", k, k, types.String, err.Error())
		}
		slice[i] = s.(string)
		i++
	}
	sort.Strings(slice)
	return slice, nil
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
