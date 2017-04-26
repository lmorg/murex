package math

import (
	"errors"
	"fmt"
	gov "github.com/Knetic/govaluate"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"regexp"
	"strings"
)

func init() {
	proc.GoFunctions["eval"] = proc.GoFunction{Func: cmdEval, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["let"] = proc.GoFunction{Func: cmdLet, TypeIn: types.Null, TypeOut: types.Boolean}
}

var rxLet *regexp.Regexp = regexp.MustCompile(`^([-_a-zA-Z0-9]+)\s*=(.*)`)

func cmdEval(p *proc.Process) (err error) {
	if p.Parameters.Len() == 0 {
		return errors.New("Missing expression.")
	}
	value, err := evaluate(p, p.Parameters.AllString())
	if err != nil {
		return
	}

	_, err = p.Stdout.Write([]byte(value))
	return
}

func cmdLet(p *proc.Process) (err error) {
	defer func() {
		if r := recover(); r != nil {
			p.ExitNum = 2
			err = errors.New(fmt.Sprint("Panic caught: ", r))
		}
	}()

	params := p.Parameters.AllString()

	if !rxLet.MatchString(params) {
		return errors.New("Invalid syntax for `let`. Should be `let variable-name = expression`.")
	}

	match := rxLet.FindAllStringSubmatch(params, -1)

	value, err := evaluate(p, match[0][2])
	if err != nil {
		return err
	}

	err = proc.GlobalVars.Set(match[0][1], value, types.Float)

	return err
}

func evaluate(p *proc.Process, expression string) (value string, err error) {
	defer func() {
		if r := recover(); r != nil {
			p.ExitNum = 2
			err = errors.New(fmt.Sprint("Panic caught: ", r))
		}
	}()

	eval, err := gov.NewEvaluableExpression(expression)
	if err != nil {
		return
	}

	result, err := eval.Evaluate(nil)
	if err != nil {
		return
	}

	goType, err := types.ConvertGoType(result, types.String)
	if err != nil {
		return
	}

	value = strings.Replace(goType.(string), ".000000", "", 1) // TODO: this is hacky!!!
	return
}
