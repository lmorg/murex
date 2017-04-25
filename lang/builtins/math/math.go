package math

import (
	"fmt"
	gov "github.com/Knetic/govaluate"
	"github.com/lmorg/murex/lang/proc"
	"github.com/lmorg/murex/lang/types"
	"strings"
)

func init() {
	proc.GoFunctions["eval"] = proc.GoFunction{Func: cmdEval, TypeIn: types.Null, TypeOut: types.Generic}
	proc.GoFunctions["let"] = proc.GoFunction{Func: cmdLet, TypeIn: types.Null, TypeOut: types.Boolean}
}

func cmdEval(p *proc.Process) error {
	defer func() {
		if r := recover(); r != nil {
			p.ExitNum = 2
			msg := fmt.Sprint("Panic caught: ", r)
			p.Stderr.Writeln([]byte(msg))
		}
	}()

	expression, err := gov.NewEvaluableExpression(p.Parameters.AllString())
	if err != nil {
		return err
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		return err
	}

	cgt, err := types.ConvertGoType(result, types.String)
	if err != nil {
		return err
	}

	s := strings.Replace(cgt.(string), ".000000", "", 1) // TODO: this is hacky!!!
	_, err = p.Stdout.Write([]byte(s))
	return err
}

func cmdLet(p *proc.Process) error {
	params := p.Parameters.AllString()
	// TODO
	defer func() {
		if r := recover(); r != nil {
			p.ExitNum = 2
			msg := fmt.Sprint("Panic caught: ", r)
			p.Stderr.Writeln([]byte(msg))
		}
	}()

	expression, err := gov.NewEvaluableExpression(params)
	if err != nil {
		return err
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		return err
	}

	cgt, err := types.ConvertGoType(result, types.String)
	if err != nil {
		return err
	}

	s := strings.Replace(cgt.(string), ".000000", "", 1) // TODO: this is hacky!!!
	_, err = p.Stdout.Write([]byte(s))
	return err
}
