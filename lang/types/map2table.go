package types

import (
	"fmt"
	"reflect"
	"sort"
)

func MapToTable(v []any, callback func([]string) error) error {
	if len(v) == 0 {
		return nil
	}

	if reflect.TypeOf(v[0]).Kind() != reflect.Map {
		return fmt.Errorf("expecting map on row %d, instead got a %s", 0, reflect.TypeOf(v[0]).Kind().String())
	}

	headings, err := getMapKeys(v[0].(map[string]any))
	if err != nil {
		return err
	}

	//table := make([][]string, len(v)+1)
	//table[0] = headings
	err = callback(headings)
	if err != nil {
		return err
	}

	lenHeadings := len(headings)
	slice := make([]string, lenHeadings)
	var j int

	for i := range v {
		//if reflect.TypeOf(v[i]).Kind() != reflect.Map {
		//	return nil, fmt.Errorf("expecting map on row %d, instead got a %s", i, reflect.TypeOf(v[i]).Kind().String())
		//}

		m, ok := v[i].(map[string]any)
		if !ok {
			return fmt.Errorf("expecting map on row %d, instead got a %s", i, reflect.TypeOf(v[i]).Kind().String())
		}

		if len(m) != len(headings) {
			return fmt.Errorf("row %d has a different number of records to the first row:\nrow 0 == %d records,\nrow %d == %d records",
				i, lenHeadings, i, len(v))
		}

		for j = 0; j < lenHeadings; j++ {
			val, ok := v[i].(map[string]any)[headings[j]]
			if !ok {
				return fmt.Errorf("row %d is missing a record name found in the first row: '%s'", i, headings[j])
			}
			s, err := ConvertGoType(val, String)
			if err != nil {
				return fmt.Errorf("cannot convert a %T (%v) to a %s in record %d: %s", val, val, String, i, err.Error())
			}
			slice[j] = s.(string)
		}

		err = callback(slice)
		if err != nil {
			return err
		}
	}

	return nil
}

func getMapKeys(v map[string]any) ([]string, error) {
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
