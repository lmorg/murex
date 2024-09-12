package typemgmt

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Knetic/govaluate"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	lang.DefineMethod("=", cmdEqu, types.Math, types.Math)
}

func cmdEqu(p *lang.Process) (err error) {
	lang.FeatureDeprecatedBuiltin(p)

	if p.Parameters.Len() == 0 {
		return errors.New("missing expression")
	}

	var leftSide string

	if p.IsMethod {
		if !debug.Enabled {
			defer func() {
				if r := recover(); r != nil {
					p.ExitNum = 2
					err = errors.New(fmt.Sprint("Panic caught: ", r))
				}
			}()
		}

		dt := p.Stdin.GetDataType()
		b, err := p.Stdin.ReadAll()
		if err != nil {
			return err
		}

		v, err := types.ConvertGoType(b, dt)
		if err != nil {
			return err
		}

		switch dt {
		case types.Integer:
			leftSide = strconv.Itoa(v.(int))

		case types.Float, types.Number:
			leftSide = types.FloatToString(v.(float64))

		case types.Boolean:
			if v.(bool) {
				leftSide = "true"
			} else {
				leftSide = "false"
			}

		default:
			leftSide = `"` + v.(string) + `"`
		}
	}

	value, dt, err := evaluate(p, leftSide+p.Parameters.StringAll())
	if err != nil {
		return err
	}

	s, err := types.ConvertGoType(value, types.String)
	if err != nil {
		return fmt.Errorf("unable to convert result to text: %s", err.Error())
	}

	p.Stdout.SetDataType(dt)
	_, err = p.Stdout.Write([]byte(s.(string)))
	return err
}

func evaluate(p *lang.Process, expression string) (value interface{}, dataType string, err error) {
	if !debug.Enabled {
		defer func() {
			if r := recover(); r != nil {
				p.ExitNum = 2
				err = errors.New(fmt.Sprint("Panic caught: ", r))
			}
		}()
	}

	eval, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return
	}

	value, err = eval.Evaluate(lang.DumpVariables(p))
	if err != nil {
		return
	}

	dataType = types.DataTypeFromInterface(value)
	if dataType == types.String {
		// this is purely here for backwards compatibility. All other consumers of DataTypeFromInterface should expect a string instead
		dataType = types.Generic
	}

	return
}
