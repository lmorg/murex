package typemgmt

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

type jsonInterface map[interface{}]interface{}

func init() {
	proc.GoFunctions["["] = proc.GoFunction{Func: indexJson, TypeIn: types.Generic, TypeOut: types.Generic}
	//proc.GoFunctions["@["] = proc.GoFunction{Func: indexJson, TypeIn: types.Generic, TypeOut: types.Generic}
}

func indexJson(p *proc.Process) (err error) {
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

		for _, field := range params {
			switch t := jInterface.(type) {
			case map[string]interface{}:
				jInterface = t[field]

			case string:
				jInterface = t

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
