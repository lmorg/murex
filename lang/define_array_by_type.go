package lang

import (
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

// ArrayByTypeTemplate is a template function for reading arrays from marshalled data
func ArrayByTypeTemplate(dataType string, marshal func(interface{}) ([]byte, error), unmarshal func([]byte, interface{}) error, read stdio.Io, callback func([]byte, string)) error {
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
		return readArrayByTypeByString(v.(string), callback)

	case []string:
		return readArrayByTypeBySliceString(v.([]string), callback)

	case []interface{}:
		return readArrayByTypeBySliceInterface(dataType, marshal, v.([]interface{}), callback)

	/*case map[string]string:
		return readArrayByTypeByMapStrStr(v.(map[string]string), callback)

	case map[string]interface{}:
		return readArrayByTypeByMapStrIface(marshal, v.(map[string]interface{}), callback)

	case map[interface{}]string:
		return readArrayByTypeByMapIfaceStr(v.(map[interface{}]string), callback)

	case map[interface{}]interface{}:
		return readArrayByTypeByMapIfaceIface(marshal, v.(map[interface{}]interface{}), callback)
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

func readArrayByTypeByString(v string, callback func([]byte, string)) error {
	callback([]byte(v), types.String)

	return nil
}

func readArrayByTypeBySliceString(v []string, callback func([]byte, string)) error {
	for i := range v {
		callback([]byte(v[i]), types.String)
	}

	return nil
}

func readArrayByTypeBySliceInterface(dataType string, marshal func(interface{}) ([]byte, error), v []interface{}, callback func([]byte, string)) error {
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

/*func readArrayByTypeByMapIfaceIface(marshal func(interface{}) ([]byte, error), v map[interface{}]interface{}, callback func([]byte, string)) error {
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

func readArrayByTypeByMapStrStr(v map[string]string, callback func([]byte, string)) error {
	for key, val := range v {

		callback([]byte(key + ": " + val))
	}

	return nil
}

func readArrayByTypeByMapStrIface(marshal func(interface{}) ([]byte, error), v map[string]interface{}, callback func([]byte, string)) error {
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

func readArrayByTypeByMapIfaceStr(v map[interface{}]string, callback func([]byte, string)) error {
	for key, val := range v {

		callback([]byte(fmt.Sprint(key) + ": " + val))
	}

	return nil
}
*/
