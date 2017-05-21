package typemgmt

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"regexp"
)

func init() {
	proc.GoFunctions["eval"] = proc.GoFunction{Func: cmdEval, TypeIn: types.Null, TypeOut: types.Number}
	proc.GoFunctions["let"] = proc.GoFunction{Func: cmdLet, TypeIn: types.Null, TypeOut: types.Null}

	/*proc.GlobalConf.Define("murex", "Auto-Declare", config.Properties{
		Description: "If a number is undefined then auto-declare it as zero.",
		Default:     true,
		DataType:    types.Boolean,
	})*/
}

var (
	rxLet   *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)\s*=(.*)$`)
	rxMinus *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)--$`)
	rxPlus  *regexp.Regexp = regexp.MustCompile(`^([_a-zA-Z0-9]+)\+\+$`)
)

func cmdEval(p *proc.Process) (err error) {
	if p.Parameters.Len() == 0 {
		return errors.New("Missing expression.")
	}
	value, err := evaluate(p, p.Parameters.StringAll())
	if err != nil {
		return
	}

	_, err = p.Stdout.Write([]byte(value))
	return
}

func cmdLet(p *proc.Process) (err error) {
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

	err = proc.GlobalVars.Set(variable, value, types.Number)

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

	result, err := eval.Evaluate(proc.GlobalVars.DumpMap())
	if err != nil {
		return
	}

	/*
		autoDeclare, err := proc.GlobalConf.Get("murex", "Auto-Declare", types.Boolean)
		if err != nil {
			return
		}
		p.Stderr.Writeln([]byte(fmt.Sprint(autoDeclare)))
		if autoDeclare.(bool) {
			f, err := types.ConvertGoType(result, types.Float)
			if err != nil {
				return "", err
			}
			value = types.FloatToString(f.(float64))
			p.Stderr.Writeln([]byte(value))
			return value, err
		}
	*/

	s, err := types.ConvertGoType(result, types.String)
	if err == nil {
		value = s.(string)
	}
	return
}
