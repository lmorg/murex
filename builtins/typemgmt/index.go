package typemgmt

import (
	"errors"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/data"
	"strings"
)

type jsonInterface map[interface{}]interface{}

func init() {
	proc.GoFunctions["["] = index
}

func index(p *proc.Process) (err error) {
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

	dt := p.Stdin.GetDataType()

	if data.ReadIndexes[dt] != nil {
		p.Stdout.SetDataType(dt)
		return data.ReadIndexes[dt](p, params)
	}

	p.Stdout.SetDataType(types.Null)
	err = errors.New("I don't know how to get an index from this data type: `" + dt + "`")

	return err
}
