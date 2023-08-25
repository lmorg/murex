package types

import (
	"fmt"
	"reflect"
	"sort"
)

func MapToTable(v []interface{}) ([][]string, error) {
	if len(v) == 0 {
		return nil, nil
	}

	if reflect.TypeOf(v[0]).Kind() != reflect.Map {
		return nil, fmt.Errorf("expecting map on row %d, instead got a %s", 0, reflect.TypeOf(v[0]).Kind().String())
	}

	headings, err := getMapKeys(v[0].(map[string]any))
	if err != nil {
		return nil, err
	}

	table := make([][]string, len(v)+1)
	table[0] = headings

	lenHeadings := len(headings)
	slice := make([]string, lenHeadings)
	var j int

	for i := range v {
		if reflect.TypeOf(v[i]).Kind() != reflect.Map {
			return nil, fmt.Errorf("expecting map on row %d, instead got a %s", i, reflect.TypeOf(v[i]).Kind().String())
		}

		if len(v[i].(map[string]any)) != len(headings) {
			return nil, fmt.Errorf("row %d has a different number of records to the first row:\nrow 0 == %d records,\nrow %d == %d records",
				i, lenHeadings, i, len(v))
		}

		for j = 0; j < lenHeadings; j++ {
			val, ok := v[i].(map[string]any)[headings[j]]
			if !ok {
				return nil, fmt.Errorf("row %d is missing a record name found in the first row: '%s'", i, headings[j])
			}
			s, err := ConvertGoType(val, String)
			if err != nil {
				return nil, fmt.Errorf("cannot convert a %T (%v) to a %s in record %d: %s", val, val, String, i, err.Error())
			}
			slice[j] = s.(string)
		}

		table[i+1] = slice
	}

	return table, nil
}

func getMapKeys[T comparable](v map[string]T) ([]string, error) {
	slice := make([]string, len(v))
	var i int

	for k := range v {
		s, err := ConvertGoType(k, String)
		if err != nil {
			return nil, fmt.Errorf("cannot convert a %T (%v) to a %s: %s", k, k, String, err.Error())
		}
		slice[i] = s.(string)
		i++
	}
	sort.Strings(slice)
	return slice, nil
}
