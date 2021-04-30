package lang

import (
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

// ArrayWithTypeTemplate is a template function for reading arrays from marshalled data
func ArrayWithTypeTemplate(dataType string, marshal func(interface{}) ([]byte, error), unmarshal func([]byte, interface{}) error, read stdio.Io, callback func([]byte, string)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var v interface{}
	err = unmarshal(b, &v)

	if err != nil {
		return err
	}

	switch v.(type) {
	case string:
		return readArrayWithTypeByString(v.(string), callback)

	case []string:
		return readArrayWithTypeBySliceString(v.([]string), callback)

	case []interface{}:
		return readArrayWithTypeBySliceInterface(dataType, marshal, v.([]interface{}), callback)

	/*case map[string]string:
		return readArrayWithTypeByMapStrStr(v.(map[string]string), callback)

	case map[string]interface{}:
		return readArrayWithTypeByMapStrIface(marshal, v.(map[string]interface{}), callback)

	case map[interface{}]string:
		return readArrayWithTypeByMapIfaceStr(v.(map[interface{}]string), callback)

	case map[interface{}]interface{}:
		return readArrayWithTypeByMapIfaceIface(marshal, v.(map[interface{}]interface{}), callback)
	*/
	default:
		jBytes, err := marshal(v)
		if err != nil {

			return err
		}

		callback(jBytes, dataType)

		return nil
	}
}

func readArrayWithTypeByString(v string, callback func([]byte, string)) error {
	callback([]byte(v), types.String)

	return nil
}

func readArrayWithTypeBySliceString(v []string, callback func([]byte, string)) error {
	for i := range v {
		callback([]byte(v[i]), types.String)
	}

	return nil
}

func readArrayWithTypeBySliceInterface(dataType string, marshal func(interface{}) ([]byte, error), v []interface{}, callback func([]byte, string)) error {
	if len(v) == 0 {
		return nil
	}

	switch v[0].(type) {
	case string:
		for i := range v {
			callback([]byte(v[i].(string)), types.String)
		}

	case []byte:
		for i := range v {
			callback(v[i].([]byte), types.String)
		}

	default:
		for i := range v {

			jBytes, err := marshal(v[i])
			if err != nil {
				return err
			}
			callback(jBytes, dataType)

		}
	}

	return nil
}

/*func readArrayWithTypeByMapIfaceIface(marshal func(interface{}) ([]byte, error), v map[interface{}]interface{}, callback func([]byte, string)) error {
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

func readArrayWithTypeByMapStrStr(v map[string]string, callback func([]byte, string)) error {
	for key, val := range v {

		callback([]byte(key + ": " + val))
	}

	return nil
}

func readArrayWithTypeByMapStrIface(marshal func(interface{}) ([]byte, error), v map[string]interface{}, callback func([]byte, string)) error {
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

func readArrayWithTypeByMapIfaceStr(v map[interface{}]string, callback func([]byte, string)) error {
	for key, val := range v {

		callback([]byte(fmt.Sprint(key) + ": " + val))
	}

	return nil
}
*/
