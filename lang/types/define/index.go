package define

import (
	"errors"
	"fmt"
	"github.com/lmorg/murex/lang/proc"
	"strconv"
)

// IndexTemplateObject is a handy standard indexer you can use in your custom data types for structured object types.
// The point of this is to minimize code rewriting and standardising the behavior of the indexer.
func IndexTemplateObject(p *proc.Process, params []string, object *interface{}, marshaller func(interface{}) ([]byte, error)) error {
	if p.IsNot {
		switch v := (*object).(type) {
		case []interface{}:
			var objArray []interface{}
			not := make(map[int]bool)
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

				not[i] = true
			}

			for i := range v {
				if !not[i] {
					objArray = append(objArray, v[i])
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
			objMap := make(map[string]interface{})
			not := make(map[string]bool)
			for _, key := range params {
				not[key] = true
			}

			for s := range v {
				if !not[s] {
					objMap[s] = v[s]
				}
			}

			if len(objMap) > 0 {
				b, err := marshaller(objMap)
				if err != nil {
					return err
				}
				p.Stdout.Writeln(b)
			}
			return nil

		case map[interface{}]interface{}:
			objMap := make(map[interface{}]interface{})
			not := make(map[string]bool)
			for _, key := range params {
				not[key] = true
			}

			for iface := range v {
				s := fmt.Sprint(iface)
				if !not[s] {
					objMap[iface] = v[iface]
				}
			}

			if len(objMap) > 0 {
				b, err := marshaller(objMap)
				if err != nil {
					return err
				}
				p.Stdout.Writeln(b)
			}
			return nil

		default:
			return errors.New("Object cannot be !indexed.")
		}
	}

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
