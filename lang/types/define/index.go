package define

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"strconv"
)

func IndexTemplateObject(p *proc.Process, params []string, object *interface{}, marshaller func(interface{}) ([]byte, error)) error {
	var objArray []interface{}
	switch v := (*object).(type) {
	case []interface{}:
		for _, key := range params {
			i, err := strconv.Atoi(key)
			if err != nil {
				return err
			}
			if i < 0 {
				return errors.New("Cannot have negative keys in array.")
			}
			if i >= len(v) {
				return errors.New("Key '" + key + "' greater than number of items in array.")
			}

			if len(params) > 1 {
				objArray = append(objArray, v[i])

			} else {
				switch v[i].(type) {
				case string:
					p.Stdout.Write([]byte(v[i].(string)))
				default:
					b, err := marshaller(v[i])
					if err != nil {
						return err
					}
					p.Stdout.Writeln(b)
				}
			}
		}
		if len(objArray) > 0 {
			b, err := marshaller(objArray)
			if err != nil {
				return err
			}
			p.Stdout.Writeln(b)
		}
		return nil

	case map[string]interface{}:
		for _, key := range params {
			if v[key] == nil {
				return errors.New("Key '" + key + "' not found.")
			}

			if len(params) > 1 {
				objArray = append(objArray, v[key])

			} else {
				switch v[key].(type) {
				case string:
					p.Stdout.Write([]byte(v[key].(string)))
				default:
					b, err := marshaller(v[key])
					if err != nil {
						return err
					}
					p.Stdout.Writeln(b)
				}
			}
		}
		if len(objArray) > 0 {
			b, err := marshaller(objArray)
			if err != nil {
				return err
			}
			p.Stdout.Writeln(b)
		}
		return nil

	case map[interface{}]interface{}:
		for _, key := range params {
			if v[key] == nil {
				return errors.New("Key '" + key + "' not found.")
			}

			if len(params) > 1 {
				objArray = append(objArray, v[key])

			} else {
				switch v[key].(type) {
				case string:
					p.Stdout.Write([]byte(v[key].(string)))
				default:
					b, err := marshaller(v[key])
					if err != nil {
						return err
					}
					p.Stdout.Writeln(b)
				}
			}
		}
		if len(objArray) > 0 {
			b, err := marshaller(objArray)
			if err != nil {
				return err
			}
			p.Stdout.Writeln(b)
		}
		return nil

	default:
		return errors.New("Object cannot be indexed.")
	}
}
