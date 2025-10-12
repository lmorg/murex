package markdown

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func marshal(p *lang.Process, iface any) ([]byte, error) {
	var (
		buf = bytes.NewBuffer([]byte{})
		w   = tabwriter.NewWriter(buf, 1, 2, 1, ' ', tabwriter.Debug)
		err error
	)

switchCondition:
	switch v := iface.(type) {
	/*case []string:
	for i := range v {
		err = w.Write([]string{v[i]})
		if err != nil {
			return buf.Bytes(), err
		}
	}*/

	case [][]string:
		for i := range v {
			_, err = w.Write([]byte("| " + strings.Join(v[i], "\t ") + " |\n"))
			if err != nil {
				return buf.Bytes(), err
			}
		}

	case []any:
		if len(v) == 0 {
			return buf.Bytes(), nil
		}

		if reflect.TypeOf(v[0]).Kind() != reflect.Map {
			panic("TODO")
		}

		err = types.MapToTable_Any(v, func(s []string) error {
			_, err := w.Write([]byte("| " + strings.Join(s, "\t ") + "|\n"))
			return err
		})
		if err != nil {
			return buf.Bytes(), err
		}

	case []map[string]any:
		err = types.MapToTable_MapStringAny(v, func(s []string) error {
			_, err := w.Write([]byte("| " + strings.Join(s, "\t ") + "|\n"))
			return err
		})
		if err != nil {
			return buf.Bytes(), err
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

	err = w.Flush()
	return buf.Bytes(), err
}

/*func unmarshal(p *lang.Process) (any, error) {
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
}*/
