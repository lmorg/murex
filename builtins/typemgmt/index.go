package typemgmt

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/csv"
	"strconv"
)

type jsonInterface map[interface{}]interface{}

func init() {
	proc.GoFunctions["["] = proc.GoFunction{Func: array, TypeIn: types.Generic, TypeOut: types.Generic}
}

func array(p *proc.Process) (err error) {
	end, err := p.Parameters.String(p.Parameters.Len() - 1)
	if err != nil {
		return err
	}
	if end != "]" {
		return errors.New("Missing closing bracket, ` ]`")
	}

	params := p.Parameters.StringArray()[:p.Parameters.Len()-1]

	switch p.Stdin.GetDataType() {
	case types.Json:
		p.Stdout.SetDataType(types.Json)

		var jInterface interface{}

		if err = json.Unmarshal(p.Stdin.ReadAll(), &jInterface); err != nil {
			return
		}

		var jArray []interface{}
		switch v := jInterface.(type) {
		case map[string]interface{}:
			for _, key := range params {
				if v[key] == nil {
					return errors.New("Key '" + key + "' not found.")
				}

				if len(params) > 1 {
					jArray = append(jArray, v[key])

				} else {
					switch v[key].(type) {
					case string:
						p.Stdout.Write([]byte(v[key].(string)))
					default:
						b, err := json.Marshal(v[key])
						if err != nil {
							return err
						}
						p.Stdout.Writeln(b)
					}
				}
			}
			if len(jArray) > 0 {
				b, err := json.Marshal(jArray)
				if err != nil {
					return err
				}
				p.Stdout.Writeln(b)
			}
			return nil

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
					jArray = append(jArray, v[i])

				} else {
					switch v[i].(type) {
					case string:
						p.Stdout.Write([]byte(v[i].(string)))
					default:
						b, err := json.Marshal(v[i])
						if err != nil {
							return err
						}
						p.Stdout.Writeln(b)
					}
				}
			}
			if len(jArray) > 0 {
				b, err := json.Marshal(jArray)
				if err != nil {
					return err
				}
				p.Stdout.Writeln(b)
			}
			return nil

		default:
			return errors.New("JSON object cannot be indexed by array")
		}

	case types.Csv:
		p.Stdout.SetDataType(types.Csv)

		match := make(map[string]int)
		for i := range params {
			match[params[i]] = i + 1
		}

		csvParser, err := csv.NewParser(nil, &proc.GlobalConf)
		if err != nil {
			return err
		}
		records := make([]string, len(params)+1)
		var matched bool

		p.Stdin.ReadMap(&proc.GlobalConf, func(key, value string, last bool) {
			if match[key] != 0 {
				matched = true
				records[match[key]] = value
			}

			if last && matched {
				p.Stdout.Writeln(csvParser.ArrayToCsv(records[1:]))
				matched = false
				records = make([]string, len(params)+1)
			}
		})

	case types.String, types.Generic:
		p.Stdout.SetDataType(types.String)

		match := make(map[string]bool)
		for i := range params {
			match[params[i]] = true
		}

		p.Stdin.ReadMap(&proc.GlobalConf, func(key, value string, last bool) {
			if match[key] {
				p.Stdout.Writeln([]byte(value))
			}
		})

	default:
		p.Stdout.SetDataType(types.Null)
		err = errors.New("I don't know how to get an index from this data type")
	}

	return err
}
