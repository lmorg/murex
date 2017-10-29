package typemgmt

import (
	"errors"
	"fmt"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
	"strings"
)

type jsonInterface map[interface{}]interface{}

func init() {
	proc.GoFunctions["["] = index

	proc.GlobalConf.Define("index", "silent", config.Properties{
		Description: "Don't report error if an index in [ ] does not exist",
		Default:     false,
		DataType:    types.Boolean,
	})
}

func index(p *proc.Process) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic caught: %s", r)
		}
	}()

	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	params := p.Parameters.StringArray()
	l := len(params) - 1
	if l < 0 {
		return errors.New("Missing parameters. Please select 1 or more indexes.")
	}
	switch {
	case params[l] == "]":
		params = params[:l]
	case strings.HasSuffix(params[l], "]"):
		params[l] = params[l][:len(params[l])-1]
	default:
		return errors.New("Missing closing bracket, ` ]`")
	}

	if define.ReadIndexes[dt] != nil {
		silent, err := proc.GlobalConf.Get("index", "silent", types.Boolean)
		if err != nil {
			silent = false
		}

		err = define.ReadIndexes[dt](p, params)
		if silent.(bool) {
			return nil
		}
		return err

	}

	return errors.New("I don't know how to get an index from this data type: `" + dt + "`")
}
