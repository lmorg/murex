package csv

import (
	"bytes"
	enc "encoding/csv"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func marshal(p *lang.Process, iface any) ([]byte, error) {
	inputType := p.Stdin.GetDataType()

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

	leftTrimV, err := p.Config.Get("csv", "left-trim", types.Boolean)
	if err != nil {
		return nil, err
	}

	leftTrim := leftTrimV.(bool) && (inputType == types.Generic || inputType == types.String)
switchCondition:
	switch v := iface.(type) {
	case []string:
		for i := range v {
			err = w.Write([]string{v[i]})
			if err != nil {
				return buf.Bytes(), err
			}
		}

	case [][]string:
		if len(v) == 0 || len(v[0]) == 0 ||
			(len(v[0][0]) != 0 && v[0][0][0] != ' ' && v[0][0][0] != '\t') {
			leftTrim = false
		}
		for i := range v {
			if leftTrim {
				v[i][0] = strings.TrimSpace(v[i][0])
				if len(v[i][0]) == 0 {
					v[i] = v[i][1:]
				}
			}
			err = w.Write(v[i])
			if err != nil {
				return buf.Bytes(), err
			}
		}

	case []any:
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
			break
		}

		err = types.MapToTable_Any(v, func(s []string) error { return w.Write(s) })
		if err != nil {
			return nil, err
		}

	case []map[string]any:
		err = types.MapToTable_MapStringAny(v, func(s []string) error { return w.Write(s) })
		if err != nil {
			return nil, err
		}

	case map[string]any:
		const xmlDefaultElement = "list"
		if len(v) != 1 {
			return buf.Bytes(), fmt.Errorf("cannot marshal %T data into a `%s`", v, typeName)
		}
		el, ok := v[xmlDefaultElement]
		if !ok {
			return buf.Bytes(), fmt.Errorf("cannot marshal %T data into a `%s`\nmissing %s element", v, typeName, xmlDefaultElement)
		}

		switch t := el.(type) {
		case []any, []map[string]any:
			iface = t
			goto switchCondition
		default:
			return buf.Bytes(), fmt.Errorf("cannot marshal %T data into a `%s`\n%s element is not an array", v, typeName, xmlDefaultElement)
		}

	default:
		return buf.Bytes(), fmt.Errorf("cannot marshal %T data into a `%s`", v, typeName)
	}

	w.Flush()
	err = w.Error()
	return buf.Bytes(), err
}

func unmarshal(p *lang.Process) (any, error) {
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
