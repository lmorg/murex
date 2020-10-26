package index

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.GoFunctions["["] = index
	lang.GoFunctions["!["] = index

	config.InitConf.Define("index", "silent", config.Properties{
		Description: "Don't report error if an index in [ ] does not exist",
		Default:     false,
		DataType:    types.Boolean,
	})
}

func index(p *lang.Process) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Panic caught: %s", r)
		}
	}()

	dt := p.Stdin.GetDataType()
	p.Stdout.SetDataType(dt)

	if err := p.ErrIfNotAMethod(); err != nil {
		return err
	}

	params := p.Parameters.StringArray()
	l := len(params) - 1
	if l < 0 {
		return errors.New("Missing parameters. Please select 1 or more indexes")
	}
	switch {
	case params[l] == "]":
		params = params[:l]
	case strings.HasSuffix(params[l], "]"):
		params[l] = params[l][:len(params[l])-1]
	default:
		return errors.New("Missing closing bracket, ` ]`")
	}

	var f func(p *lang.Process, params []string) error
	if p.IsNot {
		f = lang.ReadNotIndexes[dt]
		if f == nil {
			return errors.New("I don't know how to get an !index from this data type: `" + dt + "`")
		}
	} else {
		f = lang.ReadIndexes[dt]
		if f == nil {
			return errors.New("I don't know how to get an index from this data type: `" + dt + "`")
		}
	}

	silent, err := p.Config.Get("index", "silent", types.Boolean)
	if err != nil {
		silent = false
	}

	err = f(p, params)
	if silent.(bool) {
		return nil
	}

	return err
}
