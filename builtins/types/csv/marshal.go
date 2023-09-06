package csv

import (
	"bytes"
	enc "encoding/csv"
	"fmt"
	"io"
	"reflect"

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
		if len(v) == 0 {
			w.Flush()
			return buf.Bytes(), w.Error()
		}

		if reflect.TypeOf(v[0]).Kind() != reflect.Map {
			for i := range v {
				err = w.Write([]string{fmt.Sprint(v[i])})
				if err != nil {
					return buf.Bytes(), err
				}
			}
			w.Flush()
			return buf.Bytes(), w.Error()
		}

		err = types.MapToTable(v, func(s []string) error { return w.Write(s) })
		if err != nil {
			return nil, err
		}
		w.Flush()
		return buf.Bytes(), w.Error()

	default:
		err = fmt.Errorf("cannot marshal %T data into a `%s`", v, typeName)
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
