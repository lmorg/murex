package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/utils"
)

// deinterface is used to fudge around the lack of support for `map[interface{}]interface{}` in Go's JSON marshaller.
func deinterface(v interface{}) interface{} {
	switch t := v.(type) {
	case map[interface{}]interface{}:
		newV := make(map[string]interface{})
		for key := range t {
			newV[fmt.Sprint(key)] = deinterface(t[key])
		}
		fmt.Printf("%T\n", newV)
		return newV

	default:
		fmt.Printf("%T\n", t)
		return v
	}
}

func marshalJson(p *proc.Process, v interface{}) ([]byte, error) {
	//newV := deinterface(v)
	//fmt.Printf("--> %T\n", v)
	//fmt.Printf("==> %T\n", newV)
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
