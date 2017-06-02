package textmanip

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

type jsonInterface map[interface{}]interface{}

func init() {
	proc.GoFunctions["json"] = proc.GoFunction{Func: cmdJson, TypeIn: types.Json, TypeOut: types.Generic}
	proc.GoFunctions["prettify"] = proc.GoFunction{Func: cmdPrettify, TypeIn: types.Json, TypeOut: types.String}
}

func cmdJson(p *proc.Process) (err error) {
	var jInterface interface{}

	if err = json.Unmarshal(p.Stdin.ReadAll(), &jInterface); err != nil {
		return
	}

	for _, field := range p.Parameters.StringArray() {
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

func cmdPrettify(p *proc.Process) (err error) {
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, p.Stdin.ReadAll(), "", "\t")
	p.Stdout.Write(prettyJSON.Bytes())

	return
}
