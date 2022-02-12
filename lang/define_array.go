package lang

import (
	"fmt"

	"github.com/lmorg/murex/lang/stdio"
)

// ArrayTemplate is a template function for reading arrays from marshalled data
func ArrayTemplate(marshal func(interface{}) ([]byte, error), unmarshal func([]byte, interface{}) error, read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var v interface{}
	err = unmarshal(b, &v)

	if err != nil {
		return err
	}

	switch v := v.(type) {
	case string:
		return readArrayByString(v, callback)

	case []string:
		return readArrayBySliceString(v, callback)

	case []interface{}:
		return readArrayBySliceInterface(marshal, v, callback)

	case map[string]string:
		return readArrayByMapStrStr(v, callback)

	case map[string]interface{}:
		return readArrayByMapStrIface(marshal, v, callback)

	case map[interface{}]string:
		return readArrayByMapIfaceStr(v, callback)

	case map[interface{}]interface{}:
		return readArrayByMapIfaceIface(marshal, v, callback)

	default:
		jBytes, err := marshal(v)
		if err != nil {

			return err
		}

		callback(jBytes)

		return nil
	}
}

func readArrayByString(v string, callback func([]byte)) error {
	callback([]byte(v))

	return nil
}

func readArrayBySliceString(v []string, callback func([]byte)) error {
	for i := range v {
		callback([]byte(v[i]))
	}

	return nil
}

func readArrayBySliceInterface(marshal func(interface{}) ([]byte, error), v []interface{}, callback func([]byte)) error {
	if len(v) == 0 {
		return nil
	}

	for i := range v {
		switch v := v[i].(type) {
		case string:
			callback([]byte(v))

		case []byte:
			callback(v)

		default:
			jBytes, err := marshal(v)
			if err != nil {
				return err
			}
			callback(jBytes)
		}
	}

	return nil
}

func readArrayByMapIfaceIface(marshal func(interface{}) ([]byte, error), v map[interface{}]interface{}, callback func([]byte)) error {
	for key, val := range v {

		bKey := []byte(fmt.Sprint(key) + ": ")
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}

	return nil
}

func readArrayByMapStrStr(v map[string]string, callback func([]byte)) error {
	for key, val := range v {

		callback([]byte(key + ": " + val))
	}

	return nil
}

func readArrayByMapStrIface(marshal func(interface{}) ([]byte, error), v map[string]interface{}, callback func([]byte)) error {
	for key, val := range v {

		bKey := []byte(key + ": ")
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}

	return nil
}

func readArrayByMapIfaceStr(v map[interface{}]string, callback func([]byte)) error {
	for key, val := range v {

		callback([]byte(fmt.Sprint(key) + ": " + val))
	}

	return nil
}
