package define

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang"
)

// IndexTemplateObject is a handy standard indexer you can use in your custom data types for structured object types.
// The point of this is to minimize code rewriting and standardising the behavior of the indexer.
func IndexTemplateObject(p *lang.Process, params []string, object *interface{}, marshaller func(interface{}) ([]byte, error)) error {
	if p.IsNot {
		return itoNot(p, params, object, marshaller)
	}
	return itoIndex(p, params, object, marshaller)
}

// itoIndex allow
func itoIndex(p *lang.Process, params []string, object *interface{}, marshaller func(interface{}) ([]byte, error)) error {
	var objArray []interface{}
	switch v := (*object).(type) {
	case []interface{}:
		for _, key := range params {
			i, err := strconv.Atoi(key)
			if err != nil {
				return err
			}
			if i < 0 {
				//return errors.New("Cannot have negative keys in array")
				i = len(v) + i
			}
			if i >= len(v) {
				return errors.New("Key '" + key + "' greater than number of items in array")
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
		for i := range params {
			switch {
			case v[params[i]] != nil:
			case v[strings.Title(params[i])] != nil:
				params[i] = strings.Title(params[i])
			case v[strings.ToLower(params[i])] != nil:
				params[i] = strings.ToLower(params[i])
			case v[strings.ToUpper(params[i])] != nil:
				params[i] = strings.ToUpper(params[i])
			//case v[strings.ToTitle(params[i])] != nil:
			//	params[i] = strings.ToTitle(params[i])
			default:
				return errors.New("Key '" + params[i] + "' not found")
			}

			if len(params) > 1 {
				objArray = append(objArray, v[params[i]])

			} else {
				switch v[params[i]].(type) {
				case string:
					p.Stdout.Write([]byte(v[params[i]].(string)))
				default:
					b, err := marshaller(v[params[i]])
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
		for i := range params {
			//if v[key] == nil {
			//	return errors.New("Key '" + key + "' not found.")
			//}
			switch {
			case v[params[i]] != nil:
			case v[strings.Title(params[i])] != nil:
				params[i] = strings.Title(params[i])
			case v[strings.ToLower(params[i])] != nil:
				params[i] = strings.ToLower(params[i])
			case v[strings.ToUpper(params[i])] != nil:
				params[i] = strings.ToUpper(params[i])
			//case v[strings.ToTitle(params[i])] != nil:
			//	params[i] = strings.ToTitle(params[i])
			default:
				return errors.New("Key '" + params[i] + "' not found")
			}

			if len(params) > 1 {
				objArray = append(objArray, v[params[i]])

			} else {
				switch v[params[i]].(type) {
				case string:
					p.Stdout.Write([]byte(v[params[i]].(string)))
				default:
					b, err := marshaller(v[params[i]])
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
		return errors.New("Object cannot be indexed")
	}
}

// itoNot requires the indexes to be explicit
func itoNot(p *lang.Process, params []string, object *interface{}, marshaller func(interface{}) ([]byte, error)) error {
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
				return errors.New("Cannot have negative keys in array")
			}
			if i >= len(v) {
				return errors.New("Key '" + key + "' greater than number of items in array")
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
			not[strings.Title(key)] = true
			not[strings.ToLower(key)] = true
			not[strings.ToUpper(key)] = true
			//not[strings.ToTitle(key)] = true
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
			not[strings.Title(key)] = true
			not[strings.ToLower(key)] = true
			not[strings.ToUpper(key)] = true
			//not[strings.ToTitle(key)] = true
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
		return errors.New("Object cannot be !indexed")
	}
}
