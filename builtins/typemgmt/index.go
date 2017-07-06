package typemgmt

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"strconv"
	"strings"
)

type jsonInterface map[interface{}]interface{}

func init() {
	proc.GoFunctions["["] = proc.GoFunction{Func: array, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["table"] = proc.GoFunction{Func: cmdTable, TypeIn: types.Generic, TypeOut: types.Csv}
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
				//b, err := json.Marshal(v[key])
				//if err != nil {
				//	return err
				//}

				if len(params) > 1 {
					jArray = append(jArray, v[key])
				} else {
					switch v[key].(type) {
					case string:
						p.Stdout.Write([]byte(v[key].(string)))
					default:
						b, err := json.Marshal(jArray)
						if err != nil {
							return err
						}
						p.Stdout.Write(b)
					}
				}
			}
			if len(jArray) > 0 {
				b, err := json.Marshal(jArray)
				if err != nil {
					return err
				}
				p.Stdout.Write(b)
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
				//b, err := utils.JsonMarshal(v[i])
				//if err != nil {
				//	return err
				//}
				//p.Stdout.Writeln(b)

				if len(params) > 1 {
					jArray = append(jArray, v[i])
				} else {
					b, err := json.Marshal(v[i])
					if err != nil {
						return err
					}
					p.Stdout.Write(b)
				}
			}
			if len(jArray) > 0 {
				b, err := json.Marshal(jArray)
				if err != nil {
					return err
				}
				p.Stdout.Write(b)
			}
			return nil

		default:
			return errors.New("JSON object cannot be indexed by array")
		}

	case types.Csv:
		p.Stdout.SetDataType(types.Csv)

		match := make(map[string]bool)
		for i := range params {
			match[params[i]] = true
		}

		//v, _ := proc.GlobalConf.Get("shell", "Csv-Headings", types.Boolean)
		//useHeadings := v.(bool)

		v, _ := proc.GlobalConf.Get("shell", "Csv-Separator", types.String)
		separator := v.(string)[:1]

		p.Stdin.ReadMap(&proc.GlobalConf, func(key, value string, last bool) {
			if match[key] {
				if !last {
					p.Stdout.Write([]byte(value + separator))
				} else {
					p.Stdout.Write([]byte(value + "\n"))
				}
			}
		})

	default:
		p.Stdout.SetDataType(types.Null)
		err = errors.New("I don't know how to get an index from this data type")
	}

	return err
}

func cmdTable(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Csv)

	separator, err := p.Parameters.String(0)
	if err != nil {
		return
	}

	var (
		a []string
		s string
	)

	join := func(b []byte) {
		a = append(a, string(b))
	}

	if p.IsMethod {
		p.Stdin.ReadArray(join)
		s = strings.Join(a, separator)
	} else {
		s = strings.Join(p.Parameters.StringArray()[1:], string(separator))
	}

	_, err = p.Stdout.Writeln([]byte(s))
	return
}
