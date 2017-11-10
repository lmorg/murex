package string

import (
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"net/url"
	"strconv"
)

func marshal(_ *proc.Process, iface interface{}) (b []byte, err error) {
	var qs url.Values

	switch v := iface.(type) {
	case []string:
		for i := range v {
			qs.Add(strconv.Itoa(i), v[i])
		}
		b = []byte(qs.Encode())

	case []interface{}:
		for i := range v {
			qs.Add(strconv.Itoa(i), fmt.Sprint(v[i]))
		}
		b = []byte(qs.Encode())

	case map[string]string:
		for s := range v {
			qs.Add(s, v[s])
		}
		b = []byte(qs.Encode())

	case map[string]interface{}:
		for s := range v {
			qs.Add(s, fmt.Sprint(v[s]))
		}
		b = []byte(qs.Encode())

	case map[interface{}]interface{}:
		for s := range v {
			qs.Add(fmt.Sprint(s), fmt.Sprint(v[s]))
		}
		b = []byte(qs.Encode())

	case map[interface{}]string:
		for s := range v {
			qs.Add(fmt.Sprint(s), v[s])
		}
		b = []byte(qs.Encode())

	case interface{}:
		qs.Add(fmt.Sprint(v), "")

	default:
		err = errors.New("I don't know how to marshal that data into a `" + dataType + "`. Data possibly too complex?")
	}

	return
}

func unmarshal(p *proc.Process) (interface{}, error) {
	b, err := p.Stdin.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(b) == 0 {
		return nil, nil
	}

	if b[0] == '?' {
		if len(b) == 1 {
			return nil, nil
		}
		b = b[1:]
	}

	values, err := url.ParseQuery(string(b))
	if err != nil {
		return nil, err
	}

	qs := make(map[string]interface{})
	for s := range values {
		if len(values[s]) == 1 {
			float, tnErr := toNumber(values[s][0])
			if tnErr != nil {
				qs[s] = values[s][0]
				continue
			}
			qs[s] = float

		} else {
			qs[s] = values[s]
		}
	}

	return qs, nil
}

func toNumber(s string) (f float64, err error) {
	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return
	}

	if s != strconv.FormatFloat(f, 'f', -1, 64) {
		err = errors.New("Input doesn't match converted output. Possibly due to padding?")
	}

	return
}
