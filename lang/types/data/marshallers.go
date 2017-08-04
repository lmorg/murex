package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
	"github.com/lmorg/murex/utils/csv"
	"strings"
)

func marshalCsv(_ *proc.Process, iface interface{}) (b []byte, err error) {
	w, err := csv.NewParser(nil, &proc.GlobalConf)
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

	case []interface{}:
		for i := range v {
			s := strings.TrimSpace(fmt.Sprintln(v[i]))
			b = append(b, w.ArrayToCsv([]string{s})...)
			b = append(b, utils.NewLineByte...)
		}
		return

	/*case map[string]string:
	var
	var array []string
	for s := range v {
		array=append(array,v[s])
	}
	sort.Strings(array)
	for i:= range
	b = w.ArrayToCsv()
	for s := range v {
		b = append(b, []byte(s+": "+v[s]+utils.NewLineString)...)
	}
	return*/

	/*case map[string]interface{}:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", s, fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return

	case map[interface{}]interface{}:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", fmt.Sprint(s), fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return

	case map[interface{}]string:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", fmt.Sprint(s), v[s], utils.NewLineString))...)
		}
		return

	case interface{}:
		return []byte(fmt.Sprintln(iface)), nil*/

	default:
		err = errors.New("I don't know how to marshal that data into a `csv`. Data possibly too complex?")
		return
	}
}

func unmarshalCsv(p *proc.Process) (interface{}, error) {
	csvReader, err := csv.NewParser(p.Stdin, &proc.GlobalConf)
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

func marshalJson(p *proc.Process, v interface{}) ([]byte, error) {
	return utils.JsonMarshal(v, p.Stdout.IsTTY())
}

func unmarshalJson(p *proc.Process) (v interface{}, err error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &v)
	return
}

func marshalString(_ *proc.Process, iface interface{}) (b []byte, err error) {
	switch v := iface.(type) {
	case []string:
		for i := range v {
			b = append(b, []byte(v[i]+utils.NewLineString)...)
		}
		return

	case []interface{}:
		for i := range v {
			b = append(b, []byte(fmt.Sprintln(v[i]))...)
		}
		return

	case map[string]string:
		for s := range v {
			b = append(b, []byte(s+": "+v[s]+utils.NewLineString)...)
		}
		return

	case map[string]interface{}:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", s, fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return

	case map[interface{}]interface{}:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", fmt.Sprint(s), fmt.Sprint(v[s]), utils.NewLineString))...)
		}
		return

	case map[interface{}]string:
		for s := range v {
			b = append(b, []byte(fmt.Sprintf("%s: %s%s", fmt.Sprint(s), v[s], utils.NewLineString))...)
		}
		return

	case interface{}:
		return []byte(fmt.Sprintln(iface)), nil

	default:
		err = errors.New("I don't know how to marshal that data into a `str`. Data possibly too complex?")
		return
	}
}

func unmarshalString(p *proc.Process) (interface{}, error) {
	s := make([]string, 0)
	err := p.Stdin.ReadLine(func(b []byte) {
		s = append(s, string(b))
	})

	return s, err
}
