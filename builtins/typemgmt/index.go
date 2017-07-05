package typemgmt

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils"
	"strconv"
)

type jsonInterface map[interface{}]interface{}

func init() {
	proc.GoFunctions["$["] = proc.GoFunction{Func: index, TypeIn: types.Generic, TypeOut: types.Generic}
	proc.GoFunctions["@["] = proc.GoFunction{Func: array, TypeIn: types.Generic, TypeOut: types.Generic}
}

func index(p *proc.Process) (err error) {
	dt := p.Stdin.GetDataType()

	end, err := p.Parameters.String(p.Parameters.Len() - 1)
	if err != nil {
		return err
	}
	if end != "]" {
		return errors.New("Missing closing bracket, ` ]`")
	}

	params := p.Parameters.StringArray()[:p.Parameters.Len()-1]

	switch dt {
	case types.Json:
		p.Stdout.SetDataType(types.Json)

		var jInterface interface{}

		if err = json.Unmarshal(p.Stdin.ReadAll(), &jInterface); err != nil {
			return
		}

		for _, field := range params {
			switch v := jInterface.(type) {
			case map[string]interface{}:
				jInterface = v[field]

			case string:
				jInterface = v

			default:
				errors.New("Unable to find " + p.Parameters.StringAll() + " in JSON.")
				return
			}
		}

		var b []byte
		b, err = json.MarshalIndent(jInterface, "", "\t")
		p.Stdout.Write(b)

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

func array(p *proc.Process) (err error) {
	dt := p.Stdin.GetDataType()

	end, err := p.Parameters.String(p.Parameters.Len() - 1)
	if err != nil {
		return err
	}
	if end != "]" {
		return errors.New("Missing closing bracket, ` ]`")
	}

	params := p.Parameters.StringArray()[:p.Parameters.Len()-1]

	switch dt {
	case types.Json:
		p.Stdout.SetDataType(types.Json)

		var jInterface interface{}

		if err = json.Unmarshal(p.Stdin.ReadAll(), &jInterface); err != nil {
			return
		}

		switch v := jInterface.(type) {
		case map[string]interface{}:
			for _, key := range params {
				if v[key] == nil {
					return errors.New("Key '" + key + "' not found.")
				}
				b, err := utils.JsonMarshal(v[key])
				if err != nil {
					return err
				}
				p.Stdout.Writeln(b)
			}

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
				b, err := utils.JsonMarshal(v[i])
				if err != nil {
					return err
				}
				p.Stdout.Writeln(b)
			}

		default:
			errors.New("JSON object cannot be indexed by array")
			return
		}

		var b []byte
		b, err = json.MarshalIndent(jInterface, "", "\t")
		p.Stdout.Write(b)

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
