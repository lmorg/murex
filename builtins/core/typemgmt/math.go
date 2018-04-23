package typemgmt

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/Knetic/govaluate"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	proc.GoFunctions["eval"] = cmdEval
	proc.GoFunctions["="] = cmdEval
	proc.GoFunctions["let"] = cmdLet
}

var (
	rxLet   *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)\s*=(.*)$`)
	rxMinus *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)--$`)
	rxPlus  *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)\+\+$`)
)

func cmdEval(p *proc.Process) error {
	p.Stdout.SetDataType(types.Generic)

	if p.Parameters.Len() == 0 {
		return errors.New("Missing expression.")
	}

	var leftSide string

	if p.IsMethod {
		if debug.Enable == false {
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

	value, err := evaluate(p, leftSide+p.Parameters.StringAll())
	if err != nil {
		return err
	}

	_, err = p.Stdout.Write([]byte(value))
	return err
}

func cmdLet(p *proc.Process) (err error) {
	p.Stdout.SetDataType(types.Null)

	if debug.Enable == false {
		defer func() {
			if r := recover(); r != nil {
				p.ExitNum = 2
				err = errors.New(fmt.Sprint("Panic caught: ", r))
			}
		}()
	}

	params := p.Parameters.StringAll()
	var variable, expression string

	switch {
	case rxLet.MatchString(params):
		match := rxLet.FindAllStringSubmatch(params, -1)
		variable = match[0][1]
		expression = match[0][2]

	case rxPlus.MatchString(params):
		match := rxPlus.FindAllStringSubmatch(params, -1)
		variable = match[0][1]
		expression = variable + "+1"

	case rxMinus.MatchString(params):
		match := rxMinus.FindAllStringSubmatch(params, -1)
		variable = match[0][1]
		expression = variable + "-1"

	default:
		return errors.New("Invalid syntax for `let`. Should be `let variable-name = expression`.")
	}

	value, err := evaluate(p, expression)
	if err != nil {
		return err
	}

	err = p.Variables.Set(variable, value, types.Number)

	return err
}

func evaluate(p *proc.Process, expression string) (value string, err error) {
	if debug.Enable == false {
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

	result, err := eval.Evaluate(p.Variables.DumpMap())
	if err != nil {
		return
	}

	s, err := types.ConvertGoType(result, types.String)
	if err == nil {
		value = s.(string)
	}
	return
}
