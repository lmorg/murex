package typemgmt

import (
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"strconv"
	"strings"
)

func init() {
	proc.GoFunctions["append"] = cmdAppend
	proc.GoFunctions["prepend"] = cmdPrepend
	proc.GoFunctions["update"] = cmdUpdate
}

func cmdPrepend(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	var array []string

	err := p.Stdin.ReadArray(func(b []byte) {
		array = append(array, string(b))
	})

	if err != nil {
		return err
	}

	array = append(p.Parameters.StringArray(), array...)

	b, err := define.MarshalData(p, dt, array)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdAppend(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	var array []string

	err := p.Stdin.ReadArray(func(b []byte) {
		array = append(array, string(b))
	})

	if err != nil {
		return err
	}

	array = append(array, p.Parameters.StringArray()...)

	b, err := define.MarshalData(p, dt, array)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}

func cmdUpdate(p *proc.Process) error {
	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	v, err := define.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	s, err := p.Parameters.String(0)
	if err != nil {
		return err
	}

	new, err := p.Parameters.String(1)
	if err != nil {
		return err
	}

	path := strings.Split(s, string(s[0]))
	if len(path) == 0 || (len(path) == 1 && path[0] == "") {
		return errors.New("Empty path.")
	}

	if path[0] == "" {
		path = path[1:]
	}

	var loop func(interface{}, int) (interface{}, error)
	loop = func(v interface{}, i int) (retIface interface{}, retErr error) {
		switch {
		case i < len(path):
			switch t := v.(type) {
			case map[interface{}]interface{}:
				retIface, retErr = loop(v.(map[interface{}]interface{})[path[i]], i+1)
				if err == nil {
					v.(map[interface{}]interface{})[path[i]] = retIface
					retIface = v
				}

			case map[string]interface{}:
				retIface, retErr = loop(v.(map[string]interface{})[path[i]], i+1)
				if err == nil {
					v.(map[string]interface{})[path[i]] = retIface
					retIface = v
				}

			case map[interface{}]string:
				retIface, retErr = loop(v.(map[interface{}]string)[path[i]], i+1)
				if err == nil {
					v.(map[interface{}]string)[path[i]] = retIface.(string)
					retIface = v
				}

			default:
				return nil, fmt.Errorf("murex code error: No condition is made for `%T`.", t)
			}

		case i == len(path):
			switch t := v.(type) {
			case string:
				retIface = new

			case int:
				num, err := strconv.Atoi(new)
				if err != nil {
					return nil, err
				}
				retIface = num

			case float64:
				num, err := strconv.ParseFloat(new, 64)
				if err != nil {
					return nil, err
				}
				retIface = num

			case bool:
				retIface = types.IsTrue([]byte(new), 0)

			default:
				return nil, fmt.Errorf("Cannot locate `%s` in object path or no condition is made for `%T`.", path[i-1], t)
			}

		default:
			return nil, errors.New("I don't know how I got here!")
		}

		return retIface, retErr
	}

	v, err = loop(v, 0)
	if err != nil {
		return err
	}

	b, err := define.MarshalData(p, dt, v)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
