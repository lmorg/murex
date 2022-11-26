package expressions

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/expressions"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineFunction("exp", cmdExpressions, types.Any)
}

func cmdExpressions(p *lang.Process) error {
	getVar := func(name string) (interface{}, string, error) {
		value := p.Variables.GetValue(name)
		dataType := p.Variables.GetDataType(name)
		return value, dataType, nil
	}

	setVar := func(name string, value interface{}, dataType string) error {
		return p.Variables.Set(p, name, value, dataType)
	}

	expression := []rune(p.Parameters.StringAll())

	result, err := expressions.Execute(expression, getVar, setVar)
	if err != nil {
		return err
	}

	dt := result.DataType()
	p.Stdout.SetDataType(dt)

	b, err := lang.MarshalData(p, dt, result.Value)
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write(b)
	return err
}
