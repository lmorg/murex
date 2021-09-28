package element

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("[[", element, types.Unmarshal, types.Marshal)
}

func element(p *lang.Process) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic caught: %s", r)
		}
	}()

	dt := p.Stdin.GetDataType()
	//p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	params := p.Parameters.StringArray()

	switch len(params) {
	case 0:
		return errors.New("Missing parameter. Requires nested index")

	case 1:
		if strings.HasSuffix(params[0], "]]") {
			params[0] = params[0][0 : len(params[0])-2]
		} else {
			return errors.New("Missing closing brackets, ` ]]`")
		}

	case 2:
		last := len(params) - 1
		if strings.HasSuffix(params[last], "]]") {
			if len(params[last]) != 2 {
				return errors.New("Too many parameters")
			}
		} else {
			return errors.New("Missing closing brackets, ` ]]`")
		}

	default:
		return errors.New("Too many parameters")
	}

	path := strings.Split(params[0], params[0][0:1])

	obj, err := lang.UnmarshalData(p, dt)
	if err != nil {
		return err
	}

	for i := 1; i < len(path); i++ {
		if len(path[i]) == 0 {
			if i == len(path)-1 {
				break
			} else {
				return fmt.Errorf("Path element %d is a zero length string: '%s'", i-1, strings.Join(params, "/"))
			}
		}

		obj, err = recursiveLookup(path, i, obj)
		if err != nil {
			return err
		}
	}

	switch v := obj.(type) {
	case string:
		p.Stdout.SetDataType(types.String)
		_, err = p.Stdout.Write([]byte(v))
	case []byte:
		p.Stdout.SetDataType(types.String)
		_, err = p.Stdout.Write(v)
	case int:
		p.Stdout.SetDataType(types.Integer)
		_, err = p.Stdout.Write([]byte(strconv.Itoa(v)))
	case int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		p.Stdout.SetDataType(types.Integer)
		_, err = fmt.Fprint(p.Stdout, v)
	case float32:
		p.Stdout.SetDataType(types.Float)
		_, err = p.Stdout.Write([]byte(types.FloatToString(float64(v))))
	case float64:
		p.Stdout.SetDataType(types.Float)
		_, err = p.Stdout.Write([]byte(types.FloatToString(v)))
	case bool:
		p.Stdout.SetDataType(types.Boolean)
		if v {
			_, err = p.Stdout.Write(types.TrueByte)
		} else {
			_, err = p.Stdout.Write(types.FalseByte)
		}
	default:
		p.Stdout.SetDataType(dt)
		b, err := lang.MarshalData(p, dt, obj)
		if err != nil {
			return err
		}
		_, err = p.Stdout.Write(b)
		if err != nil {
			return err
		}
	}

	return err
}

func recursiveLookup(path []string, i int, obj interface{}) (interface{}, error) {
	switch v := obj.(type) {
	case []string:
		i, err := isValidIndex(path[i], len(v))
		if err != nil {
			return nil, err
		}
		return v[i], nil

	case []interface{}:
		i, err := isValidIndex(path[i], len(v))
		if err != nil {
			return nil, err
		}
		return v[i], nil

	case map[string]interface{}:
		switch {
		case v[path[i]] != nil:
			return v[path[i]], nil
		case v[strings.Title(path[i])] != nil:
			return v[strings.Title(path[i])], nil
		case v[strings.ToLower(path[i])] != nil:
			return v[strings.ToLower(path[i])], nil
		case v[strings.ToUpper(path[i])] != nil:
			return v[strings.ToUpper(path[i])], nil
			//case v[strings.ToTitle(params[i])] != nil:
			//	return v[strings.ToTitle(path[i])], nil
		default:
			return nil, fmt.Errorf("Key '%s' not found", path[i])
		}

	case map[interface{}]interface{}:
		switch {
		case v[path[i]] != nil:
			return v[path[i]], nil
		case v[strings.Title(path[i])] != nil:
			return v[strings.Title(path[i])], nil
		case v[strings.ToLower(path[i])] != nil:
			return v[strings.ToLower(path[i])], nil
		case v[strings.ToUpper(path[i])] != nil:
			return v[strings.ToUpper(path[i])], nil
			//case v[strings.ToTitle(params[i])] != nil:
			//	return v[strings.ToTitle(path[i])], nil
		default:
			return nil, fmt.Errorf("Key '%s' not found", path[i])
		}

	default:
		return nil, fmt.Errorf("I don't know how to lookup %T (please file a bug with on the murex Github page: https://lmorg/murex)", v)
	}
}

func isValidIndex(key string, length int) (int, error) {
	i, err := strconv.Atoi(key)
	if err != nil {
		return 0, fmt.Errorf("Element is an array however supplied key, '%s', is not an integer", key)
	}

	if i < 0 {
		return 0, fmt.Errorf("Negative keys are not allowed for arrays: %s", key)
	}

	if i >= length {
		return 0, fmt.Errorf("Element is an array however key is greater than the length: %s", key)
	}

	return i, nil
}
