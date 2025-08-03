package lang

import (
	"strconv"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

func MapTemplate(dataType string, marshal func(any) ([]byte, error), unmarshal func([]byte, any) error, read stdio.Io, callback func(*stdio.Map)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var obj any
	err = unmarshal(b, &obj)
	if err != nil {
		return err
	}

	switch t := obj.(type) {
	case []any:
		for i := range t {
			b, err := marshal(t[i])
			if err != nil {
				return err
			}
			callback(&stdio.Map{
				Key:      strconv.Itoa(i),
				Value:    string(b),
				DataType: dataType,
				Last:     i != len(t),
			})
		}

	case map[string]string:
		return mapTemplateStringOfPrimitives(t, callback)
	case map[string]float64:
		return mapTemplateStringOfPrimitives(t, callback)
	case map[string]int:
		return mapTemplateStringOfPrimitives(t, callback)
	case map[string]bool:
		return mapTemplateStringOfPrimitives(t, callback)

	case map[string]any:
		return mapTemplateStringOfObjects(marshal, t, dataType, callback)

	case map[any]any:
		return readMapAsInterfaceOfInterfaces(marshal, t, dataType, callback)

	default:
		if debug.Enabled {
			panic(t)
		}
	}
	return nil
}

func mapTemplateStringOfPrimitives[V string | int | float64 | bool](t map[string]V, callback func(*stdio.Map)) error {
	i := 1
	for key, val := range t {
		callback(&stdio.Map{
			Key:      key,
			Value:    val,
			DataType: types.DataTypeFromInterface(val),
			Last:     i != len(t),
		})

		i++
	}
	return nil
}

func mapTemplateStringOfObjects(marshal func(any) ([]byte, error), t map[string]any, dataType string, callback func(*stdio.Map)) error {
	i := 1
	for key, val := range t {
		switch val.(type) {
		case string, int, float64, bool, nil:
			callback(&stdio.Map{
				Key:      key,
				Value:    val,
				DataType: types.DataTypeFromInterface(val),
				Last:     i != len(t),
			})

		default:
			b, err := marshal(val)
			if err != nil {
				return err
			}

			callback(&stdio.Map{
				Key:      key,
				Value:    string(b),
				DataType: dataType,
				Last:     i != len(t),
			})
		}

		i++
	}
	return nil
}

func readMapAsInterfaceOfInterfaces(marshal func(any) ([]byte, error), t map[any]any, dataType string, callback func(*stdio.Map)) error {
	i := 1
	for key, val := range t {
		s, err := types.ConvertGoType(key, types.String)
		if err != nil {
			return err
		}

		switch val.(type) {
		case string, int, float64, bool, nil:
			callback(&stdio.Map{
				Key:      s.(string),
				Value:    val,
				DataType: types.DataTypeFromInterface(val),
				Last:     i != len(t),
			})

		default:
			b, err := marshal(val)
			if err != nil {
				return err
			}

			callback(&stdio.Map{
				Key:      s.(string),
				Value:    string(b),
				DataType: dataType,
				Last:     i != len(t),
			})
		}

		i++
	}
	return nil
}
