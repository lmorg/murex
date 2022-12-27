package lang

import (
	"context"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
)

// ArrayWithTypeTemplate is a template function for reading arrays from marshalled data
func ArrayWithTypeTemplate(ctx context.Context, dataType string, marshal func(interface{}) ([]byte, error), unmarshal func([]byte, interface{}) error, read stdio.Io, callback func(interface{}, string)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	if len(utils.CrLfTrim(b)) == 0 {
		return nil
	}

	var v interface{}
	err = unmarshal(b, &v)

	if err != nil {
		return err
	}

	switch v := v.(type) {
	case []interface{}:
		return readArrayWithTypeBySliceInterface(ctx, dataType, marshal, v, callback)

	case []string:
		return readArrayWithTypeBySliceString(ctx, v, callback)

	case []float64:
		return readArrayWithTypeBySliceFloat(ctx, v, callback)

	case []int:
		return readArrayWithTypeBySliceInt(ctx, v, callback)

	case string:
		return readArrayWithTypeByString(v, callback)

	case []byte:
		return readArrayWithTypeByString(string(v), callback)

	case []rune:
		return readArrayWithTypeByString(string(v), callback)

	case []bool:
		return readArrayWithTypeBySliceBool(ctx, v, callback)

	/*case map[string]string:
		return readArrayWithTypeByMapStrStr(v, callback)

	case map[string]interface{}:
		return readArrayWithTypeByMapStrIface(marshal, v, callback)

	case map[interface{}]string:
		return readArrayWithTypeByMapIfaceStr(v, callback)

	case map[interface{}]interface{}:
		return readArrayWithTypeByMapIfaceIface(marshal, v, callback)
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

func readArrayWithTypeByString(v string, callback func(interface{}, string)) error {
	callback(v, types.String)

	return nil
}

func readArrayWithTypeBySliceInt(ctx context.Context, v []int, callback func(interface{}, string)) error {
	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback(v[i], types.Integer)
		}
	}

	return nil
}

func readArrayWithTypeBySliceFloat(ctx context.Context, v []float64, callback func(interface{}, string)) error {
	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback(v[i], types.Number)
		}
	}

	return nil
}

func readArrayWithTypeBySliceBool(ctx context.Context, v []bool, callback func(interface{}, string)) error {
	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback(v[i], types.Boolean)

		}
	}

	return nil
}

func readArrayWithTypeBySliceString(ctx context.Context, v []string, callback func(interface{}, string)) error {
	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			callback(v[i], types.String)
		}
	}

	return nil
}

func readArrayWithTypeBySliceInterface(ctx context.Context, dataType string, marshal func(interface{}) ([]byte, error), v []interface{}, callback func(interface{}, string)) error {
	if len(v) == 0 {
		return nil
	}

	for i := range v {
		select {
		case <-ctx.Done():
			return nil

		default:
			switch v[i].(type) {

			case string:
				callback((v[i].(string)), types.String)

			case float64:
				callback(v[i].(float64), types.Number)

			case int:
				callback(v[i].(int), types.Integer)

			case bool:
				if v[i].(bool) {
					callback(true, types.Boolean)
				} else {
					callback(false, types.Boolean)
				}

			case []byte:
				callback(string(v[i].([]byte)), types.String)

			case nil:
				callback(nil, types.Null)

			default:
				jBytes, err := marshal(v[i])
				if err != nil {
					return err
				}
				callback(jBytes, dataType)
			}
		}
	}

	return nil
}

/*func readArrayWithTypeByMapIfaceIface(marshal func(interface{}) ([]byte, error), v map[interface{}]interface{}, callback func([]byte, string)) error {
	for key, val := range v {
	select {
		case <-ctx.Done():
			return nil
		default:
		bKey := []byte(fmt.Sprint(key) + ": ")
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}}

	return nil
}

func readArrayWithTypeByMapStrStr(v map[string]string, callback func([]byte, string)) error {
	for key, val := range v {
	select {
		case <-ctx.Done():
			return nil
		default:
		callback([]byte(key + ": " + val))
	}}

	return nil
}

func readArrayWithTypeByMapStrIface(marshal func(interface{}) ([]byte, error), v map[string]interface{}, callback func([]byte, string)) error {
	for key, val := range v {
	select {
		case <-ctx.Done():
			return nil
		default:
		bKey := []byte(key + ": ")
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}}

	return nil
}

func readArrayWithTypeByMapIfaceStr(v map[interface{}]string, callback func([]byte, string)) error {
	for key, val := range v {
	select {
		case <-ctx.Done():
			return nil
		default:
		callback([]byte(fmt.Sprint(key) + ": " + val))
	}}

	return nil
}
*/
