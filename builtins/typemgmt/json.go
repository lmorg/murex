package typemgmt

import (
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

type jsonInterface map[interface{}]interface{}

func init() {
	proc.GoFunctions["["] = proc.GoFunction{Func: indexJson, TypeIn: types.Json, TypeOut: types.Generic}
}

func indexJson(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Generic)

	var jInterface interface{}

	end, err := p.Parameters.String(p.Parameters.Len() - 1)
	if err != nil {
		return err
	}
	if end != "]" {
		return errors.New("Missing closing bracket, ` ]`")
	}

	if err = json.Unmarshal(p.Stdin.ReadAll(), &jInterface); err != nil {
		return
	}

	for _, field := range p.Parameters.StringArray()[:p.Parameters.Len()-1] {
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

	b, err := json.MarshalIndent(jInterface, "", "\t")
	p.Stdout.Write(b)
	return err
}
